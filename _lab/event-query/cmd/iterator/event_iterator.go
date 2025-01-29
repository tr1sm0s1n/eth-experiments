package main

import (
	cmn "_lab/event-query/common"
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type blockSpan struct {
	start uint64
	end   uint64
}

var (
	// Slice to store 'Stored' events
	events []cmn.DatastoreStored
	// To prevent race condition
	mu sync.Mutex
)

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
	payloadChan := make(chan blockSpan)
	resultChan := make(chan bool)
	errorsChan := make(chan error)

	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start worker pool
	var wg sync.WaitGroup
	for i := 0; i < cmn.MaxWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, instance, payloadChan, resultChan, errorsChan)
	}

	// Error handling goroutine
	go func() {
		for err := range errorsChan {
			log.Printf("Error: %v", err)
		}
	}()

	span := blockSpan{start: dataRange.Start.Uint64()}
	for {
		if span.start > dataRange.End.Uint64() {
			break
		}
		span.end = min(span.start+uint64(cmn.BlockRange), dataRange.End.Uint64())
		payloadChan <- span
		<-resultChan
		span.start = span.end + 1
	}

	// Clean up
	close(payloadChan)
	wg.Wait()
	close(resultChan)
	close(errorsChan)

	log.Printf("Retrieved \033[1;34m%d\033[0m event logs!!", len(events))
}

func worker(ctx context.Context, wg *sync.WaitGroup, instance *cmn.Datastore, payload <-chan blockSpan, result chan<- bool, errors chan<- error) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case span, ok := <-payload:
			if !ok {
				return
			}

			iterator, err := instance.FilterStored(&bind.FilterOpts{
				Start: span.start,
				End:   &span.end,
			}, []string{cmn.ExamTitle})
			if err != nil {
				errors <- fmt.Errorf("failed to create iterator: %v", err)
			}
			defer iterator.Close()

			for {
				if !iterator.Next() {
					if iterator.Error() != nil {
						errors <- fmt.Errorf("failed to proceed further: %v", err)
					}
					break
				}

				parsed, err := instance.ParseStored(iterator.Event.Raw)
				if err != nil {
					errors <- fmt.Errorf("failed to parse logs: %v", err)
				}
				mu.Lock()
				events = append(events, *parsed)
				mu.Unlock()
			}

			result <- true
		}
	}
}
