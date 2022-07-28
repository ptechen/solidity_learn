// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// EventMetaData contains all meta data concerning the Event contract.
var EventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"IndexedLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Message\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"example\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"54353f2f": "example()",
		"de6f24bb": "sendMessage(address,string)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610226806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806354353f2f1461003b578063de6f24bb14610045575b600080fd5b610043610058565b005b610043610053366004610130565b6100de565b6040805181815260038183015262666f6f60e81b60608201526104d2602082015290517fdd970dd9b5bfe707922155b058a407655cb18288b807e2216442bca8ad83d6b59181900360800190a1604051610315815233907f919a4e0615aef45fd56105241973bd25f287f1040399e7c9f97fa2ca1c2b51749060200160405180910390a2565b826001600160a01b0316336001600160a01b03167fb3dbe9e9894ca2c11cb6c80bd0b0bccb9f5b41d612dbeeda0d5474de40b874fe84846040516101239291906101c1565b60405180910390a3505050565b60008060006040848603121561014557600080fd5b83356001600160a01b038116811461015c57600080fd5b9250602084013567ffffffffffffffff8082111561017957600080fd5b818601915086601f83011261018d57600080fd5b81358181111561019c57600080fd5b8760208285010111156101ae57600080fd5b6020830194508093505050509250925092565b60208152816020820152818360408301376000818301604090810191909152601f909201601f1916010191905056fea26469706673582212203b52cc6f3a33da2f54dfac48663a62001850bed6283dfe532634f446bed4a55364736f6c634300080f0033",
}

// EventABI is the input ABI used to generate the binding from.
// Deprecated: Use EventMetaData.ABI instead.
var EventABI = EventMetaData.ABI

// Deprecated: Use EventMetaData.Sigs instead.
// EventFuncSigs maps the 4-byte function signature to its string representation.
var EventFuncSigs = EventMetaData.Sigs

// EventBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventMetaData.Bin instead.
var EventBin = EventMetaData.Bin

// DeployEvent deploys a new Ethereum contract, binding an instance of Event to it.
func DeployEvent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Event, error) {
	parsed, err := EventMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Event{EventCaller: EventCaller{contract: contract}, EventTransactor: EventTransactor{contract: contract}, EventFilterer: EventFilterer{contract: contract}}, nil
}

// Event is an auto generated Go binding around an Ethereum contract.
type Event struct {
	EventCaller     // Read-only binding to the contract
	EventTransactor // Write-only binding to the contract
	EventFilterer   // Log filterer for contract events
}

// EventCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventSession struct {
	Contract     *Event            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventCallerSession struct {
	Contract *EventCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventTransactorSession struct {
	Contract     *EventTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventRaw struct {
	Contract *Event // Generic contract binding to access the raw methods on
}

// EventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventCallerRaw struct {
	Contract *EventCaller // Generic read-only contract binding to access the raw methods on
}

// EventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventTransactorRaw struct {
	Contract *EventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvent creates a new instance of Event, bound to a specific deployed contract.
func NewEvent(address common.Address, backend bind.ContractBackend) (*Event, error) {
	contract, err := bindEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Event{EventCaller: EventCaller{contract: contract}, EventTransactor: EventTransactor{contract: contract}, EventFilterer: EventFilterer{contract: contract}}, nil
}

// NewEventCaller creates a new read-only instance of Event, bound to a specific deployed contract.
func NewEventCaller(address common.Address, caller bind.ContractCaller) (*EventCaller, error) {
	contract, err := bindEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventCaller{contract: contract}, nil
}

// NewEventTransactor creates a new write-only instance of Event, bound to a specific deployed contract.
func NewEventTransactor(address common.Address, transactor bind.ContractTransactor) (*EventTransactor, error) {
	contract, err := bindEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventTransactor{contract: contract}, nil
}

// NewEventFilterer creates a new log filterer instance of Event, bound to a specific deployed contract.
func NewEventFilterer(address common.Address, filterer bind.ContractFilterer) (*EventFilterer, error) {
	contract, err := bindEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventFilterer{contract: contract}, nil
}

// bindEvent binds a generic wrapper to an already deployed contract.
func bindEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Event *EventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Event.Contract.EventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Event *EventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Event.Contract.EventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Event *EventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Event.Contract.EventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Event *EventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Event.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Event *EventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Event.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Event *EventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Event.Contract.contract.Transact(opts, method, params...)
}

// Example is a paid mutator transaction binding the contract method 0x54353f2f.
//
// Solidity: function example() returns()
func (_Event *EventTransactor) Example(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Event.contract.Transact(opts, "example")
}

// Example is a paid mutator transaction binding the contract method 0x54353f2f.
//
// Solidity: function example() returns()
func (_Event *EventSession) Example() (*types.Transaction, error) {
	return _Event.Contract.Example(&_Event.TransactOpts)
}

// Example is a paid mutator transaction binding the contract method 0x54353f2f.
//
// Solidity: function example() returns()
func (_Event *EventTransactorSession) Example() (*types.Transaction, error) {
	return _Event.Contract.Example(&_Event.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0xde6f24bb.
//
// Solidity: function sendMessage(address _to, string message) returns()
func (_Event *EventTransactor) SendMessage(opts *bind.TransactOpts, _to common.Address, message string) (*types.Transaction, error) {
	return _Event.contract.Transact(opts, "sendMessage", _to, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xde6f24bb.
//
// Solidity: function sendMessage(address _to, string message) returns()
func (_Event *EventSession) SendMessage(_to common.Address, message string) (*types.Transaction, error) {
	return _Event.Contract.SendMessage(&_Event.TransactOpts, _to, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xde6f24bb.
//
// Solidity: function sendMessage(address _to, string message) returns()
func (_Event *EventTransactorSession) SendMessage(_to common.Address, message string) (*types.Transaction, error) {
	return _Event.Contract.SendMessage(&_Event.TransactOpts, _to, message)
}

// EventIndexedLogIterator is returned from FilterIndexedLog and is used to iterate over the raw logs and unpacked data for IndexedLog events raised by the Event contract.
type EventIndexedLogIterator struct {
	Event *EventIndexedLog // Event containing the contract specifics and raw log

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
func (it *EventIndexedLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventIndexedLog)
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
		it.Event = new(EventIndexedLog)
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
func (it *EventIndexedLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventIndexedLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventIndexedLog represents a IndexedLog event raised by the Event contract.
type EventIndexedLog struct {
	Sender common.Address
	Val    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIndexedLog is a free log retrieval operation binding the contract event 0x919a4e0615aef45fd56105241973bd25f287f1040399e7c9f97fa2ca1c2b5174.
//
// Solidity: event IndexedLog(address indexed sender, uint256 val)
func (_Event *EventFilterer) FilterIndexedLog(opts *bind.FilterOpts, sender []common.Address) (*EventIndexedLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "IndexedLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &EventIndexedLogIterator{contract: _Event.contract, event: "IndexedLog", logs: logs, sub: sub}, nil
}

// WatchIndexedLog is a free log subscription operation binding the contract event 0x919a4e0615aef45fd56105241973bd25f287f1040399e7c9f97fa2ca1c2b5174.
//
// Solidity: event IndexedLog(address indexed sender, uint256 val)
func (_Event *EventFilterer) WatchIndexedLog(opts *bind.WatchOpts, sink chan<- *EventIndexedLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "IndexedLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventIndexedLog)
				if err := _Event.contract.UnpackLog(event, "IndexedLog", log); err != nil {
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

// ParseIndexedLog is a log parse operation binding the contract event 0x919a4e0615aef45fd56105241973bd25f287f1040399e7c9f97fa2ca1c2b5174.
//
// Solidity: event IndexedLog(address indexed sender, uint256 val)
func (_Event *EventFilterer) ParseIndexedLog(log types.Log) (*EventIndexedLog, error) {
	event := new(EventIndexedLog)
	if err := _Event.contract.UnpackLog(event, "IndexedLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the Event contract.
type EventLogIterator struct {
	Event *EventLog // Event containing the contract specifics and raw log

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
func (it *EventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventLog)
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
		it.Event = new(EventLog)
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
func (it *EventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventLog represents a Log event raised by the Event contract.
type EventLog struct {
	Message string
	Val     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0xdd970dd9b5bfe707922155b058a407655cb18288b807e2216442bca8ad83d6b5.
//
// Solidity: event Log(string message, uint256 val)
func (_Event *EventFilterer) FilterLog(opts *bind.FilterOpts) (*EventLogIterator, error) {

	logs, sub, err := _Event.contract.FilterLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return &EventLogIterator{contract: _Event.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0xdd970dd9b5bfe707922155b058a407655cb18288b807e2216442bca8ad83d6b5.
//
// Solidity: event Log(string message, uint256 val)
func (_Event *EventFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *EventLog) (event.Subscription, error) {

	logs, sub, err := _Event.contract.WatchLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventLog)
				if err := _Event.contract.UnpackLog(event, "Log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0xdd970dd9b5bfe707922155b058a407655cb18288b807e2216442bca8ad83d6b5.
//
// Solidity: event Log(string message, uint256 val)
func (_Event *EventFilterer) ParseLog(log types.Log) (*EventLog, error) {
	event := new(EventLog)
	if err := _Event.contract.UnpackLog(event, "Log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventMessageIterator is returned from FilterMessage and is used to iterate over the raw logs and unpacked data for Message events raised by the Event contract.
type EventMessageIterator struct {
	Event *EventMessage // Event containing the contract specifics and raw log

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
func (it *EventMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventMessage)
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
		it.Event = new(EventMessage)
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
func (it *EventMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventMessage represents a Message event raised by the Event contract.
type EventMessage struct {
	From    common.Address
	To      common.Address
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMessage is a free log retrieval operation binding the contract event 0xb3dbe9e9894ca2c11cb6c80bd0b0bccb9f5b41d612dbeeda0d5474de40b874fe.
//
// Solidity: event Message(address indexed _from, address indexed _to, string message)
func (_Event *EventFilterer) FilterMessage(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*EventMessageIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "Message", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &EventMessageIterator{contract: _Event.contract, event: "Message", logs: logs, sub: sub}, nil
}

// WatchMessage is a free log subscription operation binding the contract event 0xb3dbe9e9894ca2c11cb6c80bd0b0bccb9f5b41d612dbeeda0d5474de40b874fe.
//
// Solidity: event Message(address indexed _from, address indexed _to, string message)
func (_Event *EventFilterer) WatchMessage(opts *bind.WatchOpts, sink chan<- *EventMessage, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "Message", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventMessage)
				if err := _Event.contract.UnpackLog(event, "Message", log); err != nil {
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

// ParseMessage is a log parse operation binding the contract event 0xb3dbe9e9894ca2c11cb6c80bd0b0bccb9f5b41d612dbeeda0d5474de40b874fe.
//
// Solidity: event Message(address indexed _from, address indexed _to, string message)
func (_Event *EventFilterer) ParseMessage(log types.Log) (*EventMessage, error) {
	event := new(EventMessage)
	if err := _Event.contract.UnpackLog(event, "Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
