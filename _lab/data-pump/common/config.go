package common

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ProviderURL     string
	ContractAddress common.Address
	Transactors     []Transactor
	EntryPerTx      int
	LoopBound       int
	ReceiptInterval time.Duration
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
	t.Auth.GasLimit = uint64(9300000)
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
	id := envToInt("CHAIN_ID")
	chainID := big.NewInt(int64(id))

	EntryPerTx = envToInt("ENTRY_PER_TX")
	LoopBound = envToInt("LOOP_BOUND")
	ReceiptInterval = time.Duration(envToInt("RECEIPT_INTERVAL")) * time.Second

	ProviderURL = os.Getenv("CHAIN_URL")
	ContractAddress = common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	privateKeys := strings.SplitSeq(os.Getenv("PRIVATE_KEYS"), ",")
	for k := range privateKeys {
		privateKey, err := crypto.HexToECDSA(k)
		if err != nil {
			log.Fatalf("\033[31m[ERR]\033[0m Failed to parse private key: %v\n", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatalln("\033[31m[ERR]\033[0m Error casting public key to ECDSA")
		}

		Transactors = append(Transactors,
			Transactor{
				Address: crypto.PubkeyToAddress(*publicKeyECDSA),
				Auth:    bind.NewKeyedTransactor(privateKey, chainID),
			},
		)
	}
}

func envToInt(env string) int {
	v := os.Getenv(env)
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to parse %s: %v\n", env, err)
	}
	return i
}
