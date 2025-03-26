package middlewares

import (
	"_lab/event-query/common"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AuthGenerator(client *ethclient.Client, t common.Transactor) (*bind.TransactOpts, error) {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(t.Key, chainID)
	nonce, err := client.PendingNonceAt(context.Background(), t.Address)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(9270000)
	auth.GasPrice = big.NewInt(0)

	/* In case of no-gas network, comment the following code. */
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth.GasPrice = gasPrice

	return auth, nil

}
