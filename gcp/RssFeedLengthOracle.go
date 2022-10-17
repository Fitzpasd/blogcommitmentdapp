// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// RssFeedLengthOracleFeedInfo is an auto generated low-level Go binding around an user-defined struct.
type RssFeedLengthOracleFeedInfo struct {
	Date   *big.Int
	Length *big.Int
}

// RssFeedLengthOracleMetaData contains all meta data concerning the RssFeedLengthOracle contract.
var RssFeedLengthOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DatesNotTracked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"FeedNotTracked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughDataSaved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UrlIsEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"getAllEntriesForUrl\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"internalType\":\"structRssFeedLengthOracle.FeedInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getLengthChangeWitinRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"getLengthOfFeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"other\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"setCurrentLengthOfFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RssFeedLengthOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use RssFeedLengthOracleMetaData.ABI instead.
var RssFeedLengthOracleABI = RssFeedLengthOracleMetaData.ABI

// RssFeedLengthOracle is an auto generated Go binding around an Ethereum contract.
type RssFeedLengthOracle struct {
	RssFeedLengthOracleCaller     // Read-only binding to the contract
	RssFeedLengthOracleTransactor // Write-only binding to the contract
	RssFeedLengthOracleFilterer   // Log filterer for contract events
}

// RssFeedLengthOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RssFeedLengthOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RssFeedLengthOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RssFeedLengthOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RssFeedLengthOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RssFeedLengthOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RssFeedLengthOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RssFeedLengthOracleSession struct {
	Contract     *RssFeedLengthOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RssFeedLengthOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RssFeedLengthOracleCallerSession struct {
	Contract *RssFeedLengthOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RssFeedLengthOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RssFeedLengthOracleTransactorSession struct {
	Contract     *RssFeedLengthOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RssFeedLengthOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RssFeedLengthOracleRaw struct {
	Contract *RssFeedLengthOracle // Generic contract binding to access the raw methods on
}

// RssFeedLengthOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RssFeedLengthOracleCallerRaw struct {
	Contract *RssFeedLengthOracleCaller // Generic read-only contract binding to access the raw methods on
}

// RssFeedLengthOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RssFeedLengthOracleTransactorRaw struct {
	Contract *RssFeedLengthOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRssFeedLengthOracle creates a new instance of RssFeedLengthOracle, bound to a specific deployed contract.
func NewRssFeedLengthOracle(address common.Address, backend bind.ContractBackend) (*RssFeedLengthOracle, error) {
	contract, err := bindRssFeedLengthOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RssFeedLengthOracle{RssFeedLengthOracleCaller: RssFeedLengthOracleCaller{contract: contract}, RssFeedLengthOracleTransactor: RssFeedLengthOracleTransactor{contract: contract}, RssFeedLengthOracleFilterer: RssFeedLengthOracleFilterer{contract: contract}}, nil
}

// NewRssFeedLengthOracleCaller creates a new read-only instance of RssFeedLengthOracle, bound to a specific deployed contract.
func NewRssFeedLengthOracleCaller(address common.Address, caller bind.ContractCaller) (*RssFeedLengthOracleCaller, error) {
	contract, err := bindRssFeedLengthOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RssFeedLengthOracleCaller{contract: contract}, nil
}

// NewRssFeedLengthOracleTransactor creates a new write-only instance of RssFeedLengthOracle, bound to a specific deployed contract.
func NewRssFeedLengthOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*RssFeedLengthOracleTransactor, error) {
	contract, err := bindRssFeedLengthOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RssFeedLengthOracleTransactor{contract: contract}, nil
}

// NewRssFeedLengthOracleFilterer creates a new log filterer instance of RssFeedLengthOracle, bound to a specific deployed contract.
func NewRssFeedLengthOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*RssFeedLengthOracleFilterer, error) {
	contract, err := bindRssFeedLengthOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RssFeedLengthOracleFilterer{contract: contract}, nil
}

// bindRssFeedLengthOracle binds a generic wrapper to an already deployed contract.
func bindRssFeedLengthOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RssFeedLengthOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RssFeedLengthOracle *RssFeedLengthOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RssFeedLengthOracle.Contract.RssFeedLengthOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RssFeedLengthOracle *RssFeedLengthOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RssFeedLengthOracle.Contract.RssFeedLengthOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RssFeedLengthOracle *RssFeedLengthOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RssFeedLengthOracle.Contract.RssFeedLengthOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RssFeedLengthOracle *RssFeedLengthOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RssFeedLengthOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RssFeedLengthOracle *RssFeedLengthOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RssFeedLengthOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RssFeedLengthOracle *RssFeedLengthOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RssFeedLengthOracle.Contract.contract.Transact(opts, method, params...)
}

// GetAllEntriesForUrl is a free data retrieval call binding the contract method 0xb63b8336.
//
// Solidity: function getAllEntriesForUrl(string url) view returns((uint256,uint256)[])
func (_RssFeedLengthOracle *RssFeedLengthOracleCaller) GetAllEntriesForUrl(opts *bind.CallOpts, url string) ([]RssFeedLengthOracleFeedInfo, error) {
	var out []interface{}
	err := _RssFeedLengthOracle.contract.Call(opts, &out, "getAllEntriesForUrl", url)

	if err != nil {
		return *new([]RssFeedLengthOracleFeedInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]RssFeedLengthOracleFeedInfo)).(*[]RssFeedLengthOracleFeedInfo)

	return out0, err

}

// GetAllEntriesForUrl is a free data retrieval call binding the contract method 0xb63b8336.
//
// Solidity: function getAllEntriesForUrl(string url) view returns((uint256,uint256)[])
func (_RssFeedLengthOracle *RssFeedLengthOracleSession) GetAllEntriesForUrl(url string) ([]RssFeedLengthOracleFeedInfo, error) {
	return _RssFeedLengthOracle.Contract.GetAllEntriesForUrl(&_RssFeedLengthOracle.CallOpts, url)
}

// GetAllEntriesForUrl is a free data retrieval call binding the contract method 0xb63b8336.
//
// Solidity: function getAllEntriesForUrl(string url) view returns((uint256,uint256)[])
func (_RssFeedLengthOracle *RssFeedLengthOracleCallerSession) GetAllEntriesForUrl(url string) ([]RssFeedLengthOracleFeedInfo, error) {
	return _RssFeedLengthOracle.Contract.GetAllEntriesForUrl(&_RssFeedLengthOracle.CallOpts, url)
}

// GetLengthChangeWitinRange is a free data retrieval call binding the contract method 0x9237016a.
//
// Solidity: function getLengthChangeWitinRange(string url, uint256 start, uint256 end) view returns(uint256)
func (_RssFeedLengthOracle *RssFeedLengthOracleCaller) GetLengthChangeWitinRange(opts *bind.CallOpts, url string, start *big.Int, end *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RssFeedLengthOracle.contract.Call(opts, &out, "getLengthChangeWitinRange", url, start, end)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLengthChangeWitinRange is a free data retrieval call binding the contract method 0x9237016a.
//
// Solidity: function getLengthChangeWitinRange(string url, uint256 start, uint256 end) view returns(uint256)
func (_RssFeedLengthOracle *RssFeedLengthOracleSession) GetLengthChangeWitinRange(url string, start *big.Int, end *big.Int) (*big.Int, error) {
	return _RssFeedLengthOracle.Contract.GetLengthChangeWitinRange(&_RssFeedLengthOracle.CallOpts, url, start, end)
}

// GetLengthChangeWitinRange is a free data retrieval call binding the contract method 0x9237016a.
//
// Solidity: function getLengthChangeWitinRange(string url, uint256 start, uint256 end) view returns(uint256)
func (_RssFeedLengthOracle *RssFeedLengthOracleCallerSession) GetLengthChangeWitinRange(url string, start *big.Int, end *big.Int) (*big.Int, error) {
	return _RssFeedLengthOracle.Contract.GetLengthChangeWitinRange(&_RssFeedLengthOracle.CallOpts, url, start, end)
}

// GetLengthOfFeed is a free data retrieval call binding the contract method 0x6c6eabed.
//
// Solidity: function getLengthOfFeed(string url) view returns(uint256)
func (_RssFeedLengthOracle *RssFeedLengthOracleCaller) GetLengthOfFeed(opts *bind.CallOpts, url string) (*big.Int, error) {
	var out []interface{}
	err := _RssFeedLengthOracle.contract.Call(opts, &out, "getLengthOfFeed", url)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLengthOfFeed is a free data retrieval call binding the contract method 0x6c6eabed.
//
// Solidity: function getLengthOfFeed(string url) view returns(uint256)
func (_RssFeedLengthOracle *RssFeedLengthOracleSession) GetLengthOfFeed(url string) (*big.Int, error) {
	return _RssFeedLengthOracle.Contract.GetLengthOfFeed(&_RssFeedLengthOracle.CallOpts, url)
}

// GetLengthOfFeed is a free data retrieval call binding the contract method 0x6c6eabed.
//
// Solidity: function getLengthOfFeed(string url) view returns(uint256)
func (_RssFeedLengthOracle *RssFeedLengthOracleCallerSession) GetLengthOfFeed(url string) (*big.Int, error) {
	return _RssFeedLengthOracle.Contract.GetLengthOfFeed(&_RssFeedLengthOracle.CallOpts, url)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address other, bool val) returns()
func (_RssFeedLengthOracle *RssFeedLengthOracleTransactor) SetAdmin(opts *bind.TransactOpts, other common.Address, val bool) (*types.Transaction, error) {
	return _RssFeedLengthOracle.contract.Transact(opts, "setAdmin", other, val)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address other, bool val) returns()
func (_RssFeedLengthOracle *RssFeedLengthOracleSession) SetAdmin(other common.Address, val bool) (*types.Transaction, error) {
	return _RssFeedLengthOracle.Contract.SetAdmin(&_RssFeedLengthOracle.TransactOpts, other, val)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address other, bool val) returns()
func (_RssFeedLengthOracle *RssFeedLengthOracleTransactorSession) SetAdmin(other common.Address, val bool) (*types.Transaction, error) {
	return _RssFeedLengthOracle.Contract.SetAdmin(&_RssFeedLengthOracle.TransactOpts, other, val)
}

// SetCurrentLengthOfFeed is a paid mutator transaction binding the contract method 0x20f02d47.
//
// Solidity: function setCurrentLengthOfFeed(string url, uint256 length) returns()
func (_RssFeedLengthOracle *RssFeedLengthOracleTransactor) SetCurrentLengthOfFeed(opts *bind.TransactOpts, url string, length *big.Int) (*types.Transaction, error) {
	return _RssFeedLengthOracle.contract.Transact(opts, "setCurrentLengthOfFeed", url, length)
}

// SetCurrentLengthOfFeed is a paid mutator transaction binding the contract method 0x20f02d47.
//
// Solidity: function setCurrentLengthOfFeed(string url, uint256 length) returns()
func (_RssFeedLengthOracle *RssFeedLengthOracleSession) SetCurrentLengthOfFeed(url string, length *big.Int) (*types.Transaction, error) {
	return _RssFeedLengthOracle.Contract.SetCurrentLengthOfFeed(&_RssFeedLengthOracle.TransactOpts, url, length)
}

// SetCurrentLengthOfFeed is a paid mutator transaction binding the contract method 0x20f02d47.
//
// Solidity: function setCurrentLengthOfFeed(string url, uint256 length) returns()
func (_RssFeedLengthOracle *RssFeedLengthOracleTransactorSession) SetCurrentLengthOfFeed(url string, length *big.Int) (*types.Transaction, error) {
	return _RssFeedLengthOracle.Contract.SetCurrentLengthOfFeed(&_RssFeedLengthOracle.TransactOpts, url, length)
}
