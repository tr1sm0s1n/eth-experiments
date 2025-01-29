package common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// Make sure aetherguild/druid is running.
	ProviderURL = "http://127.0.0.1:8545"
	// WebSocket URL for listener.
	WebSocketURL = "ws://127.0.0.1:8546"
	// Private key of aetherguild faucet.
	PrivateKey = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	// Replace after deployment.
	ContractAddress = common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	// Location of .csv file.
	CSVFile = "./test.csv"
	// Number of workers.
	MaxWorkers = 10
	// Number of rows taken for one transaction.
	BatchSize = 1000
	// Block range for event log processing.
	BlockRange int64 = 200
	// Signature hash of the major event.
	EventSignature = crypto.Keccak256Hash([]byte("Stored(string,string[])"))
	// To query the block range and apply filtering.
	ExamTitle = "TEST01"
	// Topic hash for the exam filter.
	FilterTopic = crypto.Keccak256Hash([]byte(ExamTitle))
)
