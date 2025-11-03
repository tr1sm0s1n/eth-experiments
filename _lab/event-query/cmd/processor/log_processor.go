package main

import (
	"_lab/event-query/artifacts"
	"_lab/event-query/config"
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
)

type blockSpan struct {
	start int64
	end   int64
}

var (
	// Slice to store 'Stored' events
	events []artifacts.DataStoreStored
	// To prevent race condition
	mu sync.Mutex
)

func main() {
	// f, _ := os.Create("cpu.prof")
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	dataRange, err := config.Transactors[0].GetRange(config.ExamTitle)
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
	for range config.MaxWorkers {
		wg.Add(1)
		go worker(ctx, &wg, config.Transactors[1], payloadChan, errorsChan)
	}

	// Error handling goroutine
	go func() {
		for err := range errorsChan {
			log.Printf("Error: %v", err)
		}
	}()

	log.Printf("Data Range: [\033[1;36m%d\033[0m] -> [\033[1;36m%d\033[0m]\n", dataRange.Start.Int64(), dataRange.End.Int64())
	span := blockSpan{start: dataRange.Start.Int64()}

	for span.start <= dataRange.End.Int64() {
		span.end = min(span.start+config.BlockRange, dataRange.End.Int64())
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

func worker(ctx context.Context, wg *sync.WaitGroup, t *config.Transactor, payload <-chan blockSpan, errors chan<- error) {
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
					config.ContractAddress,
				},
			}

			logs, err := t.Backend.FilterLogs(context.Background(), query)
			if err != nil {
				errors <- fmt.Errorf("failed to filter logs: %v", err)
			}

			for _, l := range logs {
				switch l.Topics[1] {
				case config.FilterTopic:
					parsed, err := t.DataStore.UnpackStoredEvent(&l)
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
