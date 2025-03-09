package utils

import (
	"context"
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DialClient(rawurl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func AddressGenerator(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("error casting public key to ECDSA")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

// will remove this after PR is accepted
func ClientVersion(eth *ethclient.Client, ctx context.Context) (string, error) {
	var version string
	if err := eth.Client().CallContext(ctx, &version, "web3_clientVersion"); err != nil {
		return "", err
	}
	return version, nil
}
