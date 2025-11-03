package main

import (
	"_lab/event-query/config"
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"
)

type payload struct {
	records [][]string
	count   int
}

func main() {
	// Open CSV file
	file, err := os.Open(config.CSVFile)
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
	payloadChan := make(chan payload)
	errorsChan := make(chan error)

	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start worker pool
	var wg sync.WaitGroup
	for _, t := range config.Transactors {
		t.Wg = &wg
		t.Ctx = ctx
		wg.Add(1)
		go worker(t, payloadChan, errorsChan)
	}

	// Error handling goroutine
	go func() {
		for err := range errorsChan {
			log.Printf("Error: %v\n", err)
		}
	}()

	processCount := 0
	currentBatch := make([][]string, 0, config.BatchSize)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading row: %v\n", err)
			continue
		}

		currentBatch = append(currentBatch, row)
		processCount++

		if len(currentBatch) >= config.BatchSize {
			// Send batch to workers
			payloadChan <- payload{records: currentBatch, count: processCount}
			currentBatch = make([][]string, 0, config.BatchSize)
		}
	}

	// Send remaining payload
	if len(currentBatch) > 0 {
		payloadChan <- payload{records: currentBatch, count: processCount}
	}

	// Clean up
	close(payloadChan)
	wg.Wait()
	close(errorsChan)

	log.Printf("Processed \033[1;36m%d\033[0m payload!!", processCount)
}

func worker(t *config.Transactor, payload <-chan payload, errors chan<- error) {
	defer t.Wg.Done()

	for {
		select {
		case <-t.Ctx.Done():
			return
		case data, ok := <-payload:
			if !ok {
				return
			}
			log.Printf("Processing: [\033[1;32m%d\033[0m] -> [\033[1;31m%d\033[0m]\n", data.count-len(data.records), data.count)

			if err := t.StoreData(data.records); err != nil {
				errors <- err
			}

			log.Printf("Completed: [\033[1;32m%d\033[0m] -> [\033[1;31m%d\033[0m]\n", data.count-len(data.records), data.count)
		}
	}
}
