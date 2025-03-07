package main

import (
	"_lab/event-query/common"
	"_lab/event-query/middlewares"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(common.ProviderURL)
	if err != nil {
		log.Fatalf("Failed to connect client: %v", err)
	}

	instance, err := common.NewDatastore(common.ContractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}

	// Open CSV file
	file, err := os.Open(common.CSVFile)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header
	if _, err := reader.Read(); err != nil {
		log.Fatalf("Failed to read CSV header: %v", err)
	}

	// Create channels for processing
	payloadChan := make(chan [][]string, common.BatchSize)
	errorsChan := make(chan error)

	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start worker pool
	var wg sync.WaitGroup
	for _, t := range common.Transactors {
		wg.Add(1)
		go worker(ctx, &wg, client, instance, t, payloadChan, errorsChan)
	}

	// Error handling goroutine
	go func() {
		for err := range errorsChan {
			log.Printf("Error: %v", err)
		}
	}()

	processCount := 0
	currentBatch := make([][]string, 0, common.BatchSize)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading row: %v", err)
			continue
		}

		currentBatch = append(currentBatch, row)
		processCount++

		if len(currentBatch) >= common.BatchSize {
			// Send batch to workers
			log.Printf("Processing: [\033[1;32m%d\033[0m] -> [\033[1;31m%d\033[0m]\n", processCount-common.BatchSize, processCount)
			payloadChan <- currentBatch
			currentBatch = make([][]string, 0, common.BatchSize)
		}
	}

	// Send remaining payload
	if len(currentBatch) > 0 {
		payloadChan <- currentBatch
	}

	// Clean up
	close(payloadChan)
	wg.Wait()
	close(errorsChan)

	log.Printf("Processed \033[45m%d\033[0m payload!!", processCount)
}

func worker(ctx context.Context, wg *sync.WaitGroup, client *ethclient.Client, instance *common.Datastore, tnr common.Transactor, payload <-chan [][]string, errors chan<- error) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-payload:
			if !ok {
				return
			}

			auth, err := middlewares.AuthGenerator(client, tnr)
			if err != nil {
				errors <- fmt.Errorf("failed to generate auth: %v", err)
			}

			trx, err := instance.StoreData(auth, data)
			if err != nil {
				errors <- fmt.Errorf("failed to store data: %v", err)
			}

			if err := middlewares.WaitForReceipt(client, trx); err != nil {
				errors <- fmt.Errorf("failed to fetch receipt: %v", err)
			}
		}
	}
}
