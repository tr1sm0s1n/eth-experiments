package main

import (
	"_lab/event-query/config"
	"encoding/csv"
	"log"
	"os"
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
	contractAddress, err := config.Transactors[0].DeployContract()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Printf("Contract Address: \033[32m%s\033[0m\n", contractAddress)
	log.Println("Update '\033[33mContractAddress\033[0m' in 'common/config.go'.")

}

func writeContract() {
	data, err := readCSV(config.CSVFile)
	if err != nil {
		log.Println("Error reading CSV:", err)
		return
	}

	chunkSize := min(len(data), 10)
	chunk := data[:chunkSize]

	if err := config.Transactors[0].StoreData(chunk); err != nil {
		log.Fatalf("Error: %v\n", err)
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
