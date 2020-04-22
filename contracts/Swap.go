// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// SwapABI is the input ABI used to generate the binding from.
const SwapABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_dagAddress\",\"type\":\"string\"}],\"name\":\"mapAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ethAddress\",\"type\":\"address\"}],\"name\":\"removeMappedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getKeyList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMappedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"mappedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SwapBin is the compiled bytecode used for deploying new contracts.
var SwapBin = "0x608060405234801561001057600080fd5b50600436106100935760003560e01c80638f32d59b116100665780638f32d59b146100f8578063d73683c114610116578063eab3c92d14610134578063f2fde38b14610152578063f7ae48c41461016e57610093565b8063401105441461009857806354d3cf5d146100b4578063715018a6146100d05780638da5cb5b146100da575b600080fd5b6100b260048036036100ad9190810190610be0565b61019e565b005b6100ce60048036036100c99190810190610c09565b610285565b005b6100d86103b6565b005b6100e26104bc565b6040516100ef9190610e86565b60405180910390f35b6101006104e5565b60405161010d9190610ee5565b60405180910390f35b61011e610543565b60405161012b9190610ec3565b60405180910390f35b61013c6106d0565b6040516101499190610ea1565b60405180910390f35b61016c60048036036101679190810190610be0565b6107b0565b005b61018860048036036101839190810190610be0565b610803565b6040516101959190610f00565b60405180910390f35b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546001816001161561010002031660029004905014156101ff57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461023757600080fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006102829190610a8b565b50565b60003390506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805460018160011615610100020316600290049050146102ea57600080fd5b6102f3816108b3565b61035e5760028190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b81600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090805190602001906103b1929190610ad3565b505050565b6103be6104e5565b6103fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f490610f42565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610527610955565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b60608060028054905060405190808252806020026020018201604052801561057f57816020015b606081526020019060019003908161056a5790505b50905060008090505b6002805490508110156106c85760016000600283815481106105a657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561069f5780601f106106745761010080835404028352916020019161069f565b820191906000526020600020905b81548152906001019060200180831161068257829003601f168201915b50505050508282815181106106b057fe5b60200260200101819052508080600101915050610588565b508091505090565b6060806002805490506040519080825280602002602001820160405280156107075781602001602082028038833980820191505090505b50905060008090505b6002805490508110156107a8576002818154811061072a57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682828151811061076157fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610710565b508091505090565b6107b86104e5565b6107f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ee90610f42565b60405180910390fd5b6108008161095d565b50565b60016020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108ab5780601f10610880576101008083540402835291602001916108ab565b820191906000526020600020905b81548152906001019060200180831161088e57829003601f168201915b505050505081565b600080600090505b60028054905081101561094a578273ffffffffffffffffffffffffffffffffffffffff16600282815481106108ec57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561093d576001915050610950565b80806001019150506108bb565b50600090505b919050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156109cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c490610f22565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b50805460018160011615610100020316600290046000825580601f10610ab15750610ad0565b601f016020900490600052602060002090810190610acf9190610b53565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b1457805160ff1916838001178555610b42565b82800160010185558215610b42579182015b82811115610b41578251825591602001919060010190610b26565b5b509050610b4f9190610b53565b5090565b610b7591905b80821115610b71576000816000905550600101610b59565b5090565b90565b6000610b848235611092565b905092915050565b600082601f830112610b9d57600080fd5b8135610bb0610bab82610f8f565b610f62565b91508082526020830160208301858383011115610bcc57600080fd5b610bd78382846110c4565b50505092915050565b600060208284031215610bf257600080fd5b6000610c0084828501610b78565b91505092915050565b600060208284031215610c1b57600080fd5b600082013567ffffffffffffffff811115610c3557600080fd5b610c4184828501610b8c565b91505092915050565b6000610c568383610c76565b60208301905092915050565b6000610c6e8383610d6e565b905092915050565b610c7f81611054565b82525050565b610c8e81611054565b82525050565b6000610c9f82610fd5565b610ca98185611010565b9350610cb483610fbb565b60005b82811015610ce257610cca868351610c4a565b9550610cd582610ff6565b9150600181019050610cb7565b50849250505092915050565b6000610cf982610fe0565b610d038185611021565b935083602082028501610d1585610fc8565b60005b84811015610d4e578383038852610d30838351610c62565b9250610d3b82611003565b9150602088019750600181019050610d18565b508196508694505050505092915050565b610d6881611066565b82525050565b6000610d7982610feb565b610d838185611032565b9350610d938185602086016110d3565b610d9c81611106565b840191505092915050565b6000610db282610feb565b610dbc8185611043565b9350610dcc8185602086016110d3565b610dd581611106565b840191505092915050565b6000610ded602683611043565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610e53602083611043565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000602082019050610e9b6000830184610c85565b92915050565b60006020820190508181036000830152610ebb8184610c94565b905092915050565b60006020820190508181036000830152610edd8184610cee565b905092915050565b6000602082019050610efa6000830184610d5f565b92915050565b60006020820190508181036000830152610f1a8184610da7565b905092915050565b60006020820190508181036000830152610f3b81610de0565b9050919050565b60006020820190508181036000830152610f5b81610e46565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715610f8557600080fd5b8060405250919050565b600067ffffffffffffffff821115610fa657600080fd5b601f19601f8301169050602081019050919050565b6000602082019050919050565b6000602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061105f82611072565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061109d826110a4565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b82818337600083830152505050565b60005b838110156110f15780820151818401526020810190506110d6565b83811115611100576000848401525b50505050565b6000601f19601f830116905091905056fea265627a7a72305820c360b34ff259c3480da8e7458e9ae12d3ff59369cdd7d1025e1bf9a24f7582766c6578706572696d656e74616cf50037"

// DeploySwap deploys a new Ethereum contract, binding an instance of Swap to it.
func DeploySwap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swap, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// Swap is an auto generated Go binding around an Ethereum contract.
type Swap struct {
	SwapCaller     // Read-only binding to the contract
	SwapTransactor // Write-only binding to the contract
	SwapFilterer   // Log filterer for contract events
}

// SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapSession struct {
	Contract     *Swap             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapCallerSession struct {
	Contract *SwapCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTransactorSession struct {
	Contract     *SwapTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRaw struct {
	Contract *Swap // Generic contract binding to access the raw methods on
}

// SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapCallerRaw struct {
	Contract *SwapCaller // Generic read-only contract binding to access the raw methods on
}

// SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTransactorRaw struct {
	Contract *SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwap creates a new instance of Swap, bound to a specific deployed contract.
func NewSwap(address common.Address, backend bind.ContractBackend) (*Swap, error) {
	contract, err := bindSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// NewSwapCaller creates a new read-only instance of Swap, bound to a specific deployed contract.
func NewSwapCaller(address common.Address, caller bind.ContractCaller) (*SwapCaller, error) {
	contract, err := bindSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapCaller{contract: contract}, nil
}

// NewSwapTransactor creates a new write-only instance of Swap, bound to a specific deployed contract.
func NewSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapTransactor, error) {
	contract, err := bindSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTransactor{contract: contract}, nil
}

// NewSwapFilterer creates a new log filterer instance of Swap, bound to a specific deployed contract.
func NewSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFilterer, error) {
	contract, err := bindSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFilterer{contract: contract}, nil
}

// bindSwap binds a generic wrapper to an already deployed contract.
func bindSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transact(opts, method, params...)
}

// GetKeyList is a free data retrieval call binding the contract method 0xeab3c92d.
//
// Solidity: function getKeyList() view returns(address[])
func (_Swap *SwapCaller) GetKeyList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Swap.contract.Call(opts, out, "getKeyList")
	return *ret0, err
}

// GetKeyList is a free data retrieval call binding the contract method 0xeab3c92d.
//
// Solidity: function getKeyList() view returns(address[])
func (_Swap *SwapSession) GetKeyList() ([]common.Address, error) {
	return _Swap.Contract.GetKeyList(&_Swap.CallOpts)
}

// GetKeyList is a free data retrieval call binding the contract method 0xeab3c92d.
//
// Solidity: function getKeyList() view returns(address[])
func (_Swap *SwapCallerSession) GetKeyList() ([]common.Address, error) {
	return _Swap.Contract.GetKeyList(&_Swap.CallOpts)
}

// GetMappedAddresses is a free data retrieval call binding the contract method 0xd73683c1.
//
// Solidity: function getMappedAddresses() view returns(string[])
func (_Swap *SwapCaller) GetMappedAddresses(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Swap.contract.Call(opts, out, "getMappedAddresses")
	return *ret0, err
}

// GetMappedAddresses is a free data retrieval call binding the contract method 0xd73683c1.
//
// Solidity: function getMappedAddresses() view returns(string[])
func (_Swap *SwapSession) GetMappedAddresses() ([]string, error) {
	return _Swap.Contract.GetMappedAddresses(&_Swap.CallOpts)
}

// GetMappedAddresses is a free data retrieval call binding the contract method 0xd73683c1.
//
// Solidity: function getMappedAddresses() view returns(string[])
func (_Swap *SwapCallerSession) GetMappedAddresses() ([]string, error) {
	return _Swap.Contract.GetMappedAddresses(&_Swap.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Swap *SwapCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Swap.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Swap *SwapSession) IsOwner() (bool, error) {
	return _Swap.Contract.IsOwner(&_Swap.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Swap *SwapCallerSession) IsOwner() (bool, error) {
	return _Swap.Contract.IsOwner(&_Swap.CallOpts)
}

// MappedAddresses is a free data retrieval call binding the contract method 0xf7ae48c4.
//
// Solidity: function mappedAddresses(address ) view returns(string)
func (_Swap *SwapCaller) MappedAddresses(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Swap.contract.Call(opts, out, "mappedAddresses", arg0)
	return *ret0, err
}

// MappedAddresses is a free data retrieval call binding the contract method 0xf7ae48c4.
//
// Solidity: function mappedAddresses(address ) view returns(string)
func (_Swap *SwapSession) MappedAddresses(arg0 common.Address) (string, error) {
	return _Swap.Contract.MappedAddresses(&_Swap.CallOpts, arg0)
}

// MappedAddresses is a free data retrieval call binding the contract method 0xf7ae48c4.
//
// Solidity: function mappedAddresses(address ) view returns(string)
func (_Swap *SwapCallerSession) MappedAddresses(arg0 common.Address) (string, error) {
	return _Swap.Contract.MappedAddresses(&_Swap.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Swap.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapSession) Owner() (common.Address, error) {
	return _Swap.Contract.Owner(&_Swap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapCallerSession) Owner() (common.Address, error) {
	return _Swap.Contract.Owner(&_Swap.CallOpts)
}

// MapAddress is a paid mutator transaction binding the contract method 0x54d3cf5d.
//
// Solidity: function mapAddress(string _dagAddress) returns()
func (_Swap *SwapTransactor) MapAddress(opts *bind.TransactOpts, _dagAddress string) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "mapAddress", _dagAddress)
}

// MapAddress is a paid mutator transaction binding the contract method 0x54d3cf5d.
//
// Solidity: function mapAddress(string _dagAddress) returns()
func (_Swap *SwapSession) MapAddress(_dagAddress string) (*types.Transaction, error) {
	return _Swap.Contract.MapAddress(&_Swap.TransactOpts, _dagAddress)
}

// MapAddress is a paid mutator transaction binding the contract method 0x54d3cf5d.
//
// Solidity: function mapAddress(string _dagAddress) returns()
func (_Swap *SwapTransactorSession) MapAddress(_dagAddress string) (*types.Transaction, error) {
	return _Swap.Contract.MapAddress(&_Swap.TransactOpts, _dagAddress)
}

// RemoveMappedAddress is a paid mutator transaction binding the contract method 0x40110544.
//
// Solidity: function removeMappedAddress(address _ethAddress) returns()
func (_Swap *SwapTransactor) RemoveMappedAddress(opts *bind.TransactOpts, _ethAddress common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "removeMappedAddress", _ethAddress)
}

// RemoveMappedAddress is a paid mutator transaction binding the contract method 0x40110544.
//
// Solidity: function removeMappedAddress(address _ethAddress) returns()
func (_Swap *SwapSession) RemoveMappedAddress(_ethAddress common.Address) (*types.Transaction, error) {
	return _Swap.Contract.RemoveMappedAddress(&_Swap.TransactOpts, _ethAddress)
}

// RemoveMappedAddress is a paid mutator transaction binding the contract method 0x40110544.
//
// Solidity: function removeMappedAddress(address _ethAddress) returns()
func (_Swap *SwapTransactorSession) RemoveMappedAddress(_ethAddress common.Address) (*types.Transaction, error) {
	return _Swap.Contract.RemoveMappedAddress(&_Swap.TransactOpts, _ethAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swap *SwapTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swap *SwapSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swap.Contract.RenounceOwnership(&_Swap.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swap *SwapTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swap.Contract.RenounceOwnership(&_Swap.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swap *SwapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swap *SwapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swap.Contract.TransferOwnership(&_Swap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swap *SwapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swap.Contract.TransferOwnership(&_Swap.TransactOpts, newOwner)
}

// SwapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Swap contract.
type SwapOwnershipTransferredIterator struct {
	Event *SwapOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapOwnershipTransferred)
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
		it.Event = new(SwapOwnershipTransferred)
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
func (it *SwapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapOwnershipTransferred represents a OwnershipTransferred event raised by the Swap contract.
type SwapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swap *SwapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapOwnershipTransferredIterator{contract: _Swap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swap *SwapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapOwnershipTransferred)
				if err := _Swap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swap *SwapFilterer) ParseOwnershipTransferred(log types.Log) (*SwapOwnershipTransferred, error) {
	event := new(SwapOwnershipTransferred)
	if err := _Swap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
