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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"exam_no\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"data\",\"type\":\"string[]\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"eventCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"first\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllData\",\"outputs\":[{\"internalType\":\"string[][]\",\"name\":\"\",\"type\":\"string[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDataArray\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_exam_no\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"getDataMap\",\"outputs\":[{\"internalType\":\"string[][]\",\"name\":\"\",\"type\":\"string[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[][]\",\"name\":\"newStrings\",\"type\":\"string[][]\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[][]\",\"name\":\"newStrings\",\"type\":\"string[][]\"}],\"name\":\"storeDataMap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[][]\",\"name\":\"newStrings\",\"type\":\"string[][]\"}],\"name\":\"storeDataWithEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611952806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806355e0ade21161006657806355e0ade21461010a57806373d4a13a1461013a5780639f64d2a314610158578063a1b8127414610174578063ec1f18ba146101a457610093565b80632d10fa28146100985780633808ee05146100b657806347064d6a146100d25780634f7de7f8146100ee575b600080fd5b6100a06101d6565b6040516100ad9190610e9a565b60405180910390f35b6100d060048036038101906100cb91906111cc565b6102f6565b005b6100ec60048036038101906100e79190611215565b6105ea565b005b610108600480360381019061010391906111cc565b6105fd565b005b610124600480360381019061011f9190611294565b610617565b6040516101319190611347565b60405180910390f35b61014261070e565b60405161014f91906113b3565b60405180910390f35b610172600480360381019061016d91906111cc565b61079c565b005b61018e600480360381019061018991906113d5565b6108b1565b60405161019b9190610e9a565b60405180910390f35b6101be60048036038101906101b99190611215565b610aba565b6040516101cd9392919061146e565b60405180910390f35b60606001805480602002602001604051908101604052809291908181526020016000905b828210156102ed57838290600052602060002001805480602002602001604051908101604052809291908181526020016000905b828210156102da57838290600052602060002001805461024d906114d4565b80601f0160208091040260200160405190810160405280929190818152602001828054610279906114d4565b80156102c65780601f1061029b576101008083540402835291602001916102c6565b820191906000526020600020905b8154815290600101906020018083116102a957829003601f168201915b50505050508152602001906001019061022e565b50505050815260200190600101906101fa565b50505050905090565b60005b81518110156105e657600282828151811061031757610316611505565b5b602002602001015160038151811061033257610331611505565b5b60200260200101516040516103479190611570565b908152602001604051809103902082828151811061036857610367611505565b5b60200260200101519080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906103a8929190610b07565b508181815181106103bc576103bb611505565b5b60200260200101516003815181106103d7576103d6611505565b5b60200260200101516040516103ec9190611570565b60405180910390207fff9d822953cc36b9d20eb70835ee8f63607bffb5a7cf7ed83e03acaf5acf299183838151811061042857610427611505565b5b602002602001015160405161043d9190611347565b60405180910390a243600483838151811061045b5761045a611505565b5b602002602001015160038151811061047657610475611505565b5b602002602001015160405161048b9190611570565b90815260200160405180910390206001018190555060048282815181106104b5576104b4611505565b5b60200260200101516003815181106104d0576104cf611505565b5b60200260200101516040516104e59190611570565b908152602001604051809103902060020160009054906101000a900460ff166105d35743600483838151811061051e5761051d611505565b5b602002602001015160038151811061053957610538611505565b5b602002602001015160405161054e9190611570565b9081526020016040518091039020600001819055506001600483838151811061057a57610579611505565b5b602002602001015160038151811061059557610594611505565b5b60200260200101516040516105aa9190611570565b908152602001604051809103902060020160006101000a81548160ff0219169083151502179055505b80806105de906115b6565b9150506102f9565b5050565b80600090816105f991906117aa565b5050565b8060019080519060200190610613929190610b60565b5050565b60606001828154811061062d5761062c611505565b5b90600052602060002001805480602002602001604051908101604052809291908181526020016000905b82821015610703578382906000526020600020018054610676906114d4565b80601f01602080910402602001604051908101604052809291908181526020018280546106a2906114d4565b80156106ef5780601f106106c4576101008083540402835291602001916106ef565b820191906000526020600020905b8154815290600101906020018083116106d257829003601f168201915b505050505081526020019060010190610657565b505050509050919050565b6000805461071b906114d4565b80601f0160208091040260200160405190810160405280929190818152602001828054610747906114d4565b80156107945780601f1061076957610100808354040283529160200191610794565b820191906000526020600020905b81548152906001019060200180831161077757829003601f168201915b505050505081565b60005b81518110156108ad5760018282815181106107bd576107bc611505565b5b60200260200101519080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906107fd929190610b07565b5081818151811061081157610810611505565b5b602002602001015160038151811061082c5761082b611505565b5b60200260200101516040516108419190611570565b60405180910390207fff9d822953cc36b9d20eb70835ee8f63607bffb5a7cf7ed83e03acaf5acf299183838151811061087d5761087c611505565b5b60200260200101516040516108929190611347565b60405180910390a280806108a5906115b6565b91505061079f565b5050565b606060006002856040516108c59190611570565b90815260200160405180910390209050808054905083856108e6919061187c565b1115610927576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091e906118fc565b60405180910390fd5b60008367ffffffffffffffff81111561094357610942610ed5565b5b60405190808252806020026020018201604052801561097657816020015b60608152602001906001900390816109615790505b50905060005b84811015610aad57828187610991919061187c565b815481106109a2576109a1611505565b5b90600052602060002001805480602002602001604051908101604052809291908181526020016000905b82821015610a785783829060005260206000200180546109eb906114d4565b80601f0160208091040260200160405190810160405280929190818152602001828054610a17906114d4565b8015610a645780601f10610a3957610100808354040283529160200191610a64565b820191906000526020600020905b815481529060010190602001808311610a4757829003601f168201915b5050505050815260200190600101906109cc565b50505050828281518110610a8f57610a8e611505565b5b60200260200101819052508080610aa5906115b6565b91505061097c565b5080925050509392505050565b6004818051602081018201805184825260208301602085012081835280955050505050506000915090508060000154908060010154908060020160009054906101000a900460ff16905083565b828054828255906000526020600020908101928215610b4f579160200282015b82811115610b4e578251829081610b3e91906117aa565b5091602001919060010190610b27565b5b509050610b5c9190610bc0565b5090565b828054828255906000526020600020908101928215610baf579160200282015b82811115610bae578251829080519060200190610b9e929190610b07565b5091602001919060010190610b80565b5b509050610bbc9190610be4565b5090565b5b80821115610be05760008181610bd79190610c08565b50600101610bc1565b5090565b5b80821115610c045760008181610bfb9190610c48565b50600101610be5565b5090565b508054610c14906114d4565b6000825580601f10610c265750610c45565b601f016020900490600052602060002090810190610c449190610c69565b5b50565b5080546000825590600052602060002090810190610c669190610bc0565b50565b5b80821115610c82576000816000905550600101610c6a565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610d18578082015181840152602081019050610cfd565b60008484015250505050565b6000601f19601f8301169050919050565b6000610d4082610cde565b610d4a8185610ce9565b9350610d5a818560208601610cfa565b610d6381610d24565b840191505092915050565b6000610d7a8383610d35565b905092915050565b6000602082019050919050565b6000610d9a82610cb2565b610da48185610cbd565b935083602082028501610db685610cce565b8060005b85811015610df25784840389528151610dd38582610d6e565b9450610dde83610d82565b925060208a01995050600181019050610dba565b50829750879550505050505092915050565b6000610e108383610d8f565b905092915050565b6000602082019050919050565b6000610e3082610c86565b610e3a8185610c91565b935083602082028501610e4c85610ca2565b8060005b85811015610e885784840389528151610e698582610e04565b9450610e7483610e18565b925060208a01995050600181019050610e50565b50829750879550505050505092915050565b60006020820190508181036000830152610eb48184610e25565b905092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f0d82610d24565b810181811067ffffffffffffffff82111715610f2c57610f2b610ed5565b5b80604052505050565b6000610f3f610ebc565b9050610f4b8282610f04565b919050565b600067ffffffffffffffff821115610f6b57610f6a610ed5565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff821115610f9c57610f9b610ed5565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff821115610fcd57610fcc610ed5565b5b610fd682610d24565b9050602081019050919050565b82818337600083830152505050565b600061100561100084610fb2565b610f35565b90508281526020810184848401111561102157611020610fad565b5b61102c848285610fe3565b509392505050565b600082601f83011261104957611048610ed0565b5b8135611059848260208601610ff2565b91505092915050565b600061107561107084610f81565b610f35565b9050808382526020820190506020840283018581111561109857611097610f7c565b5b835b818110156110df57803567ffffffffffffffff8111156110bd576110bc610ed0565b5b8086016110ca8982611034565b8552602085019450505060208101905061109a565b5050509392505050565b600082601f8301126110fe576110fd610ed0565b5b813561110e848260208601611062565b91505092915050565b600061112a61112584610f50565b610f35565b9050808382526020820190506020840283018581111561114d5761114c610f7c565b5b835b8181101561119457803567ffffffffffffffff81111561117257611171610ed0565b5b80860161117f89826110e9565b8552602085019450505060208101905061114f565b5050509392505050565b600082601f8301126111b3576111b2610ed0565b5b81356111c3848260208601611117565b91505092915050565b6000602082840312156111e2576111e1610ec6565b5b600082013567ffffffffffffffff811115611200576111ff610ecb565b5b61120c8482850161119e565b91505092915050565b60006020828403121561122b5761122a610ec6565b5b600082013567ffffffffffffffff81111561124957611248610ecb565b5b61125584828501611034565b91505092915050565b6000819050919050565b6112718161125e565b811461127c57600080fd5b50565b60008135905061128e81611268565b92915050565b6000602082840312156112aa576112a9610ec6565b5b60006112b88482850161127f565b91505092915050565b600082825260208201905092915050565b60006112dd82610cb2565b6112e781856112c1565b9350836020820285016112f985610cce565b8060005b8581101561133557848403895281516113168582610d6e565b945061132183610d82565b925060208a019950506001810190506112fd565b50829750879550505050505092915050565b6000602082019050818103600083015261136181846112d2565b905092915050565b600082825260208201905092915050565b600061138582610cde565b61138f8185611369565b935061139f818560208601610cfa565b6113a881610d24565b840191505092915050565b600060208201905081810360008301526113cd818461137a565b905092915050565b6000806000606084860312156113ee576113ed610ec6565b5b600084013567ffffffffffffffff81111561140c5761140b610ecb565b5b61141886828701611034565b93505060206114298682870161127f565b925050604061143a8682870161127f565b9150509250925092565b61144d8161125e565b82525050565b60008115159050919050565b61146881611453565b82525050565b60006060820190506114836000830186611444565b6114906020830185611444565b61149d604083018461145f565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806114ec57607f821691505b6020821081036114ff576114fe6114a5565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081905092915050565b600061154a82610cde565b6115548185611534565b9350611564818560208601610cfa565b80840191505092915050565b600061157c828461153f565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006115c18261125e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036115f3576115f2611587565b5b600182019050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026116607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611623565b61166a8683611623565b95508019841693508086168417925050509392505050565b6000819050919050565b60006116a76116a261169d8461125e565b611682565b61125e565b9050919050565b6000819050919050565b6116c18361168c565b6116d56116cd826116ae565b848454611630565b825550505050565b600090565b6116ea6116dd565b6116f58184846116b8565b505050565b5b818110156117195761170e6000826116e2565b6001810190506116fb565b5050565b601f82111561175e5761172f816115fe565b61173884611613565b81016020851015611747578190505b61175b61175385611613565b8301826116fa565b50505b505050565b600082821c905092915050565b600061178160001984600802611763565b1980831691505092915050565b600061179a8383611770565b9150826002028217905092915050565b6117b382610cde565b67ffffffffffffffff8111156117cc576117cb610ed5565b5b6117d682546114d4565b6117e182828561171d565b600060209050601f8311600181146118145760008415611802578287015190505b61180c858261178e565b865550611874565b601f198416611822866115fe565b60005b8281101561184a57848901518255600182019150602085019450602081019050611825565b868310156118675784890151611863601f891682611770565b8355505b6001600288020188555050505b505050505050565b60006118878261125e565b91506118928361125e565b92508282019050808211156118aa576118a9611587565b5b92915050565b7f4f7574206f6620626f756e647300000000000000000000000000000000000000600082015250565b60006118e6600d83611369565b91506118f1826118b0565b602082019050919050565b60006020820190508181036000830152611915816118d9565b905091905056fea26469706673582212205dba1fe76109436365d7e1969b6f3867bb1b031a1d0c5821c263e5f2714902ef64736f6c63430008120033",
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

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(string)
func (_Datastore *DatastoreCaller) Data(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "data")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(string)
func (_Datastore *DatastoreSession) Data() (string, error) {
	return _Datastore.Contract.Data(&_Datastore.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(string)
func (_Datastore *DatastoreCallerSession) Data() (string, error) {
	return _Datastore.Contract.Data(&_Datastore.CallOpts)
}

// EventCount is a free data retrieval call binding the contract method 0xec1f18ba.
//
// Solidity: function eventCount(string ) view returns(uint256 start, uint256 end, bool first)
func (_Datastore *DatastoreCaller) EventCount(opts *bind.CallOpts, arg0 string) (struct {
	Start *big.Int
	End   *big.Int
	First bool
}, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "eventCount", arg0)

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

// EventCount is a free data retrieval call binding the contract method 0xec1f18ba.
//
// Solidity: function eventCount(string ) view returns(uint256 start, uint256 end, bool first)
func (_Datastore *DatastoreSession) EventCount(arg0 string) (struct {
	Start *big.Int
	End   *big.Int
	First bool
}, error) {
	return _Datastore.Contract.EventCount(&_Datastore.CallOpts, arg0)
}

// EventCount is a free data retrieval call binding the contract method 0xec1f18ba.
//
// Solidity: function eventCount(string ) view returns(uint256 start, uint256 end, bool first)
func (_Datastore *DatastoreCallerSession) EventCount(arg0 string) (struct {
	Start *big.Int
	End   *big.Int
	First bool
}, error) {
	return _Datastore.Contract.EventCount(&_Datastore.CallOpts, arg0)
}

// GetAllData is a free data retrieval call binding the contract method 0x2d10fa28.
//
// Solidity: function getAllData() view returns(string[][])
func (_Datastore *DatastoreCaller) GetAllData(opts *bind.CallOpts) ([][]string, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getAllData")

	if err != nil {
		return *new([][]string), err
	}

	out0 := *abi.ConvertType(out[0], new([][]string)).(*[][]string)

	return out0, err

}

// GetAllData is a free data retrieval call binding the contract method 0x2d10fa28.
//
// Solidity: function getAllData() view returns(string[][])
func (_Datastore *DatastoreSession) GetAllData() ([][]string, error) {
	return _Datastore.Contract.GetAllData(&_Datastore.CallOpts)
}

// GetAllData is a free data retrieval call binding the contract method 0x2d10fa28.
//
// Solidity: function getAllData() view returns(string[][])
func (_Datastore *DatastoreCallerSession) GetAllData() ([][]string, error) {
	return _Datastore.Contract.GetAllData(&_Datastore.CallOpts)
}

// GetDataArray is a free data retrieval call binding the contract method 0x55e0ade2.
//
// Solidity: function getDataArray(uint256 index) view returns(string[])
func (_Datastore *DatastoreCaller) GetDataArray(opts *bind.CallOpts, index *big.Int) ([]string, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getDataArray", index)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetDataArray is a free data retrieval call binding the contract method 0x55e0ade2.
//
// Solidity: function getDataArray(uint256 index) view returns(string[])
func (_Datastore *DatastoreSession) GetDataArray(index *big.Int) ([]string, error) {
	return _Datastore.Contract.GetDataArray(&_Datastore.CallOpts, index)
}

// GetDataArray is a free data retrieval call binding the contract method 0x55e0ade2.
//
// Solidity: function getDataArray(uint256 index) view returns(string[])
func (_Datastore *DatastoreCallerSession) GetDataArray(index *big.Int) ([]string, error) {
	return _Datastore.Contract.GetDataArray(&_Datastore.CallOpts, index)
}

// GetDataMap is a free data retrieval call binding the contract method 0xa1b81274.
//
// Solidity: function getDataMap(string _exam_no, uint256 _start, uint256 _number) view returns(string[][])
func (_Datastore *DatastoreCaller) GetDataMap(opts *bind.CallOpts, _exam_no string, _start *big.Int, _number *big.Int) ([][]string, error) {
	var out []interface{}
	err := _Datastore.contract.Call(opts, &out, "getDataMap", _exam_no, _start, _number)

	if err != nil {
		return *new([][]string), err
	}

	out0 := *abi.ConvertType(out[0], new([][]string)).(*[][]string)

	return out0, err

}

// GetDataMap is a free data retrieval call binding the contract method 0xa1b81274.
//
// Solidity: function getDataMap(string _exam_no, uint256 _start, uint256 _number) view returns(string[][])
func (_Datastore *DatastoreSession) GetDataMap(_exam_no string, _start *big.Int, _number *big.Int) ([][]string, error) {
	return _Datastore.Contract.GetDataMap(&_Datastore.CallOpts, _exam_no, _start, _number)
}

// GetDataMap is a free data retrieval call binding the contract method 0xa1b81274.
//
// Solidity: function getDataMap(string _exam_no, uint256 _start, uint256 _number) view returns(string[][])
func (_Datastore *DatastoreCallerSession) GetDataMap(_exam_no string, _start *big.Int, _number *big.Int) ([][]string, error) {
	return _Datastore.Contract.GetDataMap(&_Datastore.CallOpts, _exam_no, _start, _number)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(string _data) returns()
func (_Datastore *DatastoreTransactor) SetData(opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "setData", _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(string _data) returns()
func (_Datastore *DatastoreSession) SetData(_data string) (*types.Transaction, error) {
	return _Datastore.Contract.SetData(&_Datastore.TransactOpts, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(string _data) returns()
func (_Datastore *DatastoreTransactorSession) SetData(_data string) (*types.Transaction, error) {
	return _Datastore.Contract.SetData(&_Datastore.TransactOpts, _data)
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

// StoreDataMap is a paid mutator transaction binding the contract method 0x3808ee05.
//
// Solidity: function storeDataMap(string[][] newStrings) returns()
func (_Datastore *DatastoreTransactor) StoreDataMap(opts *bind.TransactOpts, newStrings [][]string) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "storeDataMap", newStrings)
}

// StoreDataMap is a paid mutator transaction binding the contract method 0x3808ee05.
//
// Solidity: function storeDataMap(string[][] newStrings) returns()
func (_Datastore *DatastoreSession) StoreDataMap(newStrings [][]string) (*types.Transaction, error) {
	return _Datastore.Contract.StoreDataMap(&_Datastore.TransactOpts, newStrings)
}

// StoreDataMap is a paid mutator transaction binding the contract method 0x3808ee05.
//
// Solidity: function storeDataMap(string[][] newStrings) returns()
func (_Datastore *DatastoreTransactorSession) StoreDataMap(newStrings [][]string) (*types.Transaction, error) {
	return _Datastore.Contract.StoreDataMap(&_Datastore.TransactOpts, newStrings)
}

// StoreDataWithEvent is a paid mutator transaction binding the contract method 0x9f64d2a3.
//
// Solidity: function storeDataWithEvent(string[][] newStrings) returns()
func (_Datastore *DatastoreTransactor) StoreDataWithEvent(opts *bind.TransactOpts, newStrings [][]string) (*types.Transaction, error) {
	return _Datastore.contract.Transact(opts, "storeDataWithEvent", newStrings)
}

// StoreDataWithEvent is a paid mutator transaction binding the contract method 0x9f64d2a3.
//
// Solidity: function storeDataWithEvent(string[][] newStrings) returns()
func (_Datastore *DatastoreSession) StoreDataWithEvent(newStrings [][]string) (*types.Transaction, error) {
	return _Datastore.Contract.StoreDataWithEvent(&_Datastore.TransactOpts, newStrings)
}

// StoreDataWithEvent is a paid mutator transaction binding the contract method 0x9f64d2a3.
//
// Solidity: function storeDataWithEvent(string[][] newStrings) returns()
func (_Datastore *DatastoreTransactorSession) StoreDataWithEvent(newStrings [][]string) (*types.Transaction, error) {
	return _Datastore.Contract.StoreDataWithEvent(&_Datastore.TransactOpts, newStrings)
}

// DatastoreDataStoredIterator is returned from FilterDataStored and is used to iterate over the raw logs and unpacked data for DataStored events raised by the Datastore contract.
type DatastoreDataStoredIterator struct {
	Event *DatastoreDataStored // Event containing the contract specifics and raw log

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
func (it *DatastoreDataStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatastoreDataStored)
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
		it.Event = new(DatastoreDataStored)
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
func (it *DatastoreDataStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatastoreDataStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatastoreDataStored represents a DataStored event raised by the Datastore contract.
type DatastoreDataStored struct {
	ExamNo common.Hash
	Data   []string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDataStored is a free log retrieval operation binding the contract event 0xff9d822953cc36b9d20eb70835ee8f63607bffb5a7cf7ed83e03acaf5acf2991.
//
// Solidity: event DataStored(string indexed exam_no, string[] data)
func (_Datastore *DatastoreFilterer) FilterDataStored(opts *bind.FilterOpts, exam_no []string) (*DatastoreDataStoredIterator, error) {

	var exam_noRule []interface{}
	for _, exam_noItem := range exam_no {
		exam_noRule = append(exam_noRule, exam_noItem)
	}

	logs, sub, err := _Datastore.contract.FilterLogs(opts, "DataStored", exam_noRule)
	if err != nil {
		return nil, err
	}
	return &DatastoreDataStoredIterator{contract: _Datastore.contract, event: "DataStored", logs: logs, sub: sub}, nil
}

// WatchDataStored is a free log subscription operation binding the contract event 0xff9d822953cc36b9d20eb70835ee8f63607bffb5a7cf7ed83e03acaf5acf2991.
//
// Solidity: event DataStored(string indexed exam_no, string[] data)
func (_Datastore *DatastoreFilterer) WatchDataStored(opts *bind.WatchOpts, sink chan<- *DatastoreDataStored, exam_no []string) (event.Subscription, error) {

	var exam_noRule []interface{}
	for _, exam_noItem := range exam_no {
		exam_noRule = append(exam_noRule, exam_noItem)
	}

	logs, sub, err := _Datastore.contract.WatchLogs(opts, "DataStored", exam_noRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatastoreDataStored)
				if err := _Datastore.contract.UnpackLog(event, "DataStored", log); err != nil {
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

// ParseDataStored is a log parse operation binding the contract event 0xff9d822953cc36b9d20eb70835ee8f63607bffb5a7cf7ed83e03acaf5acf2991.
//
// Solidity: event DataStored(string indexed exam_no, string[] data)
func (_Datastore *DatastoreFilterer) ParseDataStored(log types.Log) (*DatastoreDataStored, error) {
	event := new(DatastoreDataStored)
	if err := _Datastore.contract.UnpackLog(event, "DataStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
