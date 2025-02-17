package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
		log.Fatal("Failed to dial client:", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Failed to retrieve the chain ID:", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal("Failed to parse the private key:", err)
	}

	from, err := utils.AddressGenerator(privateKey)
	if err != nil {
		log.Fatal("Failed to generate adress:", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal("Failed to return nonce:", err)
	}

	to := common.Address{0x03, 0x04, 0x05}

	sidecar := &types.BlobTxSidecar{
		Blobs:       []kzg4844.Blob{*myBlob, *myBlob, *myBlob, *myBlob, *myBlob, *myBlob},
		Commitments: []kzg4844.Commitment{myBlobCommit, myBlobCommit, myBlobCommit, myBlobCommit, myBlobCommit, myBlobCommit},
		Proofs:      []kzg4844.Proof{myBlobProof, myBlobProof, myBlobProof, myBlobProof, myBlobProof, myBlobProof},
	}

	signedTx, _ := types.SignNewTx(privateKey, types.LatestSignerForChainID(chainID), &types.BlobTx{
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
		log.Fatal("Failed to send trx:", err)
	}

	log.Println("Trx Hash:", signedTx.Hash())

	for {
		r, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
		if err != nil {
			if err == ethereum.NotFound {
				log.Println("Receipt not available")
				time.Sleep(5 * time.Second)
				continue
			} else {
				log.Fatal("Failed to get receipt:", err)
			}
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
		log.Fatal("Failed to fetch trx:", err)
	}

	if btx.BlobHashes()[0] != sidecar.BlobHashes()[0] {
		log.Fatal("Failed to verify blob hashes")
	}

	log.Println(">>> Type 0x3 Transaction: END <<<")
}

func main() {
	blobTx()
}
