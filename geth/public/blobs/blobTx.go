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
	log.Println("\033[1;36m>>> Type 0x3 Transaction: BEGIN <<<\033[0m")

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
		log.Fatal("Failed to generate address:", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal("Failed to return nonce:", err)
	}

	to := common.HexToAddress("0x09778b53bbDFd17438c9e111995728ca80f6c5b1")

	sidecar := types.BlobTxSidecar{}
	for range 6 {
		sidecar.Blobs = append(sidecar.Blobs, *myBlob)
		sidecar.Commitments = append(sidecar.Commitments, myBlobCommit)
		sidecar.Proofs = append(sidecar.Proofs, myBlobProof)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to get suggested gas price:", err)
	}

	gasTip, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Fatal("Failed to get suggested gas tip:", err)
	}

	blobFee, err := client.BlobBaseFee(context.Background())
	if err != nil {
		log.Fatal("Failed to get blob base fee:", err)
	}

	signedTx, _ := types.SignNewTx(privateKey, types.LatestSignerForChainID(chainID), &types.BlobTx{
		Nonce:      nonce,
		To:         to,
		GasTipCap:  uint256.MustFromBig(gasTip),
		GasFeeCap:  uint256.MustFromBig(gasPrice),
		Gas:        21000,
		BlobFeeCap: uint256.MustFromBig(blobFee),
		BlobHashes: sidecar.BlobHashes(),
		Sidecar:    &sidecar,
	})

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("Failed to send trx:", err)
	}

	log.Printf("Blob Tx: \033[1;34m%s\033[0m\n", signedTx.Hash())

	for {
		r, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
		if err != nil {
			if err == ethereum.NotFound {
				log.Println("Receipt not available. Will try after 5s..")
				time.Sleep(5 * time.Second)
				continue
			} else {
				log.Fatal("Failed to get receipt:", err)
			}
		}

		if r.Status == 1 {
			log.Println("\033[1;32mTransaction succeeded!!\033[0m")
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

	log.Println("\033[1;36m>>> Type 0x3 Transaction: END <<<\033[0m")
}

func main() {
	blobTx()
}
