package helpers

import (
	"errors"
	"fmt"
	"io"
	"time"
)

type RetryConfig struct {
	maxRetries    int
	initialDelay  time.Duration
	backoffFactor float64
}

func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		maxRetries:    3,
		initialDelay:  1 * time.Second,
		backoffFactor: 2.0,
	}
}

func RetryOnEOF[T any](f func() (T, error), config RetryConfig) (T, error) {
	var result T
	var err error

	delay := config.initialDelay

	for attempt := 0; attempt <= config.maxRetries; attempt++ {
		result, err = f()

		if err == nil {
			return result, err
		}

		if !errors.Is(err, io.EOF) {
			return result, err
		}

		if attempt == config.maxRetries {
			return result, fmt.Errorf("failed to retry after %d attempts", config.maxRetries+1)
		}

		time.Sleep(delay)
		delay = time.Duration(float64(delay) * config.backoffFactor)
	}

	return result, err
}

func RetryOnEOFVoid(fn func() error, config RetryConfig) error {
	_, err := RetryOnEOF(func() (struct{}, error) {
		return struct{}{}, fn()
	}, config)
	return err
}
