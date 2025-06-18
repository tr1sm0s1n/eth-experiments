package helpers

import (
	"_lab/data-pump/common"
	"context"
	"errors"
	"io"
	"log"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ReceiptManager(client *ethclient.Client, trx *types.Transaction) error {
	for {
		r, err := client.TransactionReceipt(context.Background(), trx.Hash())
		if err != nil {
			switch err {
			case ethereum.NotFound:
				log.Printf("\033[32m[INF]\033[0m Receipt isn't available. Will check after the interval.\n")
				time.Sleep(common.ReceiptInterval)
				continue
			case &url.Error{Err: io.EOF}:
				log.Printf("\033[32m[WRN]\033[0m Chain isn't responding. Will check after the interval.\n")
				time.Sleep(common.ReceiptInterval)
				continue
			default:
				return err
			}
		}

		if r.Status == types.ReceiptStatusSuccessful {
			log.Printf("\033[32m[INF]\033[0m Transaction has been committed. Hash: \033[1;35m%v\033[0m\n", r.TxHash)
			return nil
		}
		return errors.New("transaction execution failed")
	}
}
