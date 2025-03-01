// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package common

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DatastoreMetaData contains all meta data concerning the Datastore contract.
var DatastoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"exam_no\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"data\",\"type\":\"string[]\"}],\"name\":\"Stored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"EventCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"first\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[][]\",\"name\":\"newStrings\",\"type\":\"string[][]\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506109d3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634f7de7f81461003b578063ea17fc7c14610057575b600080fd5b61005560048036038101906100509190610644565b610089565b005b610071600480360381019061006c919061068d565b6102d6565b6040516100809392919061070a565b60405180910390f35b60005b81518110156102d2578181815181106100a8576100a7610741565b5b60200260200101516003815181106100c3576100c2610741565b5b60200260200101516040516100d891906107e1565b60405180910390207f54036692cf01024a8aa507eab9ba1f9bdea128475be185471f57c70ab4f3a87883838151811061011457610113610741565b5b60200260200101516040516101299190610904565b60405180910390a243600083838151811061014757610146610741565b5b602002602001015160038151811061016257610161610741565b5b602002602001015160405161017791906107e1565b90815260200160405180910390206001018190555060008282815181106101a1576101a0610741565b5b60200260200101516003815181106101bc576101bb610741565b5b60200260200101516040516101d191906107e1565b908152602001604051809103902060020160009054906101000a900460ff166102bf5743600083838151811061020a57610209610741565b5b602002602001015160038151811061022557610224610741565b5b602002602001015160405161023a91906107e1565b9081526020016040518091039020600001819055506001600083838151811061026657610265610741565b5b602002602001015160038151811061028157610280610741565b5b602002602001015160405161029691906107e1565b908152602001604051809103902060020160006101000a81548160ff0219169083151502179055505b80806102ca90610955565b91505061008c565b5050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000154908060010154908060020160009054906101000a900460ff16905083565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103858261033c565b810181811067ffffffffffffffff821117156103a4576103a361034d565b5b80604052505050565b60006103b7610323565b90506103c3828261037c565b919050565b600067ffffffffffffffff8211156103e3576103e261034d565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff8211156104145761041361034d565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff8211156104455761044461034d565b5b61044e8261033c565b9050602081019050919050565b82818337600083830152505050565b600061047d6104788461042a565b6103ad565b90508281526020810184848401111561049957610498610425565b5b6104a484828561045b565b509392505050565b600082601f8301126104c1576104c0610337565b5b81356104d184826020860161046a565b91505092915050565b60006104ed6104e8846103f9565b6103ad565b905080838252602082019050602084028301858111156105105761050f6103f4565b5b835b8181101561055757803567ffffffffffffffff81111561053557610534610337565b5b80860161054289826104ac565b85526020850194505050602081019050610512565b5050509392505050565b600082601f83011261057657610575610337565b5b81356105868482602086016104da565b91505092915050565b60006105a261059d846103c8565b6103ad565b905080838252602082019050602084028301858111156105c5576105c46103f4565b5b835b8181101561060c57803567ffffffffffffffff8111156105ea576105e9610337565b5b8086016105f78982610561565b855260208501945050506020810190506105c7565b5050509392505050565b600082601f83011261062b5761062a610337565b5b813561063b84826020860161058f565b91505092915050565b60006020828403121561065a5761065961032d565b5b600082013567ffffffffffffffff81111561067857610677610332565b5b61068484828501610616565b91505092915050565b6000602082840312156106a3576106a261032d565b5b600082013567ffffffffffffffff8111156106c1576106c0610332565b5b6106cd848285016104ac565b91505092915050565b6000819050919050565b6106e9816106d6565b82525050565b60008115159050919050565b610704816106ef565b82525050565b600060608201905061071f60008301866106e0565b61072c60208301856106e0565b61073960408301846106fb565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081519050919050565b600081905092915050565b60005b838110156107a4578082015181840152602081019050610789565b60008484015250505050565b60006107bb82610770565b6107c5818561077b565b93506107d5818560208601610786565b80840191505092915050565b60006107ed82846107b0565b915081905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b600061084082610770565b61084a8185610824565b935061085a818560208601610786565b6108638161033c565b840191505092915050565b600061087a8383610835565b905092915050565b6000602082019050919050565b600061089a826107f8565b6108a48185610803565b9350836020820285016108b685610814565b8060005b858110156108f257848403895281516108d3858261086e565b94506108de83610882565b925060208a019950506001810190506108ba565b50829750879550505050505092915050565b6000602082019050818103600083015261091e818461088f565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610960826106d6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361099257610991610926565b5b60018201905091905056fea2646970667358221220cdce876501823f549b9952b758af9be27a4b20ddbaab41b1fe23f7c96379233164736f6c63430008120033",
}

// DatastoreABI is the input ABI used to generate the binding from.
// Deprecated: Use DatastoreMetaData.ABI instead.
var DatastoreABI = DatastoreMetaData.ABI

// DatastoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DatastoreMetaData.Bin instead.
var DatastoreBin = DatastoreMetaData.Bin

// DeployDatastore deploys a new Ethereum contract, binding an instance of Datastore to it.
func DeployDatastore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Datastore, error) {
	parsed, err := DatastoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DatastoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Datastore{DatastoreCaller: DatastoreCaller{contract: contract}, DatastoreTransactor: DatastoreTransactor{contract: contract}, DatastoreFilterer: DatastoreFilterer{contract: contract}}, nil
}

// Datastore is an auto generated Go binding around an Ethereum contract.
type Datastore struct {
	DatastoreCaller     // Read-only binding to the contract
	DatastoreTransactor // Write-only binding to the contract
	DatastoreFilterer   // Log filterer for contract events
}

// DatastoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type DatastoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DatastoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DatastoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DatastoreSession struct {
	Contract     *Datastore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatastoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DatastoreCallerSession struct {
	Contract *DatastoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DatastoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DatastoreTransactorSession struct {
	Contract     *DatastoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DatastoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type DatastoreRaw struct {
	Contract *Datastore // Generic contract binding to access the raw methods on
}

// DatastoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DatastoreCallerRaw struct {
	Contract *DatastoreCaller // Generic read-only contract binding to access the raw methods on
}

// DatastoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DatastoreTransactorRaw struct {
	Contract *DatastoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDatastore creates a new instance of Datastore, bound to a specific deployed contract.
func NewDatastore(address common.Address, backend bind.ContractBackend) (*Datastore, error) {
	contract, err := bindDatastore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Datastore{DatastoreCaller: DatastoreCaller{contract: contract}, DatastoreTransactor: DatastoreTransactor{contract: contract}, DatastoreFilterer: DatastoreFilterer{contract: contract}}, nil
}

// NewDatastoreCaller creates a new read-only instance of Datastore, bound to a specific deployed contract.
func NewDatastoreCaller(address common.Address, caller bind.ContractCaller) (*DatastoreCaller, error) {
	contract, err := bindDatastore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DatastoreCaller{contract: contract}, nil
}

// NewDatastoreTransactor creates a new write-only instance of Datastore, bound to a specific deployed contract.
func NewDatastoreTransactor(address common.Address, transactor bind.ContractTransactor) (*DatastoreTransactor, error) {
	contract, err := bindDatastore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DatastoreTransactor{contract: contract}, nil
}

// NewDatastoreFilterer creates a new log filterer instance of Datastore, bound to a specific deployed contract.
func NewDatastoreFilterer(address common.Address, filterer bind.ContractFilterer) (*DatastoreFilterer, error) {
	contract, err := bindDatastore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DatastoreFilterer{contract: contract}, nil
}

// bindDatastore binds a generic wrapper to an already deployed contract.
func bindDatastore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DatastoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datastore *DatastoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Datastore.Contract.DatastoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datastore *DatastoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datastore.Contract.DatastoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datastore *DatastoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datastore.Contract.DatastoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datastore *DatastoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Datastore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datastore *DatastoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datastore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datastore *DatastoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datastore.Contract.contract.Transact(opts, method, params...)
}

// EventCount is a free data retrieval call binding the contract method 0xea17fc7c.
//
// Solidity: function EventCount(string ) view returns(uint256 start, uint256 end, bool first)
func (_Datastore *DatastoreCaller) EventCount(opts *bind.CallOpts, arg0 string) (struct {
	Start *big.Int
	End   *big.Int
	First bool
}, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "EventCount", arg0)

	outstruct := new(struct {
		Start *big.Int
		End   *big.Int
		First bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Start = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.First = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// EventCount is a free data retrieval call binding the contract method 0xea17fc7c.
//
// Solidity: function EventCount(string ) view returns(uint256 start, uint256 end, bool first)
func (_Datastore *DatastoreSession) EventCount(arg0 string) (struct {
	Start *big.Int
	End   *big.Int
	First bool
}, error) {
	return _Datastore.Contract.EventCount(&_Datastore.CallOpts, arg0)
}

// EventCount is a free data retrieval call binding the contract method 0xea17fc7c.
//
// Solidity: function EventCount(string ) view returns(uint256 start, uint256 end, bool first)
func (_Datastore *DatastoreCallerSession) EventCount(arg0 string) (struct {
	Start *big.Int
	End   *big.Int
	First bool
}, error) {
	return _Datastore.Contract.EventCount(&_Datastore.CallOpts, arg0)
}

// StoreData is a paid mutator transaction binding the contract method 0x4f7de7f8.
//
// Solidity: function storeData(string[][] newStrings) returns()
func (_Datastore *DatastoreTransactor) StoreData(opts *bind.TransactOpts, newStrings [][]string) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "storeData", newStrings)
}

// StoreData is a paid mutator transaction binding the contract method 0x4f7de7f8.
//
// Solidity: function storeData(string[][] newStrings) returns()
func (_Datastore *DatastoreSession) StoreData(newStrings [][]string) (*types.Transaction, error) {
	return _Datastore.Contract.StoreData(&_Datastore.TransactOpts, newStrings)
}

// StoreData is a paid mutator transaction binding the contract method 0x4f7de7f8.
//
// Solidity: function storeData(string[][] newStrings) returns()
func (_Datastore *DatastoreTransactorSession) StoreData(newStrings [][]string) (*types.Transaction, error) {
	return _Datastore.Contract.StoreData(&_Datastore.TransactOpts, newStrings)
}

// DatastoreStoredIterator is returned from FilterStored and is used to iterate over the raw logs and unpacked data for Stored events raised by the Datastore contract.
type DatastoreStoredIterator struct {
	Event *DatastoreStored // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DatastoreStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatastoreStored)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DatastoreStored)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DatastoreStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatastoreStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatastoreStored represents a Stored event raised by the Datastore contract.
type DatastoreStored struct {
	ExamNo common.Hash
	Data   []string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStored is a free log retrieval operation binding the contract event 0x54036692cf01024a8aa507eab9ba1f9bdea128475be185471f57c70ab4f3a878.
//
// Solidity: event Stored(string indexed exam_no, string[] data)
func (_Datastore *DatastoreFilterer) FilterStored(opts *bind.FilterOpts, exam_no []string) (*DatastoreStoredIterator, error) {

	var exam_noRule []interface{}
	for _, exam_noItem := range exam_no {
		exam_noRule = append(exam_noRule, exam_noItem)
	}

	logs, sub, err := _Datastore.contract.FilterLogs(opts, "Stored", exam_noRule)
	if err != nil {
		return nil, err
	}
	return &DatastoreStoredIterator{contract: _Datastore.contract, event: "Stored", logs: logs, sub: sub}, nil
}

// WatchStored is a free log subscription operation binding the contract event 0x54036692cf01024a8aa507eab9ba1f9bdea128475be185471f57c70ab4f3a878.
//
// Solidity: event Stored(string indexed exam_no, string[] data)
func (_Datastore *DatastoreFilterer) WatchStored(opts *bind.WatchOpts, sink chan<- *DatastoreStored, exam_no []string) (event.Subscription, error) {

	var exam_noRule []interface{}
	for _, exam_noItem := range exam_no {
		exam_noRule = append(exam_noRule, exam_noItem)
	}

	logs, sub, err := _Datastore.contract.WatchLogs(opts, "Stored", exam_noRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatastoreStored)
				if err := _Datastore.contract.UnpackLog(event, "Stored", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStored is a log parse operation binding the contract event 0x54036692cf01024a8aa507eab9ba1f9bdea128475be185471f57c70ab4f3a878.
//
// Solidity: event Stored(string indexed exam_no, string[] data)
func (_Datastore *DatastoreFilterer) ParseStored(log types.Log) (*DatastoreStored, error) {
	event := new(DatastoreStored)
	if err := _Datastore.contract.UnpackLog(event, "Stored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
