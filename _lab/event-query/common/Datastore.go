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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"exam_no\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"Stored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"EventCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"first\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[][]\",\"name\":\"newStrings\",\"type\":\"string[][]\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506108618061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80634f7de7f814610038578063ea17fc7c14610054575b5f5ffd5b610052600480360381019061004d9190610639565b610086565b005b61006e60048036038101906100699190610680565b6102e3565b60405161007d939291906106f9565b60405180910390f35b5f5f90505b81518110156102df578181815181106100a7576100a661072e565b5b60200260200101516003815181106100c2576100c161072e565b5b60200260200101516040516100d791906107ad565b60405180910390207f05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae998383815181106101135761011261072e565b5b602002602001015160018151811061012e5761012d61072e565b5b6020026020010151604051610143919061080b565b60405180910390a2435f8383815181106101605761015f61072e565b5b602002602001015160038151811061017b5761017a61072e565b5b602002602001015160405161019091906107ad565b9081526020016040518091039020600101819055505f8282815181106101b9576101b861072e565b5b60200260200101516003815181106101d4576101d361072e565b5b60200260200101516040516101e991906107ad565b90815260200160405180910390206002015f9054906101000a900460ff166102d257435f8383815181106102205761021f61072e565b5b602002602001015160038151811061023b5761023a61072e565b5b602002602001015160405161025091906107ad565b90815260200160405180910390205f018190555060015f83838151811061027a5761027961072e565b5b60200260200101516003815181106102955761029461072e565b5b60200260200101516040516102aa91906107ad565b90815260200160405180910390206002015f6101000a81548160ff0219169083151502179055505b808060010191505061008b565b5050565b5f818051602081018201805184825260208301602085012081835280955050505050505f91509050805f015490806001015490806002015f9054906101000a900460ff16905083565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61038782610341565b810181811067ffffffffffffffff821117156103a6576103a5610351565b5b80604052505050565b5f6103b861032c565b90506103c4828261037e565b919050565b5f67ffffffffffffffff8211156103e3576103e2610351565b5b602082029050602081019050919050565b5f5ffd5b5f67ffffffffffffffff82111561041257610411610351565b5b602082029050602081019050919050565b5f5ffd5b5f67ffffffffffffffff82111561044157610440610351565b5b61044a82610341565b9050602081019050919050565b828183375f83830152505050565b5f61047761047284610427565b6103af565b90508281526020810184848401111561049357610492610423565b5b61049e848285610457565b509392505050565b5f82601f8301126104ba576104b961033d565b5b81356104ca848260208601610465565b91505092915050565b5f6104e56104e0846103f8565b6103af565b90508083825260208201905060208402830185811115610508576105076103f4565b5b835b8181101561054f57803567ffffffffffffffff81111561052d5761052c61033d565b5b80860161053a89826104a6565b8552602085019450505060208101905061050a565b5050509392505050565b5f82601f83011261056d5761056c61033d565b5b813561057d8482602086016104d3565b91505092915050565b5f610598610593846103c9565b6103af565b905080838252602082019050602084028301858111156105bb576105ba6103f4565b5b835b8181101561060257803567ffffffffffffffff8111156105e0576105df61033d565b5b8086016105ed8982610559565b855260208501945050506020810190506105bd565b5050509392505050565b5f82601f8301126106205761061f61033d565b5b8135610630848260208601610586565b91505092915050565b5f6020828403121561064e5761064d610335565b5b5f82013567ffffffffffffffff81111561066b5761066a610339565b5b6106778482850161060c565b91505092915050565b5f6020828403121561069557610694610335565b5b5f82013567ffffffffffffffff8111156106b2576106b1610339565b5b6106be848285016104a6565b91505092915050565b5f819050919050565b6106d9816106c7565b82525050565b5f8115159050919050565b6106f3816106df565b82525050565b5f60608201905061070c5f8301866106d0565b61071960208301856106d0565b61072660408301846106ea565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f6107878261075b565b6107918185610765565b93506107a181856020860161076f565b80840191505092915050565b5f6107b8828461077d565b915081905092915050565b5f82825260208201905092915050565b5f6107dd8261075b565b6107e781856107c3565b93506107f781856020860161076f565b61080081610341565b840191505092915050565b5f6020820190508181035f83015261082381846107d3565b90509291505056fea2646970667358221220d1a5476a3facc289ca8386daf766cf956ebecf2c2427b370b61822d144bb401164736f6c634300081b0033",
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
	Data   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStored is a free log retrieval operation binding the contract event 0x05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae99.
//
// Solidity: event Stored(string indexed exam_no, string data)
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

// WatchStored is a free log subscription operation binding the contract event 0x05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae99.
//
// Solidity: event Stored(string indexed exam_no, string data)
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

// ParseStored is a log parse operation binding the contract event 0x05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae99.
//
// Solidity: event Stored(string indexed exam_no, string data)
func (_Datastore *DatastoreFilterer) ParseStored(log types.Log) (*DatastoreStored, error) {
	event := new(DatastoreStored)
	if err := _Datastore.contract.UnpackLog(event, "Stored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
