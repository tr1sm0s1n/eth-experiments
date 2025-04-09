package main

import (
	"_lab/event-query/common"
	"_lab/event-query/middlewares"
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
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

	auth, err := middlewares.AuthGenerator(client, common.Transactors[0])
	if err != nil {
		log.Fatal(err)
	}

	deployParams := bind.DeploymentParams{
		Contracts: []*bind.MetaData{&common.DatastoreMetaData},
	}

	deployer := bind.DefaultDeployer(auth, client)
	result, err := bind.LinkAndDeploy(&deployParams, deployer)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := bind.WaitDeployed(context.Background(), client, result.Txs[common.DatastoreMetaData.ID].Hash()); err != nil {
		log.Fatal(err)
	}

	log.Printf("Contract Address: \033[32m%s\033[0m\n", result.Addresses[common.DatastoreMetaData.ID].Hex())
	log.Println("Update '\033[33mContractAddress\033[0m' in 'common/config.go'.")

}

func writeContract() {
	client, err := ethclient.Dial(common.ProviderURL)
	if err != nil {
		log.Fatal(err)
	}

	ds := common.NewDatastore()
	instance := ds.Instance(client, common.ContractAddress)

	data, err := readCSV(common.CSVFile)
	if err != nil {
		log.Println("Error reading CSV:", err)
		return
	}

	chunkSize := min(len(data), 10)
	chunk := data[:chunkSize]

	auth, err := middlewares.AuthGenerator(client, common.Transactors[0])
	if err != nil {
		log.Fatal(err)
	}

	tx, err := bind.Transact(instance, auth, ds.PackStoreData(chunk))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := bind.WaitMined(context.Background(), client, tx.Hash()); err != nil {
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
