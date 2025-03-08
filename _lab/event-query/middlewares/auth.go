package middlewares

import (
	"_lab/event-query/common"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AuthGenerator(client *ethclient.Client, tnr common.Transactor) (*bind.TransactOpts, error) {
	nonce, err := client.PendingNonceAt(context.Background(), tnr.Address)
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(9270000)
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return nil, err
	// }

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(tnr.Key, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = big.NewInt(0)

	return auth, nil
}
