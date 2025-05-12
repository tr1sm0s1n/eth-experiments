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
	intID       int64 = 31337
)

var (
	privateKeys = []string{
		"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		"59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
		"5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
		"7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
		"47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
	}

	ContractAddress = common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	Transactors     []Transactor
)

type Transactor struct {
	Address common.Address
	Auth    *bind.TransactOpts
}

func (t *Transactor) NewAuth(client *ethclient.Client) (*bind.TransactOpts, error) {
	nonce, err := client.PendingNonceAt(context.Background(), t.Address)
	if err != nil {
		return nil, err
	}

	t.Auth.Nonce = big.NewInt(int64(nonce))
	t.Auth.Value = big.NewInt(0)
	t.Auth.GasLimit = uint64(927000)
	t.Auth.GasPrice = big.NewInt(0)

	/* In case of no-gas network, comment the following code. */
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	t.Auth.GasPrice = gasPrice

	return t.Auth, nil
}

func init() {
	chainID := big.NewInt(intID)

	for _, k := range privateKeys {
		privateKey, err := crypto.HexToECDSA(k)
		if err != nil {
			log.Fatalf("Failed to parse private key: %v\n", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatalln("Error casting public key to ECDSA")
		}

		Transactors = append(Transactors,
			Transactor{
				Address: crypto.PubkeyToAddress(*publicKeyECDSA),
				Auth:    bind.NewKeyedTransactor(privateKey, chainID),
			},
		)
	}
}
