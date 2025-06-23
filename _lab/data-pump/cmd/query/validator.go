package main

import (
	"_lab/data-pump/common"
	"_lab/data-pump/db"
	"_lab/data-pump/models"
	"encoding/json"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	registry = common.NewRegistry()
)

func main() {
	client, err := ethclient.Dial(common.ProviderURL)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to connect client: %v", err)
	}

	instance := registry.Instance(client, common.ContractAddress)

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal("\033[31m[ERR]\033[0m Failed to connect the database")
	}

	var dbEntry models.Entry
	if err := dbConn.Order("RANDOM()").Preload("Ownership").Preload("Properties").First(&dbEntry).Error; err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to fetch an entry: %v", err)
	}

	log.Printf("DB Data: \033[1;36m%v\033[0m\n", dbEntry)

	data, err := bind.Call(instance, nil, registry.PackGetLatestProperty(dbEntry.ID), registry.UnpackGetLatestProperty)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to unpack latest property: %v", err)
	}

	var bcEntry models.Entry
	if err := json.Unmarshal([]byte(data), &bcEntry); err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to unmarshal data: %v", err)
	}

	log.Printf("BC Data: \033[1;36m%v\033[0m\n", bcEntry)
}
