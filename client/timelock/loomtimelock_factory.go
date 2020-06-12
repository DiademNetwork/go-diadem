// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package timelock

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// diademTimelockFactoryABI is the input ABI used to generate the binding from.
const diademTimelockFactoryABI = "[{\"inputs\":[{\"name\":\"_diadem\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"validatorEthAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timelockContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"validatorName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"validatorPublicKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_duration\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_releaseTime\",\"type\":\"uint256\"}],\"name\":\"diademTimeLockCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"validatorEthAddress\",\"type\":\"address\"},{\"name\":\"validatorName\",\"type\":\"string\"},{\"name\":\"validatorPublicKey\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"deployTimeLock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// diademTimelockFactory is an auto generated Go binding around an Ethereum contract.
type diademTimelockFactory struct {
	diademTimelockFactoryCaller     // Read-only binding to the contract
	diademTimelockFactoryTransactor // Write-only binding to the contract
	diademTimelockFactoryFilterer   // Log filterer for contract events
}

// diademTimelockFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type diademTimelockFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// diademTimelockFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type diademTimelockFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// diademTimelockFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type diademTimelockFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// diademTimelockFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type diademTimelockFactorySession struct {
	Contract     *diademTimelockFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// diademTimelockFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type diademTimelockFactoryCallerSession struct {
	Contract *diademTimelockFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// diademTimelockFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type diademTimelockFactoryTransactorSession struct {
	Contract     *diademTimelockFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// diademTimelockFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type diademTimelockFactoryRaw struct {
	Contract *diademTimelockFactory // Generic contract binding to access the raw methods on
}

// diademTimelockFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type diademTimelockFactoryCallerRaw struct {
	Contract *diademTimelockFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// diademTimelockFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type diademTimelockFactoryTransactorRaw struct {
	Contract *diademTimelockFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewdiademTimelockFactory creates a new instance of diademTimelockFactory, bound to a specific deployed contract.
func NewdiademTimelockFactory(address common.Address, backend bind.ContractBackend) (*diademTimelockFactory, error) {
	contract, err := binddiademTimelockFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &diademTimelockFactory{diademTimelockFactoryCaller: diademTimelockFactoryCaller{contract: contract}, diademTimelockFactoryTransactor: diademTimelockFactoryTransactor{contract: contract}, diademTimelockFactoryFilterer: diademTimelockFactoryFilterer{contract: contract}}, nil
}

// NewdiademTimelockFactoryCaller creates a new read-only instance of diademTimelockFactory, bound to a specific deployed contract.
func NewdiademTimelockFactoryCaller(address common.Address, caller bind.ContractCaller) (*diademTimelockFactoryCaller, error) {
	contract, err := binddiademTimelockFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &diademTimelockFactoryCaller{contract: contract}, nil
}

// NewdiademTimelockFactoryTransactor creates a new write-only instance of diademTimelockFactory, bound to a specific deployed contract.
func NewdiademTimelockFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*diademTimelockFactoryTransactor, error) {
	contract, err := binddiademTimelockFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &diademTimelockFactoryTransactor{contract: contract}, nil
}

// NewdiademTimelockFactoryFilterer creates a new log filterer instance of diademTimelockFactory, bound to a specific deployed contract.
func NewdiademTimelockFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*diademTimelockFactoryFilterer, error) {
	contract, err := binddiademTimelockFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &diademTimelockFactoryFilterer{contract: contract}, nil
}

// binddiademTimelockFactory binds a generic wrapper to an already deployed contract.
func binddiademTimelockFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(diademTimelockFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_diademTimelockFactory *diademTimelockFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _diademTimelockFactory.Contract.diademTimelockFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_diademTimelockFactory *diademTimelockFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _diademTimelockFactory.Contract.diademTimelockFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_diademTimelockFactory *diademTimelockFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _diademTimelockFactory.Contract.diademTimelockFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_diademTimelockFactory *diademTimelockFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _diademTimelockFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_diademTimelockFactory *diademTimelockFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _diademTimelockFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_diademTimelockFactory *diademTimelockFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _diademTimelockFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployTimeLock is a paid mutator transaction binding the contract method 0xee66ca2b.
//
// Solidity: function deployTimeLock(validatorEthAddress address, validatorName string, validatorPublicKey string, amount uint256, duration uint256) returns()
func (_diademTimelockFactory *diademTimelockFactoryTransactor) DeployTimeLock(opts *bind.TransactOpts, validatorEthAddress common.Address, validatorName string, validatorPublicKey string, amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _diademTimelockFactory.contract.Transact(opts, "deployTimeLock", validatorEthAddress, validatorName, validatorPublicKey, amount, duration)
}

// DeployTimeLock is a paid mutator transaction binding the contract method 0xee66ca2b.
//
// Solidity: function deployTimeLock(validatorEthAddress address, validatorName string, validatorPublicKey string, amount uint256, duration uint256) returns()
func (_diademTimelockFactory *diademTimelockFactorySession) DeployTimeLock(validatorEthAddress common.Address, validatorName string, validatorPublicKey string, amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _diademTimelockFactory.Contract.DeployTimeLock(&_diademTimelockFactory.TransactOpts, validatorEthAddress, validatorName, validatorPublicKey, amount, duration)
}

// DeployTimeLock is a paid mutator transaction binding the contract method 0xee66ca2b.
//
// Solidity: function deployTimeLock(validatorEthAddress address, validatorName string, validatorPublicKey string, amount uint256, duration uint256) returns()
func (_diademTimelockFactory *diademTimelockFactoryTransactorSession) DeployTimeLock(validatorEthAddress common.Address, validatorName string, validatorPublicKey string, amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _diademTimelockFactory.Contract.DeployTimeLock(&_diademTimelockFactory.TransactOpts, validatorEthAddress, validatorName, validatorPublicKey, amount, duration)
}

// diademTimelockFactorydiademTimeLockCreatedIterator is returned from FilterdiademTimeLockCreated and is used to iterate over the raw logs and unpacked data for diademTimeLockCreated events raised by the diademTimelockFactory contract.
type diademTimelockFactorydiademTimeLockCreatedIterator struct {
	Event *diademTimelockFactorydiademTimeLockCreated // Event containing the contract specifics and raw log

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
func (it *diademTimelockFactorydiademTimeLockCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(diademTimelockFactorydiademTimeLockCreated)
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
		it.Event = new(diademTimelockFactorydiademTimeLockCreated)
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
func (it *diademTimelockFactorydiademTimeLockCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *diademTimelockFactorydiademTimeLockCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// diademTimelockFactorydiademTimeLockCreated represents a diademTimeLockCreated event raised by the diademTimelockFactory contract.
type diademTimelockFactorydiademTimeLockCreated struct {
	ValidatorEthAddress     common.Address
	TimelockContractAddress common.Address
	ValidatorName           string
	ValidatorPublicKey      string
	Amount                  *big.Int
	Duration                *big.Int
	ReleaseTime             *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterdiademTimeLockCreated is a free log retrieval operation binding the contract event 0x920eb344f07704601d8be5e215805a8f8bf98fd7b7c0b73d48a0f61ac61bffd6.
//
// Solidity: e diademTimeLockCreated(validatorEthAddress address, timelockContractAddress address, validatorName string, validatorPublicKey string, _amount uint256, _duration uint256, _releaseTime uint256)
func (_diademTimelockFactory *diademTimelockFactoryFilterer) FilterdiademTimeLockCreated(opts *bind.FilterOpts) (*diademTimelockFactorydiademTimeLockCreatedIterator, error) {

	logs, sub, err := _diademTimelockFactory.contract.FilterLogs(opts, "diademTimeLockCreated")
	if err != nil {
		return nil, err
	}
	return &diademTimelockFactorydiademTimeLockCreatedIterator{contract: _diademTimelockFactory.contract, event: "diademTimeLockCreated", logs: logs, sub: sub}, nil
}

// WatchdiademTimeLockCreated is a free log subscription operation binding the contract event 0x920eb344f07704601d8be5e215805a8f8bf98fd7b7c0b73d48a0f61ac61bffd6.
//
// Solidity: e diademTimeLockCreated(validatorEthAddress address, timelockContractAddress address, validatorName string, validatorPublicKey string, _amount uint256, _duration uint256, _releaseTime uint256)
func (_diademTimelockFactory *diademTimelockFactoryFilterer) WatchdiademTimeLockCreated(opts *bind.WatchOpts, sink chan<- *diademTimelockFactorydiademTimeLockCreated) (event.Subscription, error) {

	logs, sub, err := _diademTimelockFactory.contract.WatchLogs(opts, "diademTimeLockCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(diademTimelockFactorydiademTimeLockCreated)
				if err := _diademTimelockFactory.contract.UnpackLog(event, "diademTimeLockCreated", log); err != nil {
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
