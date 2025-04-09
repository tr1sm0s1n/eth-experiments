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
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type blockSpan struct {
	start int64
	end   int64
}

var (
	// Initialize the contract instance
	ds = cmn.NewDataStore()
	// Slice to store 'Stored' events
	events []cmn.DataStoreStored
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

	instance := ds.Instance(client, cmn.ContractAddress)
	dataRange, err := bind.Call(instance, nil, ds.PackEventCount(cmn.ExamTitle), ds.UnpackEventCount)
	if err != nil {
		log.Fatalf("Failed to query event count: %v", err)
	}

	// Create channels for processing
	payloadChan := make(chan blockSpan)
	errorsChan := make(chan error)

	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start worker pool
	var wg sync.WaitGroup
	for range cmn.MaxWorkers {
		wg.Add(1)
		go worker(ctx, &wg, client, payloadChan, errorsChan)
	}

	// Error handling goroutine
	go func() {
		for err := range errorsChan {
			log.Printf("Error: %v", err)
		}
	}()

	log.Printf("Data Range: [\033[1;36m%d\033[0m] -> [\033[1;36m%d\033[0m]\n", dataRange.Start.Int64(), dataRange.End.Int64())
	span := blockSpan{start: dataRange.Start.Int64()}

	for {
		if span.start > dataRange.End.Int64() {
			break
		}
		span.end = min(span.start+cmn.BlockRange, dataRange.End.Int64())
		log.Printf("Processing: [\033[1;32m%d\033[0m] -> [\033[1;31m%d\033[0m]\n", span.start, span.end)
		payloadChan <- span
		span.start = span.end + 1
	}

	// Clean up
	close(payloadChan)
	wg.Wait()
	close(errorsChan)

	log.Printf("Retrieved \033[1;34m%d\033[0m event logs!!", len(events))
}

func worker(ctx context.Context, wg *sync.WaitGroup, client *ethclient.Client, payload <-chan blockSpan, errors chan<- error) {
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
				case l.Topics[1] == cmn.FilterTopic:
					parsed, err := ds.UnpackStoredEvent(&l)
					if err != nil {
						errors <- fmt.Errorf("failed to parse logs: %v", err)
					}

					mu.Lock()
					events = append(events, *parsed)
					mu.Unlock()
				}
			}
		}
	}
}
