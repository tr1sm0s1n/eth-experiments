package main

import (
	cmn "_lab/event-query/common"
	"context"
	"fmt"
	"log"
	"math/big"
	_ "net/http/pprof"
	_ "os"
	_ "runtime/pprof"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type blockSpan struct {
	start int64
	end   int64
}

var (
	// Slice to store 'Stored' events
	events []cmn.DatastoreStored
	// To prevent race condition
	mu sync.Mutex
)

func main() {
	// f, _ := os.Create("cpu.prof")
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

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
		go worker(ctx, &wg, client, instance, payloadChan, resultChan, errorsChan)
	}

	// Error handling goroutine
	go func() {
		for err := range errorsChan {
			log.Printf("Error: %v", err)
		}
	}()

	span := blockSpan{start: dataRange.Start.Int64(), end: dataRange.Start.Int64() + cmn.BlockRange}

	for {
		if span.end >= dataRange.End.Int64() {
			break
		}
		payloadChan <- span
		<-resultChan
		span.end++
		span.start = span.end
		span.end = span.end + cmn.BlockRange
	}

	span.end = dataRange.End.Int64()
	payloadChan <- span
	<-resultChan

	// Clean up
	close(payloadChan)
	wg.Wait()
	close(resultChan)
	close(errorsChan)

	log.Printf("Retrieved \033[45m%d\033[0m event logs!!", len(events))
}

func worker(ctx context.Context, wg *sync.WaitGroup, client *ethclient.Client, instance *cmn.Datastore, payload <-chan blockSpan, result chan<- bool, errors chan<- error) {
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

					mu.Lock()
					events = append(events, *parsed)
					mu.Unlock()
				}
			}
			result <- true
		}
	}
}
