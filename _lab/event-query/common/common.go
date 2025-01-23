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
	PrivateKey = "0xb71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291"
	// Replace after deployment.
	ContractAddress = common.HexToAddress("0x3A220f351252089D385b29beca14e27F204c296A")
	// Location of .csv file.
	CSVFile = "./marklist.csv"
	// Number of workers.
	MaxWorkers = 4
	// Number of rows taken for one transaction.
	BatchSize = 1000
	// Block range for event log processing.
	BlockRange int64 = 200
	// Signature hash of the major event.
	EventSignature = crypto.Keccak256Hash([]byte("Stored(string,string[])"))
	// Topic hash for the exam filter.
	FilterTopic = crypto.Keccak256Hash([]byte("TENK"))
)
