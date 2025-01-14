package main

import (
	cmn "_lab/event-query/common"
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getBlocksLimit(instance *cmn.Datastore, topic string) (*big.Int, *big.Int, error) {
	r, err := instance.EventCount(nil, topic)
	if err != nil {
		return nil, nil, err
	}
	return r.Start, r.End, nil
}

func fetchLogs(instance *cmn.Datastore, start, end *big.Int, examNo string) (*cmn.DatastoreDataStored, error) {
	e := end.Uint64()
	eventIterator, err := instance.FilterDataStored(&bind.FilterOpts{
		Start: start.Uint64(),
		End:   &e,
	}, []string{examNo})
	if err != nil {
		return nil, err
	}

	events, err := instance.ParseDataStored(eventIterator.Event.Raw)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func main() {
	providerUrl, ok := os.LookupEnv("RPC_URL")
	if !ok {
		log.Fatal("RPC_URL is not found")
	}

	contractAddress, ok := os.LookupEnv("CONTRACT_ADDRESS")
	if !ok {
		log.Fatal("CONTRACT_ADDRESS is not found")
	}

	addressHex := common.HexToAddress(contractAddress)

	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := cmn.NewDatastore(addressHex, client)
	if err != nil {
		log.Fatal(err)
	}

	latest, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Latest block:", latest)

	start, end, err := getBlocksLimit(instance, "topic")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fetchLogs(instance, start, end, "examNo")
	if err != nil {
		log.Fatal(err)
	}
}
