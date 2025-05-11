package common

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ProviderURL       = "http://127.0.0.1:8545"
	keyString         = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	intID       int64 = 31337
)

var (
	auth           *bind.TransactOpts
	privateKey     *ecdsa.PrivateKey
	accountAddress common.Address
)

func init() {
	privateKey, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v\n", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("Error casting public key to ECDSA")
	}

	accountAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	chainID := big.NewInt(intID)
	auth = bind.NewKeyedTransactor(privateKey, chainID)
}

func GenerateAuth(client *ethclient.Client) (*bind.TransactOpts, error) {
	nonce, err := client.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(927000)
	auth.GasPrice = big.NewInt(0)

	/* In case of no-gas network, comment the following code. */
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth.GasPrice = gasPrice

	return auth, nil
}
