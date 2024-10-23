package main

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/holiman/uint256"
	"github.com/joho/godotenv"
	"github.com/tr1sm0s1n/eth-experiments/utils"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func blobTx() {
	log.Println(">>> Type 0x3 Transaction: BEGIN <<<")

	myBlob := new(kzg4844.Blob)
	copy(myBlob[:], "Hello, World!")

	myBlobCommit, _ := kzg4844.BlobToCommitment(myBlob)
	myBlobProof, _ := kzg4844.ComputeBlobProof(myBlob, myBlobCommit)

	client, err := utils.DialClient()
	if err != nil {
		log.Println("Failed to dial client:", err)
		return
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("Failed to retrieve the chain ID:", err)
		return
	}

	auth, key, err := utils.AuthGenerator(client)
	if err != nil {
		log.Println("Failed to generate auth:", err)
		return
	}

	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Println("Failed to return nonce:", err)
		return
	}

	to := common.Address{0x03, 0x04, 0x05}

	sidecar := &types.BlobTxSidecar{
		Blobs:       []kzg4844.Blob{*myBlob},
		Commitments: []kzg4844.Commitment{myBlobCommit},
		Proofs:      []kzg4844.Proof{myBlobProof},
	}

	signedTx, _ := types.SignNewTx(key, types.LatestSignerForChainID(chainID), &types.BlobTx{
		Nonce:      nonce,
		To:         to,
		GasTipCap:  uint256.NewInt(1000000),
		GasFeeCap:  uint256.NewInt(1000000000),
		Gas:        21000,
		BlobFeeCap: uint256.NewInt(15),
		BlobHashes: sidecar.BlobHashes(),
		Sidecar:    sidecar,
	})

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println("Failed to send trx:", err)
		return
	}

	log.Println("Trx Hash:", signedTx.Hash())

	for {
		r, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
		if err != nil {
			log.Println("Receipt not available")
			time.Sleep(5 * time.Second)
			continue
		}

		if r.Status == 1 {
			log.Println("Transaction success for", r.TxHash)
			break
		}

		log.Println("Status not committed")
		time.Sleep(5 * time.Second)
	}

	btx, _, err := client.TransactionByHash(context.Background(), signedTx.Hash())
	if err != nil {
		log.Println("Failed to fetch trx:", err)
		return
	}

	if btx.BlobHashes()[0] != sidecar.BlobHashes()[0] {
		log.Println("Failed to verify blob hashes")
		return
	}

	log.Println(">>> Type 0x3 Transaction: END <<<")
}

func main() {
	blobTx()
}
