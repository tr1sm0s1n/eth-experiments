package config

import (
	"_lab/event-query/artifacts"
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	// Make sure aetherguild/druid is running.
	providerURL = "http://127.0.0.1:8545"
	// WebSocket URL for listener.
	webSocketURL = "ws://127.0.0.1:8546"
	// Chain ID of the network
	chainID = big.NewInt(31337)
	// Set this true if the chain doesn't concern with gas.
	noGasNetwork = false
	// Sleep time between querying for the receipt.
	receiptInterval = 10 * time.Second
	// Replace after deployment.
	ContractAddress = common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	// Location of .csv file.
	CSVFile = "./test.csv"
	// Number of workers.
	MaxWorkers = 7
	// Number of rows taken for one transaction.
	BatchSize = 200
	// Block range for event log processing.
	BlockRange int64 = 500
	// Signature hash of the major event.
	EventSignature = crypto.Keccak256Hash([]byte("Stored(string,string[])"))
	// To query the block range and apply filtering.
	ExamTitle = "TEST01"
	// Topic hash for the exam filter.
	FilterTopic = crypto.Keccak256Hash([]byte(ExamTitle))
	// Private keys with enough balance.
	privateKeys = []string{
		"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		"59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
		"5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
		"7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
		"47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
	}
	// Transactors
	Transactors []*Transactor
)

func init() {
	ec, err := ethclient.Dial(providerURL)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to dial chain: %v\n", err)
	}
	ds := artifacts.NewDataStore()
	for _, k := range privateKeys {
		privateKey, err := crypto.HexToECDSA(k)
		if err != nil {
			log.Fatalf("\033[31m[ERR]\033[0m Failed to parse private key: %v", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatalln("\033[31m[ERR]\033[0m Error casting public key to ECDSA")
		}

		Transactors = append(Transactors, &Transactor{
			Ctx:       context.Background(),
			Auth:      bind.NewKeyedTransactor(privateKey, chainID),
			Backend:   ec,
			Address:   crypto.PubkeyToAddress(*publicKeyECDSA),
			Instance:  ds.Instance(ec, ContractAddress),
			DataStore: ds,
		})
	}
}
