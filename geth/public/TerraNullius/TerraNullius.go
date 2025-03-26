// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// TerraNulliusMetaData contains all meta data concerning the TerraNullius contract.
var TerraNulliusMetaData = bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"number_of_claims\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\"}],\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claims\",\"outputs\":[{\"name\":\"claimant\",\"type\":\"address\"},{\"name\":\"message\",\"type\":\"string\"},{\"name\":\"block_number\",\"type\":\"uint256\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"message\",\"type\":\"string\"}],\"name\":\"claim\",\"outputs\":[],\"type\":\"function\"}]",
	ID:  "TerraNullius",
}

// TerraNullius is an auto generated Go binding around an Ethereum contract.
type TerraNullius struct {
	abi abi.ABI
}

// NewTerraNullius creates a new instance of TerraNullius.
func NewTerraNullius() *TerraNullius {
	parsed, err := TerraNulliusMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TerraNullius{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TerraNullius) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackClaim is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3fe12c9.
//
// Solidity: function claim(string message) returns()
func (terraNullius *TerraNullius) PackClaim(message string) []byte {
	enc, err := terraNullius.abi.Pack("claim", message)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackClaims is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa888c2cd.
//
// Solidity: function claims(uint256 ) returns(address claimant, string message, uint256 block_number)
func (terraNullius *TerraNullius) PackClaims(arg0 *big.Int) []byte {
	enc, err := terraNullius.abi.Pack("claims", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// ClaimsOutput serves as a container for the return parameters of contract
// method Claims.
type ClaimsOutput struct {
	Claimant    common.Address
	Message     string
	BlockNumber *big.Int
}

// UnpackClaims is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa888c2cd.
//
// Solidity: function claims(uint256 ) returns(address claimant, string message, uint256 block_number)
func (terraNullius *TerraNullius) UnpackClaims(data []byte) (ClaimsOutput, error) {
	out, err := terraNullius.abi.Unpack("claims", data)
	outstruct := new(ClaimsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Claimant = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Message = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.BlockNumber = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackNumberOfClaims is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6bd5084a.
//
// Solidity: function number_of_claims() returns(uint256 result)
func (terraNullius *TerraNullius) PackNumberOfClaims() []byte {
	enc, err := terraNullius.abi.Pack("number_of_claims")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNumberOfClaims is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6bd5084a.
//
// Solidity: function number_of_claims() returns(uint256 result)
func (terraNullius *TerraNullius) UnpackNumberOfClaims(data []byte) (*big.Int, error) {
	out, err := terraNullius.abi.Unpack("number_of_claims", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}
