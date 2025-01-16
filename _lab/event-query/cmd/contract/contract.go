package main

import (
	cmn "_lab/event-query/common"
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"errors"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	// Make sure aetherguild/druid is running.
	providerUrl = "http://127.0.0.1:8545"
	// Priavte key of aetherguild faucet.
	privateKey = "0xb71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291"
	// Replace after deployment.
	contractAddress = common.HexToAddress("0x3A220f351252089D385b29beca14e27F204c296A")
	// Location of .csv file.
	csvFile = "./marklist.csv"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Not enough! 👾")
		return
	}

	switch os.Args[1] {
	case "d":
		deployContract()
	case "w":
		writeContract()
	default:
		log.Println("Not right! 👾")
	}
}

func deployContract() {
	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := authGenerator(client, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress, trx, _, err := cmn.DeployDatastore(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	if err := waitForReceipt(client, trx); err != nil {
		log.Fatal(err)
	}

	log.Printf("Contract Address: \033[32m%s\033[0m\n", contractAddress.String())
	log.Printf("Transaction Hash: \033[32m%s\033[0m\n", trx.Hash())
}

func writeContract() {
	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := cmn.NewDatastore(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	data, err := readCSV(csvFile)
	if err != nil {
		log.Println("Error reading CSV:", err)
		return
	}

	chunkSize := 10
	if len(data) < chunkSize {
		chunkSize = len(data)
	}
	chunk := data[:chunkSize]

	auth, err := authGenerator(client, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	trx, err := instance.StoreData(auth, chunk)
	if err != nil {
		log.Fatal(err)
	}

	if err := waitForReceipt(client, trx); err != nil {
		log.Fatal(err)
	}

	log.Printf("Transaction Hash: \033[32m%s\033[0m\n", trx.Hash())
}

func authGenerator(client *ethclient.Client, key string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(key[2:])
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(9270000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}

func waitForReceipt(client *ethclient.Client, trx *types.Transaction) error {
	for {
		r, err := client.TransactionReceipt(context.Background(), trx.Hash())
		if err != nil {
			if err == ethereum.NotFound {
				time.Sleep(time.Second)
				continue
			}
			return err
		}

		if r.Status == types.ReceiptStatusSuccessful {
			log.Println("Transaction has been committed!!")
			break
		}

		log.Println("Transaction is pending...")
		time.Sleep(time.Second)
	}
	return nil
}

func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var results [][]string

	// Skip the header row.
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}
