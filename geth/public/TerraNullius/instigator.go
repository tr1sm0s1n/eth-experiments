package main

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatal("Failed to dial client:", err)
	}

	contractAddress := common.HexToAddress("0x6e38A457C722C6011B2dfa06d49240e797844d66")

	instance, err := NewTerraNullius(contractAddress, client)
	if err != nil {
		log.Fatal("Failed to dial client:", err)
	}

	claim, err := instance.Claims(nil, big.NewInt(404))
	if err != nil {
		log.Fatal("Failed to dial client:", err)
	}

	if claim.Message != "Probably Nothing" {
		log.Fatal("Oops!! You've got 'Probably Something'.")
	}

	log.Println("Hooray!! You've got 'Probably Nothing'.")
}