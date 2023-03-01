// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package game_abi

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

// BidwarBindMetaData contains all meta data concerning the BidwarBind contract.
var BidwarBindMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_increment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_commissionPercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalTimestamp\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Timeout\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAdress\",\"type\":\"address\"}],\"name\":\"finalise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"newBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BidwarBindABI is the input ABI used to generate the binding from.
// Deprecated: Use BidwarBindMetaData.ABI instead.
var BidwarBindABI = BidwarBindMetaData.ABI

// BidwarBind is an auto generated Go binding around an Ethereum contract.
type BidwarBind struct {
	BidwarBindCaller     // Read-only binding to the contract
	BidwarBindTransactor // Write-only binding to the contract
	BidwarBindFilterer   // Log filterer for contract events
}

// BidwarBindCaller is an auto generated read-only Go binding around an Ethereum contract.
type BidwarBindCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BidwarBindTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BidwarBindTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BidwarBindFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BidwarBindFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BidwarBindSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BidwarBindSession struct {
	Contract     *BidwarBind       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BidwarBindCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BidwarBindCallerSession struct {
	Contract *BidwarBindCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BidwarBindTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BidwarBindTransactorSession struct {
	Contract     *BidwarBindTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BidwarBindRaw is an auto generated low-level Go binding around an Ethereum contract.
type BidwarBindRaw struct {
	Contract *BidwarBind // Generic contract binding to access the raw methods on
}

// BidwarBindCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BidwarBindCallerRaw struct {
	Contract *BidwarBindCaller // Generic read-only contract binding to access the raw methods on
}

// BidwarBindTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BidwarBindTransactorRaw struct {
	Contract *BidwarBindTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBidwarBind creates a new instance of BidwarBind, bound to a specific deployed contract.
func NewBidwarBind(address common.Address, backend bind.ContractBackend) (*BidwarBind, error) {
	contract, err := bindBidwarBind(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BidwarBind{BidwarBindCaller: BidwarBindCaller{contract: contract}, BidwarBindTransactor: BidwarBindTransactor{contract: contract}, BidwarBindFilterer: BidwarBindFilterer{contract: contract}}, nil
}

// NewBidwarBindCaller creates a new read-only instance of BidwarBind, bound to a specific deployed contract.
func NewBidwarBindCaller(address common.Address, caller bind.ContractCaller) (*BidwarBindCaller, error) {
	contract, err := bindBidwarBind(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BidwarBindCaller{contract: contract}, nil
}

// NewBidwarBindTransactor creates a new write-only instance of BidwarBind, bound to a specific deployed contract.
func NewBidwarBindTransactor(address common.Address, transactor bind.ContractTransactor) (*BidwarBindTransactor, error) {
	contract, err := bindBidwarBind(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BidwarBindTransactor{contract: contract}, nil
}

// NewBidwarBindFilterer creates a new log filterer instance of BidwarBind, bound to a specific deployed contract.
func NewBidwarBindFilterer(address common.Address, filterer bind.ContractFilterer) (*BidwarBindFilterer, error) {
	contract, err := bindBidwarBind(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BidwarBindFilterer{contract: contract}, nil
}

// bindBidwarBind binds a generic wrapper to an already deployed contract.
func bindBidwarBind(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BidwarBindABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BidwarBind *BidwarBindRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BidwarBind.Contract.BidwarBindCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BidwarBind *BidwarBindRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BidwarBind.Contract.BidwarBindTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BidwarBind *BidwarBindRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BidwarBind.Contract.BidwarBindTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BidwarBind *BidwarBindCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BidwarBind.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BidwarBind *BidwarBindTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BidwarBind.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BidwarBind *BidwarBindTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BidwarBind.Contract.contract.Transact(opts, method, params...)
}

// Finalise is a paid mutator transaction binding the contract method 0x330c3fe5.
//
// Solidity: function finalise(address tokenAdress) returns()
func (_BidwarBind *BidwarBindTransactor) Finalise(opts *bind.TransactOpts, tokenAdress common.Address) (*types.Transaction, error) {
	return _BidwarBind.contract.Transact(opts, "finalise", tokenAdress)
}

// Finalise is a paid mutator transaction binding the contract method 0x330c3fe5.
//
// Solidity: function finalise(address tokenAdress) returns()
func (_BidwarBind *BidwarBindSession) Finalise(tokenAdress common.Address) (*types.Transaction, error) {
	return _BidwarBind.Contract.Finalise(&_BidwarBind.TransactOpts, tokenAdress)
}

// Finalise is a paid mutator transaction binding the contract method 0x330c3fe5.
//
// Solidity: function finalise(address tokenAdress) returns()
func (_BidwarBind *BidwarBindTransactorSession) Finalise(tokenAdress common.Address) (*types.Transaction, error) {
	return _BidwarBind.Contract.Finalise(&_BidwarBind.TransactOpts, tokenAdress)
}

// NewBid is a paid mutator transaction binding the contract method 0xb7d1f8cf.
//
// Solidity: function newBid(uint256 amount, address tokenAddress) returns()
func (_BidwarBind *BidwarBindTransactor) NewBid(opts *bind.TransactOpts, amount *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _BidwarBind.contract.Transact(opts, "newBid", amount, tokenAddress)
}

// NewBid is a paid mutator transaction binding the contract method 0xb7d1f8cf.
//
// Solidity: function newBid(uint256 amount, address tokenAddress) returns()
func (_BidwarBind *BidwarBindSession) NewBid(amount *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _BidwarBind.Contract.NewBid(&_BidwarBind.TransactOpts, amount, tokenAddress)
}

// NewBid is a paid mutator transaction binding the contract method 0xb7d1f8cf.
//
// Solidity: function newBid(uint256 amount, address tokenAddress) returns()
func (_BidwarBind *BidwarBindTransactorSession) NewBid(amount *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _BidwarBind.Contract.NewBid(&_BidwarBind.TransactOpts, amount, tokenAddress)
}

// BidwarBindNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the BidwarBind contract.
type BidwarBindNewBidIterator struct {
	Event *BidwarBindNewBid // Event containing the contract specifics and raw log

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
func (it *BidwarBindNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidwarBindNewBid)
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
		it.Event = new(BidwarBindNewBid)
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
func (it *BidwarBindNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidwarBindNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidwarBindNewBid represents a NewBid event raised by the BidwarBind contract.
type BidwarBindNewBid struct {
	Bidder         common.Address
	Amount         *big.Int
	FinalTimestamp *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0x87f5b2fe112bf269d30fb8ca9dc0bde0afd0cc39258e13fafd75fe794795bf0e.
//
// Solidity: event NewBid(address bidder, uint256 amount, uint256 finalTimestamp)
func (_BidwarBind *BidwarBindFilterer) FilterNewBid(opts *bind.FilterOpts) (*BidwarBindNewBidIterator, error) {

	logs, sub, err := _BidwarBind.contract.FilterLogs(opts, "NewBid")
	if err != nil {
		return nil, err
	}
	return &BidwarBindNewBidIterator{contract: _BidwarBind.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0x87f5b2fe112bf269d30fb8ca9dc0bde0afd0cc39258e13fafd75fe794795bf0e.
//
// Solidity: event NewBid(address bidder, uint256 amount, uint256 finalTimestamp)
func (_BidwarBind *BidwarBindFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *BidwarBindNewBid) (event.Subscription, error) {

	logs, sub, err := _BidwarBind.contract.WatchLogs(opts, "NewBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidwarBindNewBid)
				if err := _BidwarBind.contract.UnpackLog(event, "NewBid", log); err != nil {
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

// ParseNewBid is a log parse operation binding the contract event 0x87f5b2fe112bf269d30fb8ca9dc0bde0afd0cc39258e13fafd75fe794795bf0e.
//
// Solidity: event NewBid(address bidder, uint256 amount, uint256 finalTimestamp)
func (_BidwarBind *BidwarBindFilterer) ParseNewBid(log types.Log) (*BidwarBindNewBid, error) {
	event := new(BidwarBindNewBid)
	if err := _BidwarBind.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BidwarBindTimeoutIterator is returned from FilterTimeout and is used to iterate over the raw logs and unpacked data for Timeout events raised by the BidwarBind contract.
type BidwarBindTimeoutIterator struct {
	Event *BidwarBindTimeout // Event containing the contract specifics and raw log

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
func (it *BidwarBindTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidwarBindTimeout)
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
		it.Event = new(BidwarBindTimeout)
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
func (it *BidwarBindTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidwarBindTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidwarBindTimeout represents a Timeout event raised by the BidwarBind contract.
type BidwarBindTimeout struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTimeout is a free log retrieval operation binding the contract event 0x48f41610c1ed8e474ea2df3e4c778c8e152b4b5d0f4f91a992a4901a94825661.
//
// Solidity: event Timeout(address bidder, uint256 amount)
func (_BidwarBind *BidwarBindFilterer) FilterTimeout(opts *bind.FilterOpts) (*BidwarBindTimeoutIterator, error) {

	logs, sub, err := _BidwarBind.contract.FilterLogs(opts, "Timeout")
	if err != nil {
		return nil, err
	}
	return &BidwarBindTimeoutIterator{contract: _BidwarBind.contract, event: "Timeout", logs: logs, sub: sub}, nil
}

// WatchTimeout is a free log subscription operation binding the contract event 0x48f41610c1ed8e474ea2df3e4c778c8e152b4b5d0f4f91a992a4901a94825661.
//
// Solidity: event Timeout(address bidder, uint256 amount)
func (_BidwarBind *BidwarBindFilterer) WatchTimeout(opts *bind.WatchOpts, sink chan<- *BidwarBindTimeout) (event.Subscription, error) {

	logs, sub, err := _BidwarBind.contract.WatchLogs(opts, "Timeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidwarBindTimeout)
				if err := _BidwarBind.contract.UnpackLog(event, "Timeout", log); err != nil {
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

// ParseTimeout is a log parse operation binding the contract event 0x48f41610c1ed8e474ea2df3e4c778c8e152b4b5d0f4f91a992a4901a94825661.
//
// Solidity: event Timeout(address bidder, uint256 amount)
func (_BidwarBind *BidwarBindFilterer) ParseTimeout(log types.Log) (*BidwarBindTimeout, error) {
	event := new(BidwarBindTimeout)
	if err := _BidwarBind.contract.UnpackLog(event, "Timeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
