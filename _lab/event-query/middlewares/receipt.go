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
				log.Println("Receipt isn't available")
				time.Sleep(10 * time.Second)
				continue
			}
			return err
		}

		if r.Status == types.ReceiptStatusSuccessful {
			log.Println("Transaction has been committed!!")
			log.Printf("Transaction Hash: \033[35m%s\033[0m\n", r.TxHash)
			break
		}

		log.Println("Transaction is pending...")
		time.Sleep(10 * time.Second)
	}
	return nil
}
