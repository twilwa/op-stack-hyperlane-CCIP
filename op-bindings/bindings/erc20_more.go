// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const ERC20StorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol:ERC20\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1001,\"contract\":\"node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol:ERC20\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1002,\"contract\":\"node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol:ERC20\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol:ERC20\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1004,\"contract\":\"node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol:ERC20\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var ERC20StorageLayout = new(solc.StorageLayout)

var ERC20DeployedBin = "0x608060405234801561001057600080fd5b50600436106100c95760003560e01c80633950935111610081578063a457c2d71161005b578063a457c2d714610194578063a9059cbb146101a7578063dd62ed3e146101ba57600080fd5b8063395093511461014357806370a082311461015657806395d89b411461018c57600080fd5b806318160ddd116100b257806318160ddd1461010f57806323b872dd14610121578063313ce5671461013457600080fd5b806306fdde03146100ce578063095ea7b3146100ec575b600080fd5b6100d6610200565b6040516100e3919061094a565b60405180910390f35b6100ff6100fa3660046109e6565b610292565b60405190151581526020016100e3565b6002545b6040519081526020016100e3565b6100ff61012f366004610a10565b6102aa565b604051601281526020016100e3565b6100ff6101513660046109e6565b6102ce565b610113610164366004610a4c565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100d661031a565b6100ff6101a23660046109e6565b610329565b6100ff6101b53660046109e6565b6103ff565b6101136101c8366004610a6e565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461020f90610aa1565b80601f016020809104026020016040519081016040528092919081815260200182805461023b90610aa1565b80156102885780601f1061025d57610100808354040283529160200191610288565b820191906000526020600020905b81548152906001019060200180831161026b57829003601f168201915b5050505050905090565b6000336102a081858561040d565b5060019392505050565b6000336102b88582856105c0565b6102c3858585610697565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906102a09082908690610315908790610af4565b61040d565b60606004805461020f90610aa1565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156103f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102c3828686840361040d565b6000336102a0818585610697565b73ffffffffffffffffffffffffffffffffffffffff83166104af576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff8216610552576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146106915781811015610684576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103e9565b610691848484840361040d565b50505050565b73ffffffffffffffffffffffffffffffffffffffff831661073a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff82166107dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610893576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152602081905260408082208585039055918516815290812080548492906108d7908490610af4565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161093d91815260200190565b60405180910390a3610691565b600060208083528351808285015260005b818110156109775785810183015185820160400152820161095b565b81811115610989576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146109e157600080fd5b919050565b600080604083850312156109f957600080fd5b610a02836109bd565b946020939093013593505050565b600080600060608486031215610a2557600080fd5b610a2e846109bd565b9250610a3c602085016109bd565b9150604084013590509250925092565b600060208284031215610a5e57600080fd5b610a67826109bd565b9392505050565b60008060408385031215610a8157600080fd5b610a8a836109bd565b9150610a98602084016109bd565b90509250929050565b600181811c90821680610ab557607f821691505b602082108103610aee577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60008219821115610b2e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50019056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(ERC20StorageLayoutJSON), ERC20StorageLayout); err != nil {
		panic(err)
	}

	layouts["ERC20"] = ERC20StorageLayout
	deployedBytecodes["ERC20"] = ERC20DeployedBin
}
