package main

import (
	"_lab/data-pump/common"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(common.ProviderURL)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to connect client: %v", err)
	}

	tr := common.Transactors[0]
	auth, err := tr.NewAuth(client)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to create auth: %v", err)
	}

	deployParams := bind.DeploymentParams{
		Contracts: []*bind.MetaData{&common.RegistryMetaData},
	}

	deployer := bind.DefaultDeployer(auth, client)
	result, err := bind.LinkAndDeploy(&deployParams, deployer)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to deploy contract: %v", err)
	}

	if _, err := bind.WaitDeployed(context.Background(), client, result.Txs[common.RegistryMetaData.ID].Hash()); err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to fetch deployment receipt: %v", err)
	}

	log.Printf("\033[32m[INF]\033[0m Contract Address: \033[32m%s\033[0m\n", result.Addresses[common.RegistryMetaData.ID].Hex())
}
