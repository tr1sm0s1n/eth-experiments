package main

import (
	"_lab/data-pump/common"
	"_lab/data-pump/db"
	"_lab/data-pump/helpers"
	"_lab/data-pump/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type payload struct {
	entries []models.Entry
	count   int
}

var (
	registry = common.NewRegistry()
	dbConn   *gorm.DB
)

func main() {
	client, err := ethclient.Dial(common.ProviderURL)
	if err != nil {
		log.Fatalf("Failed to connect client: %v", err)
	}

	instance := registry.Instance(client, common.ContractAddress)

	dbConn, err = db.Connect()
	if err != nil {
		log.Fatal("Failed to connect the database")
	}

	dbConn.AutoMigrate(&models.Entry{})
	dbConn.AutoMigrate(&models.Owner{})
	dbConn.AutoMigrate(&models.Property{})

	// Create channels for processing
	payloadChan := make(chan payload)
	errorsChan := make(chan error)

	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start worker pool
	var wg sync.WaitGroup
	for _, t := range common.Transactors {
		wg.Add(1)
		go worker(ctx, &wg, client, instance, &t, payloadChan, errorsChan)
	}

	// Error handling goroutine
	go func() {
		for err := range errorsChan {
			log.Fatalf("Error: %v", err)
		}
	}()

	processCount := 0
	entryBatch := make([]models.Entry, 0)

	for range common.LoopBound {
		for range common.EntryPerTx {
			entryBatch = append(entryBatch, models.RandomEntry())
			processCount++
		}
		payloadChan <- payload{entries: entryBatch, count: processCount}
		entryBatch = make([]models.Entry, 0)
	}

	close(payloadChan)
	wg.Wait()
	close(errorsChan)

	log.Printf("Processed \033[1;36m%d\033[0m payload!!", processCount)
}

func worker(ctx context.Context, wg *sync.WaitGroup, client *ethclient.Client, instance *bind.BoundContract, tnr *common.Transactor, payload <-chan payload, errors chan<- error) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-payload:
			if !ok {
				return
			}
			log.Printf("Processing: [\033[1;32m%d\033[0m] -> [\033[1;31m%d\033[0m]\n", data.count-len(data.entries)+1, data.count)

			auth, err := tnr.NewAuth(client)
			if err != nil {
				errors <- fmt.Errorf("failed to generate auth: %v", err)
			}

			ids := make([]string, len(data.entries))
			entries := make([]string, len(data.entries))
			for _, e := range data.entries {
				ids = append(ids, e.ID)
				eb, err := json.Marshal(e)
				if err != nil {
					errors <- fmt.Errorf("failed to marshal entries: %v", err)
				}
				entries = append(entries, string(eb))
			}

			trx, err := bind.Transact(instance, auth, registry.PackAddProperty(ids, entries))
			if err != nil {
				errors <- fmt.Errorf("failed to add property: %v", err)
			}

			if err := helpers.ReceiptManager(client, trx); err != nil {
				errors <- fmt.Errorf("failed to fetch receipt: %v", err)
			}

			if res := dbConn.Create(&data.entries); res.Error != nil {
				errors <- fmt.Errorf("failed to store in db: %v", err)
			}

			log.Printf("Completed: [\033[1;32m%d\033[0m] -> [\033[1;31m%d\033[0m]\n", data.count-len(data.entries)+1, data.count)
		}
	}
}
