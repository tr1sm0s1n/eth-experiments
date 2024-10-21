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

var (
	emptyBlob          = new(kzg4844.Blob)
	emptyBlobCommit, _ = kzg4844.BlobToCommitment(emptyBlob)
	emptyBlobProof, _  = kzg4844.ComputeBlobProof(emptyBlob, emptyBlobCommit)
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func blobTx() {
	log.Println(">>> Type 0x3 Transaction: BEGIN <<<")

	sim, err := utils.DialClient()
	if err != nil {
		log.Println("Failed to dial client:", err)
		return
	}

	chainID, err := sim.ChainID(context.Background())
	if err != nil {
		log.Println("Failed to retrieve the chain ID:", err)
		return
	}

	auth, key, err := utils.AuthGenerator(sim)
	if err != nil {
		log.Println("Failed to generate auth:", err)
		return
	}

	nonce, err := sim.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Println("Failed to return nonce:", err)
		return
	}

	to := common.Address{0x03, 0x04, 0x05}

	sidecar := &types.BlobTxSidecar{
		Blobs:       []kzg4844.Blob{*emptyBlob},
		Commitments: []kzg4844.Commitment{emptyBlobCommit},
		Proofs:      []kzg4844.Proof{emptyBlobProof},
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

	err = sim.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println("Failed to send trx:", err)
		return
	}

	log.Println("Trx Hash:", signedTx.Hash())

	for {
		r, err := sim.TransactionReceipt(context.Background(), signedTx.Hash())
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

	btx, _, err := sim.TransactionByHash(context.Background(), signedTx.Hash())
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
