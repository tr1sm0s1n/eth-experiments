package config

import (
	"_lab/event-query/artifacts"
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Transactor struct {
	Wg        *sync.WaitGroup
	Ctx       context.Context
	Auth      *bind.TransactOpts
	Backend   bind.Backend
	Address   common.Address
	Instance  *bind.BoundContract
	DataStore *artifacts.DataStore
}

type trxResult struct {
	trx *types.Transaction
	err error
}

func (t *Transactor) DeployContract() (common.Address, error) {
	if err := t.newAuth(); err != nil {
		return common.Address{}, fmt.Errorf("failed to generate transaction auth: %v", err)
	}

	deployer := bind.DefaultDeployer(t.Auth, t.Backend)
	result, err := bind.LinkAndDeploy(&bind.DeploymentParams{
		Contracts: []*bind.MetaData{&artifacts.DataStoreMetaData},
	}, deployer)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy contract: %v", err)
	}

	if _, err := bind.WaitDeployed(t.Ctx, t.Backend, result.Txs[artifacts.DataStoreMetaData.ID].Hash()); err != nil {
		return common.Address{}, fmt.Errorf("failed to fetch receipt: %v", err)
	}

	return result.Addresses[artifacts.DataStoreMetaData.ID], nil
}

func (t *Transactor) StoreData(data [][]string) error {
	if err := t.newAuth(); err != nil {
		return err
	}

	trxChan := make(chan trxResult, 1)
	go func() {
		defer close(trxChan)
		trx, err := bind.Transact(t.Instance, t.Auth, t.DataStore.PackStoreData(data))
		trxChan <- trxResult{trx: trx, err: err}
	}()

	var trx *types.Transaction
	select {
	case <-t.Ctx.Done():
		return t.Ctx.Err()
	case result := <-trxChan:
		if result.err != nil {
			return result.err
		}
		trx = result.trx
	}

	if err := t.receiptManager(trx); err != nil {
		return err
	}

	return nil
}

func (t *Transactor) GetRange(param string) (*artifacts.EventCountOutput, error) {
	dataRange, err := bind.Call(t.Instance, nil, t.DataStore.PackEventCount(param), t.DataStore.UnpackEventCount)
	if err != nil {
		return nil, err
	}

	return &dataRange, nil
}

func (t *Transactor) newAuth() error {
	select {
	case <-t.Ctx.Done():
		return t.Ctx.Err()
	default:
	}

	ctx, cancel := context.WithTimeout(t.Ctx, 5*time.Second)
	defer cancel()

	nonce, err := t.Backend.PendingNonceAt(ctx, t.Address)
	if err != nil {
		return err
	}
	t.Auth.Nonce = big.NewInt(int64(nonce))

	if noGasNetwork {
		t.Auth.GasPrice = big.NewInt(0)
	} else {
		gasPrice, err := t.Backend.SuggestGasPrice(ctx)
		if err != nil {
			return err
		}
		t.Auth.GasPrice = gasPrice
	}

	return nil
}

func (t *Transactor) receiptManager(trx *types.Transaction) error {
	ticker := time.NewTicker(receiptInterval)
	defer ticker.Stop()

	for {
		select {
		case <-t.Ctx.Done():
			return t.Ctx.Err()
		case <-ticker.C:
			receiptCtx, cancel := context.WithTimeout(t.Ctx, 5*time.Second)
			receipt, err := t.Backend.TransactionReceipt(receiptCtx, trx.Hash())
			cancel()

			switch {
			case errors.Is(err, ethereum.NotFound):
				log.Printf("\033[32m[INF]\033[0m Receipt not found. Will retry.\n")
				continue
			case err != nil:
				log.Printf("\033[31m[ERR]\033[0m Error getting receipt: %v\n", err)
				return err
			case receipt.Status == types.ReceiptStatusSuccessful:
				log.Printf("\033[32m[INF]\033[0m Transaction committed: \033[1;35m%v\033[0m\n", receipt.TxHash)
				return nil
			default:
				log.Printf("\033[31m[ERR]\033[0m Transaction failed: \033[1;35m%v\033[0m\n", receipt.TxHash)
				return fmt.Errorf("transaction execution reverted: status=%d", receipt.Status)
			}
		}
	}
}
