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
	receiptChan := make(chan bool)
	errorsChan := make(chan error)

	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start worker pool
	var wg sync.WaitGroup
	// for i := 0; i < common.MaxWorkers; i++ {
	wg.Add(1)
	go worker(ctx, &wg, client, instance, payloadChan, receiptChan, errorsChan)
	// }

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
			payloadChan <- currentBatch
			currentBatch = make([][]string, 0, common.BatchSize)
			<-receiptChan
		}
	}

	// Send remaining payload
	if len(currentBatch) > 0 {
		payloadChan <- currentBatch
		<-receiptChan
	}

	// Clean up
	close(payloadChan)
	wg.Wait()
	close(receiptChan)
	close(errorsChan)

	log.Printf("Processed \033[45m%d\033[0m payload!!", processCount)
}

func worker(ctx context.Context, wg *sync.WaitGroup, client *ethclient.Client, instance *common.Datastore, payload <-chan [][]string, receipt chan<- bool, errors chan<- error) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-payload:
			if !ok {
				return
			}

			auth, err := middlewares.AuthGenerator(client, common.PrivateKey)
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

			receipt <- true
		}
	}
}
