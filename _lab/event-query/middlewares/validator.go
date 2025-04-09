package middlewares

import (
	cmn "_lab/event-query/common"
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"reflect"
	"sync"
)

func CSVValidator(lookupMap map[string]cmn.DataStoreStored, idColumn int, numWorkers int) (int64, int64, error) {
	file, err := os.Open(cmn.CSVFile)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	// Create a buffered reader for better performance
	reader := bufio.NewReaderSize(file, 1024*1024) // 1MB buffer
	csvReader := csv.NewReader(reader)

	// Create channels for workers and results
	rowChan := make(chan []string, 10000)
	resultChan := make(chan struct{ matches, mismatches int64 }, numWorkers)
	errorChan := make(chan error, 1)

	// Start worker goroutines
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for range numWorkers {
		go func() {
			defer wg.Done()
			var matches, mismatches int64

			for row := range rowChan {
				// Ensure row has enough columns
				if len(row) <= idColumn {
					mismatches++
					continue
				}

				// Extract ID from CSV row
				id := row[idColumn]

				// Look up stored item by ID
				stored, exists := lookupMap[id]

				if exists {
					// Use reflect.DeepEqual for comprehensive comparison
					if reflect.DeepEqual(stored.Data, row) {
						matches++
					} else {
						// Further verification with bytes.Equal on each element
						allMatch := true
						if len(stored.Data) == len(row) {
							for i := range stored.Data {
								if !bytes.Equal([]byte(stored.Data[i]), []byte(row[i])) {
									allMatch = false
									break
								}
							}
						} else {
							allMatch = false
						}

						if allMatch {
							matches++
						} else {
							mismatches++
						}
					}
				} else {
					mismatches++
				}
			}

			resultChan <- struct{ matches, mismatches int64 }{matches, mismatches}
		}()
	}

	// Read CSV and send to workers
	go func() {
		defer close(rowChan)
		csvReader.Read()

		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				errorChan <- err
				return
			}
			rowChan <- record
		}
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Check for errors
	select {
	case err := <-errorChan:
		return 0, 0, err
	default:
		// No errors
	}

	var totalMatches, totalMismatches int64
	for result := range resultChan {
		totalMatches += result.matches
		totalMismatches += result.mismatches
	}

	return totalMatches, totalMismatches, nil
}
