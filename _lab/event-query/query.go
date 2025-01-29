package main

import (
	cmn "_lab/event-query/common"
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Not enough! 👾")
		return
	}

	switch os.Args[1] {
	case "1":
		fetchComplex()
	case "2":
		fetchNormal()
	default:
		log.Println("Not right! 👾")
	}
}

func fetchComplex() {
	client, err := ethclient.Dial(cmn.ProviderURL)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := cmn.NewDatastore(cmn.ContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	latest, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Latest block:", latest)

	start, end, err := getBlocksLimit(instance, "TENK")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Range:", start, end)

	events, err := fetchLogs(instance, start, end, "TENK")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Retrieved \033[1;34m%d\033[0m event logs!!", len(events))
}

func fetchNormal() {
	client, err := ethclient.Dial(cmn.ProviderURL)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := cmn.NewDatastore(cmn.ContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	latest, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Latest block:", latest)

	start, end, err := getBlocksLimit(instance, "TENK")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Range:", start, end)

	query := ethereum.FilterQuery{
		FromBlock: start,
		ToBlock:   end,
		Addresses: []common.Address{
			cmn.ContractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	eventSigHash := crypto.Keccak256Hash([]byte("Stored(string,string[])"))
	examNoHash := crypto.Keccak256Hash([]byte("TENK"))
	logEvents := []cmn.DatastoreStored{}

	for _, vLog := range logs {
		switch {
		case vLog.Topics[0].Hex() == eventSigHash.Hex() && vLog.Topics[1] == examNoHash:
			vv, err := instance.ParseStored(vLog)
			if err != nil {
				log.Fatal(err)
			}
			logEvents = append(logEvents, *vv)
		}
	}

	log.Printf("Retrieved \033[1;34m%d\033[0m event logs!!", len(logEvents))
}

func getBlocksLimit(instance *cmn.Datastore, topic string) (*big.Int, *big.Int, error) {
	r, err := instance.EventCount(nil, topic)
	if err != nil {
		return nil, nil, err
	}
	return r.Start, r.End, nil
}

func fetchLogs(instance *cmn.Datastore, start, end *big.Int, examNo string) ([]cmn.DatastoreStored, error) {
	logEvents := make([]cmn.DatastoreStored, 0)
	e := end.Uint64()
	eventIterator, err := instance.FilterStored(&bind.FilterOpts{
		Start: start.Uint64(),
		End:   &e,
	}, []string{examNo})
	if err != nil {
		return nil, err
	}
	defer eventIterator.Close()

	for {
		if !eventIterator.Next() {
			if eventIterator.Error() != nil {
				return nil, eventIterator.Error()
			}
			break
		}
		ev, err := instance.ParseStored(eventIterator.Event.Raw)
		if err != nil {
			return nil, err
		}
		logEvents = append(logEvents, *ev)
	}

	return logEvents, nil
}
