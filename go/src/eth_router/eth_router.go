// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EthRouter

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

// EthRouterABI is the input ABI used to generate the binding from.
const EthRouterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addrs\",\"type\":\"address[]\"},{\"name\":\"_eth\",\"type\":\"uint256[]\"}],\"name\":\"sendEth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addrs\",\"type\":\"address[]\"},{\"name\":\"_ethPerAddress\",\"type\":\"uint256\"}],\"name\":\"sendEth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EthSent\",\"type\":\"event\"}]"

// EthRouterBin is the compiled bytecode used for deploying new contracts.
const EthRouterBin = `6060604052341561000f57600080fd5b6103cc8061001e6000396000f30060606040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166325245b268114610050578063feaf653d146100e8575b600080fd5b6100d460046024813581810190830135806020818102016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284375094965061012e95505050505050565b604051901515815260200160405180910390f35b6100d46004602481358181019083013580602081810201604051908101604052809392919081815260200183836020028082843750949650509335935061029292505050565b6000803481901161013e57600080fd5b5060005b835181101561028857600084828151811061015957fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff161415801561019c5750600083828151811061019057fe5b90602001906020020151115b15610280578381815181106101ad57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff166108fc8483815181106101dc57fe5b906020019060200201519081150290604051600060405180830381858888f19350505050151561020b57600080fd5b83818151811061021757fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff167f78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c84838151811061026457fe5b9060200190602002015160405190815260200160405180910390a25b600101610142565b5060019392505050565b600080828451023410156102a557600080fd5b5060005b83518110156102885760008482815181106102c057fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1614610398578381815181106102f157fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1683156108fc0284604051600060405180830381858888f19350505050151561033857600080fd5b83818151811061034457fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff167f78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c8460405190815260200160405180910390a25b6001016102a95600a165627a7a72305820665e28a3dbe0362c82902a41003226faea4fa6a4f7e5fc5b5515acd9385cd5640029`

// DeployEthRouter deploys a new Ethereum contract, binding an instance of EthRouter to it.
func DeployEthRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthRouter, error) {
	parsed, err := abi.JSON(strings.NewReader(EthRouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthRouter{EthRouterCaller: EthRouterCaller{contract: contract}, EthRouterTransactor: EthRouterTransactor{contract: contract}, EthRouterFilterer: EthRouterFilterer{contract: contract}}, nil
}

// EthRouter is an auto generated Go binding around an Ethereum contract.
type EthRouter struct {
	EthRouterCaller     // Read-only binding to the contract
	EthRouterTransactor // Write-only binding to the contract
	EthRouterFilterer   // Log filterer for contract events
}

// EthRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthRouterSession struct {
	Contract     *EthRouter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthRouterCallerSession struct {
	Contract *EthRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthRouterTransactorSession struct {
	Contract     *EthRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthRouterRaw struct {
	Contract *EthRouter // Generic contract binding to access the raw methods on
}

// EthRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthRouterCallerRaw struct {
	Contract *EthRouterCaller // Generic read-only contract binding to access the raw methods on
}

// EthRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthRouterTransactorRaw struct {
	Contract *EthRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthRouter creates a new instance of EthRouter, bound to a specific deployed contract.
func NewEthRouter(address common.Address, backend bind.ContractBackend) (*EthRouter, error) {
	contract, err := bindEthRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthRouter{EthRouterCaller: EthRouterCaller{contract: contract}, EthRouterTransactor: EthRouterTransactor{contract: contract}, EthRouterFilterer: EthRouterFilterer{contract: contract}}, nil
}

// NewEthRouterCaller creates a new read-only instance of EthRouter, bound to a specific deployed contract.
func NewEthRouterCaller(address common.Address, caller bind.ContractCaller) (*EthRouterCaller, error) {
	contract, err := bindEthRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthRouterCaller{contract: contract}, nil
}

// NewEthRouterTransactor creates a new write-only instance of EthRouter, bound to a specific deployed contract.
func NewEthRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*EthRouterTransactor, error) {
	contract, err := bindEthRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthRouterTransactor{contract: contract}, nil
}

// NewEthRouterFilterer creates a new log filterer instance of EthRouter, bound to a specific deployed contract.
func NewEthRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*EthRouterFilterer, error) {
	contract, err := bindEthRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthRouterFilterer{contract: contract}, nil
}

// bindEthRouter binds a generic wrapper to an already deployed contract.
func bindEthRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthRouter *EthRouterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthRouter.Contract.EthRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthRouter *EthRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthRouter.Contract.EthRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthRouter *EthRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthRouter.Contract.EthRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthRouter *EthRouterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthRouter *EthRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthRouter *EthRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthRouter.Contract.contract.Transact(opts, method, params...)
}

// SendEth is a paid mutator transaction binding the contract method 0xfeaf653d.
//
// Solidity: function sendEth(_addrs address[], _ethPerAddress uint256) returns(bool)
func (_EthRouter *EthRouterTransactor) SendEth(opts *bind.TransactOpts, _addrs []common.Address, _ethPerAddress *big.Int) (*types.Transaction, error) {
	return _EthRouter.contract.Transact(opts, "sendEth", _addrs, _ethPerAddress)
}

// SendEth is a paid mutator transaction binding the contract method 0xfeaf653d.
//
// Solidity: function sendEth(_addrs address[], _ethPerAddress uint256) returns(bool)
func (_EthRouter *EthRouterSession) SendEth(_addrs []common.Address, _ethPerAddress *big.Int) (*types.Transaction, error) {
	return _EthRouter.Contract.SendEth(&_EthRouter.TransactOpts, _addrs, _ethPerAddress)
}

// SendEth is a paid mutator transaction binding the contract method 0xfeaf653d.
//
// Solidity: function sendEth(_addrs address[], _ethPerAddress uint256) returns(bool)
func (_EthRouter *EthRouterTransactorSession) SendEth(_addrs []common.Address, _ethPerAddress *big.Int) (*types.Transaction, error) {
	return _EthRouter.Contract.SendEth(&_EthRouter.TransactOpts, _addrs, _ethPerAddress)
}

// EthRouterEthSentIterator is returned from FilterEthSent and is used to iterate over the raw logs and unpacked data for EthSent events raised by the EthRouter contract.
type EthRouterEthSentIterator struct {
	Event *EthRouterEthSent // Event containing the contract specifics and raw log

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
func (it *EthRouterEthSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthRouterEthSent)
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
		it.Event = new(EthRouterEthSent)
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
func (it *EthRouterEthSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthRouterEthSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthRouterEthSent represents a EthSent event raised by the EthRouter contract.
type EthRouterEthSent struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthSent is a free log retrieval operation binding the contract event 0x78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c.
//
// Solidity: event EthSent(_recipient indexed address, _amount uint256)
func (_EthRouter *EthRouterFilterer) FilterEthSent(opts *bind.FilterOpts, _recipient []common.Address) (*EthRouterEthSentIterator, error) {

	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _EthRouter.contract.FilterLogs(opts, "EthSent", _recipientRule)
	if err != nil {
		return nil, err
	}
	return &EthRouterEthSentIterator{contract: _EthRouter.contract, event: "EthSent", logs: logs, sub: sub}, nil
}

// WatchEthSent is a free log subscription operation binding the contract event 0x78f5cdad99320ec2ba57132d7dffb1d125775c823239e60ff5e9300fd4ac898c.
//
// Solidity: event EthSent(_recipient indexed address, _amount uint256)
func (_EthRouter *EthRouterFilterer) WatchEthSent(opts *bind.WatchOpts, sink chan<- *EthRouterEthSent, _recipient []common.Address) (event.Subscription, error) {

	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _EthRouter.contract.WatchLogs(opts, "EthSent", _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthRouterEthSent)
				if err := _EthRouter.contract.UnpackLog(event, "EthSent", log); err != nil {
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
