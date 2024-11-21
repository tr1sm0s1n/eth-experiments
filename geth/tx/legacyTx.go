package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tr1sm0s1n/eth-experiments/utils"
)

func legacyTx() {
	fmt.Println(">>> Type 0x0 Transaction: BEGIN <<<")

	key, from, err := utils.GenerateAccount()
	if err != nil {
		fmt.Println("Failed to generate key:", err)
		return
	}

	sim, chainID := utils.GenerateBackend(from)

	nonce, err := sim.Client().PendingNonceAt(context.Background(), *from)
	if err != nil {
		fmt.Println("Failed to fetch nonce:", err)
		return
	}

	to := common.HexToAddress("0x09778b53bbDFd17438c9e111995728ca80f6c5b1")
	signedTx, _ := types.SignNewTx(key, types.LatestSignerForChainID(chainID), &types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		GasPrice: big.NewInt(1000000000),
		Gas:      21000,
		Value:    new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil), // 1 ETH
		Data:     nil,
	})

	err = sim.Client().SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("Failed to send trx:", err)
		return
	}

	sim.Commit()

	r, err := sim.Client().TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		fmt.Println("Failed to generate receipt:", err)
		return
	}

	if r.Status != 1 {
		fmt.Println("Failed to commit trx")
		return
	}

	br, _ := sim.Client().BalanceAt(context.Background(), to, nil)
	bs, _ := sim.Client().BalanceAt(context.Background(), *from, nil)
	fmt.Println("    Balance of receiver:", br)
	fmt.Println("    Balance of sender:", bs)

	fmt.Println(">>> Type 0x0 Transaction: END <<<")
}
