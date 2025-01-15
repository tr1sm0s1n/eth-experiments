package main

import (
	cmn "_lab/event-query/common"
	"context"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x3A220f351252089D385b29beca14e27F204c296A")

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	certABI, err := abi.JSON(strings.NewReader(cmn.DatastoreMetaData.ABI))
	if err != nil {
		panic(err)
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		panic(err)
	}

	log.Println("Listening for events...")
	log.Println("-----------------------")

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case vLog := <-logs:
			var event cmn.DatastoreStored

			certABI.UnpackIntoInterface(&event, "Stored", vLog.Data)
			log.Println("Event occured!!")
			log.Println("--------------------")
			log.Printf("No: \033[34m%s\033[0m\n", event.ExamNo)
			log.Printf("Data: \033[34m%s\033[0m\n", event.Data)
			log.Printf("Raw log: \033[32m%v\033[0m\n", event.Raw)
			log.Println("--------------------")
		}
	}
}
