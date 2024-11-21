package utils

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DialClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func AuthGenerator(client *ethclient.Client) (*bind.TransactOpts, *ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, errors.New("error casting")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, nil, err
	}

	gasLimit := uint64(927000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, privateKey, nil
}
