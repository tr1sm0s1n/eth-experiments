package main

import (
	"context"
	"fmt"

	"github.com/tr1sm0s1n/eth-experiments/utils"
)

func setCodeTx() {
	fmt.Println(">>> Type 0x4 Transaction: BEGIN <<<")

	_, from, err := utils.GenerateAccount()
	if err != nil {
		fmt.Println("Failed to generate key:", err)
		return
	}

	sim, _ := utils.GenerateBackend(from)

	_, err = sim.Client().PendingNonceAt(context.Background(), *from)
	if err != nil {
		fmt.Println("Failed to fetch nonce:", err)
		return
	}

	fmt.Println(">>> Type 0x2 Transaction: END <<<")
}
