package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_CONTAINER"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	var db *gorm.DB
	var err error
	start := time.Now()

	for {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}

		if time.Since(start) > 10*time.Second {
			break
		}

		log.Println("Waiting for DB to be ready...")
		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to DB after 10 seconds: %w", err)
}
