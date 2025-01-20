package main

import (
	"_lab/event-query/common"
	"_lab/event-query/middlewares"
	"encoding/csv"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
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
	client, err := ethclient.Dial(common.ProviderURL)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := middlewares.AuthGenerator(client, common.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress, trx, _, err := common.DeployDatastore(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	if err := middlewares.WaitForReceipt(client, trx); err != nil {
		log.Fatal(err)
	}

	log.Printf("Contract Address: \033[32m%s\033[0m\n", contractAddress.String())
}

func writeContract() {
	client, err := ethclient.Dial(common.ProviderURL)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := common.NewDatastore(common.ContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	data, err := readCSV(common.CSVFile)
	if err != nil {
		log.Println("Error reading CSV:", err)
		return
	}

	chunkSize := 10
	if len(data) < chunkSize {
		chunkSize = len(data)
	}
	chunk := data[:chunkSize]

	auth, err := middlewares.AuthGenerator(client, common.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	trx, err := instance.StoreData(auth, chunk)
	if err != nil {
		log.Fatal(err)
	}

	if err := middlewares.WaitForReceipt(client, trx); err != nil {
		log.Fatal(err)
	}
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
