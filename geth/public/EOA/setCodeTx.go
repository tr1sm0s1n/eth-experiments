package main

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	"github.com/joho/godotenv"
	"github.com/tr1sm0s1n/eth-experiments/utils"
)

/* Log.sol
// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

contract Log {
    event Hello();
    event World();

    function emitHello() public {
        emit Hello();
    }

    function emitWorld() public {
        emit World();
    }
}
*/

var (
	ABI      = "[{\"inputs\": [],\"name\":\"emitHello\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitWorld\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Hello\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"World\",\"type\":\"event\"}]"
	bytecode = "6080604052348015600e575f5ffd5b5060d680601a5f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c80637b3ab2d01460345780639ee1a44014603c575b5f5ffd5b603a6044565b005b60426072565b005b7fbcdfe0d5b27dd186282e187525415c57ea3077c34efb39148111e4d342e7ab0e60405160405180910390a1565b7f2d67bb91f17bca05af6764ab411e86f4ddf757adb89fcec59a7d21c525d4171260405160405180910390a156fea26469706673582212204390e23ddadb60a9a2425dbae44932179a87c21f9e05bdac9b7335032396fe3b64736f6c634300081b0033"
	gasLimit = uint64(927000)
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setCodeTx() {
	log.Println("\033[1;36m>>> Type 0x4 Transaction: BEGIN <<<\033[0m")

	rpcURL := os.Getenv("RPC_URL")
	log.Printf("RPC URL: \033[1;31m%s\033[0m\n", rpcURL)

	client, err := utils.DialClient(rpcURL)
	if err != nil {
		log.Fatal("Failed to dial client:", err)
	}

	version, err := utils.ClientVersion(client, context.Background())
	if err != nil {
		log.Fatal("Failed to fetch version:", err)
	}
	log.Printf("Client: \033[1;31m%s\033[0m\n", version)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Failed to retrieve the chain ID:", err)
	}

	firstKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY_1"))
	if err != nil {
		log.Fatal("Failed to parse the private key:", err)
	}

	firstAcc, err := utils.AddressGenerator(firstKey)
	if err != nil {
		log.Fatal("Failed to generate address:", err)
	}

	firstNonce, err := client.PendingNonceAt(context.Background(), firstAcc)
	if err != nil {
		log.Fatal("Failed to fetch nonce:", err)
	}

	secondKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY_2"))
	if err != nil {
		log.Fatal("Failed to parse the private key:", err)
	}

	secondAcc, err := utils.AddressGenerator(secondKey)
	if err != nil {
		log.Fatal("Failed to generate address:", err)
	}

	secondNonce, err := client.PendingNonceAt(context.Background(), secondAcc)
	if err != nil {
		log.Fatal("Failed to fetch nonce:", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to get gas price:", err)
	}

	opts := bind.NewKeyedTransactor(secondKey, chainID)
	opts.Nonce = big.NewInt(int64(secondNonce))
	opts.Value = big.NewInt(0)
	opts.GasLimit = gasLimit
	opts.GasPrice = gasPrice

	var parsedABI abi.ABI
	parsedABI, err = abi.JSON(strings.NewReader(ABI))
	if err != nil {
		log.Fatal("Failed to parse ABI:", err)
	}

	contractAddress, trx, err := bind.DeployContract(opts, common.FromHex(bytecode), client, nil)
	if err != nil {
		log.Fatal("Failed to deploy contract:", err)
		return
	}

	log.Printf("Deployment Tx: \033[1;34m%s\033[0m\n", trx.Hash())

	waitForReceipt(client, trx.Hash())

	auth := types.SetCodeAuthorization{
		ChainID: *uint256.NewInt(chainID.Uint64()),
		Address: contractAddress,
		Nonce:   firstNonce,
	}

	auth, err = types.SignSetCode(firstKey, auth)
	if err != nil {
		log.Fatal("Failed to sign the authorization:", err)
	}

	input, err := parsedABI.Pack("emitHello")
	if err != nil {
		log.Fatal("Failed to pack emitHello:", err)
	}

	gasTip, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Fatal("Failed to get gas price:", err)
	}

	signedTx, err := types.SignNewTx(secondKey, types.LatestSignerForChainID(chainID), &types.SetCodeTx{
		Nonce:     secondNonce + 1,
		To:        firstAcc,
		GasTipCap: uint256.MustFromBig(gasTip),
		GasFeeCap: uint256.MustFromBig(gasPrice),
		Gas:       48000,
		Data:      input,
		AuthList:  []types.SetCodeAuthorization{auth},
	})
	if err != nil {
		log.Fatal("Failed to sign SetCodeTx:", err)
	}

	log.Printf("EOA Code Tx: \033[1;34m%s\033[0m\n", signedTx.Hash())

	if err = client.SendTransaction(context.Background(), signedTx); err != nil {
		log.Fatal("Failed to send trx:", err)
	}

	waitForReceipt(client, signedTx.Hash())

	log.Println("\033[1;36m>>> Type 0x4 Transaction: END <<<\033[0m")
}

func waitForReceipt(client *ethclient.Client, tx common.Hash) {
	for {
		r, err := client.TransactionReceipt(context.Background(), tx)
		if err != nil {
			if err == ethereum.NotFound {
				log.Println("Receipt not available. Will try after 5s..")
				time.Sleep(5 * time.Second)
				continue
			} else {
				log.Fatal("Failed to get receipt:", err)
			}
		}

		if r.Status == types.ReceiptStatusSuccessful {
			log.Println("\033[1;32mTransaction succeeded!!\033[0m")
			return
		}

		log.Println("Status not committed. Will try after 5s..")
		time.Sleep(5 * time.Second)
	}
}

func txFetcher() {
	rpcURL := os.Getenv("RPC_URL")
	log.Printf("RPC URL: \033[1;31m%s\033[0m\n", rpcURL)

	client, err := utils.DialClient(rpcURL)
	if err != nil {
		log.Fatal("Failed to dial client:", err)
	}

	version, err := utils.ClientVersion(client, context.Background())
	if err != nil {
		log.Fatal("Failed to fetch version:", err)
	}
	log.Printf("Client: \033[1;34m%s\033[0m\n", version)

	txHash := os.Getenv("TX_HASH")
	log.Printf("Transaction Hash: \033[1;35m%s\033[0m\n", txHash)
	tx, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Fatal("Failed to fetch transaction:", err)
	}

	if isPending {
		log.Println("Transaction is still pending")
		return
	}

	auth, _ := json.Marshal(tx.SetCodeAuthorizations())
	log.Printf("Auth: \033[1;36m%s\033[0m\n", auth)
}

func main() {
	setCodeTx()
}
