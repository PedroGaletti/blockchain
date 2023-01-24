// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
)

// ContractTask is an auto generated low-level Go binding around an user-defined struct.
type ContractTask struct {
	Content string
	Status  bool
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structContract.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structContract.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"toggle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611057806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100d85780639507d39a146100f6578063b0c8f9dc14610126578063f745630f146101425761007d565b80630f560cd7146100825780634cc82215146100a0578063751ef753146100bc575b600080fd5b61008a61015e565b60405161009791906107da565b60405180910390f35b6100ba60048036038101906100b59190610846565b61026a565b005b6100d660048036038101906100d19190610846565b610382565b005b6100e06103f5565b6040516100ed91906108b4565b60405180910390f35b610110600480360381019061010b9190610846565b610419565b60405161011d919061090c565b60405180910390f35b610140600480360381019061013b9190610a63565b610502565b005b61015c60048036038101906101579190610aac565b610580565b005b60606001805480602002602001604051908101604052809291908181526020016000905b8282101561026157838290600052602060002090600202016040518060400160405290816000820180546101b590610b37565b80601f01602080910402602001604051908101604052809291908181526020018280546101e190610b37565b801561022e5780601f106102035761010080835404028352916020019161022e565b820191906000526020600020905b81548152906001019060200180831161021157829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152505081526020019060010190610182565b50505050905090565b60008190505b600180805490506102819190610b97565b81101561032f57600180826102969190610bcb565b815481106102a7576102a6610bff565b5b9060005260206000209060020201600182815481106102c9576102c8610bff565b5b9060005260206000209060020201600082018160000190816102eb9190610df0565b506001820160009054906101000a900460ff168160010160006101000a81548160ff021916908315150217905550905050808061032790610ed8565b915050610270565b50600180548061034257610341610f20565b5b60019003818190600052602060002090600202016000808201600061036791906105b7565b6001820160006101000a81549060ff02191690555050905550565b6001818154811061039657610395610bff565b5b906000526020600020906002020160010160009054906101000a900460ff1615600182815481106103ca576103c9610bff565b5b906000526020600020906002020160010160006101000a81548160ff02191690831515021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6104216105f7565b6001828154811061043557610434610bff565b5b906000526020600020906002020160405180604001604052908160008201805461045e90610b37565b80601f016020809104026020016040519081016040528092919081815260200182805461048a90610b37565b80156104d75780601f106104ac576101008083540402835291602001916104d7565b820191906000526020600020905b8154815290600101906020018083116104ba57829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050919050565b60016040518060400160405280838152602001600015158152509080600181540180825580915050600190039060005260206000209060020201600090919091909150600082015181600001908161055a9190610f4f565b5060208201518160010160006101000a81548160ff021916908315150217905550505050565b806001838154811061059557610594610bff565b5b906000526020600020906002020160000190816105b29190610f4f565b505050565b5080546105c390610b37565b6000825580601f106105d557506105f4565b601f0160209004906000526020600020908101906105f39190610613565b5b50565b6040518060400160405280606081526020016000151581525090565b5b8082111561062c576000816000905550600101610614565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561069657808201518184015260208101905061067b565b60008484015250505050565b6000601f19601f8301169050919050565b60006106be8261065c565b6106c88185610667565b93506106d8818560208601610678565b6106e1816106a2565b840191505092915050565b60008115159050919050565b610701816106ec565b82525050565b6000604083016000830151848203600086015261072482826106b3565b915050602083015161073960208601826106f8565b508091505092915050565b60006107508383610707565b905092915050565b6000602082019050919050565b600061077082610630565b61077a818561063b565b93508360208202850161078c8561064c565b8060005b858110156107c857848403895281516107a98582610744565b94506107b483610758565b925060208a01995050600181019050610790565b50829750879550505050505092915050565b600060208201905081810360008301526107f48184610765565b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61082381610810565b811461082e57600080fd5b50565b6000813590506108408161081a565b92915050565b60006020828403121561085c5761085b610806565b5b600061086a84828501610831565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061089e82610873565b9050919050565b6108ae81610893565b82525050565b60006020820190506108c960008301846108a5565b92915050565b600060408301600083015184820360008601526108ec82826106b3565b915050602083015161090160208601826106f8565b508091505092915050565b6000602082019050818103600083015261092681846108cf565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610970826106a2565b810181811067ffffffffffffffff8211171561098f5761098e610938565b5b80604052505050565b60006109a26107fc565b90506109ae8282610967565b919050565b600067ffffffffffffffff8211156109ce576109cd610938565b5b6109d7826106a2565b9050602081019050919050565b82818337600083830152505050565b6000610a06610a01846109b3565b610998565b905082815260208101848484011115610a2257610a21610933565b5b610a2d8482856109e4565b509392505050565b600082601f830112610a4a57610a4961092e565b5b8135610a5a8482602086016109f3565b91505092915050565b600060208284031215610a7957610a78610806565b5b600082013567ffffffffffffffff811115610a9757610a9661080b565b5b610aa384828501610a35565b91505092915050565b60008060408385031215610ac357610ac2610806565b5b6000610ad185828601610831565b925050602083013567ffffffffffffffff811115610af257610af161080b565b5b610afe85828601610a35565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610b4f57607f821691505b602082108103610b6257610b61610b08565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ba282610810565b9150610bad83610810565b9250828203905081811115610bc557610bc4610b68565b5b92915050565b6000610bd682610810565b9150610be183610810565b9250828201905080821115610bf957610bf8610b68565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081549050610c3d81610b37565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610ca67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610c69565b610cb08683610c69565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610ced610ce8610ce384610810565b610cc8565b610810565b9050919050565b6000819050919050565b610d0783610cd2565b610d1b610d1382610cf4565b848454610c76565b825550505050565b600090565b610d30610d23565b610d3b818484610cfe565b505050565b5b81811015610d5f57610d54600082610d28565b600181019050610d41565b5050565b601f821115610da457610d7581610c44565b610d7e84610c59565b81016020851015610d8d578190505b610da1610d9985610c59565b830182610d40565b50505b505050565b600082821c905092915050565b6000610dc760001984600802610da9565b1980831691505092915050565b6000610de08383610db6565b9150826002028217905092915050565b818103610dfe575050610ed6565b610e0782610c2e565b67ffffffffffffffff811115610e2057610e1f610938565b5b610e2a8254610b37565b610e35828285610d63565b6000601f831160018114610e645760008415610e52578287015490505b610e5c8582610dd4565b865550610ecf565b601f198416610e7287610c44565b9650610e7d86610c44565b60005b82811015610ea557848901548255600182019150600185019450602081019050610e80565b86831015610ec25784890154610ebe601f891682610db6565b8355505b6001600288020188555050505b5050505050505b565b6000610ee382610810565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f1557610f14610b68565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b610f588261065c565b67ffffffffffffffff811115610f7157610f70610938565b5b610f7b8254610b37565b610f86828285610d63565b600060209050601f831160018114610fb95760008415610fa7578287015190505b610fb18582610dd4565b865550611019565b601f198416610fc786610c44565b60005b82811015610fef57848901518255600182019150602085019450602081019050610fca565b8683101561100c5784890151611008601f891682610db6565b8355505b6001600288020188555050505b50505050505056fea26469706673582212201719d82344ba1dbb8f574253e721902ee06724a41b346be146cc048163e32edd64736f6c63430008110033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Contract *ContractCaller) Get(opts *bind.CallOpts, _id *big.Int) (ContractTask, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get", _id)

	if err != nil {
		return *new(ContractTask), err
	}

	out0 := *abi.ConvertType(out[0], new(ContractTask)).(*ContractTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Contract *ContractSession) Get(_id *big.Int) (ContractTask, error) {
	return _Contract.Contract.Get(&_Contract.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Contract *ContractCallerSession) Get(_id *big.Int) (ContractTask, error) {
	return _Contract.Contract.Get(&_Contract.CallOpts, _id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Contract *ContractCaller) List(opts *bind.CallOpts) ([]ContractTask, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]ContractTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]ContractTask)).(*[]ContractTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Contract *ContractSession) List() ([]ContractTask, error) {
	return _Contract.Contract.List(&_Contract.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Contract *ContractCallerSession) List() ([]ContractTask, error) {
	return _Contract.Contract.List(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Contract *ContractTransactor) Add(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "add", _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Contract *ContractSession) Add(_content string) (*types.Transaction, error) {
	return _Contract.Contract.Add(&_Contract.TransactOpts, _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Contract *ContractTransactorSession) Add(_content string) (*types.Transaction, error) {
	return _Contract.Contract.Add(&_Contract.TransactOpts, _content)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Contract *ContractTransactor) Remove(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove", _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Contract *ContractSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Remove(&_Contract.TransactOpts, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Contract *ContractTransactorSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Remove(&_Contract.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Contract *ContractTransactor) Toggle(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "toggle", _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Contract *ContractSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Toggle(&_Contract.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Contract *ContractTransactorSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Toggle(&_Contract.TransactOpts, _id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Contract *ContractTransactor) Update(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "update", _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Contract *ContractSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Contract.Contract.Update(&_Contract.TransactOpts, _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Contract *ContractTransactorSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Contract.Contract.Update(&_Contract.TransactOpts, _id, _content)
}
