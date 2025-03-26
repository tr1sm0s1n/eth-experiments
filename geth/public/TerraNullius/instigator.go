package main

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatal("Failed to dial client:", err)
	}

	contractAddress := common.HexToAddress("0x6e38A457C722C6011B2dfa06d49240e797844d66")

	contract := NewTerraNullius()
	instance := contract.Instance(client, contractAddress)

	claim, err := bind.Call(instance, nil, contract.PackClaims(big.NewInt(404)), contract.UnpackClaims)
	if err != nil {
		log.Fatal("Failed to dial client:", err)
	}

	if claim.Message != "Probably Nothing" {
		log.Fatal("Oops!! You've got 'Probably Something'.")
	}

	log.Println("Hooray!! You've got 'Probably Nothing'.")
}
