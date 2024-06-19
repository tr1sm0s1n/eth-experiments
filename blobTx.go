package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/holiman/uint256"
)

var (
	emptyBlob          = new(kzg4844.Blob)
	emptyBlobCommit, _ = kzg4844.BlobToCommitment(emptyBlob)
	emptyBlobProof, _  = kzg4844.ComputeBlobProof(emptyBlob, emptyBlobCommit)
)

func blobTx() {
	fmt.Println(">>> Type 0x3 Transaction: BEGIN <<<")

	key, from, err := generateAccount()
	if err != nil {
		fmt.Println("Failed to generate key:", err)
		return
	}

	sim, chainID := generateBackend(from)

	nonce, err := sim.Client().PendingNonceAt(context.Background(), *from)
	if err != nil {
		fmt.Println("Failed to fetch nonce:", err)
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

	err = sim.Client().SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("Failed to send trx:", err)
		return
	}

	sim.Commit()

	r, err := sim.Client().TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		fmt.Println("Failed to generate receipt:", err)
		return
	}

	if r.Status != 1 {
		fmt.Println("Failed to commit trx")
		return
	}

	btx, _, err := sim.Client().TransactionByHash(context.Background(), signedTx.Hash())
	if err != nil {
		fmt.Println("Failed to fetch trx:", err)
		return
	}

	if btx.BlobHashes()[0] != sidecar.BlobHashes()[0] {
		fmt.Println("Failed to verify blob hashes")
		return
	}

	fmt.Println(">>> Type 0x3 Transaction: END <<<")
}
