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

// WalletFactoryMetaData contains all meta data concerning the WalletFactory contract.
var WalletFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"WalletMinted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"mintCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factoryGov\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"mintWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// WalletFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletFactoryMetaData.ABI instead.
var WalletFactoryABI = WalletFactoryMetaData.ABI

// WalletFactory is an auto generated Go binding around an Ethereum contract.
type WalletFactory struct {
	WalletFactoryCaller     // Read-only binding to the contract
	WalletFactoryTransactor // Write-only binding to the contract
	WalletFactoryFilterer   // Log filterer for contract events
}

// WalletFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletFactorySession struct {
	Contract     *WalletFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletFactoryCallerSession struct {
	Contract *WalletFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WalletFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletFactoryTransactorSession struct {
	Contract     *WalletFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WalletFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletFactoryRaw struct {
	Contract *WalletFactory // Generic contract binding to access the raw methods on
}

// WalletFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletFactoryCallerRaw struct {
	Contract *WalletFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// WalletFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletFactoryTransactorRaw struct {
	Contract *WalletFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletFactory creates a new instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactory(address common.Address, backend bind.ContractBackend) (*WalletFactory, error) {
	contract, err := bindWalletFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletFactory{WalletFactoryCaller: WalletFactoryCaller{contract: contract}, WalletFactoryTransactor: WalletFactoryTransactor{contract: contract}, WalletFactoryFilterer: WalletFactoryFilterer{contract: contract}}, nil
}

// NewWalletFactoryCaller creates a new read-only instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactoryCaller(address common.Address, caller bind.ContractCaller) (*WalletFactoryCaller, error) {
	contract, err := bindWalletFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryCaller{contract: contract}, nil
}

// NewWalletFactoryTransactor creates a new write-only instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletFactoryTransactor, error) {
	contract, err := bindWalletFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryTransactor{contract: contract}, nil
}

// NewWalletFactoryFilterer creates a new log filterer instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFactoryFilterer, error) {
	contract, err := bindWalletFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryFilterer{contract: contract}, nil
}

// bindWalletFactory binds a generic wrapper to an already deployed contract.
func bindWalletFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletFactory *WalletFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletFactory.Contract.WalletFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletFactory *WalletFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletFactory.Contract.WalletFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletFactory *WalletFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletFactory.Contract.WalletFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletFactory *WalletFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletFactory *WalletFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletFactory *WalletFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletFactory.Contract.contract.Transact(opts, method, params...)
}

// MintCost is a free data retrieval call binding the contract method 0xbdb4b848.
//
// Solidity: function mintCost() view returns(uint256)
func (_WalletFactory *WalletFactoryCaller) MintCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletFactory.contract.Call(opts, &out, "mintCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCost is a free data retrieval call binding the contract method 0xbdb4b848.
//
// Solidity: function mintCost() view returns(uint256)
func (_WalletFactory *WalletFactorySession) MintCost() (*big.Int, error) {
	return _WalletFactory.Contract.MintCost(&_WalletFactory.CallOpts)
}

// MintCost is a free data retrieval call binding the contract method 0xbdb4b848.
//
// Solidity: function mintCost() view returns(uint256)
func (_WalletFactory *WalletFactoryCallerSession) MintCost() (*big.Int, error) {
	return _WalletFactory.Contract.MintCost(&_WalletFactory.CallOpts)
}

// MintWallet is a paid mutator transaction binding the contract method 0x6ea0bab7.
//
// Solidity: function mintWallet(address _admin) payable returns(address)
func (_WalletFactory *WalletFactoryTransactor) MintWallet(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _WalletFactory.contract.Transact(opts, "mintWallet", _admin)
}

// MintWallet is a paid mutator transaction binding the contract method 0x6ea0bab7.
//
// Solidity: function mintWallet(address _admin) payable returns(address)
func (_WalletFactory *WalletFactorySession) MintWallet(_admin common.Address) (*types.Transaction, error) {
	return _WalletFactory.Contract.MintWallet(&_WalletFactory.TransactOpts, _admin)
}

// MintWallet is a paid mutator transaction binding the contract method 0x6ea0bab7.
//
// Solidity: function mintWallet(address _admin) payable returns(address)
func (_WalletFactory *WalletFactoryTransactorSession) MintWallet(_admin common.Address) (*types.Transaction, error) {
	return _WalletFactory.Contract.MintWallet(&_WalletFactory.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _factoryGov) returns()
func (_WalletFactory *WalletFactoryTransactor) SetAdmin(opts *bind.TransactOpts, _factoryGov common.Address) (*types.Transaction, error) {
	return _WalletFactory.contract.Transact(opts, "setAdmin", _factoryGov)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _factoryGov) returns()
func (_WalletFactory *WalletFactorySession) SetAdmin(_factoryGov common.Address) (*types.Transaction, error) {
	return _WalletFactory.Contract.SetAdmin(&_WalletFactory.TransactOpts, _factoryGov)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _factoryGov) returns()
func (_WalletFactory *WalletFactoryTransactorSession) SetAdmin(_factoryGov common.Address) (*types.Transaction, error) {
	return _WalletFactory.Contract.SetAdmin(&_WalletFactory.TransactOpts, _factoryGov)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_WalletFactory *WalletFactoryTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletFactory.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_WalletFactory *WalletFactorySession) Withdraw() (*types.Transaction, error) {
	return _WalletFactory.Contract.Withdraw(&_WalletFactory.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_WalletFactory *WalletFactoryTransactorSession) Withdraw() (*types.Transaction, error) {
	return _WalletFactory.Contract.Withdraw(&_WalletFactory.TransactOpts)
}

// WalletFactoryWalletMintedIterator is returned from FilterWalletMinted and is used to iterate over the raw logs and unpacked data for WalletMinted events raised by the WalletFactory contract.
type WalletFactoryWalletMintedIterator struct {
	Event *WalletFactoryWalletMinted // Event containing the contract specifics and raw log

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
func (it *WalletFactoryWalletMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletFactoryWalletMinted)
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
		it.Event = new(WalletFactoryWalletMinted)
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
func (it *WalletFactoryWalletMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletFactoryWalletMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletFactoryWalletMinted represents a WalletMinted event raised by the WalletFactory contract.
type WalletFactoryWalletMinted struct {
	Wallet common.Address
	Admin  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWalletMinted is a free log retrieval operation binding the contract event 0x73e8e01c6120c8daad5aa83307027ad4145ebd31d1a4e04c41404a0e2f3102ac.
//
// Solidity: event WalletMinted(address wallet, address admin)
func (_WalletFactory *WalletFactoryFilterer) FilterWalletMinted(opts *bind.FilterOpts) (*WalletFactoryWalletMintedIterator, error) {

	logs, sub, err := _WalletFactory.contract.FilterLogs(opts, "WalletMinted")
	if err != nil {
		return nil, err
	}
	return &WalletFactoryWalletMintedIterator{contract: _WalletFactory.contract, event: "WalletMinted", logs: logs, sub: sub}, nil
}

// WatchWalletMinted is a free log subscription operation binding the contract event 0x73e8e01c6120c8daad5aa83307027ad4145ebd31d1a4e04c41404a0e2f3102ac.
//
// Solidity: event WalletMinted(address wallet, address admin)
func (_WalletFactory *WalletFactoryFilterer) WatchWalletMinted(opts *bind.WatchOpts, sink chan<- *WalletFactoryWalletMinted) (event.Subscription, error) {

	logs, sub, err := _WalletFactory.contract.WatchLogs(opts, "WalletMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletFactoryWalletMinted)
				if err := _WalletFactory.contract.UnpackLog(event, "WalletMinted", log); err != nil {
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

// ParseWalletMinted is a log parse operation binding the contract event 0x73e8e01c6120c8daad5aa83307027ad4145ebd31d1a4e04c41404a0e2f3102ac.
//
// Solidity: event WalletMinted(address wallet, address admin)
func (_WalletFactory *WalletFactoryFilterer) ParseWalletMinted(log types.Log) (*WalletFactoryWalletMinted, error) {
	event := new(WalletFactoryWalletMinted)
	if err := _WalletFactory.contract.UnpackLog(event, "WalletMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
