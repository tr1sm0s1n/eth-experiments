package main

import (
	"_lab/data-pump/common"
	"_lab/data-pump/db"
	"_lab/data-pump/helpers"
	"_lab/data-pump/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	"gorm.io/gorm"
)

type payload struct {
	entries []models.Entry
	count   int
}

var (
	apiURL           string
	jwtToken         string
	dbConn           *gorm.DB
	maxRetries       = 3
	pingRoute        = "/ping"
	jwtRoute         = "/api/v1/jwt/generate"
	addPropertyRoute = "/api/v1/properties/add"
)

func main() {
	apiURL = os.Getenv("EXTERNAL_API")
	if len(apiURL) == 0 {
		log.Fatalln("\033[31m[ERR]\033[0m Empty EXTERNAL_API")
	}

	pingRes, err := http.Get(apiURL + pingRoute)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Error making ping request: %v", err)
	}

	pingBody, err := io.ReadAll(pingRes.Body)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Error reading ping response: %v", err)
	}

	var pingJson models.PingResponse
	if err := json.Unmarshal(pingBody, &pingJson); err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Error parsing ping JSON: %v", err)
	}

	if pingJson.Status != "alive" {
		log.Fatalln("\033[31m[ERR]\033[0m API is dead")
	}

	dbConn, err = db.Connect()
	if err != nil {
		log.Fatal("\033[31m[ERR]\033[0m Failed to connect the database")
	}

	if err = dbConn.AutoMigrate(
		&models.Entry{},
		&models.Land{},
		&models.Owner{},
		&models.Building{},
	); err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to auto migrate: %v", err)
	}

	// Create channels for processing
	payloadChan := make(chan payload)
	errorsChan := make(chan error)

	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := http.DefaultClient

	// Start worker pool
	var wg sync.WaitGroup
	for range common.Transactors {
		wg.Add(1)
		go worker(ctx, &wg, client, payloadChan, errorsChan)
	}

	// Error handling goroutine
	go func() {
		for err := range errorsChan {
			if merr := helpers.SendAlert(err); merr != nil {
				log.Printf("\033[31m[ERR]\033[0m Mail Error: %v\n", merr)
			}
			log.Fatalf("\033[31m[ERR]\033[0m Error: %v\n", err)
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

	log.Printf("\033[32m[INF]\033[0m Processed \033[1;36m%d\033[0m payload!!", processCount)
}

func worker(ctx context.Context, wg *sync.WaitGroup, client *http.Client, payload <-chan payload, errors chan<- error) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			stackTrace := string(buf[:n])

			errors <- fmt.Errorf("worker panicked, trace: %s", stackTrace)
		}
	}()

	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-payload:
			if !ok {
				return
			}

			log.Printf("\033[32m[INF]\033[0m Sending: [\033[1;32m%d\033[0m] -> [\033[1;31m%d\033[0m]\n", data.count-len(data.entries)+1, data.count)

			// Retry loop for the current data batch
			for retryCount := 0; retryCount < maxRetries; retryCount++ {
				res, err := sendRequest(client, data.entries)
				if err != nil {
					if retryCount == maxRetries-1 {
						errors <- err
						break
					}
					log.Printf("\033[33m[WRN]\033[0m Request failed (attempt %d/%d): %v. Retrying...\n", retryCount+1, maxRetries, err)
					time.Sleep(time.Duration(retryCount+1) * time.Second) // Exponential backoff
					continue
				}

				switch res.StatusCode {
				case http.StatusAccepted:
					if res := dbConn.Create(&data.entries); res.Error != nil {
						errors <- fmt.Errorf("failed to store in db: %v", err)
						break
					}
					log.Printf("\033[32m[INF]\033[0m Added: [\033[1;32m%d\033[0m] -> [\033[1;31m%d\033[0m]\n", data.count-len(data.entries)+1, data.count)
					time.Sleep(common.ReceiptInterval)
					goto nextPayload // Success, move to next payload

				case http.StatusUnauthorized:
					log.Println("\033[33m[WRN]\033[0m Expired/invalid JWT. Generating new and retrying..")

					jwtToken, err = generateJWT(client)
					if err != nil {
						errors <- err
						break
					}

					// Don't increment retryCount for JWT refresh - just retry with new token
					retryCount--                       // Compensate for the increment that will happen at loop end
					time.Sleep(500 * time.Millisecond) // Brief pause before retry
					continue

				default:
					body, err := io.ReadAll(res.Body)
					res.Body.Close()

					if err != nil {
						if retryCount == maxRetries-1 {
							errors <- fmt.Errorf("failed to read failed response: %v", err)
							break
						}
						log.Printf("\033[33m[WRN]\033[0m Failed to read response body (attempt %d/%d): %v. Retrying...\n", retryCount+1, maxRetries, err)
						time.Sleep(time.Duration(retryCount+1) * time.Second)
						continue
					}

					var failure models.FailureResponse
					if err := json.Unmarshal(body, &failure); err != nil {
						if retryCount == maxRetries-1 {
							errors <- fmt.Errorf("failed to parse failed response: %v", err)
							break
						}
						log.Printf("\033[33m[WRN]\033[0m Failed to parse response (attempt %d/%d): %v. Retrying...\n", retryCount+1, maxRetries, err)
						time.Sleep(time.Duration(retryCount+1) * time.Second)
						continue
					}

					if retryCount == maxRetries-1 {
						errors <- fmt.Errorf("failed to add property after %d attempts (HTTP %d): %s", maxRetries, res.StatusCode, failure.Error)
						break
					}

					log.Printf("\033[33m[WRN]\033[0m Request failed (attempt %d/%d, HTTP %d): %s. Retrying...\n", retryCount+1, maxRetries, res.StatusCode, failure.Error)
					time.Sleep(time.Duration(retryCount+1) * time.Second)
				}
			}

		nextPayload:
			// Continue to next payload item
		}
	}
}

func sendRequest(client *http.Client, entries []models.Entry) (*http.Response, error) {
	jsonData, err := json.Marshal(entries)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal entries to JSON: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL+addPropertyRoute, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create Property request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send Property request: %v", err)
	}

	return res, nil
}

func generateJWT(client *http.Client) (string, error) {
	req, err := http.NewRequest("POST", apiURL+jwtRoute, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create JWT request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("AUTH_TOKEN"))

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send JWT request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read JWT response: %v", err)
	}

	if resp.StatusCode == http.StatusCreated {
		var success models.TokenResponse
		if err := json.Unmarshal(body, &success); err != nil {
			return "", fmt.Errorf("failed to parse JWT success response: %v", err)
		}

		return success.Token, nil
	}

	var failure models.FailureResponse
	if err := json.Unmarshal(body, &failure); err != nil {
		return "", fmt.Errorf("failed to parse JWT failure response: %v", err)
	}
	return "", fmt.Errorf("%s", failure.Error)
}
