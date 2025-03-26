package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
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

func setCode() {
	fmt.Println(">>> Type 0x4 Transaction: BEGIN <<<")

	aliceKey, alice, err := utils.GenerateAccount()
	if err != nil {
		fmt.Println("Failed to generate key:", err)
		return
	}

	bobKey, bob, err := utils.GenerateAccount()
	if err != nil {
		fmt.Println("Failed to generate key:", err)
		return
	}

	sim, chainID := utils.GenerateBackend(bob)

	bobNonce, err := sim.Client().PendingNonceAt(context.Background(), *bob)
	if err != nil {
		fmt.Println("Failed to fetch nonce:", err)
		return
	}

	gasPrice, err := sim.Client().SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("Failed to get gas price:", err)
		return
	}

	opts := bind.NewKeyedTransactor(bobKey, chainID)
	opts.Nonce = big.NewInt(int64(bobNonce))
	opts.Value = big.NewInt(0)
	opts.GasLimit = gasLimit
	opts.GasPrice = gasPrice

	var parsedABI abi.ABI
	parsedABI, err = abi.JSON(strings.NewReader(ABI))
	if err != nil {
		fmt.Println("Failed to parse ABI:", err)
		return
	}

	contractAddress, _, err := bind.DeployContract(opts, common.FromHex(bytecode), sim.Client(), nil)
	if err != nil {
		fmt.Println("Failed to deploy contract:", err)
		return
	}

	sim.Commit()

	aliceNonce, err := sim.Client().PendingNonceAt(context.Background(), *alice)
	if err != nil {
		fmt.Println("Failed to fetch nonce:", err)
		return
	}

	auth := types.SetCodeAuthorization{
		ChainID: *uint256.NewInt(chainID.Uint64()),
		Address: contractAddress,
		Nonce:   aliceNonce,
	}

	auth, err = types.SignSetCode(aliceKey, auth)
	if err != nil {
		fmt.Println("Failed to sign the authorization:", err)
		return
	}

	bobNonce, err = sim.Client().PendingNonceAt(context.Background(), *bob)
	if err != nil {
		fmt.Println("Failed to fetch nonce:", err)
		return
	}

	input, err := parsedABI.Pack("emitHello")
	if err != nil {
		fmt.Println("Failed to pack emitHello:", err)
		return
	}

	signedTx, err := types.SignNewTx(bobKey, types.LatestSignerForChainID(chainID), &types.SetCodeTx{
		Nonce:     bobNonce,
		To:        *alice,
		GasTipCap: uint256.NewInt(1000000),
		GasFeeCap: uint256.NewInt(1000000000),
		Gas:       48000,
		Data:      input,
		AuthList:  []types.SetCodeAuthorization{auth},
	})
	if err != nil {
		fmt.Println("Failed to sign SetCodeTx:", err)
		return
	}

	if err = sim.Client().SendTransaction(context.Background(), signedTx); err != nil {
		fmt.Println("Failed to send trx:", err)
		return
	}

	sim.Commit()

	r, err := sim.Client().TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		fmt.Println("Failed to generate receipt:", err)
		return
	}

	if r.Status != types.ReceiptStatusSuccessful {
		fmt.Println("Failed to commit trx")
		return
	}

	if r.Logs[0].Address != *alice {
		fmt.Println("Failed to get signer address")
		return
	}

	fmt.Println("    Authorization signer:", alice)
	fmt.Println("    Delegation designator:", contractAddress)

	fmt.Println(">>> Type 0x4 Transaction: END <<<")
}
