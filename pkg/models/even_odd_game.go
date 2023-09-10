// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package models

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

// EvenOddGameMetaData contains all meta data concerning the EvenOddGame contract.
var EvenOddGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"gameResult\",\"type\":\"bool\"}],\"name\":\"BetResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addr_requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"randomNum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"mainBet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_choice\",\"type\":\"uint8\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"randomNum\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"callerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumEvenOddGame.Choice\",\"name\":\"choice\",\"type\":\"uint8\"},{\"internalType\":\"enumEvenOddGame.GameState\",\"name\":\"gameStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gameDoneTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// EvenOddGameABI is the input ABI used to generate the binding from.
// Deprecated: Use EvenOddGameMetaData.ABI instead.
var EvenOddGameABI = EvenOddGameMetaData.ABI

// EvenOddGame is an auto generated Go binding around an Ethereum contract.
type EvenOddGame struct {
	EvenOddGameCaller     // Read-only binding to the contract
	EvenOddGameTransactor // Write-only binding to the contract
	EvenOddGameFilterer   // Log filterer for contract events
}

// EvenOddGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type EvenOddGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvenOddGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EvenOddGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvenOddGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EvenOddGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvenOddGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EvenOddGameSession struct {
	Contract     *EvenOddGame      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EvenOddGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EvenOddGameCallerSession struct {
	Contract *EvenOddGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EvenOddGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EvenOddGameTransactorSession struct {
	Contract     *EvenOddGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EvenOddGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type EvenOddGameRaw struct {
	Contract *EvenOddGame // Generic contract binding to access the raw methods on
}

// EvenOddGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EvenOddGameCallerRaw struct {
	Contract *EvenOddGameCaller // Generic read-only contract binding to access the raw methods on
}

// EvenOddGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EvenOddGameTransactorRaw struct {
	Contract *EvenOddGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvenOddGame creates a new instance of EvenOddGame, bound to a specific deployed contract.
func NewEvenOddGame(address common.Address, backend bind.ContractBackend) (*EvenOddGame, error) {
	contract, err := bindEvenOddGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EvenOddGame{EvenOddGameCaller: EvenOddGameCaller{contract: contract}, EvenOddGameTransactor: EvenOddGameTransactor{contract: contract}, EvenOddGameFilterer: EvenOddGameFilterer{contract: contract}}, nil
}

// NewEvenOddGameCaller creates a new read-only instance of EvenOddGame, bound to a specific deployed contract.
func NewEvenOddGameCaller(address common.Address, caller bind.ContractCaller) (*EvenOddGameCaller, error) {
	contract, err := bindEvenOddGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EvenOddGameCaller{contract: contract}, nil
}

// NewEvenOddGameTransactor creates a new write-only instance of EvenOddGame, bound to a specific deployed contract.
func NewEvenOddGameTransactor(address common.Address, transactor bind.ContractTransactor) (*EvenOddGameTransactor, error) {
	contract, err := bindEvenOddGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EvenOddGameTransactor{contract: contract}, nil
}

// NewEvenOddGameFilterer creates a new log filterer instance of EvenOddGame, bound to a specific deployed contract.
func NewEvenOddGameFilterer(address common.Address, filterer bind.ContractFilterer) (*EvenOddGameFilterer, error) {
	contract, err := bindEvenOddGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EvenOddGameFilterer{contract: contract}, nil
}

// bindEvenOddGame binds a generic wrapper to an already deployed contract.
func bindEvenOddGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EvenOddGameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EvenOddGame *EvenOddGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EvenOddGame.Contract.EvenOddGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EvenOddGame *EvenOddGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvenOddGame.Contract.EvenOddGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EvenOddGame *EvenOddGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EvenOddGame.Contract.EvenOddGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EvenOddGame *EvenOddGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EvenOddGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EvenOddGame *EvenOddGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvenOddGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EvenOddGame *EvenOddGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EvenOddGame.Contract.contract.Transact(opts, method, params...)
}

// AddrRequests is a free data retrieval call binding the contract method 0x3dfefd11.
//
// Solidity: function addr_requests(address , uint256 ) view returns(uint256)
func (_EvenOddGame *EvenOddGameCaller) AddrRequests(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EvenOddGame.contract.Call(opts, &out, "addr_requests", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddrRequests is a free data retrieval call binding the contract method 0x3dfefd11.
//
// Solidity: function addr_requests(address , uint256 ) view returns(uint256)
func (_EvenOddGame *EvenOddGameSession) AddrRequests(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _EvenOddGame.Contract.AddrRequests(&_EvenOddGame.CallOpts, arg0, arg1)
}

// AddrRequests is a free data retrieval call binding the contract method 0x3dfefd11.
//
// Solidity: function addr_requests(address , uint256 ) view returns(uint256)
func (_EvenOddGame *EvenOddGameCallerSession) AddrRequests(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _EvenOddGame.Contract.AddrRequests(&_EvenOddGame.CallOpts, arg0, arg1)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, address callerAddress, uint256 randomNum)
func (_EvenOddGame *EvenOddGameCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (struct {
	Fulfilled     bool
	CallerAddress common.Address
	RandomNum     *big.Int
}, error) {
	var out []interface{}
	err := _EvenOddGame.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(struct {
		Fulfilled     bool
		CallerAddress common.Address
		RandomNum     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CallerAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.RandomNum = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, address callerAddress, uint256 randomNum)
func (_EvenOddGame *EvenOddGameSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled     bool
	CallerAddress common.Address
	RandomNum     *big.Int
}, error) {
	return _EvenOddGame.Contract.GetRequestStatus(&_EvenOddGame.CallOpts, _requestId)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, address callerAddress, uint256 randomNum)
func (_EvenOddGame *EvenOddGameCallerSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled     bool
	CallerAddress common.Address
	RandomNum     *big.Int
}, error) {
	return _EvenOddGame.Contract.GetRequestStatus(&_EvenOddGame.CallOpts, _requestId)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_EvenOddGame *EvenOddGameCaller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EvenOddGame.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_EvenOddGame *EvenOddGameSession) LastRequestId() (*big.Int, error) {
	return _EvenOddGame.Contract.LastRequestId(&_EvenOddGame.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_EvenOddGame *EvenOddGameCallerSession) LastRequestId() (*big.Int, error) {
	return _EvenOddGame.Contract.LastRequestId(&_EvenOddGame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EvenOddGame *EvenOddGameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EvenOddGame.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EvenOddGame *EvenOddGameSession) Owner() (common.Address, error) {
	return _EvenOddGame.Contract.Owner(&_EvenOddGame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EvenOddGame *EvenOddGameCallerSession) Owner() (common.Address, error) {
	return _EvenOddGame.Contract.Owner(&_EvenOddGame.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_EvenOddGame *EvenOddGameCaller) RequestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EvenOddGame.contract.Call(opts, &out, "requestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_EvenOddGame *EvenOddGameSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _EvenOddGame.Contract.RequestIds(&_EvenOddGame.CallOpts, arg0)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_EvenOddGame *EvenOddGameCallerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _EvenOddGame.Contract.RequestIds(&_EvenOddGame.CallOpts, arg0)
}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(bool fulfilled, bool exists, uint256 randomNum, address callerAddress, uint256 amount, uint8 choice, uint8 gameStatus, uint256 gameDoneTimestamp)
func (_EvenOddGame *EvenOddGameCaller) SRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Fulfilled         bool
	Exists            bool
	RandomNum         *big.Int
	CallerAddress     common.Address
	Amount            *big.Int
	Choice            uint8
	GameStatus        uint8
	GameDoneTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _EvenOddGame.contract.Call(opts, &out, "s_requests", arg0)

	outstruct := new(struct {
		Fulfilled         bool
		Exists            bool
		RandomNum         *big.Int
		CallerAddress     common.Address
		Amount            *big.Int
		Choice            uint8
		GameStatus        uint8
		GameDoneTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RandomNum = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CallerAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Choice = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.GameStatus = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.GameDoneTimestamp = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(bool fulfilled, bool exists, uint256 randomNum, address callerAddress, uint256 amount, uint8 choice, uint8 gameStatus, uint256 gameDoneTimestamp)
func (_EvenOddGame *EvenOddGameSession) SRequests(arg0 *big.Int) (struct {
	Fulfilled         bool
	Exists            bool
	RandomNum         *big.Int
	CallerAddress     common.Address
	Amount            *big.Int
	Choice            uint8
	GameStatus        uint8
	GameDoneTimestamp *big.Int
}, error) {
	return _EvenOddGame.Contract.SRequests(&_EvenOddGame.CallOpts, arg0)
}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(bool fulfilled, bool exists, uint256 randomNum, address callerAddress, uint256 amount, uint8 choice, uint8 gameStatus, uint256 gameDoneTimestamp)
func (_EvenOddGame *EvenOddGameCallerSession) SRequests(arg0 *big.Int) (struct {
	Fulfilled         bool
	Exists            bool
	RandomNum         *big.Int
	CallerAddress     common.Address
	Amount            *big.Int
	Choice            uint8
	GameStatus        uint8
	GameDoneTimestamp *big.Int
}, error) {
	return _EvenOddGame.Contract.SRequests(&_EvenOddGame.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_EvenOddGame *EvenOddGameTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvenOddGame.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_EvenOddGame *EvenOddGameSession) AcceptOwnership() (*types.Transaction, error) {
	return _EvenOddGame.Contract.AcceptOwnership(&_EvenOddGame.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_EvenOddGame *EvenOddGameTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EvenOddGame.Contract.AcceptOwnership(&_EvenOddGame.TransactOpts)
}

// MainBet is a paid mutator transaction binding the contract method 0x1cae7fc2.
//
// Solidity: function mainBet(uint256 _requestId) returns(bool result)
func (_EvenOddGame *EvenOddGameTransactor) MainBet(opts *bind.TransactOpts, _requestId *big.Int) (*types.Transaction, error) {
	return _EvenOddGame.contract.Transact(opts, "mainBet", _requestId)
}

// MainBet is a paid mutator transaction binding the contract method 0x1cae7fc2.
//
// Solidity: function mainBet(uint256 _requestId) returns(bool result)
func (_EvenOddGame *EvenOddGameSession) MainBet(_requestId *big.Int) (*types.Transaction, error) {
	return _EvenOddGame.Contract.MainBet(&_EvenOddGame.TransactOpts, _requestId)
}

// MainBet is a paid mutator transaction binding the contract method 0x1cae7fc2.
//
// Solidity: function mainBet(uint256 _requestId) returns(bool result)
func (_EvenOddGame *EvenOddGameTransactorSession) MainBet(_requestId *big.Int) (*types.Transaction, error) {
	return _EvenOddGame.Contract.MainBet(&_EvenOddGame.TransactOpts, _requestId)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_EvenOddGame *EvenOddGameTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _EvenOddGame.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_EvenOddGame *EvenOddGameSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _EvenOddGame.Contract.RawFulfillRandomWords(&_EvenOddGame.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_EvenOddGame *EvenOddGameTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _EvenOddGame.Contract.RawFulfillRandomWords(&_EvenOddGame.TransactOpts, requestId, randomWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x4b765522.
//
// Solidity: function requestRandomWords(uint8 _choice) payable returns(uint256 requestId)
func (_EvenOddGame *EvenOddGameTransactor) RequestRandomWords(opts *bind.TransactOpts, _choice uint8) (*types.Transaction, error) {
	return _EvenOddGame.contract.Transact(opts, "requestRandomWords", _choice)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x4b765522.
//
// Solidity: function requestRandomWords(uint8 _choice) payable returns(uint256 requestId)
func (_EvenOddGame *EvenOddGameSession) RequestRandomWords(_choice uint8) (*types.Transaction, error) {
	return _EvenOddGame.Contract.RequestRandomWords(&_EvenOddGame.TransactOpts, _choice)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x4b765522.
//
// Solidity: function requestRandomWords(uint8 _choice) payable returns(uint256 requestId)
func (_EvenOddGame *EvenOddGameTransactorSession) RequestRandomWords(_choice uint8) (*types.Transaction, error) {
	return _EvenOddGame.Contract.RequestRandomWords(&_EvenOddGame.TransactOpts, _choice)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_EvenOddGame *EvenOddGameTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EvenOddGame.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_EvenOddGame *EvenOddGameSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EvenOddGame.Contract.TransferOwnership(&_EvenOddGame.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_EvenOddGame *EvenOddGameTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EvenOddGame.Contract.TransferOwnership(&_EvenOddGame.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_EvenOddGame *EvenOddGameTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvenOddGame.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_EvenOddGame *EvenOddGameSession) Withdraw() (*types.Transaction, error) {
	return _EvenOddGame.Contract.Withdraw(&_EvenOddGame.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_EvenOddGame *EvenOddGameTransactorSession) Withdraw() (*types.Transaction, error) {
	return _EvenOddGame.Contract.Withdraw(&_EvenOddGame.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EvenOddGame *EvenOddGameTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvenOddGame.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EvenOddGame *EvenOddGameSession) Receive() (*types.Transaction, error) {
	return _EvenOddGame.Contract.Receive(&_EvenOddGame.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EvenOddGame *EvenOddGameTransactorSession) Receive() (*types.Transaction, error) {
	return _EvenOddGame.Contract.Receive(&_EvenOddGame.TransactOpts)
}

// EvenOddGameBetResultIterator is returned from FilterBetResult and is used to iterate over the raw logs and unpacked data for BetResult events raised by the EvenOddGame contract.
type EvenOddGameBetResultIterator struct {
	Event *EvenOddGameBetResult // Event containing the contract specifics and raw log

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
func (it *EvenOddGameBetResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvenOddGameBetResult)
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
		it.Event = new(EvenOddGameBetResult)
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
func (it *EvenOddGameBetResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvenOddGameBetResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvenOddGameBetResult represents a BetResult event raised by the EvenOddGame contract.
type EvenOddGameBetResult struct {
	RequestId     *big.Int
	CallerAddress common.Address
	GameResult    bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBetResult is a free log retrieval operation binding the contract event 0x312d795b063a2c7cb51c0f1a67049c262d805e3b9adc88de44d294fcd4f890c0.
//
// Solidity: event BetResult(uint256 requestId, address callerAddress, bool gameResult)
func (_EvenOddGame *EvenOddGameFilterer) FilterBetResult(opts *bind.FilterOpts) (*EvenOddGameBetResultIterator, error) {

	logs, sub, err := _EvenOddGame.contract.FilterLogs(opts, "BetResult")
	if err != nil {
		return nil, err
	}
	return &EvenOddGameBetResultIterator{contract: _EvenOddGame.contract, event: "BetResult", logs: logs, sub: sub}, nil
}

// WatchBetResult is a free log subscription operation binding the contract event 0x312d795b063a2c7cb51c0f1a67049c262d805e3b9adc88de44d294fcd4f890c0.
//
// Solidity: event BetResult(uint256 requestId, address callerAddress, bool gameResult)
func (_EvenOddGame *EvenOddGameFilterer) WatchBetResult(opts *bind.WatchOpts, sink chan<- *EvenOddGameBetResult) (event.Subscription, error) {

	logs, sub, err := _EvenOddGame.contract.WatchLogs(opts, "BetResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvenOddGameBetResult)
				if err := _EvenOddGame.contract.UnpackLog(event, "BetResult", log); err != nil {
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

// ParseBetResult is a log parse operation binding the contract event 0x312d795b063a2c7cb51c0f1a67049c262d805e3b9adc88de44d294fcd4f890c0.
//
// Solidity: event BetResult(uint256 requestId, address callerAddress, bool gameResult)
func (_EvenOddGame *EvenOddGameFilterer) ParseBetResult(log types.Log) (*EvenOddGameBetResult, error) {
	event := new(EvenOddGameBetResult)
	if err := _EvenOddGame.contract.UnpackLog(event, "BetResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvenOddGameOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the EvenOddGame contract.
type EvenOddGameOwnershipTransferRequestedIterator struct {
	Event *EvenOddGameOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *EvenOddGameOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvenOddGameOwnershipTransferRequested)
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
		it.Event = new(EvenOddGameOwnershipTransferRequested)
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
func (it *EvenOddGameOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvenOddGameOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvenOddGameOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the EvenOddGame contract.
type EvenOddGameOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_EvenOddGame *EvenOddGameFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EvenOddGameOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EvenOddGame.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EvenOddGameOwnershipTransferRequestedIterator{contract: _EvenOddGame.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_EvenOddGame *EvenOddGameFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EvenOddGameOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EvenOddGame.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvenOddGameOwnershipTransferRequested)
				if err := _EvenOddGame.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_EvenOddGame *EvenOddGameFilterer) ParseOwnershipTransferRequested(log types.Log) (*EvenOddGameOwnershipTransferRequested, error) {
	event := new(EvenOddGameOwnershipTransferRequested)
	if err := _EvenOddGame.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvenOddGameOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EvenOddGame contract.
type EvenOddGameOwnershipTransferredIterator struct {
	Event *EvenOddGameOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EvenOddGameOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvenOddGameOwnershipTransferred)
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
		it.Event = new(EvenOddGameOwnershipTransferred)
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
func (it *EvenOddGameOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvenOddGameOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvenOddGameOwnershipTransferred represents a OwnershipTransferred event raised by the EvenOddGame contract.
type EvenOddGameOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_EvenOddGame *EvenOddGameFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EvenOddGameOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EvenOddGame.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EvenOddGameOwnershipTransferredIterator{contract: _EvenOddGame.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_EvenOddGame *EvenOddGameFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EvenOddGameOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EvenOddGame.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvenOddGameOwnershipTransferred)
				if err := _EvenOddGame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_EvenOddGame *EvenOddGameFilterer) ParseOwnershipTransferred(log types.Log) (*EvenOddGameOwnershipTransferred, error) {
	event := new(EvenOddGameOwnershipTransferred)
	if err := _EvenOddGame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvenOddGameRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the EvenOddGame contract.
type EvenOddGameRequestFulfilledIterator struct {
	Event *EvenOddGameRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *EvenOddGameRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvenOddGameRequestFulfilled)
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
		it.Event = new(EvenOddGameRequestFulfilled)
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
func (it *EvenOddGameRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvenOddGameRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvenOddGameRequestFulfilled represents a RequestFulfilled event raised by the EvenOddGame contract.
type EvenOddGameRequestFulfilled struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0xfe2e2d779dba245964d4e3ef9b994be63856fd568bf7d3ca9e224755cb1bd54d.
//
// Solidity: event RequestFulfilled(uint256 requestId, uint256[] randomWords)
func (_EvenOddGame *EvenOddGameFilterer) FilterRequestFulfilled(opts *bind.FilterOpts) (*EvenOddGameRequestFulfilledIterator, error) {

	logs, sub, err := _EvenOddGame.contract.FilterLogs(opts, "RequestFulfilled")
	if err != nil {
		return nil, err
	}
	return &EvenOddGameRequestFulfilledIterator{contract: _EvenOddGame.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0xfe2e2d779dba245964d4e3ef9b994be63856fd568bf7d3ca9e224755cb1bd54d.
//
// Solidity: event RequestFulfilled(uint256 requestId, uint256[] randomWords)
func (_EvenOddGame *EvenOddGameFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *EvenOddGameRequestFulfilled) (event.Subscription, error) {

	logs, sub, err := _EvenOddGame.contract.WatchLogs(opts, "RequestFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvenOddGameRequestFulfilled)
				if err := _EvenOddGame.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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

// ParseRequestFulfilled is a log parse operation binding the contract event 0xfe2e2d779dba245964d4e3ef9b994be63856fd568bf7d3ca9e224755cb1bd54d.
//
// Solidity: event RequestFulfilled(uint256 requestId, uint256[] randomWords)
func (_EvenOddGame *EvenOddGameFilterer) ParseRequestFulfilled(log types.Log) (*EvenOddGameRequestFulfilled, error) {
	event := new(EvenOddGameRequestFulfilled)
	if err := _EvenOddGame.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvenOddGameRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the EvenOddGame contract.
type EvenOddGameRequestSentIterator struct {
	Event *EvenOddGameRequestSent // Event containing the contract specifics and raw log

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
func (it *EvenOddGameRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvenOddGameRequestSent)
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
		it.Event = new(EvenOddGameRequestSent)
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
func (it *EvenOddGameRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvenOddGameRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvenOddGameRequestSent represents a RequestSent event raised by the EvenOddGame contract.
type EvenOddGameRequestSent struct {
	RequestId *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0x48b98ad7a8a8dbe21cc82bf98710ad4d2cdd949ccac393692e4d9a1722c162c7.
//
// Solidity: event RequestSent(uint256 requestId, address caller)
func (_EvenOddGame *EvenOddGameFilterer) FilterRequestSent(opts *bind.FilterOpts) (*EvenOddGameRequestSentIterator, error) {

	logs, sub, err := _EvenOddGame.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &EvenOddGameRequestSentIterator{contract: _EvenOddGame.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0x48b98ad7a8a8dbe21cc82bf98710ad4d2cdd949ccac393692e4d9a1722c162c7.
//
// Solidity: event RequestSent(uint256 requestId, address caller)
func (_EvenOddGame *EvenOddGameFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *EvenOddGameRequestSent) (event.Subscription, error) {

	logs, sub, err := _EvenOddGame.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvenOddGameRequestSent)
				if err := _EvenOddGame.contract.UnpackLog(event, "RequestSent", log); err != nil {
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

// ParseRequestSent is a log parse operation binding the contract event 0x48b98ad7a8a8dbe21cc82bf98710ad4d2cdd949ccac393692e4d9a1722c162c7.
//
// Solidity: event RequestSent(uint256 requestId, address caller)
func (_EvenOddGame *EvenOddGameFilterer) ParseRequestSent(log types.Log) (*EvenOddGameRequestSent, error) {
	event := new(EvenOddGameRequestSent)
	if err := _EvenOddGame.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
