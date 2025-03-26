// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package common

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

// DatastoreMetaData contains all meta data concerning the Datastore contract.
var DatastoreMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"exam_no\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"data\",\"type\":\"string[]\"}],\"name\":\"Stored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"EventCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"first\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[][]\",\"name\":\"newStrings\",\"type\":\"string[][]\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Datastore",
	Bin: "0x608060405234801561001057600080fd5b506109d3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634f7de7f81461003b578063ea17fc7c14610057575b600080fd5b61005560048036038101906100509190610644565b610089565b005b610071600480360381019061006c919061068d565b6102d6565b6040516100809392919061070a565b60405180910390f35b60005b81518110156102d2578181815181106100a8576100a7610741565b5b60200260200101516003815181106100c3576100c2610741565b5b60200260200101516040516100d891906107e1565b60405180910390207f54036692cf01024a8aa507eab9ba1f9bdea128475be185471f57c70ab4f3a87883838151811061011457610113610741565b5b60200260200101516040516101299190610904565b60405180910390a243600083838151811061014757610146610741565b5b602002602001015160038151811061016257610161610741565b5b602002602001015160405161017791906107e1565b90815260200160405180910390206001018190555060008282815181106101a1576101a0610741565b5b60200260200101516003815181106101bc576101bb610741565b5b60200260200101516040516101d191906107e1565b908152602001604051809103902060020160009054906101000a900460ff166102bf5743600083838151811061020a57610209610741565b5b602002602001015160038151811061022557610224610741565b5b602002602001015160405161023a91906107e1565b9081526020016040518091039020600001819055506001600083838151811061026657610265610741565b5b602002602001015160038151811061028157610280610741565b5b602002602001015160405161029691906107e1565b908152602001604051809103902060020160006101000a81548160ff0219169083151502179055505b80806102ca90610955565b91505061008c565b5050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000154908060010154908060020160009054906101000a900460ff16905083565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103858261033c565b810181811067ffffffffffffffff821117156103a4576103a361034d565b5b80604052505050565b60006103b7610323565b90506103c3828261037c565b919050565b600067ffffffffffffffff8211156103e3576103e261034d565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff8211156104145761041361034d565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff8211156104455761044461034d565b5b61044e8261033c565b9050602081019050919050565b82818337600083830152505050565b600061047d6104788461042a565b6103ad565b90508281526020810184848401111561049957610498610425565b5b6104a484828561045b565b509392505050565b600082601f8301126104c1576104c0610337565b5b81356104d184826020860161046a565b91505092915050565b60006104ed6104e8846103f9565b6103ad565b905080838252602082019050602084028301858111156105105761050f6103f4565b5b835b8181101561055757803567ffffffffffffffff81111561053557610534610337565b5b80860161054289826104ac565b85526020850194505050602081019050610512565b5050509392505050565b600082601f83011261057657610575610337565b5b81356105868482602086016104da565b91505092915050565b60006105a261059d846103c8565b6103ad565b905080838252602082019050602084028301858111156105c5576105c46103f4565b5b835b8181101561060c57803567ffffffffffffffff8111156105ea576105e9610337565b5b8086016105f78982610561565b855260208501945050506020810190506105c7565b5050509392505050565b600082601f83011261062b5761062a610337565b5b813561063b84826020860161058f565b91505092915050565b60006020828403121561065a5761065961032d565b5b600082013567ffffffffffffffff81111561067857610677610332565b5b61068484828501610616565b91505092915050565b6000602082840312156106a3576106a261032d565b5b600082013567ffffffffffffffff8111156106c1576106c0610332565b5b6106cd848285016104ac565b91505092915050565b6000819050919050565b6106e9816106d6565b82525050565b60008115159050919050565b610704816106ef565b82525050565b600060608201905061071f60008301866106e0565b61072c60208301856106e0565b61073960408301846106fb565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081519050919050565b600081905092915050565b60005b838110156107a4578082015181840152602081019050610789565b60008484015250505050565b60006107bb82610770565b6107c5818561077b565b93506107d5818560208601610786565b80840191505092915050565b60006107ed82846107b0565b915081905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b600061084082610770565b61084a8185610824565b935061085a818560208601610786565b6108638161033c565b840191505092915050565b600061087a8383610835565b905092915050565b6000602082019050919050565b600061089a826107f8565b6108a48185610803565b9350836020820285016108b685610814565b8060005b858110156108f257848403895281516108d3858261086e565b94506108de83610882565b925060208a019950506001810190506108ba565b50829750879550505050505092915050565b6000602082019050818103600083015261091e818461088f565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610960826106d6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361099257610991610926565b5b60018201905091905056fea2646970667358221220cdce876501823f549b9952b758af9be27a4b20ddbaab41b1fe23f7c96379233164736f6c63430008120033",
}

// Datastore is an auto generated Go binding around an Ethereum contract.
type Datastore struct {
	abi abi.ABI
}

// NewDatastore creates a new instance of Datastore.
func NewDatastore() *Datastore {
	parsed, err := DatastoreMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Datastore{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Datastore) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackEventCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea17fc7c.
//
// Solidity: function EventCount(string ) view returns(uint256 start, uint256 end, bool first)
func (datastore *Datastore) PackEventCount(arg0 string) []byte {
	enc, err := datastore.abi.Pack("EventCount", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// EventCountOutput serves as a container for the return parameters of contract
// method EventCount.
type EventCountOutput struct {
	Start *big.Int
	End   *big.Int
	First bool
}

// UnpackEventCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xea17fc7c.
//
// Solidity: function EventCount(string ) view returns(uint256 start, uint256 end, bool first)
func (datastore *Datastore) UnpackEventCount(data []byte) (EventCountOutput, error) {
	out, err := datastore.abi.Unpack("EventCount", data)
	outstruct := new(EventCountOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Start = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.End = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.First = *abi.ConvertType(out[2], new(bool)).(*bool)
	return *outstruct, err

}

// PackStoreData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f7de7f8.
//
// Solidity: function storeData(string[][] newStrings) returns()
func (datastore *Datastore) PackStoreData(newStrings [][]string) []byte {
	enc, err := datastore.abi.Pack("storeData", newStrings)
	if err != nil {
		panic(err)
	}
	return enc
}

// DatastoreStored represents a Stored event raised by the Datastore contract.
type DatastoreStored struct {
	ExamNo common.Hash
	Data   []string
	Raw    *types.Log // Blockchain specific contextual infos
}

const DatastoreStoredEventName = "Stored"

// ContractEventName returns the user-defined event name.
func (DatastoreStored) ContractEventName() string {
	return DatastoreStoredEventName
}

// UnpackStoredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Stored(string indexed exam_no, string[] data)
func (datastore *Datastore) UnpackStoredEvent(log *types.Log) (*DatastoreStored, error) {
	event := "Stored"
	if log.Topics[0] != datastore.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DatastoreStored)
	if len(log.Data) > 0 {
		if err := datastore.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range datastore.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
