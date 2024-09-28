package main

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
)

func generateAccount() (*ecdsa.PrivateKey, *common.Address, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	address := crypto.PubkeyToAddress(key.PublicKey)
	return key, &address, nil
}

func generateBackend(faucet *common.Address) (*simulated.Backend, *big.Int) {
	chainID := params.AllDevChainProtocolChanges.ChainID

	sim := simulated.NewBackend(map[common.Address]types.Account{
		*faucet: {Balance: new(big.Int).Exp(big.NewInt(10), big.NewInt(21), nil)}, // 1000 ETH
	})

	return sim, chainID
}
