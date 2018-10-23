// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RedeemHandshakeABI is the input ABI used to generate the binding from.
const RedeemHandshakeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeems\",\"outputs\":[{\"name\":\"creator\",\"type\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"initRedeem\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rid\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"useRedeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"rid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__initRedeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"rid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__useRedeem\",\"type\":\"event\"}]"

// RedeemHandshake is an auto generated Go binding around an Ethereum contract.
type RedeemHandshake struct {
	RedeemHandshakeCaller     // Read-only binding to the contract
	RedeemHandshakeTransactor // Write-only binding to the contract
	RedeemHandshakeFilterer   // Log filterer for contract events
}

// RedeemHandshakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedeemHandshakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemHandshakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedeemHandshakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemHandshakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedeemHandshakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemHandshakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedeemHandshakeSession struct {
	Contract     *RedeemHandshake  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedeemHandshakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedeemHandshakeCallerSession struct {
	Contract *RedeemHandshakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RedeemHandshakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedeemHandshakeTransactorSession struct {
	Contract     *RedeemHandshakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RedeemHandshakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedeemHandshakeRaw struct {
	Contract *RedeemHandshake // Generic contract binding to access the raw methods on
}

// RedeemHandshakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedeemHandshakeCallerRaw struct {
	Contract *RedeemHandshakeCaller // Generic read-only contract binding to access the raw methods on
}

// RedeemHandshakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedeemHandshakeTransactorRaw struct {
	Contract *RedeemHandshakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedeemHandshake creates a new instance of RedeemHandshake, bound to a specific deployed contract.
func NewRedeemHandshake(address common.Address, backend bind.ContractBackend) (*RedeemHandshake, error) {
	contract, err := bindRedeemHandshake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RedeemHandshake{RedeemHandshakeCaller: RedeemHandshakeCaller{contract: contract}, RedeemHandshakeTransactor: RedeemHandshakeTransactor{contract: contract}, RedeemHandshakeFilterer: RedeemHandshakeFilterer{contract: contract}}, nil
}

// NewRedeemHandshakeCaller creates a new read-only instance of RedeemHandshake, bound to a specific deployed contract.
func NewRedeemHandshakeCaller(address common.Address, caller bind.ContractCaller) (*RedeemHandshakeCaller, error) {
	contract, err := bindRedeemHandshake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemHandshakeCaller{contract: contract}, nil
}

// NewRedeemHandshakeTransactor creates a new write-only instance of RedeemHandshake, bound to a specific deployed contract.
func NewRedeemHandshakeTransactor(address common.Address, transactor bind.ContractTransactor) (*RedeemHandshakeTransactor, error) {
	contract, err := bindRedeemHandshake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemHandshakeTransactor{contract: contract}, nil
}

// NewRedeemHandshakeFilterer creates a new log filterer instance of RedeemHandshake, bound to a specific deployed contract.
func NewRedeemHandshakeFilterer(address common.Address, filterer bind.ContractFilterer) (*RedeemHandshakeFilterer, error) {
	contract, err := bindRedeemHandshake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedeemHandshakeFilterer{contract: contract}, nil
}

// bindRedeemHandshake binds a generic wrapper to an already deployed contract.
func bindRedeemHandshake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RedeemHandshakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedeemHandshake *RedeemHandshakeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RedeemHandshake.Contract.RedeemHandshakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedeemHandshake *RedeemHandshakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemHandshake.Contract.RedeemHandshakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedeemHandshake *RedeemHandshakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedeemHandshake.Contract.RedeemHandshakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedeemHandshake *RedeemHandshakeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RedeemHandshake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedeemHandshake *RedeemHandshakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemHandshake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedeemHandshake *RedeemHandshakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedeemHandshake.Contract.contract.Transact(opts, method, params...)
}

// Redeems is a free data retrieval call binding the contract method 0x6b76f333.
//
// Solidity: function redeems( uint256) constant returns(creator address, stake uint256)
func (_RedeemHandshake *RedeemHandshakeCaller) Redeems(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Creator common.Address
	Stake   *big.Int
}, error) {
	ret := new(struct {
		Creator common.Address
		Stake   *big.Int
	})
	out := ret
	err := _RedeemHandshake.contract.Call(opts, out, "redeems", arg0)
	return *ret, err
}

// Redeems is a free data retrieval call binding the contract method 0x6b76f333.
//
// Solidity: function redeems( uint256) constant returns(creator address, stake uint256)
func (_RedeemHandshake *RedeemHandshakeSession) Redeems(arg0 *big.Int) (struct {
	Creator common.Address
	Stake   *big.Int
}, error) {
	return _RedeemHandshake.Contract.Redeems(&_RedeemHandshake.CallOpts, arg0)
}

// Redeems is a free data retrieval call binding the contract method 0x6b76f333.
//
// Solidity: function redeems( uint256) constant returns(creator address, stake uint256)
func (_RedeemHandshake *RedeemHandshakeCallerSession) Redeems(arg0 *big.Int) (struct {
	Creator common.Address
	Stake   *big.Int
}, error) {
	return _RedeemHandshake.Contract.Redeems(&_RedeemHandshake.CallOpts, arg0)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_RedeemHandshake *RedeemHandshakeCaller) Root(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RedeemHandshake.contract.Call(opts, out, "root")
	return *ret0, err
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_RedeemHandshake *RedeemHandshakeSession) Root() (common.Address, error) {
	return _RedeemHandshake.Contract.Root(&_RedeemHandshake.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_RedeemHandshake *RedeemHandshakeCallerSession) Root() (common.Address, error) {
	return _RedeemHandshake.Contract.Root(&_RedeemHandshake.CallOpts)
}

// InitRedeem is a paid mutator transaction binding the contract method 0x8848885e.
//
// Solidity: function initRedeem(amount uint256, fee uint256, offchain bytes32) returns()
func (_RedeemHandshake *RedeemHandshakeTransactor) InitRedeem(opts *bind.TransactOpts, amount *big.Int, fee *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _RedeemHandshake.contract.Transact(opts, "initRedeem", amount, fee, offchain)
}

// InitRedeem is a paid mutator transaction binding the contract method 0x8848885e.
//
// Solidity: function initRedeem(amount uint256, fee uint256, offchain bytes32) returns()
func (_RedeemHandshake *RedeemHandshakeSession) InitRedeem(amount *big.Int, fee *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _RedeemHandshake.Contract.InitRedeem(&_RedeemHandshake.TransactOpts, amount, fee, offchain)
}

// InitRedeem is a paid mutator transaction binding the contract method 0x8848885e.
//
// Solidity: function initRedeem(amount uint256, fee uint256, offchain bytes32) returns()
func (_RedeemHandshake *RedeemHandshakeTransactorSession) InitRedeem(amount *big.Int, fee *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _RedeemHandshake.Contract.InitRedeem(&_RedeemHandshake.TransactOpts, amount, fee, offchain)
}

// UseRedeem is a paid mutator transaction binding the contract method 0xba9ea2c6.
//
// Solidity: function useRedeem(rid uint256, amount uint256, receiver address, offchain bytes32) returns()
func (_RedeemHandshake *RedeemHandshakeTransactor) UseRedeem(opts *bind.TransactOpts, rid *big.Int, amount *big.Int, receiver common.Address, offchain [32]byte) (*types.Transaction, error) {
	return _RedeemHandshake.contract.Transact(opts, "useRedeem", rid, amount, receiver, offchain)
}

// UseRedeem is a paid mutator transaction binding the contract method 0xba9ea2c6.
//
// Solidity: function useRedeem(rid uint256, amount uint256, receiver address, offchain bytes32) returns()
func (_RedeemHandshake *RedeemHandshakeSession) UseRedeem(rid *big.Int, amount *big.Int, receiver common.Address, offchain [32]byte) (*types.Transaction, error) {
	return _RedeemHandshake.Contract.UseRedeem(&_RedeemHandshake.TransactOpts, rid, amount, receiver, offchain)
}

// UseRedeem is a paid mutator transaction binding the contract method 0xba9ea2c6.
//
// Solidity: function useRedeem(rid uint256, amount uint256, receiver address, offchain bytes32) returns()
func (_RedeemHandshake *RedeemHandshakeTransactorSession) UseRedeem(rid *big.Int, amount *big.Int, receiver common.Address, offchain [32]byte) (*types.Transaction, error) {
	return _RedeemHandshake.Contract.UseRedeem(&_RedeemHandshake.TransactOpts, rid, amount, receiver, offchain)
}

// RedeemHandshakeInitRedeemIterator is returned from FilterInitRedeem and is used to iterate over the raw logs and unpacked data for InitRedeem events raised by the RedeemHandshake contract.
type RedeemHandshakeInitRedeemIterator struct {
	Event *RedeemHandshakeInitRedeem // Event containing the contract specifics and raw log

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
func (it *RedeemHandshakeInitRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemHandshakeInitRedeem)
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
		it.Event = new(RedeemHandshakeInitRedeem)
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
func (it *RedeemHandshakeInitRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemHandshakeInitRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemHandshakeInitRedeem represents a InitRedeem event raised by the RedeemHandshake contract.
type RedeemHandshakeInitRedeem struct {
	Rid      *big.Int
	Fee      *big.Int
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitRedeem is a free log retrieval operation binding the contract event 0x8ce0189ea5ec87bced82b710181cbb27a8d753d726d9c470bae1fde88ff530af.
//
// Solidity: e __initRedeem(rid uint256, fee uint256, offchain bytes32)
func (_RedeemHandshake *RedeemHandshakeFilterer) FilterInitRedeem(opts *bind.FilterOpts) (*RedeemHandshakeInitRedeemIterator, error) {

	logs, sub, err := _RedeemHandshake.contract.FilterLogs(opts, "__initRedeem")
	if err != nil {
		return nil, err
	}
	return &RedeemHandshakeInitRedeemIterator{contract: _RedeemHandshake.contract, event: "__initRedeem", logs: logs, sub: sub}, nil
}

// WatchInitRedeem is a free log subscription operation binding the contract event 0x8ce0189ea5ec87bced82b710181cbb27a8d753d726d9c470bae1fde88ff530af.
//
// Solidity: e __initRedeem(rid uint256, fee uint256, offchain bytes32)
func (_RedeemHandshake *RedeemHandshakeFilterer) WatchInitRedeem(opts *bind.WatchOpts, sink chan<- *RedeemHandshakeInitRedeem) (event.Subscription, error) {

	logs, sub, err := _RedeemHandshake.contract.WatchLogs(opts, "__initRedeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemHandshakeInitRedeem)
				if err := _RedeemHandshake.contract.UnpackLog(event, "__initRedeem", log); err != nil {
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

// RedeemHandshakeUseRedeemIterator is returned from FilterUseRedeem and is used to iterate over the raw logs and unpacked data for UseRedeem events raised by the RedeemHandshake contract.
type RedeemHandshakeUseRedeemIterator struct {
	Event *RedeemHandshakeUseRedeem // Event containing the contract specifics and raw log

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
func (it *RedeemHandshakeUseRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemHandshakeUseRedeem)
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
		it.Event = new(RedeemHandshakeUseRedeem)
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
func (it *RedeemHandshakeUseRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemHandshakeUseRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemHandshakeUseRedeem represents a UseRedeem event raised by the RedeemHandshake contract.
type RedeemHandshakeUseRedeem struct {
	Rid      *big.Int
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUseRedeem is a free log retrieval operation binding the contract event 0x12491112acf9f0eec4456e4db4cd1e46a2256afa2181759ad1730e4003bda2b0.
//
// Solidity: e __useRedeem(rid uint256, offchain bytes32)
func (_RedeemHandshake *RedeemHandshakeFilterer) FilterUseRedeem(opts *bind.FilterOpts) (*RedeemHandshakeUseRedeemIterator, error) {

	logs, sub, err := _RedeemHandshake.contract.FilterLogs(opts, "__useRedeem")
	if err != nil {
		return nil, err
	}
	return &RedeemHandshakeUseRedeemIterator{contract: _RedeemHandshake.contract, event: "__useRedeem", logs: logs, sub: sub}, nil
}

// WatchUseRedeem is a free log subscription operation binding the contract event 0x12491112acf9f0eec4456e4db4cd1e46a2256afa2181759ad1730e4003bda2b0.
//
// Solidity: e __useRedeem(rid uint256, offchain bytes32)
func (_RedeemHandshake *RedeemHandshakeFilterer) WatchUseRedeem(opts *bind.WatchOpts, sink chan<- *RedeemHandshakeUseRedeem) (event.Subscription, error) {

	logs, sub, err := _RedeemHandshake.contract.WatchLogs(opts, "__useRedeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemHandshakeUseRedeem)
				if err := _RedeemHandshake.contract.UnpackLog(event, "__useRedeem", log); err != nil {
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
