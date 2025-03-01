package main

import (
	cmn "_lab/event-query/common"
	"context"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func listenWithClient() {
	client, err := ethclient.Dial(cmn.WebSocketURL)
	if err != nil {
		panic(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{cmn.ContractAddress},
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

			if err := certABI.UnpackIntoInterface(&event, "Stored", vLog.Data); err != nil {
				panic(err)
			}
			log.Println("Event occured!!")
			log.Println("--------------------")
			log.Printf("No: \033[34m%s\033[0m\n", event.ExamNo)
			log.Printf("Data: \033[34m%s\033[0m\n", event.Data)
			log.Printf("Raw log: \033[32m%v\033[0m\n", event.Raw)
			log.Println("--------------------")
		}
	}
}

func listenWithContract() {
	client, err := ethclient.Dial(cmn.WebSocketURL)
	if err != nil {
		panic(err)
	}

	instance, err := cmn.NewDatastore(cmn.ContractAddress, client)
	if err != nil {
		panic(err)
	}

	eventChan := make(chan *cmn.DatastoreStored)
	sub, err := instance.WatchStored(nil, eventChan, []string{})
	if err != nil {
		panic(err)
	}

	log.Println("Listening for events...")
	log.Println("-----------------------")

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case event := <-eventChan:
			log.Println("Event occured!!")
			log.Println("--------------------")
			log.Printf("No: \033[34m%s\033[0m\n", event.ExamNo)
			log.Printf("Data: \033[34m%s\033[0m\n", event.Data)

			rw, err := event.Raw.MarshalJSON()
			if err != nil {
				panic(err)
			}
			log.Printf("Raw log: \033[32m%v\033[0m\n", string(rw))
			log.Println("--------------------")
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Println("Not enough! 👾")
		return
	}

	switch os.Args[1] {
	case "1":
		listenWithClient()
	case "2":
		listenWithContract()
	default:
		log.Println("Not right! 👾")
	}
}
