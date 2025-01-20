package middlewares

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WaitForReceipt(client *ethclient.Client, trx *types.Transaction) error {
	for {
		r, err := client.TransactionReceipt(context.Background(), trx.Hash())
		if err != nil {
			if err == ethereum.NotFound {
				time.Sleep(time.Second)
				continue
			}
			return err
		}

		if r.Status == types.ReceiptStatusSuccessful {
			log.Println("Transaction has been committed!!")
			break
		}

		log.Println("Transaction is pending...")
		time.Sleep(time.Second)
	}
	return nil
}
