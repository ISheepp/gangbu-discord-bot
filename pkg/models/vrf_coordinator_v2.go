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

// VRFCoordinatorV2MetaData contains all meta data concerning the VRFCoordinatorV2 contract.
var VRFCoordinatorV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"gameResult\",\"type\":\"bool\"}],\"name\":\"BetResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addr_requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"randomNum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"mainBet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_choice\",\"type\":\"uint8\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"randomNum\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"callerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumEvenOddGame.Choice\",\"name\":\"choice\",\"type\":\"uint8\"},{\"internalType\":\"enumEvenOddGame.GameState\",\"name\":\"gameStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gameDoneTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VRFCoordinatorV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFCoordinatorV2MetaData.ABI instead.
var VRFCoordinatorV2ABI = VRFCoordinatorV2MetaData.ABI

// VRFCoordinatorV2 is an auto generated Go binding around an Ethereum contract.
type VRFCoordinatorV2 struct {
	VRFCoordinatorV2Caller     // Read-only binding to the contract
	VRFCoordinatorV2Transactor // Write-only binding to the contract
	VRFCoordinatorV2Filterer   // Log filterer for contract events
}

// VRFCoordinatorV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type VRFCoordinatorV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFCoordinatorV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFCoordinatorV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFCoordinatorV2Session struct {
	Contract     *VRFCoordinatorV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFCoordinatorV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFCoordinatorV2CallerSession struct {
	Contract *VRFCoordinatorV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// VRFCoordinatorV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFCoordinatorV2TransactorSession struct {
	Contract     *VRFCoordinatorV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// VRFCoordinatorV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type VRFCoordinatorV2Raw struct {
	Contract *VRFCoordinatorV2 // Generic contract binding to access the raw methods on
}

// VRFCoordinatorV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFCoordinatorV2CallerRaw struct {
	Contract *VRFCoordinatorV2Caller // Generic read-only contract binding to access the raw methods on
}

// VRFCoordinatorV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFCoordinatorV2TransactorRaw struct {
	Contract *VRFCoordinatorV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFCoordinatorV2 creates a new instance of VRFCoordinatorV2, bound to a specific deployed contract.
func NewVRFCoordinatorV2(address common.Address, backend bind.ContractBackend) (*VRFCoordinatorV2, error) {
	contract, err := bindVRFCoordinatorV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2{VRFCoordinatorV2Caller: VRFCoordinatorV2Caller{contract: contract}, VRFCoordinatorV2Transactor: VRFCoordinatorV2Transactor{contract: contract}, VRFCoordinatorV2Filterer: VRFCoordinatorV2Filterer{contract: contract}}, nil
}

// NewVRFCoordinatorV2Caller creates a new read-only instance of VRFCoordinatorV2, bound to a specific deployed contract.
func NewVRFCoordinatorV2Caller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorV2Caller, error) {
	contract, err := bindVRFCoordinatorV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Caller{contract: contract}, nil
}

// NewVRFCoordinatorV2Transactor creates a new write-only instance of VRFCoordinatorV2, bound to a specific deployed contract.
func NewVRFCoordinatorV2Transactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorV2Transactor, error) {
	contract, err := bindVRFCoordinatorV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Transactor{contract: contract}, nil
}

// NewVRFCoordinatorV2Filterer creates a new log filterer instance of VRFCoordinatorV2, bound to a specific deployed contract.
func NewVRFCoordinatorV2Filterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorV2Filterer, error) {
	contract, err := bindVRFCoordinatorV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Filterer{contract: contract}, nil
}

// bindVRFCoordinatorV2 binds a generic wrapper to an already deployed contract.
func bindVRFCoordinatorV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFCoordinatorV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFCoordinatorV2 *VRFCoordinatorV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2.Contract.VRFCoordinatorV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFCoordinatorV2 *VRFCoordinatorV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.VRFCoordinatorV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFCoordinatorV2 *VRFCoordinatorV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.VRFCoordinatorV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.contract.Transact(opts, method, params...)
}

// AddrRequests is a free data retrieval call binding the contract method 0x3dfefd11.
//
// Solidity: function addr_requests(address , uint256 ) view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) AddrRequests(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "addr_requests", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddrRequests is a free data retrieval call binding the contract method 0x3dfefd11.
//
// Solidity: function addr_requests(address , uint256 ) view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) AddrRequests(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.AddrRequests(&_VRFCoordinatorV2.CallOpts, arg0, arg1)
}

// AddrRequests is a free data retrieval call binding the contract method 0x3dfefd11.
//
// Solidity: function addr_requests(address , uint256 ) view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) AddrRequests(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.AddrRequests(&_VRFCoordinatorV2.CallOpts, arg0, arg1)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, address callerAddress, uint256 randomNum)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (struct {
	Fulfilled     bool
	CallerAddress common.Address
	RandomNum     *big.Int
}, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "getRequestStatus", _requestId)

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
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled     bool
	CallerAddress common.Address
	RandomNum     *big.Int
}, error) {
	return _VRFCoordinatorV2.Contract.GetRequestStatus(&_VRFCoordinatorV2.CallOpts, _requestId)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, address callerAddress, uint256 randomNum)
func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled     bool
	CallerAddress common.Address
	RandomNum     *big.Int
}, error) {
	return _VRFCoordinatorV2.Contract.GetRequestStatus(&_VRFCoordinatorV2.CallOpts, _requestId)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) LastRequestId() (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.LastRequestId(&_VRFCoordinatorV2.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) LastRequestId() (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.LastRequestId(&_VRFCoordinatorV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) Owner() (common.Address, error) {
	return _VRFCoordinatorV2.Contract.Owner(&_VRFCoordinatorV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) Owner() (common.Address, error) {
	return _VRFCoordinatorV2.Contract.Owner(&_VRFCoordinatorV2.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) RequestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "requestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.RequestIds(&_VRFCoordinatorV2.CallOpts, arg0)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _VRFCoordinatorV2.Contract.RequestIds(&_VRFCoordinatorV2.CallOpts, arg0)
}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(bool fulfilled, bool exists, uint256 randomNum, address callerAddress, uint256 amount, uint8 choice, uint8 gameStatus, uint256 gameDoneTimestamp)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Caller) SRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _VRFCoordinatorV2.contract.Call(opts, &out, "s_requests", arg0)

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
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) SRequests(arg0 *big.Int) (struct {
	Fulfilled         bool
	Exists            bool
	RandomNum         *big.Int
	CallerAddress     common.Address
	Amount            *big.Int
	Choice            uint8
	GameStatus        uint8
	GameDoneTimestamp *big.Int
}, error) {
	return _VRFCoordinatorV2.Contract.SRequests(&_VRFCoordinatorV2.CallOpts, arg0)
}

// SRequests is a free data retrieval call binding the contract method 0xa168fa89.
//
// Solidity: function s_requests(uint256 ) view returns(bool fulfilled, bool exists, uint256 randomNum, address callerAddress, uint256 amount, uint8 choice, uint8 gameStatus, uint256 gameDoneTimestamp)
func (_VRFCoordinatorV2 *VRFCoordinatorV2CallerSession) SRequests(arg0 *big.Int) (struct {
	Fulfilled         bool
	Exists            bool
	RandomNum         *big.Int
	CallerAddress     common.Address
	Amount            *big.Int
	Choice            uint8
	GameStatus        uint8
	GameDoneTimestamp *big.Int
}, error) {
	return _VRFCoordinatorV2.Contract.SRequests(&_VRFCoordinatorV2.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.AcceptOwnership(&_VRFCoordinatorV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.AcceptOwnership(&_VRFCoordinatorV2.TransactOpts)
}

// MainBet is a paid mutator transaction binding the contract method 0x1cae7fc2.
//
// Solidity: function mainBet(uint256 _requestId) returns(bool result)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) MainBet(opts *bind.TransactOpts, _requestId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "mainBet", _requestId)
}

// MainBet is a paid mutator transaction binding the contract method 0x1cae7fc2.
//
// Solidity: function mainBet(uint256 _requestId) returns(bool result)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) MainBet(_requestId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.MainBet(&_VRFCoordinatorV2.TransactOpts, _requestId)
}

// MainBet is a paid mutator transaction binding the contract method 0x1cae7fc2.
//
// Solidity: function mainBet(uint256 _requestId) returns(bool result)
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) MainBet(_requestId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.MainBet(&_VRFCoordinatorV2.TransactOpts, _requestId)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RawFulfillRandomWords(&_VRFCoordinatorV2.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RawFulfillRandomWords(&_VRFCoordinatorV2.TransactOpts, requestId, randomWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x4b765522.
//
// Solidity: function requestRandomWords(uint8 _choice) payable returns(uint256 requestId)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) RequestRandomWords(opts *bind.TransactOpts, _choice uint8) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "requestRandomWords", _choice)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x4b765522.
//
// Solidity: function requestRandomWords(uint8 _choice) payable returns(uint256 requestId)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) RequestRandomWords(_choice uint8) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RequestRandomWords(&_VRFCoordinatorV2.TransactOpts, _choice)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x4b765522.
//
// Solidity: function requestRandomWords(uint8 _choice) payable returns(uint256 requestId)
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) RequestRandomWords(_choice uint8) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.RequestRandomWords(&_VRFCoordinatorV2.TransactOpts, _choice)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.TransferOwnership(&_VRFCoordinatorV2.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.TransferOwnership(&_VRFCoordinatorV2.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) Withdraw() (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.Withdraw(&_VRFCoordinatorV2.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) Withdraw() (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.Withdraw(&_VRFCoordinatorV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2Session) Receive() (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.Receive(&_VRFCoordinatorV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VRFCoordinatorV2 *VRFCoordinatorV2TransactorSession) Receive() (*types.Transaction, error) {
	return _VRFCoordinatorV2.Contract.Receive(&_VRFCoordinatorV2.TransactOpts)
}

// VRFCoordinatorV2BetResultIterator is returned from FilterBetResult and is used to iterate over the raw logs and unpacked data for BetResult events raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2BetResultIterator struct {
	Event *VRFCoordinatorV2BetResult // Event containing the contract specifics and raw log

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
func (it *VRFCoordinatorV2BetResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2BetResult)
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
		it.Event = new(VRFCoordinatorV2BetResult)
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
func (it *VRFCoordinatorV2BetResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFCoordinatorV2BetResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFCoordinatorV2BetResult represents a BetResult event raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2BetResult struct {
	RequestId     *big.Int
	CallerAddress common.Address
	GameResult    bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBetResult is a free log retrieval operation binding the contract event 0x312d795b063a2c7cb51c0f1a67049c262d805e3b9adc88de44d294fcd4f890c0.
//
// Solidity: event BetResult(uint256 requestId, address callerAddress, bool gameResult)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterBetResult(opts *bind.FilterOpts) (*VRFCoordinatorV2BetResultIterator, error) {

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "BetResult")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2BetResultIterator{contract: _VRFCoordinatorV2.contract, event: "BetResult", logs: logs, sub: sub}, nil
}

// WatchBetResult is a free log subscription operation binding the contract event 0x312d795b063a2c7cb51c0f1a67049c262d805e3b9adc88de44d294fcd4f890c0.
//
// Solidity: event BetResult(uint256 requestId, address callerAddress, bool gameResult)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchBetResult(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2BetResult) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "BetResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFCoordinatorV2BetResult)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "BetResult", log); err != nil {
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
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseBetResult(log types.Log) (*VRFCoordinatorV2BetResult, error) {
	event := new(VRFCoordinatorV2BetResult)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "BetResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFCoordinatorV2OwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2OwnershipTransferRequestedIterator struct {
	Event *VRFCoordinatorV2OwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *VRFCoordinatorV2OwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2OwnershipTransferRequested)
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
		it.Event = new(VRFCoordinatorV2OwnershipTransferRequested)
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
func (it *VRFCoordinatorV2OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFCoordinatorV2OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFCoordinatorV2OwnershipTransferRequested represents a OwnershipTransferRequested event raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2OwnershipTransferRequestedIterator{contract: _VRFCoordinatorV2.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFCoordinatorV2OwnershipTransferRequested)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV2OwnershipTransferRequested, error) {
	event := new(VRFCoordinatorV2OwnershipTransferRequested)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFCoordinatorV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2OwnershipTransferredIterator struct {
	Event *VRFCoordinatorV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VRFCoordinatorV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2OwnershipTransferred)
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
		it.Event = new(VRFCoordinatorV2OwnershipTransferred)
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
func (it *VRFCoordinatorV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFCoordinatorV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFCoordinatorV2OwnershipTransferred represents a OwnershipTransferred event raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2OwnershipTransferredIterator{contract: _VRFCoordinatorV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFCoordinatorV2OwnershipTransferred)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV2OwnershipTransferred, error) {
	event := new(VRFCoordinatorV2OwnershipTransferred)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFCoordinatorV2RequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2RequestFulfilledIterator struct {
	Event *VRFCoordinatorV2RequestFulfilled // Event containing the contract specifics and raw log

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
func (it *VRFCoordinatorV2RequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2RequestFulfilled)
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
		it.Event = new(VRFCoordinatorV2RequestFulfilled)
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
func (it *VRFCoordinatorV2RequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFCoordinatorV2RequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFCoordinatorV2RequestFulfilled represents a RequestFulfilled event raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2RequestFulfilled struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0xfe2e2d779dba245964d4e3ef9b994be63856fd568bf7d3ca9e224755cb1bd54d.
//
// Solidity: event RequestFulfilled(uint256 requestId, uint256[] randomWords)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterRequestFulfilled(opts *bind.FilterOpts) (*VRFCoordinatorV2RequestFulfilledIterator, error) {

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "RequestFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2RequestFulfilledIterator{contract: _VRFCoordinatorV2.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0xfe2e2d779dba245964d4e3ef9b994be63856fd568bf7d3ca9e224755cb1bd54d.
//
// Solidity: event RequestFulfilled(uint256 requestId, uint256[] randomWords)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2RequestFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "RequestFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFCoordinatorV2RequestFulfilled)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseRequestFulfilled(log types.Log) (*VRFCoordinatorV2RequestFulfilled, error) {
	event := new(VRFCoordinatorV2RequestFulfilled)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFCoordinatorV2RequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2RequestSentIterator struct {
	Event *VRFCoordinatorV2RequestSent // Event containing the contract specifics and raw log

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
func (it *VRFCoordinatorV2RequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2RequestSent)
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
		it.Event = new(VRFCoordinatorV2RequestSent)
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
func (it *VRFCoordinatorV2RequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFCoordinatorV2RequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFCoordinatorV2RequestSent represents a RequestSent event raised by the VRFCoordinatorV2 contract.
type VRFCoordinatorV2RequestSent struct {
	RequestId *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0x48b98ad7a8a8dbe21cc82bf98710ad4d2cdd949ccac393692e4d9a1722c162c7.
//
// Solidity: event RequestSent(uint256 requestId, address caller)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) FilterRequestSent(opts *bind.FilterOpts) (*VRFCoordinatorV2RequestSentIterator, error) {

	logs, sub, err := _VRFCoordinatorV2.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2RequestSentIterator{contract: _VRFCoordinatorV2.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0x48b98ad7a8a8dbe21cc82bf98710ad4d2cdd949ccac393692e4d9a1722c162c7.
//
// Solidity: event RequestSent(uint256 requestId, address caller)
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2RequestSent) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFCoordinatorV2RequestSent)
				if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RequestSent", log); err != nil {
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
func (_VRFCoordinatorV2 *VRFCoordinatorV2Filterer) ParseRequestSent(log types.Log) (*VRFCoordinatorV2RequestSent, error) {
	event := new(VRFCoordinatorV2RequestSent)
	if err := _VRFCoordinatorV2.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
