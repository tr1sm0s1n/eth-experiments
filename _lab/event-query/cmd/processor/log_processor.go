package main

import (
	cmn "_lab/event-query/common"
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"
	// "time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type blockspan struct {
	start int64
	end   int64
}

func main() {
	client, err := ethclient.Dial(cmn.ProviderURL)
	if err != nil {
		log.Fatalf("Failed to connect client: %v", err)
	}

	instance, err := cmn.NewDatastore(cmn.ContractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}

	dataRange, err := instance.EventCount(nil, "TENK")
	if err != nil {
		log.Fatalf("Failed to query event count: %v", err)
	}

	// Create channels for processing
	payloadChan := make(chan blockspan)
	dataChan := make(chan *cmn.DatastoreStored)
	errorsChan := make(chan error)

	// Slice to store 'Stored' events
	events := make([]cmn.DatastoreStored, 0)

	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start worker pool
	var wg sync.WaitGroup
	for i := 0; i < cmn.MaxWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, client, instance, payloadChan, dataChan, errorsChan)
	}

	// Data handling goroutine
	go func() {
		for data := range dataChan {
			events = append(events, *data)
		}
	}()

	// Error handling goroutine
	go func() {
		for err := range errorsChan {
			log.Printf("Error: %v", err)
		}
	}()

	span := blockspan{start: dataRange.Start.Int64(), end: dataRange.Start.Int64() + 200}

	for {
		if span.end < dataRange.End.Int64() {
			payloadChan <- span
			<-dataChan
			span.start = span.end
			span.end = span.end + 200
		} else {
			span.end = dataRange.End.Int64()
			payloadChan <- span
			break
		}
	}

	// Clean up
	close(payloadChan)
	wg.Wait()
	close(dataChan)
	close(errorsChan)

	log.Printf("Retrieved \033[45m%d\033[0m event logs!!", len(events))
}

func worker(ctx context.Context, wg *sync.WaitGroup, client *ethclient.Client, instance *cmn.Datastore, payload <-chan blockspan, data chan<- *cmn.DatastoreStored, errors chan<- error) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case span, ok := <-payload:
			if !ok {
				return
			}

			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(span.start),
				ToBlock:   big.NewInt(span.end),
				Addresses: []common.Address{
					cmn.ContractAddress,
				},
			}

			logs, err := client.FilterLogs(context.Background(), query)
			if err != nil {
				errors <- fmt.Errorf("failed to filter logs: %v", err)
			}

			for _, l := range logs {
				switch {
				case l.Topics[0].Hex() == cmn.EventSignature.Hex() && l.Topics[1] == cmn.FilterTopic:
					parsed, err := instance.ParseStored(l)
					if err != nil {
						errors <- fmt.Errorf("failed to parse logs: %v", err)
					}

					data <- parsed
				}
			}
		}
	}
}
