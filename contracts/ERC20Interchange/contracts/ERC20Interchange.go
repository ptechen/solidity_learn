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

// ERC20InterchangeInterchange is an auto generated low-level Go binding around an user-defined struct.
type ERC20InterchangeInterchange struct {
	FromAddress     common.Address
	ToAddress       common.Address
	FromAmount      *big.Int
	ToAmount        *big.Int
	InterchangeType uint8
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208f1bea9bf2c2765dc370ef4e3eaa31dbbc2fd810ef61002aad9036c00a98cf4764736f6c634300080f0033",
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressUpgradeableMetaData.ABI instead.
var AddressUpgradeableABI = AddressUpgradeableMetaData.ABI

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressUpgradeableMetaData.Bin instead.
var AddressUpgradeableBin = AddressUpgradeableMetaData.Bin

// DeployAddressUpgradeable deploys a new Ethereum contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around an Ethereum contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// CountersUpgradeableMetaData contains all meta data concerning the CountersUpgradeable contract.
var CountersUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122084787f7379b7cc43b2f3391a11166983727e4a4def592ff0308b46c9804ad9ef64736f6c634300080f0033",
}

// CountersUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use CountersUpgradeableMetaData.ABI instead.
var CountersUpgradeableABI = CountersUpgradeableMetaData.ABI

// CountersUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CountersUpgradeableMetaData.Bin instead.
var CountersUpgradeableBin = CountersUpgradeableMetaData.Bin

// DeployCountersUpgradeable deploys a new Ethereum contract, binding an instance of CountersUpgradeable to it.
func DeployCountersUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CountersUpgradeable, error) {
	parsed, err := CountersUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CountersUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CountersUpgradeable{CountersUpgradeableCaller: CountersUpgradeableCaller{contract: contract}, CountersUpgradeableTransactor: CountersUpgradeableTransactor{contract: contract}, CountersUpgradeableFilterer: CountersUpgradeableFilterer{contract: contract}}, nil
}

// CountersUpgradeable is an auto generated Go binding around an Ethereum contract.
type CountersUpgradeable struct {
	CountersUpgradeableCaller     // Read-only binding to the contract
	CountersUpgradeableTransactor // Write-only binding to the contract
	CountersUpgradeableFilterer   // Log filterer for contract events
}

// CountersUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountersUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountersUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountersUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountersUpgradeableSession struct {
	Contract     *CountersUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CountersUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountersUpgradeableCallerSession struct {
	Contract *CountersUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CountersUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountersUpgradeableTransactorSession struct {
	Contract     *CountersUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CountersUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountersUpgradeableRaw struct {
	Contract *CountersUpgradeable // Generic contract binding to access the raw methods on
}

// CountersUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountersUpgradeableCallerRaw struct {
	Contract *CountersUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// CountersUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountersUpgradeableTransactorRaw struct {
	Contract *CountersUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCountersUpgradeable creates a new instance of CountersUpgradeable, bound to a specific deployed contract.
func NewCountersUpgradeable(address common.Address, backend bind.ContractBackend) (*CountersUpgradeable, error) {
	contract, err := bindCountersUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CountersUpgradeable{CountersUpgradeableCaller: CountersUpgradeableCaller{contract: contract}, CountersUpgradeableTransactor: CountersUpgradeableTransactor{contract: contract}, CountersUpgradeableFilterer: CountersUpgradeableFilterer{contract: contract}}, nil
}

// NewCountersUpgradeableCaller creates a new read-only instance of CountersUpgradeable, bound to a specific deployed contract.
func NewCountersUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*CountersUpgradeableCaller, error) {
	contract, err := bindCountersUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountersUpgradeableCaller{contract: contract}, nil
}

// NewCountersUpgradeableTransactor creates a new write-only instance of CountersUpgradeable, bound to a specific deployed contract.
func NewCountersUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*CountersUpgradeableTransactor, error) {
	contract, err := bindCountersUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountersUpgradeableTransactor{contract: contract}, nil
}

// NewCountersUpgradeableFilterer creates a new log filterer instance of CountersUpgradeable, bound to a specific deployed contract.
func NewCountersUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*CountersUpgradeableFilterer, error) {
	contract, err := bindCountersUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountersUpgradeableFilterer{contract: contract}, nil
}

// bindCountersUpgradeable binds a generic wrapper to an already deployed contract.
func bindCountersUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CountersUpgradeable *CountersUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CountersUpgradeable.Contract.CountersUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CountersUpgradeable *CountersUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CountersUpgradeable.Contract.CountersUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CountersUpgradeable *CountersUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CountersUpgradeable.Contract.CountersUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CountersUpgradeable *CountersUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CountersUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CountersUpgradeable *CountersUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CountersUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CountersUpgradeable *CountersUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CountersUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ERC20InterchangeMetaData contains all meta data concerning the ERC20Interchange contract.
var ERC20InterchangeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"interchangeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"from_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"interchange_type\",\"type\":\"uint8\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interchangeId\",\"type\":\"uint256\"}],\"name\":\"InterchangeSuccess\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"from_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"interchange_type\",\"type\":\"uint8\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"interchangeId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fromList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token_one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token_two\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interchangeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"interchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interchangeId\",\"type\":\"uint256\"}],\"name\":\"searchByInterchangeId\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"from_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"interchange_type\",\"type\":\"uint8\"}],\"internalType\":\"structERC20Interchange.Interchange\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token_one\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token_two\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8caa5614": "create(address,uint256,uint256,uint8)",
		"a021857b": "fromList()",
		"485cc955": "initialize(address,address)",
		"9e6e801f": "interchange(uint256,uint256)",
		"6db92b7e": "searchByInterchangeId(uint256)",
		"71645971": "toList()",
		"c9caf28c": "token_one()",
		"34f97823": "token_two()",
	},
	Bin: "0x608060405234801561001057600080fd5b50611185806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638caa56141161005b5780638caa5614146101cc5780639e6e801f146101ed578063a021857b14610200578063c9caf28c1461020857600080fd5b806334f978231461008d578063485cc955146100bd5780636db92b7e146100d257806371645971146101b7575b600080fd5b6003546100a0906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100d06100cb366004610ed3565b61021b565b005b6101696100e0366004610f06565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506000908152600a6020908152604091829020825160a08101845281546001600160a01b039081168252600183015416928101929092526002810154928201929092526003820154606082015260049091015460ff16608082015290565b6040516100b4919081516001600160a01b03908116825260208084015190911690820152604080830151908201526060808301519082015260809182015160ff169181019190915260a00190565b6101bf61035c565b6040516100b49190610f1f565b6101df6101da366004610f63565b610423565b6040519081526020016100b4565b6100d06101fb366004610fb1565b6107a6565b6101bf610c5b565b6002546100a0906001600160a01b031681565b600054610100900460ff161580801561023b5750600054600160ff909116105b806102555750303b158015610255575060005460ff166001145b6102bd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156102e0576000805461ff0019166101001790555b600280546001600160a01b038086166001600160a01b03199283161790925560038054928516929091169190911790558015610357576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b505050565b336000908152600760205260408120546060919067ffffffffffffffff81111561038857610388610fd3565b6040519080825280602002602001820160405280156103b1578160200160208202803683370190505b50905060005b3360009081526007602052604090205481101561041d57336000908152600560209081526040808320848452909152902054825181908490849081106103ff576103ff610fe9565b6020908102919091010152508061041581611015565b9150506103b7565b50919050565b600060028260ff16106104715760405162461bcd60e51b815260206004820152601660248201527534b73a32b931b430b733b2afba3cb8329032b93937b960511b60448201526064016102b4565b8160ff1660000361050a576002546040516370a0823160e01b815233600482015285916001600160a01b0316906370a0823190602401602060405180830381865afa1580156104c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e8919061102e565b116105055760405162461bcd60e51b81526004016102b490611047565b610593565b6003546040516370a0823160e01b815233600482015285916001600160a01b0316906370a0823190602401602060405180830381865afa158015610552573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610576919061102e565b116105935760405162461bcd60e51b81526004016102b49061108d565b6105a1600180546001019055565b6001546040805160a0810182523381526001600160a01b0388811660208084019182528385018a8152606085018a815260ff808b166080880181815260008b8152600a909652988520975188549088166001600160a01b0319918216178955955160018901805491909816961695909517909555905160028601555160038501559351600490930180549390921660ff19909316929092179055919250036106c0576002546040516323b872dd60e01b81526001600160a01b03909116906323b872dd90610677903390309089906004016110d3565b6020604051808303816000875af1158015610696573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ba91906110f7565b50610739565b6003546040516323b872dd60e01b81526001600160a01b03909116906323b872dd906106f4903390309089906004016110d3565b6020604051808303816000875af1158015610713573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073791906110f7565b505b6107438186610d1c565b604080513381526001600160a01b03871660208201529081018590526060810184905260ff8316608082015281907fad6bba013d4dd6f7b6edc281a10f8b9003c42f19c04d7b88bde202abed854b959060a00160405180910390a2949350505050565b6001548211156107f85760405162461bcd60e51b815260206004820152601760248201527f696e7465726368616e67654964206e6f7420657869737400000000000000000060448201526064016102b4565b6000828152600a6020526040902080546001600160a01b031661085d5760405162461bcd60e51b815260206004820152601760248201527f696e7465726368616e6765496420636f6d706c6574656400000000000000000060448201526064016102b4565b60018101546001600160a01b031633146108af5760405162461bcd60e51b815260206004820152601360248201527234b73a32b931b430b733b2a4b21032b93937b960691b60448201526064016102b4565b80600301548210156108fc5760405162461bcd60e51b8152602060048201526016602482015275185b5bdd5b9d080f08195e1c1958dd17d85b5bdd5b9d60521b60448201526064016102b4565b600481015460ff16600003610a92576003546040516370a0823160e01b815233600482015283916001600160a01b0316906370a0823190602401602060405180830381865afa158015610953573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610977919061102e565b10156109955760405162461bcd60e51b81526004016102b49061108d565b600280549082015460405163a9059cbb60e01b815233600482015260248101919091526001600160a01b039091169063a9059cbb906044016020604051808303816000875af11580156109ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1091906110f7565b5060035481546040516323b872dd60e01b81526001600160a01b03928316926323b872dd92610a499233929091169087906004016110d3565b6020604051808303816000875af1158015610a68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a8c91906110f7565b50610c15565b6002546040516370a0823160e01b815233600482015283916001600160a01b0316906370a0823190602401602060405180830381865afa158015610ada573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610afe919061102e565b1015610b1c5760405162461bcd60e51b81526004016102b490611047565b600354600282015460405163a9059cbb60e01b815233600482015260248101919091526001600160a01b039091169063a9059cbb906044016020604051808303816000875af1158015610b73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9791906110f7565b5060025481546040516323b872dd60e01b81526001600160a01b03928316926323b872dd92610bd09233929091169087906004016110d3565b6020604051808303816000875af1158015610bef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1391906110f7565b505b8054610c2b9084906001600160a01b0316610dd1565b6040518381527f5fea849de7674e100db2af6f37d43d6251719ec75bbdaa5d3209fd70a99e8eb79060200161034e565b336000908152600660205260408120546060919067ffffffffffffffff811115610c8757610c87610fd3565b604051908082528060200260200182016040528015610cb0578160200160208202803683370190505b50905060005b3360009081526006602052604090205481101561041d5733600090815260046020908152604080832084845290915290205482518190849084908110610cfe57610cfe610fe9565b60209081029190910101525080610d1481611015565b915050610cb6565b3360008181526004602090815260408083206006808452828520805486529184528285208890556001600160a01b038716855260058452828520600785528386208054875290855283862089905582548987526008865284872055546009855292852092909255938352905281546001929190610d9a908490611120565b90915550506001600160a01b0381166000908152600760205260408120805460019290610dc8908490611120565b90915550505050565b6001600160a01b0381166000908152600660205260408120805460019290610dfa908490611138565b9091555050336000908152600760205260408120805460019290610e1f908490611138565b90915550506000828152600860209081526040808320839055600982528083208390553383526005825280832083805282528083208390556001600160a01b0393909316825260048082528383208380528252838320839055938252600a905290812080546001600160a01b03199081168255600182018054909116905560028101829055600381019190915501805460ff19169055565b80356001600160a01b0381168114610ece57600080fd5b919050565b60008060408385031215610ee657600080fd5b610eef83610eb7565b9150610efd60208401610eb7565b90509250929050565b600060208284031215610f1857600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b81811015610f5757835183529284019291840191600101610f3b565b50909695505050505050565b60008060008060808587031215610f7957600080fd5b610f8285610eb7565b93506020850135925060408501359150606085013560ff81168114610fa657600080fd5b939692955090935050565b60008060408385031215610fc457600080fd5b50508035926020909101359150565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161102757611027610fff565b5060010190565b60006020828403121561104057600080fd5b5051919050565b60208082526026908201527f746f6b656e5f6f6e652062616c616e636520696e73756666696369656e742062604082015265616c616e636560d01b606082015260800190565b60208082526026908201527f746f6b656e5f74776f2062616c616e636520696e73756666696369656e742062604082015265616c616e636560d01b606082015260800190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b60006020828403121561110957600080fd5b8151801515811461111957600080fd5b9392505050565b6000821982111561113357611133610fff565b500190565b60008282101561114a5761114a610fff565b50039056fea2646970667358221220fc938aeb8305dde945a37ad9230ffa7de389fbc7d29a289eaf94018766d5c64a64736f6c634300080f0033",
}

// ERC20InterchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20InterchangeMetaData.ABI instead.
var ERC20InterchangeABI = ERC20InterchangeMetaData.ABI

// Deprecated: Use ERC20InterchangeMetaData.Sigs instead.
// ERC20InterchangeFuncSigs maps the 4-byte function signature to its string representation.
var ERC20InterchangeFuncSigs = ERC20InterchangeMetaData.Sigs

// ERC20InterchangeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20InterchangeMetaData.Bin instead.
var ERC20InterchangeBin = ERC20InterchangeMetaData.Bin

// DeployERC20Interchange deploys a new Ethereum contract, binding an instance of ERC20Interchange to it.
func DeployERC20Interchange(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Interchange, error) {
	parsed, err := ERC20InterchangeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20InterchangeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Interchange{ERC20InterchangeCaller: ERC20InterchangeCaller{contract: contract}, ERC20InterchangeTransactor: ERC20InterchangeTransactor{contract: contract}, ERC20InterchangeFilterer: ERC20InterchangeFilterer{contract: contract}}, nil
}

// ERC20Interchange is an auto generated Go binding around an Ethereum contract.
type ERC20Interchange struct {
	ERC20InterchangeCaller     // Read-only binding to the contract
	ERC20InterchangeTransactor // Write-only binding to the contract
	ERC20InterchangeFilterer   // Log filterer for contract events
}

// ERC20InterchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20InterchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20InterchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20InterchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20InterchangeSession struct {
	Contract     *ERC20Interchange // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20InterchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20InterchangeCallerSession struct {
	Contract *ERC20InterchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC20InterchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20InterchangeTransactorSession struct {
	Contract     *ERC20InterchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC20InterchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20InterchangeRaw struct {
	Contract *ERC20Interchange // Generic contract binding to access the raw methods on
}

// ERC20InterchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20InterchangeCallerRaw struct {
	Contract *ERC20InterchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20InterchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20InterchangeTransactorRaw struct {
	Contract *ERC20InterchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Interchange creates a new instance of ERC20Interchange, bound to a specific deployed contract.
func NewERC20Interchange(address common.Address, backend bind.ContractBackend) (*ERC20Interchange, error) {
	contract, err := bindERC20Interchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Interchange{ERC20InterchangeCaller: ERC20InterchangeCaller{contract: contract}, ERC20InterchangeTransactor: ERC20InterchangeTransactor{contract: contract}, ERC20InterchangeFilterer: ERC20InterchangeFilterer{contract: contract}}, nil
}

// NewERC20InterchangeCaller creates a new read-only instance of ERC20Interchange, bound to a specific deployed contract.
func NewERC20InterchangeCaller(address common.Address, caller bind.ContractCaller) (*ERC20InterchangeCaller, error) {
	contract, err := bindERC20Interchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterchangeCaller{contract: contract}, nil
}

// NewERC20InterchangeTransactor creates a new write-only instance of ERC20Interchange, bound to a specific deployed contract.
func NewERC20InterchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20InterchangeTransactor, error) {
	contract, err := bindERC20Interchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterchangeTransactor{contract: contract}, nil
}

// NewERC20InterchangeFilterer creates a new log filterer instance of ERC20Interchange, bound to a specific deployed contract.
func NewERC20InterchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20InterchangeFilterer, error) {
	contract, err := bindERC20Interchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20InterchangeFilterer{contract: contract}, nil
}

// bindERC20Interchange binds a generic wrapper to an already deployed contract.
func bindERC20Interchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InterchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interchange *ERC20InterchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Interchange.Contract.ERC20InterchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interchange *ERC20InterchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interchange.Contract.ERC20InterchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interchange *ERC20InterchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interchange.Contract.ERC20InterchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interchange *ERC20InterchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Interchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interchange *ERC20InterchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interchange *ERC20InterchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interchange.Contract.contract.Transact(opts, method, params...)
}

// FromList is a free data retrieval call binding the contract method 0xa021857b.
//
// Solidity: function fromList() view returns(uint256[])
func (_ERC20Interchange *ERC20InterchangeCaller) FromList(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC20Interchange.contract.Call(opts, &out, "fromList")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// FromList is a free data retrieval call binding the contract method 0xa021857b.
//
// Solidity: function fromList() view returns(uint256[])
func (_ERC20Interchange *ERC20InterchangeSession) FromList() ([]*big.Int, error) {
	return _ERC20Interchange.Contract.FromList(&_ERC20Interchange.CallOpts)
}

// FromList is a free data retrieval call binding the contract method 0xa021857b.
//
// Solidity: function fromList() view returns(uint256[])
func (_ERC20Interchange *ERC20InterchangeCallerSession) FromList() ([]*big.Int, error) {
	return _ERC20Interchange.Contract.FromList(&_ERC20Interchange.CallOpts)
}

// SearchByInterchangeId is a free data retrieval call binding the contract method 0x6db92b7e.
//
// Solidity: function searchByInterchangeId(uint256 interchangeId) view returns((address,address,uint256,uint256,uint8))
func (_ERC20Interchange *ERC20InterchangeCaller) SearchByInterchangeId(opts *bind.CallOpts, interchangeId *big.Int) (ERC20InterchangeInterchange, error) {
	var out []interface{}
	err := _ERC20Interchange.contract.Call(opts, &out, "searchByInterchangeId", interchangeId)

	if err != nil {
		return *new(ERC20InterchangeInterchange), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20InterchangeInterchange)).(*ERC20InterchangeInterchange)

	return out0, err

}

// SearchByInterchangeId is a free data retrieval call binding the contract method 0x6db92b7e.
//
// Solidity: function searchByInterchangeId(uint256 interchangeId) view returns((address,address,uint256,uint256,uint8))
func (_ERC20Interchange *ERC20InterchangeSession) SearchByInterchangeId(interchangeId *big.Int) (ERC20InterchangeInterchange, error) {
	return _ERC20Interchange.Contract.SearchByInterchangeId(&_ERC20Interchange.CallOpts, interchangeId)
}

// SearchByInterchangeId is a free data retrieval call binding the contract method 0x6db92b7e.
//
// Solidity: function searchByInterchangeId(uint256 interchangeId) view returns((address,address,uint256,uint256,uint8))
func (_ERC20Interchange *ERC20InterchangeCallerSession) SearchByInterchangeId(interchangeId *big.Int) (ERC20InterchangeInterchange, error) {
	return _ERC20Interchange.Contract.SearchByInterchangeId(&_ERC20Interchange.CallOpts, interchangeId)
}

// ToList is a free data retrieval call binding the contract method 0x71645971.
//
// Solidity: function toList() view returns(uint256[])
func (_ERC20Interchange *ERC20InterchangeCaller) ToList(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC20Interchange.contract.Call(opts, &out, "toList")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ToList is a free data retrieval call binding the contract method 0x71645971.
//
// Solidity: function toList() view returns(uint256[])
func (_ERC20Interchange *ERC20InterchangeSession) ToList() ([]*big.Int, error) {
	return _ERC20Interchange.Contract.ToList(&_ERC20Interchange.CallOpts)
}

// ToList is a free data retrieval call binding the contract method 0x71645971.
//
// Solidity: function toList() view returns(uint256[])
func (_ERC20Interchange *ERC20InterchangeCallerSession) ToList() ([]*big.Int, error) {
	return _ERC20Interchange.Contract.ToList(&_ERC20Interchange.CallOpts)
}

// TokenOne is a free data retrieval call binding the contract method 0xc9caf28c.
//
// Solidity: function token_one() view returns(address)
func (_ERC20Interchange *ERC20InterchangeCaller) TokenOne(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Interchange.contract.Call(opts, &out, "token_one")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenOne is a free data retrieval call binding the contract method 0xc9caf28c.
//
// Solidity: function token_one() view returns(address)
func (_ERC20Interchange *ERC20InterchangeSession) TokenOne() (common.Address, error) {
	return _ERC20Interchange.Contract.TokenOne(&_ERC20Interchange.CallOpts)
}

// TokenOne is a free data retrieval call binding the contract method 0xc9caf28c.
//
// Solidity: function token_one() view returns(address)
func (_ERC20Interchange *ERC20InterchangeCallerSession) TokenOne() (common.Address, error) {
	return _ERC20Interchange.Contract.TokenOne(&_ERC20Interchange.CallOpts)
}

// TokenTwo is a free data retrieval call binding the contract method 0x34f97823.
//
// Solidity: function token_two() view returns(address)
func (_ERC20Interchange *ERC20InterchangeCaller) TokenTwo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Interchange.contract.Call(opts, &out, "token_two")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenTwo is a free data retrieval call binding the contract method 0x34f97823.
//
// Solidity: function token_two() view returns(address)
func (_ERC20Interchange *ERC20InterchangeSession) TokenTwo() (common.Address, error) {
	return _ERC20Interchange.Contract.TokenTwo(&_ERC20Interchange.CallOpts)
}

// TokenTwo is a free data retrieval call binding the contract method 0x34f97823.
//
// Solidity: function token_two() view returns(address)
func (_ERC20Interchange *ERC20InterchangeCallerSession) TokenTwo() (common.Address, error) {
	return _ERC20Interchange.Contract.TokenTwo(&_ERC20Interchange.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0x8caa5614.
//
// Solidity: function create(address to_address, uint256 from_amount, uint256 to_amount, uint8 interchange_type) returns(uint256 interchangeId)
func (_ERC20Interchange *ERC20InterchangeTransactor) Create(opts *bind.TransactOpts, to_address common.Address, from_amount *big.Int, to_amount *big.Int, interchange_type uint8) (*types.Transaction, error) {
	return _ERC20Interchange.contract.Transact(opts, "create", to_address, from_amount, to_amount, interchange_type)
}

// Create is a paid mutator transaction binding the contract method 0x8caa5614.
//
// Solidity: function create(address to_address, uint256 from_amount, uint256 to_amount, uint8 interchange_type) returns(uint256 interchangeId)
func (_ERC20Interchange *ERC20InterchangeSession) Create(to_address common.Address, from_amount *big.Int, to_amount *big.Int, interchange_type uint8) (*types.Transaction, error) {
	return _ERC20Interchange.Contract.Create(&_ERC20Interchange.TransactOpts, to_address, from_amount, to_amount, interchange_type)
}

// Create is a paid mutator transaction binding the contract method 0x8caa5614.
//
// Solidity: function create(address to_address, uint256 from_amount, uint256 to_amount, uint8 interchange_type) returns(uint256 interchangeId)
func (_ERC20Interchange *ERC20InterchangeTransactorSession) Create(to_address common.Address, from_amount *big.Int, to_amount *big.Int, interchange_type uint8) (*types.Transaction, error) {
	return _ERC20Interchange.Contract.Create(&_ERC20Interchange.TransactOpts, to_address, from_amount, to_amount, interchange_type)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token_one, address _token_two) returns()
func (_ERC20Interchange *ERC20InterchangeTransactor) Initialize(opts *bind.TransactOpts, _token_one common.Address, _token_two common.Address) (*types.Transaction, error) {
	return _ERC20Interchange.contract.Transact(opts, "initialize", _token_one, _token_two)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token_one, address _token_two) returns()
func (_ERC20Interchange *ERC20InterchangeSession) Initialize(_token_one common.Address, _token_two common.Address) (*types.Transaction, error) {
	return _ERC20Interchange.Contract.Initialize(&_ERC20Interchange.TransactOpts, _token_one, _token_two)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token_one, address _token_two) returns()
func (_ERC20Interchange *ERC20InterchangeTransactorSession) Initialize(_token_one common.Address, _token_two common.Address) (*types.Transaction, error) {
	return _ERC20Interchange.Contract.Initialize(&_ERC20Interchange.TransactOpts, _token_one, _token_two)
}

// Interchange is a paid mutator transaction binding the contract method 0x9e6e801f.
//
// Solidity: function interchange(uint256 interchangeId, uint256 amount) returns()
func (_ERC20Interchange *ERC20InterchangeTransactor) Interchange(opts *bind.TransactOpts, interchangeId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Interchange.contract.Transact(opts, "interchange", interchangeId, amount)
}

// Interchange is a paid mutator transaction binding the contract method 0x9e6e801f.
//
// Solidity: function interchange(uint256 interchangeId, uint256 amount) returns()
func (_ERC20Interchange *ERC20InterchangeSession) Interchange(interchangeId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Interchange.Contract.Interchange(&_ERC20Interchange.TransactOpts, interchangeId, amount)
}

// Interchange is a paid mutator transaction binding the contract method 0x9e6e801f.
//
// Solidity: function interchange(uint256 interchangeId, uint256 amount) returns()
func (_ERC20Interchange *ERC20InterchangeTransactorSession) Interchange(interchangeId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Interchange.Contract.Interchange(&_ERC20Interchange.TransactOpts, interchangeId, amount)
}

// ERC20InterchangeCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the ERC20Interchange contract.
type ERC20InterchangeCreateIterator struct {
	Event *ERC20InterchangeCreate // Event containing the contract specifics and raw log

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
func (it *ERC20InterchangeCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InterchangeCreate)
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
		it.Event = new(ERC20InterchangeCreate)
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
func (it *ERC20InterchangeCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InterchangeCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InterchangeCreate represents a Create event raised by the ERC20Interchange contract.
type ERC20InterchangeCreate struct {
	InterchangeId   *big.Int
	From            common.Address
	To              common.Address
	FromAmount      *big.Int
	ToAmount        *big.Int
	InterchangeType uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0xad6bba013d4dd6f7b6edc281a10f8b9003c42f19c04d7b88bde202abed854b95.
//
// Solidity: event Create(uint256 indexed interchangeId, address from, address to, uint256 from_amount, uint256 to_amount, uint8 interchange_type)
func (_ERC20Interchange *ERC20InterchangeFilterer) FilterCreate(opts *bind.FilterOpts, interchangeId []*big.Int) (*ERC20InterchangeCreateIterator, error) {

	var interchangeIdRule []interface{}
	for _, interchangeIdItem := range interchangeId {
		interchangeIdRule = append(interchangeIdRule, interchangeIdItem)
	}

	logs, sub, err := _ERC20Interchange.contract.FilterLogs(opts, "Create", interchangeIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC20InterchangeCreateIterator{contract: _ERC20Interchange.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0xad6bba013d4dd6f7b6edc281a10f8b9003c42f19c04d7b88bde202abed854b95.
//
// Solidity: event Create(uint256 indexed interchangeId, address from, address to, uint256 from_amount, uint256 to_amount, uint8 interchange_type)
func (_ERC20Interchange *ERC20InterchangeFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *ERC20InterchangeCreate, interchangeId []*big.Int) (event.Subscription, error) {

	var interchangeIdRule []interface{}
	for _, interchangeIdItem := range interchangeId {
		interchangeIdRule = append(interchangeIdRule, interchangeIdItem)
	}

	logs, sub, err := _ERC20Interchange.contract.WatchLogs(opts, "Create", interchangeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InterchangeCreate)
				if err := _ERC20Interchange.contract.UnpackLog(event, "Create", log); err != nil {
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

// ParseCreate is a log parse operation binding the contract event 0xad6bba013d4dd6f7b6edc281a10f8b9003c42f19c04d7b88bde202abed854b95.
//
// Solidity: event Create(uint256 indexed interchangeId, address from, address to, uint256 from_amount, uint256 to_amount, uint8 interchange_type)
func (_ERC20Interchange *ERC20InterchangeFilterer) ParseCreate(log types.Log) (*ERC20InterchangeCreate, error) {
	event := new(ERC20InterchangeCreate)
	if err := _ERC20Interchange.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20InterchangeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20Interchange contract.
type ERC20InterchangeInitializedIterator struct {
	Event *ERC20InterchangeInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20InterchangeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InterchangeInitialized)
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
		it.Event = new(ERC20InterchangeInitialized)
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
func (it *ERC20InterchangeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InterchangeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InterchangeInitialized represents a Initialized event raised by the ERC20Interchange contract.
type ERC20InterchangeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20Interchange *ERC20InterchangeFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20InterchangeInitializedIterator, error) {

	logs, sub, err := _ERC20Interchange.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20InterchangeInitializedIterator{contract: _ERC20Interchange.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20Interchange *ERC20InterchangeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20InterchangeInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20Interchange.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InterchangeInitialized)
				if err := _ERC20Interchange.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20Interchange *ERC20InterchangeFilterer) ParseInitialized(log types.Log) (*ERC20InterchangeInitialized, error) {
	event := new(ERC20InterchangeInitialized)
	if err := _ERC20Interchange.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20InterchangeInterchangeSuccessIterator is returned from FilterInterchangeSuccess and is used to iterate over the raw logs and unpacked data for InterchangeSuccess events raised by the ERC20Interchange contract.
type ERC20InterchangeInterchangeSuccessIterator struct {
	Event *ERC20InterchangeInterchangeSuccess // Event containing the contract specifics and raw log

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
func (it *ERC20InterchangeInterchangeSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InterchangeInterchangeSuccess)
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
		it.Event = new(ERC20InterchangeInterchangeSuccess)
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
func (it *ERC20InterchangeInterchangeSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InterchangeInterchangeSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InterchangeInterchangeSuccess represents a InterchangeSuccess event raised by the ERC20Interchange contract.
type ERC20InterchangeInterchangeSuccess struct {
	InterchangeId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterchangeSuccess is a free log retrieval operation binding the contract event 0x5fea849de7674e100db2af6f37d43d6251719ec75bbdaa5d3209fd70a99e8eb7.
//
// Solidity: event InterchangeSuccess(uint256 interchangeId)
func (_ERC20Interchange *ERC20InterchangeFilterer) FilterInterchangeSuccess(opts *bind.FilterOpts) (*ERC20InterchangeInterchangeSuccessIterator, error) {

	logs, sub, err := _ERC20Interchange.contract.FilterLogs(opts, "InterchangeSuccess")
	if err != nil {
		return nil, err
	}
	return &ERC20InterchangeInterchangeSuccessIterator{contract: _ERC20Interchange.contract, event: "InterchangeSuccess", logs: logs, sub: sub}, nil
}

// WatchInterchangeSuccess is a free log subscription operation binding the contract event 0x5fea849de7674e100db2af6f37d43d6251719ec75bbdaa5d3209fd70a99e8eb7.
//
// Solidity: event InterchangeSuccess(uint256 interchangeId)
func (_ERC20Interchange *ERC20InterchangeFilterer) WatchInterchangeSuccess(opts *bind.WatchOpts, sink chan<- *ERC20InterchangeInterchangeSuccess) (event.Subscription, error) {

	logs, sub, err := _ERC20Interchange.contract.WatchLogs(opts, "InterchangeSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InterchangeInterchangeSuccess)
				if err := _ERC20Interchange.contract.UnpackLog(event, "InterchangeSuccess", log); err != nil {
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

// ParseInterchangeSuccess is a log parse operation binding the contract event 0x5fea849de7674e100db2af6f37d43d6251719ec75bbdaa5d3209fd70a99e8eb7.
//
// Solidity: event InterchangeSuccess(uint256 interchangeId)
func (_ERC20Interchange *ERC20InterchangeFilterer) ParseInterchangeSuccess(log types.Log) (*ERC20InterchangeInterchangeSuccess, error) {
	event := new(ERC20InterchangeInterchangeSuccess)
	if err := _ERC20Interchange.contract.UnpackLog(event, "InterchangeSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20UpgradeableMetaData contains all meta data concerning the IERC20Upgradeable contract.
var IERC20UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20UpgradeableMetaData.ABI instead.
var IERC20UpgradeableABI = IERC20UpgradeableMetaData.ABI

// Deprecated: Use IERC20UpgradeableMetaData.Sigs instead.
// IERC20UpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IERC20UpgradeableFuncSigs = IERC20UpgradeableMetaData.Sigs

// IERC20Upgradeable is an auto generated Go binding around an Ethereum contract.
type IERC20Upgradeable struct {
	IERC20UpgradeableCaller     // Read-only binding to the contract
	IERC20UpgradeableTransactor // Write-only binding to the contract
	IERC20UpgradeableFilterer   // Log filterer for contract events
}

// IERC20UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20UpgradeableSession struct {
	Contract     *IERC20Upgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC20UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20UpgradeableCallerSession struct {
	Contract *IERC20UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IERC20UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20UpgradeableTransactorSession struct {
	Contract     *IERC20UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IERC20UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20UpgradeableRaw struct {
	Contract *IERC20Upgradeable // Generic contract binding to access the raw methods on
}

// IERC20UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20UpgradeableCallerRaw struct {
	Contract *IERC20UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20UpgradeableTransactorRaw struct {
	Contract *IERC20UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Upgradeable creates a new instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20Upgradeable(address common.Address, backend bind.ContractBackend) (*IERC20Upgradeable, error) {
	contract, err := bindIERC20Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Upgradeable{IERC20UpgradeableCaller: IERC20UpgradeableCaller{contract: contract}, IERC20UpgradeableTransactor: IERC20UpgradeableTransactor{contract: contract}, IERC20UpgradeableFilterer: IERC20UpgradeableFilterer{contract: contract}}, nil
}

// NewIERC20UpgradeableCaller creates a new read-only instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IERC20UpgradeableCaller, error) {
	contract, err := bindIERC20Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableCaller{contract: contract}, nil
}

// NewIERC20UpgradeableTransactor creates a new write-only instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20UpgradeableTransactor, error) {
	contract, err := bindIERC20Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableTransactor{contract: contract}, nil
}

// NewIERC20UpgradeableFilterer creates a new log filterer instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20UpgradeableFilterer, error) {
	contract, err := bindIERC20Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableFilterer{contract: contract}, nil
}

// bindIERC20Upgradeable binds a generic wrapper to an already deployed contract.
func bindIERC20Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20UpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Upgradeable *IERC20UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Upgradeable.Contract.IERC20UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Upgradeable *IERC20UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.IERC20UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Upgradeable *IERC20UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.IERC20UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Upgradeable *IERC20UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Upgradeable *IERC20UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Upgradeable *IERC20UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Upgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.Allowance(&_IERC20Upgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.Allowance(&_IERC20Upgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Upgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.BalanceOf(&_IERC20Upgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.BalanceOf(&_IERC20Upgradeable.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Upgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableSession) TotalSupply() (*big.Int, error) {
	return _IERC20Upgradeable.Contract.TotalSupply(&_IERC20Upgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Upgradeable.Contract.TotalSupply(&_IERC20Upgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Approve(&_IERC20Upgradeable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Approve(&_IERC20Upgradeable.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Transfer(&_IERC20Upgradeable.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Transfer(&_IERC20Upgradeable.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.TransferFrom(&_IERC20Upgradeable.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.TransferFrom(&_IERC20Upgradeable.TransactOpts, from, to, amount)
}

// IERC20UpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Upgradeable contract.
type IERC20UpgradeableApprovalIterator struct {
	Event *IERC20UpgradeableApproval // Event containing the contract specifics and raw log

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
func (it *IERC20UpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20UpgradeableApproval)
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
		it.Event = new(IERC20UpgradeableApproval)
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
func (it *IERC20UpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20UpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20UpgradeableApproval represents a Approval event raised by the IERC20Upgradeable contract.
type IERC20UpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20UpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableApprovalIterator{contract: _IERC20Upgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20UpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20UpgradeableApproval)
				if err := _IERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) ParseApproval(log types.Log) (*IERC20UpgradeableApproval, error) {
	event := new(IERC20UpgradeableApproval)
	if err := _IERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20UpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Upgradeable contract.
type IERC20UpgradeableTransferIterator struct {
	Event *IERC20UpgradeableTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20UpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20UpgradeableTransfer)
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
		it.Event = new(IERC20UpgradeableTransfer)
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
func (it *IERC20UpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20UpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20UpgradeableTransfer represents a Transfer event raised by the IERC20Upgradeable contract.
type IERC20UpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20UpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableTransferIterator{contract: _IERC20Upgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20UpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20UpgradeableTransfer)
				if err := _IERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) ParseTransfer(log types.Log) (*IERC20UpgradeableTransfer, error) {
	event := new(IERC20UpgradeableTransfer)
	if err := _IERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// InitializableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Initializable contract.
type InitializableInitializedIterator struct {
	Event *InitializableInitialized // Event containing the contract specifics and raw log

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
func (it *InitializableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableInitialized)
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
		it.Event = new(InitializableInitialized)
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
func (it *InitializableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableInitialized represents a Initialized event raised by the Initializable contract.
type InitializableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) FilterInitialized(opts *bind.FilterOpts) (*InitializableInitializedIterator, error) {

	logs, sub, err := _Initializable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InitializableInitializedIterator{contract: _Initializable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InitializableInitialized) (event.Subscription, error) {

	logs, sub, err := _Initializable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableInitialized)
				if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) ParseInitialized(log types.Log) (*InitializableInitialized, error) {
	event := new(InitializableInitialized)
	if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
