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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"exam_no\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"Stored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"EventCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"first\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[][]\",\"name\":\"newStrings\",\"type\":\"string[][]\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506108628061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80634f7de7f814610038578063ea17fc7c14610054575b5f5ffd5b610052600480360381019061004d9190610625565b610086565b005b61006e6004803603810190610069919061066c565b6102cf565b60405161007d939291906106e5565b60405180910390f35b5f5f90505b81518110156102cb577f05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae998282815181106100c8576100c761071a565b5b60200260200101516003815181106100e3576100e261071a565b5b60200260200101518383815181106100fe576100fd61071a565b5b60200260200101516001815181106101195761011861071a565b5b602002602001015160405161012f9291906107a7565b60405180910390a1435f83838151811061014c5761014b61071a565b5b60200260200101516003815181106101675761016661071a565b5b602002602001015160405161017c9190610816565b9081526020016040518091039020600101819055505f8282815181106101a5576101a461071a565b5b60200260200101516003815181106101c0576101bf61071a565b5b60200260200101516040516101d59190610816565b90815260200160405180910390206002015f9054906101000a900460ff166102be57435f83838151811061020c5761020b61071a565b5b60200260200101516003815181106102275761022661071a565b5b602002602001015160405161023c9190610816565b90815260200160405180910390205f018190555060015f8383815181106102665761026561071a565b5b60200260200101516003815181106102815761028061071a565b5b60200260200101516040516102969190610816565b90815260200160405180910390206002015f6101000a81548160ff0219169083151502179055505b808060010191505061008b565b5050565b5f818051602081018201805184825260208301602085012081835280955050505050505f91509050805f015490806001015490806002015f9054906101000a900460ff16905083565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6103738261032d565b810181811067ffffffffffffffff821117156103925761039161033d565b5b80604052505050565b5f6103a4610318565b90506103b0828261036a565b919050565b5f67ffffffffffffffff8211156103cf576103ce61033d565b5b602082029050602081019050919050565b5f5ffd5b5f67ffffffffffffffff8211156103fe576103fd61033d565b5b602082029050602081019050919050565b5f5ffd5b5f67ffffffffffffffff82111561042d5761042c61033d565b5b6104368261032d565b9050602081019050919050565b828183375f83830152505050565b5f61046361045e84610413565b61039b565b90508281526020810184848401111561047f5761047e61040f565b5b61048a848285610443565b509392505050565b5f82601f8301126104a6576104a5610329565b5b81356104b6848260208601610451565b91505092915050565b5f6104d16104cc846103e4565b61039b565b905080838252602082019050602084028301858111156104f4576104f36103e0565b5b835b8181101561053b57803567ffffffffffffffff81111561051957610518610329565b5b8086016105268982610492565b855260208501945050506020810190506104f6565b5050509392505050565b5f82601f83011261055957610558610329565b5b81356105698482602086016104bf565b91505092915050565b5f61058461057f846103b5565b61039b565b905080838252602082019050602084028301858111156105a7576105a66103e0565b5b835b818110156105ee57803567ffffffffffffffff8111156105cc576105cb610329565b5b8086016105d98982610545565b855260208501945050506020810190506105a9565b5050509392505050565b5f82601f83011261060c5761060b610329565b5b813561061c848260208601610572565b91505092915050565b5f6020828403121561063a57610639610321565b5b5f82013567ffffffffffffffff81111561065757610656610325565b5b610663848285016105f8565b91505092915050565b5f6020828403121561068157610680610321565b5b5f82013567ffffffffffffffff81111561069e5761069d610325565b5b6106aa84828501610492565b91505092915050565b5f819050919050565b6106c5816106b3565b82525050565b5f8115159050919050565b6106df816106cb565b82525050565b5f6060820190506106f85f8301866106bc565b61070560208301856106bc565b61071260408301846106d6565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61077982610747565b6107838185610751565b9350610793818560208601610761565b61079c8161032d565b840191505092915050565b5f6040820190508181035f8301526107bf818561076f565b905081810360208301526107d3818461076f565b90509392505050565b5f81905092915050565b5f6107f082610747565b6107fa81856107dc565b935061080a818560208601610761565b80840191505092915050565b5f61082182846107e6565b91508190509291505056fea26469706673582212202192c984d5fb839dc7278a58162968181b0b64ea4c434fac81995a77841039ce64736f6c634300081b0033",
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
	ExamNo string
	Data   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStored is a free log retrieval operation binding the contract event 0x05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae99.
//
// Solidity: event Stored(string exam_no, string data)
func (_Datastore *DatastoreFilterer) FilterStored(opts *bind.FilterOpts) (*DatastoreStoredIterator, error) {

	logs, sub, err := _Datastore.contract.FilterLogs(opts, "Stored")
	if err != nil {
		return nil, err
	}
	return &DatastoreStoredIterator{contract: _Datastore.contract, event: "Stored", logs: logs, sub: sub}, nil
}

// WatchStored is a free log subscription operation binding the contract event 0x05cf4538ca45b8f0999a133df2f77a0e5648ca2349b8df20169cbc6328fdae99.
//
// Solidity: event Stored(string exam_no, string data)
func (_Datastore *DatastoreFilterer) WatchStored(opts *bind.WatchOpts, sink chan<- *DatastoreStored) (event.Subscription, error) {

	logs, sub, err := _Datastore.contract.WatchLogs(opts, "Stored")
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
// Solidity: event Stored(string exam_no, string data)
func (_Datastore *DatastoreFilterer) ParseStored(log types.Log) (*DatastoreStored, error) {
	event := new(DatastoreStored)
	if err := _Datastore.contract.UnpackLog(event, "Stored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
