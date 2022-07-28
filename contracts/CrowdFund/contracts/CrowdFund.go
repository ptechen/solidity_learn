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

// CrowdFundMetaData contains all meta data concerning the CrowdFund contract.
var CrowdFundMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startedAt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endAt\",\"type\":\"uint32\"}],\"name\":\"Launch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Pledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"UnPledge\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaigns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pledged\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startedAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endAt\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_goal\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_startAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_endAt\",\"type\":\"uint32\"}],\"name\":\"launch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pledgedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unPledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"141961bc": "campaigns(uint256)",
		"40e58ee5": "cancel(uint256)",
		"379607f5": "claim(uint256)",
		"06661abd": "count()",
		"2c63f146": "launch(uint256,uint32,uint32)",
		"fde327be": "pledge(uint256,uint256)",
		"aa4fb63a": "pledgedAmount(uint256,address)",
		"278ecde1": "refund(uint256)",
		"fc0c546a": "token()",
		"a3a0e936": "unPledge(uint256,uint256)",
	},
	Bin: "0x60a060405234801561001057600080fd5b50604051610e5a380380610e5a83398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051610db36100a7600039600081816101f2015281816103290152818161071e015281816109de0152610b990152610db36000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806340e58ee51161006657806340e58ee51461019c578063a3a0e936146101af578063aa4fb63a146101c2578063fc0c546a146101ed578063fde327be1461022c57600080fd5b806306661abd146100a3578063141961bc146100bf578063278ecde1146101615780632c63f14614610176578063379607f514610189575b600080fd5b6100ac60005481565b6040519081526020015b60405180910390f35b61011f6100cd366004610c43565b600160208190526000918252604090912080549181015460028201546003909201546001600160a01b039093169290919063ffffffff80821691600160201b810490911690600160401b900460ff1686565b604080516001600160a01b03909716875260208701959095529385019290925263ffffffff9081166060850152166080830152151560a082015260c0016100b6565b61017461016f366004610c43565b61023f565b005b610174610184366004610c75565b61039c565b610174610197366004610c43565b6105af565b6101746101aa366004610c43565b6107c4565b6101746101bd366004610cb1565b610924565b6100ac6101d0366004610cd3565b600260209081526000928352604080842090915290825290205481565b6102147f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100b6565b61017461023a366004610cb1565b610a91565b60008181526001602052604090206003810154600160201b900463ffffffff16421161029e5760405162461bcd60e51b81526020600482015260096024820152681b9bdd08195b99195960ba1b60448201526064015b60405180910390fd5b80600101548160020154106102e65760405162461bcd60e51b815260206004820152600e60248201526d1c1b195919d959080f0819dbd85b60921b6044820152606401610295565b600082815260026020908152604080832033808552925280832080549390555163a9059cbb60e01b81526004810191909152602481018290526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af1158015610372573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103969190610d0f565b50505050565b428263ffffffff1610156103e35760405162461bcd60e51b815260206004820152600e60248201526d7374617274206174203c206e6f7760901b6044820152606401610295565b8163ffffffff168163ffffffff1610156104335760405162461bcd60e51b8152602060048201526011602482015270195b9d08185d080f081cdd185c9d08185d607a1b6044820152606401610295565b610440426276a700610d4e565b8163ffffffff16111561048d5760405162461bcd60e51b815260206004820152601560248201527432b7321030ba101f1036b0bc10323ab930ba34b7b760591b6044820152606401610295565b600160008082825461049f9190610d4e565b90915550506040805160c081018252338082526020808301878152600084860181815263ffffffff89811660608089018281528b84166080808c0182815260a08d0189815289548a526001808d528f8b209e518f546001600160a01b0319166001600160a01b03909116178f559a519a8e019a909a55965160028d015591516003909b018054965198519b861667ffffffffffffffff1990971696909617600160201b98909516979097029390931760ff60401b1916600160401b9915159990990298909817909255915487519081529384018a90528387015293820152925190927f0601cd0d650b473037e838a2696d41e654433d065b3f56b28d1d3302e44a304f92908290030190a2505050565b600081815260016020526040902080546001600160a01b031633146106045760405162461bcd60e51b815260206004820152600b60248201526a3737ba1031b932b0ba37b960a91b6044820152606401610295565b6003810154600160201b900463ffffffff1642116106505760405162461bcd60e51b81526020600482015260096024820152681b9bdd08195b99195960ba1b6044820152606401610295565b8060010154816002015410156106995760405162461bcd60e51b815260206004820152600e60248201526d1c1b195919d959080f0819dbd85b60921b6044820152606401610295565b6003810154600160401b900460ff16156106df5760405162461bcd60e51b815260206004820152600760248201526618db185a5b595960ca1b6044820152606401610295565b60038101805460ff60401b1916600160401b179055600281015460405163a9059cbb60e01b815233600482015260248101919091526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af1158015610767573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078b9190610d0f565b506040518281527f7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7906020015b60405180910390a15050565b600081815260016020818152604092839020835160c08101855281546001600160a01b0316808252938201549281019290925260028101549382019390935260039092015463ffffffff8082166060850152600160201b8204166080840152600160401b900460ff16151560a083015233146108705760405162461bcd60e51b815260206004820152600b60248201526a3737ba1031b932b0ba37b960a91b6044820152606401610295565b806060015163ffffffff164210156108b45760405162461bcd60e51b81526020600482015260076024820152661cdd185c9d195960ca1b6044820152606401610295565b600082815260016020819052604080832080546001600160a01b03191681559182018390556002820192909255600301805468ffffffffffffffffff19169055517f8bf30e7ff26833413be5f69e1d373744864d600b664204b4a2f9844a8eedb9ed906107b89084815260200190565b60008281526001602052604090206003810154600160201b900463ffffffff1642111561097b5760405162461bcd60e51b8152602060048201526005602482015264195b99195960da1b6044820152606401610295565b8181600201600082825461098f9190610d66565b90915550506000838152600260209081526040808320338452909152812080548492906109bd908490610d66565b909155505060405163a9059cbb60e01b8152336004820152602481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a9059cbb906044016020604051808303816000875af1158015610a2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a539190610d0f565b50604051828152339084907fdc63160518c9c709fe39309a0e6f3d567887986a4625cd35090a6689e56e34c0906020015b60405180910390a3505050565b6000828152600160205260409020600381015463ffffffff16421015610ae75760405162461bcd60e51b815260206004820152600b60248201526a1b9bdd081cdd185c9d195960aa1b6044820152606401610295565b6003810154600160201b900463ffffffff16421115610b305760405162461bcd60e51b8152602060048201526005602482015264195b99195960da1b6044820152606401610295565b81816002016000828254610b449190610d4e565b9091555050600083815260026020908152604080832033845290915281208054849290610b72908490610d4e565b90915550506040516323b872dd60e01b8152336004820152306024820152604481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610bea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0e9190610d0f565b50604051828152339084907f06bdb975df800a73232998e71ed585d536222f1dfeaa622d7f62a23ada686c8290602001610a84565b600060208284031215610c5557600080fd5b5035919050565b803563ffffffff81168114610c7057600080fd5b919050565b600080600060608486031215610c8a57600080fd5b83359250610c9a60208501610c5c565b9150610ca860408501610c5c565b90509250925092565b60008060408385031215610cc457600080fd5b50508035926020909101359150565b60008060408385031215610ce657600080fd5b8235915060208301356001600160a01b0381168114610d0457600080fd5b809150509250929050565b600060208284031215610d2157600080fd5b81518015158114610d3157600080fd5b9392505050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610d6157610d61610d38565b500190565b600082821015610d7857610d78610d38565b50039056fea26469706673582212205afe286e7c03332827094b166a673684f41e16ef32adf39b364dc5562205fd3964736f6c634300080f0033",
}

// CrowdFundABI is the input ABI used to generate the binding from.
// Deprecated: Use CrowdFundMetaData.ABI instead.
var CrowdFundABI = CrowdFundMetaData.ABI

// Deprecated: Use CrowdFundMetaData.Sigs instead.
// CrowdFundFuncSigs maps the 4-byte function signature to its string representation.
var CrowdFundFuncSigs = CrowdFundMetaData.Sigs

// CrowdFundBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrowdFundMetaData.Bin instead.
var CrowdFundBin = CrowdFundMetaData.Bin

// DeployCrowdFund deploys a new Ethereum contract, binding an instance of CrowdFund to it.
func DeployCrowdFund(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *CrowdFund, error) {
	parsed, err := CrowdFundMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrowdFundBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrowdFund{CrowdFundCaller: CrowdFundCaller{contract: contract}, CrowdFundTransactor: CrowdFundTransactor{contract: contract}, CrowdFundFilterer: CrowdFundFilterer{contract: contract}}, nil
}

// CrowdFund is an auto generated Go binding around an Ethereum contract.
type CrowdFund struct {
	CrowdFundCaller     // Read-only binding to the contract
	CrowdFundTransactor // Write-only binding to the contract
	CrowdFundFilterer   // Log filterer for contract events
}

// CrowdFundCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdFundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdFundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdFundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdFundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrowdFundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdFundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdFundSession struct {
	Contract     *CrowdFund        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdFundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdFundCallerSession struct {
	Contract *CrowdFundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrowdFundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdFundTransactorSession struct {
	Contract     *CrowdFundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrowdFundRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdFundRaw struct {
	Contract *CrowdFund // Generic contract binding to access the raw methods on
}

// CrowdFundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdFundCallerRaw struct {
	Contract *CrowdFundCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdFundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdFundTransactorRaw struct {
	Contract *CrowdFundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdFund creates a new instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFund(address common.Address, backend bind.ContractBackend) (*CrowdFund, error) {
	contract, err := bindCrowdFund(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrowdFund{CrowdFundCaller: CrowdFundCaller{contract: contract}, CrowdFundTransactor: CrowdFundTransactor{contract: contract}, CrowdFundFilterer: CrowdFundFilterer{contract: contract}}, nil
}

// NewCrowdFundCaller creates a new read-only instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFundCaller(address common.Address, caller bind.ContractCaller) (*CrowdFundCaller, error) {
	contract, err := bindCrowdFund(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdFundCaller{contract: contract}, nil
}

// NewCrowdFundTransactor creates a new write-only instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFundTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdFundTransactor, error) {
	contract, err := bindCrowdFund(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdFundTransactor{contract: contract}, nil
}

// NewCrowdFundFilterer creates a new log filterer instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFundFilterer(address common.Address, filterer bind.ContractFilterer) (*CrowdFundFilterer, error) {
	contract, err := bindCrowdFund(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrowdFundFilterer{contract: contract}, nil
}

// bindCrowdFund binds a generic wrapper to an already deployed contract.
func bindCrowdFund(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdFundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdFund *CrowdFundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdFund.Contract.CrowdFundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdFund *CrowdFundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdFund.Contract.CrowdFundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdFund *CrowdFundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdFund.Contract.CrowdFundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdFund *CrowdFundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdFund.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdFund *CrowdFundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdFund.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdFund *CrowdFundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdFund.Contract.contract.Transact(opts, method, params...)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address creator, uint256 goal, uint256 pledged, uint32 startedAt, uint32 endAt, bool claimed)
func (_CrowdFund *CrowdFundCaller) Campaigns(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Creator   common.Address
	Goal      *big.Int
	Pledged   *big.Int
	StartedAt uint32
	EndAt     uint32
	Claimed   bool
}, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "campaigns", arg0)

	outstruct := new(struct {
		Creator   common.Address
		Goal      *big.Int
		Pledged   *big.Int
		StartedAt uint32
		EndAt     uint32
		Claimed   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Creator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Goal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Pledged = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.EndAt = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.Claimed = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address creator, uint256 goal, uint256 pledged, uint32 startedAt, uint32 endAt, bool claimed)
func (_CrowdFund *CrowdFundSession) Campaigns(arg0 *big.Int) (struct {
	Creator   common.Address
	Goal      *big.Int
	Pledged   *big.Int
	StartedAt uint32
	EndAt     uint32
	Claimed   bool
}, error) {
	return _CrowdFund.Contract.Campaigns(&_CrowdFund.CallOpts, arg0)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address creator, uint256 goal, uint256 pledged, uint32 startedAt, uint32 endAt, bool claimed)
func (_CrowdFund *CrowdFundCallerSession) Campaigns(arg0 *big.Int) (struct {
	Creator   common.Address
	Goal      *big.Int
	Pledged   *big.Int
	StartedAt uint32
	EndAt     uint32
	Claimed   bool
}, error) {
	return _CrowdFund.Contract.Campaigns(&_CrowdFund.CallOpts, arg0)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CrowdFund *CrowdFundCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CrowdFund *CrowdFundSession) Count() (*big.Int, error) {
	return _CrowdFund.Contract.Count(&_CrowdFund.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CrowdFund *CrowdFundCallerSession) Count() (*big.Int, error) {
	return _CrowdFund.Contract.Count(&_CrowdFund.CallOpts)
}

// PledgedAmount is a free data retrieval call binding the contract method 0xaa4fb63a.
//
// Solidity: function pledgedAmount(uint256 , address ) view returns(uint256)
func (_CrowdFund *CrowdFundCaller) PledgedAmount(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "pledgedAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgedAmount is a free data retrieval call binding the contract method 0xaa4fb63a.
//
// Solidity: function pledgedAmount(uint256 , address ) view returns(uint256)
func (_CrowdFund *CrowdFundSession) PledgedAmount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _CrowdFund.Contract.PledgedAmount(&_CrowdFund.CallOpts, arg0, arg1)
}

// PledgedAmount is a free data retrieval call binding the contract method 0xaa4fb63a.
//
// Solidity: function pledgedAmount(uint256 , address ) view returns(uint256)
func (_CrowdFund *CrowdFundCallerSession) PledgedAmount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _CrowdFund.Contract.PledgedAmount(&_CrowdFund.CallOpts, arg0, arg1)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdFund *CrowdFundCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdFund *CrowdFundSession) Token() (common.Address, error) {
	return _CrowdFund.Contract.Token(&_CrowdFund.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdFund *CrowdFundCallerSession) Token() (common.Address, error) {
	return _CrowdFund.Contract.Token(&_CrowdFund.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactor) Cancel(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "cancel", _id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _id) returns()
func (_CrowdFund *CrowdFundSession) Cancel(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Cancel(&_CrowdFund.TransactOpts, _id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactorSession) Cancel(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Cancel(&_CrowdFund.TransactOpts, _id)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactor) Claim(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "claim", _id)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _id) returns()
func (_CrowdFund *CrowdFundSession) Claim(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Claim(&_CrowdFund.TransactOpts, _id)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactorSession) Claim(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Claim(&_CrowdFund.TransactOpts, _id)
}

// Launch is a paid mutator transaction binding the contract method 0x2c63f146.
//
// Solidity: function launch(uint256 _goal, uint32 _startAt, uint32 _endAt) returns()
func (_CrowdFund *CrowdFundTransactor) Launch(opts *bind.TransactOpts, _goal *big.Int, _startAt uint32, _endAt uint32) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "launch", _goal, _startAt, _endAt)
}

// Launch is a paid mutator transaction binding the contract method 0x2c63f146.
//
// Solidity: function launch(uint256 _goal, uint32 _startAt, uint32 _endAt) returns()
func (_CrowdFund *CrowdFundSession) Launch(_goal *big.Int, _startAt uint32, _endAt uint32) (*types.Transaction, error) {
	return _CrowdFund.Contract.Launch(&_CrowdFund.TransactOpts, _goal, _startAt, _endAt)
}

// Launch is a paid mutator transaction binding the contract method 0x2c63f146.
//
// Solidity: function launch(uint256 _goal, uint32 _startAt, uint32 _endAt) returns()
func (_CrowdFund *CrowdFundTransactorSession) Launch(_goal *big.Int, _startAt uint32, _endAt uint32) (*types.Transaction, error) {
	return _CrowdFund.Contract.Launch(&_CrowdFund.TransactOpts, _goal, _startAt, _endAt)
}

// Pledge is a paid mutator transaction binding the contract method 0xfde327be.
//
// Solidity: function pledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundTransactor) Pledge(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "pledge", _id, _amount)
}

// Pledge is a paid mutator transaction binding the contract method 0xfde327be.
//
// Solidity: function pledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundSession) Pledge(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Pledge(&_CrowdFund.TransactOpts, _id, _amount)
}

// Pledge is a paid mutator transaction binding the contract method 0xfde327be.
//
// Solidity: function pledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundTransactorSession) Pledge(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Pledge(&_CrowdFund.TransactOpts, _id, _amount)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactor) Refund(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "refund", _id)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 _id) returns()
func (_CrowdFund *CrowdFundSession) Refund(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Refund(&_CrowdFund.TransactOpts, _id)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactorSession) Refund(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Refund(&_CrowdFund.TransactOpts, _id)
}

// UnPledge is a paid mutator transaction binding the contract method 0xa3a0e936.
//
// Solidity: function unPledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundTransactor) UnPledge(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "unPledge", _id, _amount)
}

// UnPledge is a paid mutator transaction binding the contract method 0xa3a0e936.
//
// Solidity: function unPledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundSession) UnPledge(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.UnPledge(&_CrowdFund.TransactOpts, _id, _amount)
}

// UnPledge is a paid mutator transaction binding the contract method 0xa3a0e936.
//
// Solidity: function unPledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundTransactorSession) UnPledge(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.UnPledge(&_CrowdFund.TransactOpts, _id, _amount)
}

// CrowdFundCancelIterator is returned from FilterCancel and is used to iterate over the raw logs and unpacked data for Cancel events raised by the CrowdFund contract.
type CrowdFundCancelIterator struct {
	Event *CrowdFundCancel // Event containing the contract specifics and raw log

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
func (it *CrowdFundCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundCancel)
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
		it.Event = new(CrowdFundCancel)
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
func (it *CrowdFundCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundCancel represents a Cancel event raised by the CrowdFund contract.
type CrowdFundCancel struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancel is a free log retrieval operation binding the contract event 0x8bf30e7ff26833413be5f69e1d373744864d600b664204b4a2f9844a8eedb9ed.
//
// Solidity: event Cancel(uint256 _id)
func (_CrowdFund *CrowdFundFilterer) FilterCancel(opts *bind.FilterOpts) (*CrowdFundCancelIterator, error) {

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Cancel")
	if err != nil {
		return nil, err
	}
	return &CrowdFundCancelIterator{contract: _CrowdFund.contract, event: "Cancel", logs: logs, sub: sub}, nil
}

// WatchCancel is a free log subscription operation binding the contract event 0x8bf30e7ff26833413be5f69e1d373744864d600b664204b4a2f9844a8eedb9ed.
//
// Solidity: event Cancel(uint256 _id)
func (_CrowdFund *CrowdFundFilterer) WatchCancel(opts *bind.WatchOpts, sink chan<- *CrowdFundCancel) (event.Subscription, error) {

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Cancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundCancel)
				if err := _CrowdFund.contract.UnpackLog(event, "Cancel", log); err != nil {
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

// ParseCancel is a log parse operation binding the contract event 0x8bf30e7ff26833413be5f69e1d373744864d600b664204b4a2f9844a8eedb9ed.
//
// Solidity: event Cancel(uint256 _id)
func (_CrowdFund *CrowdFundFilterer) ParseCancel(log types.Log) (*CrowdFundCancel, error) {
	event := new(CrowdFundCancel)
	if err := _CrowdFund.contract.UnpackLog(event, "Cancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the CrowdFund contract.
type CrowdFundClaimIterator struct {
	Event *CrowdFundClaim // Event containing the contract specifics and raw log

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
func (it *CrowdFundClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundClaim)
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
		it.Event = new(CrowdFundClaim)
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
func (it *CrowdFundClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundClaim represents a Claim event raised by the CrowdFund contract.
type CrowdFundClaim struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7.
//
// Solidity: event Claim(uint256 _id)
func (_CrowdFund *CrowdFundFilterer) FilterClaim(opts *bind.FilterOpts) (*CrowdFundClaimIterator, error) {

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &CrowdFundClaimIterator{contract: _CrowdFund.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7.
//
// Solidity: event Claim(uint256 _id)
func (_CrowdFund *CrowdFundFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *CrowdFundClaim) (event.Subscription, error) {

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundClaim)
				if err := _CrowdFund.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7.
//
// Solidity: event Claim(uint256 _id)
func (_CrowdFund *CrowdFundFilterer) ParseClaim(log types.Log) (*CrowdFundClaim, error) {
	event := new(CrowdFundClaim)
	if err := _CrowdFund.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundLaunchIterator is returned from FilterLaunch and is used to iterate over the raw logs and unpacked data for Launch events raised by the CrowdFund contract.
type CrowdFundLaunchIterator struct {
	Event *CrowdFundLaunch // Event containing the contract specifics and raw log

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
func (it *CrowdFundLaunchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundLaunch)
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
		it.Event = new(CrowdFundLaunch)
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
func (it *CrowdFundLaunchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundLaunchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundLaunch represents a Launch event raised by the CrowdFund contract.
type CrowdFundLaunch struct {
	Id        *big.Int
	Creator   common.Address
	Goal      *big.Int
	StartedAt uint32
	EndAt     uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLaunch is a free log retrieval operation binding the contract event 0x0601cd0d650b473037e838a2696d41e654433d065b3f56b28d1d3302e44a304f.
//
// Solidity: event Launch(uint256 id, address indexed creator, uint256 goal, uint32 startedAt, uint32 endAt)
func (_CrowdFund *CrowdFundFilterer) FilterLaunch(opts *bind.FilterOpts, creator []common.Address) (*CrowdFundLaunchIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Launch", creatorRule)
	if err != nil {
		return nil, err
	}
	return &CrowdFundLaunchIterator{contract: _CrowdFund.contract, event: "Launch", logs: logs, sub: sub}, nil
}

// WatchLaunch is a free log subscription operation binding the contract event 0x0601cd0d650b473037e838a2696d41e654433d065b3f56b28d1d3302e44a304f.
//
// Solidity: event Launch(uint256 id, address indexed creator, uint256 goal, uint32 startedAt, uint32 endAt)
func (_CrowdFund *CrowdFundFilterer) WatchLaunch(opts *bind.WatchOpts, sink chan<- *CrowdFundLaunch, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Launch", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundLaunch)
				if err := _CrowdFund.contract.UnpackLog(event, "Launch", log); err != nil {
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

// ParseLaunch is a log parse operation binding the contract event 0x0601cd0d650b473037e838a2696d41e654433d065b3f56b28d1d3302e44a304f.
//
// Solidity: event Launch(uint256 id, address indexed creator, uint256 goal, uint32 startedAt, uint32 endAt)
func (_CrowdFund *CrowdFundFilterer) ParseLaunch(log types.Log) (*CrowdFundLaunch, error) {
	event := new(CrowdFundLaunch)
	if err := _CrowdFund.contract.UnpackLog(event, "Launch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundPledgeIterator is returned from FilterPledge and is used to iterate over the raw logs and unpacked data for Pledge events raised by the CrowdFund contract.
type CrowdFundPledgeIterator struct {
	Event *CrowdFundPledge // Event containing the contract specifics and raw log

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
func (it *CrowdFundPledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundPledge)
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
		it.Event = new(CrowdFundPledge)
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
func (it *CrowdFundPledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundPledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundPledge represents a Pledge event raised by the CrowdFund contract.
type CrowdFundPledge struct {
	Id     *big.Int
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPledge is a free log retrieval operation binding the contract event 0x06bdb975df800a73232998e71ed585d536222f1dfeaa622d7f62a23ada686c82.
//
// Solidity: event Pledge(uint256 indexed _id, address indexed caller, uint256 _amount)
func (_CrowdFund *CrowdFundFilterer) FilterPledge(opts *bind.FilterOpts, _id []*big.Int, caller []common.Address) (*CrowdFundPledgeIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Pledge", _idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdFundPledgeIterator{contract: _CrowdFund.contract, event: "Pledge", logs: logs, sub: sub}, nil
}

// WatchPledge is a free log subscription operation binding the contract event 0x06bdb975df800a73232998e71ed585d536222f1dfeaa622d7f62a23ada686c82.
//
// Solidity: event Pledge(uint256 indexed _id, address indexed caller, uint256 _amount)
func (_CrowdFund *CrowdFundFilterer) WatchPledge(opts *bind.WatchOpts, sink chan<- *CrowdFundPledge, _id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Pledge", _idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundPledge)
				if err := _CrowdFund.contract.UnpackLog(event, "Pledge", log); err != nil {
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

// ParsePledge is a log parse operation binding the contract event 0x06bdb975df800a73232998e71ed585d536222f1dfeaa622d7f62a23ada686c82.
//
// Solidity: event Pledge(uint256 indexed _id, address indexed caller, uint256 _amount)
func (_CrowdFund *CrowdFundFilterer) ParsePledge(log types.Log) (*CrowdFundPledge, error) {
	event := new(CrowdFundPledge)
	if err := _CrowdFund.contract.UnpackLog(event, "Pledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the CrowdFund contract.
type CrowdFundRefundIterator struct {
	Event *CrowdFundRefund // Event containing the contract specifics and raw log

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
func (it *CrowdFundRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundRefund)
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
		it.Event = new(CrowdFundRefund)
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
func (it *CrowdFundRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundRefund represents a Refund event raised by the CrowdFund contract.
type CrowdFundRefund struct {
	Id     *big.Int
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x21e12a7cad0da5928167e1084ea4d5fdf8d9af66657a2543a9ac76a0ca081477.
//
// Solidity: event Refund(uint256 indexed id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) FilterRefund(opts *bind.FilterOpts, id []*big.Int, caller []common.Address) (*CrowdFundRefundIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Refund", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdFundRefundIterator{contract: _CrowdFund.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x21e12a7cad0da5928167e1084ea4d5fdf8d9af66657a2543a9ac76a0ca081477.
//
// Solidity: event Refund(uint256 indexed id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *CrowdFundRefund, id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Refund", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundRefund)
				if err := _CrowdFund.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x21e12a7cad0da5928167e1084ea4d5fdf8d9af66657a2543a9ac76a0ca081477.
//
// Solidity: event Refund(uint256 indexed id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) ParseRefund(log types.Log) (*CrowdFundRefund, error) {
	event := new(CrowdFundRefund)
	if err := _CrowdFund.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundUnPledgeIterator is returned from FilterUnPledge and is used to iterate over the raw logs and unpacked data for UnPledge events raised by the CrowdFund contract.
type CrowdFundUnPledgeIterator struct {
	Event *CrowdFundUnPledge // Event containing the contract specifics and raw log

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
func (it *CrowdFundUnPledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundUnPledge)
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
		it.Event = new(CrowdFundUnPledge)
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
func (it *CrowdFundUnPledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundUnPledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundUnPledge represents a UnPledge event raised by the CrowdFund contract.
type CrowdFundUnPledge struct {
	Id     *big.Int
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnPledge is a free log retrieval operation binding the contract event 0xdc63160518c9c709fe39309a0e6f3d567887986a4625cd35090a6689e56e34c0.
//
// Solidity: event UnPledge(uint256 indexed _id, address indexed caller, uint256 _amount)
func (_CrowdFund *CrowdFundFilterer) FilterUnPledge(opts *bind.FilterOpts, _id []*big.Int, caller []common.Address) (*CrowdFundUnPledgeIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "UnPledge", _idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdFundUnPledgeIterator{contract: _CrowdFund.contract, event: "UnPledge", logs: logs, sub: sub}, nil
}

// WatchUnPledge is a free log subscription operation binding the contract event 0xdc63160518c9c709fe39309a0e6f3d567887986a4625cd35090a6689e56e34c0.
//
// Solidity: event UnPledge(uint256 indexed _id, address indexed caller, uint256 _amount)
func (_CrowdFund *CrowdFundFilterer) WatchUnPledge(opts *bind.WatchOpts, sink chan<- *CrowdFundUnPledge, _id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "UnPledge", _idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundUnPledge)
				if err := _CrowdFund.contract.UnpackLog(event, "UnPledge", log); err != nil {
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

// ParseUnPledge is a log parse operation binding the contract event 0xdc63160518c9c709fe39309a0e6f3d567887986a4625cd35090a6689e56e34c0.
//
// Solidity: event UnPledge(uint256 indexed _id, address indexed caller, uint256 _amount)
func (_CrowdFund *CrowdFundFilterer) ParseUnPledge(log types.Log) (*CrowdFundUnPledge, error) {
	event := new(CrowdFundUnPledge)
	if err := _CrowdFund.contract.UnpackLog(event, "UnPledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
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

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
