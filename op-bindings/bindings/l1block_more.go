// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1BlockStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"number\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1001,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"timestamp\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1002,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"basefee\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"hash\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_bytes32\"},{\"astId\":1004,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"sequenceNumber\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint64\"},{\"astId\":1005,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_bytes32\"},{\"astId\":1006,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"l1FeeOverhead\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"l1FeeScalar\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_uint256\"}],\"types\":{\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"}}}"

var L1BlockStorageLayout = new(solc.StorageLayout)

var L1BlockDeployedBin = "0x608060405234801561001057600080fd5b50600436106100c95760003560e01c80638381f58a11610081578063b80777ea1161005b578063b80777ea14610170578063e591b28214610190578063e81b2c6d146101d057600080fd5b80638381f58a1461014a5780638b239f731461015e5780639e8c49661461016757600080fd5b806354fd4d50116100b257806354fd4d50146100ff5780635cf249691461011457806364ca23ef1461011d57600080fd5b8063015d8eb9146100ce57806309bd5a60146100e3575b600080fd5b6100e16100dc366004610515565b6101d9565b005b6100ec60025481565b6040519081526020015b60405180910390f35b610107610318565b6040516100f691906105b7565b6100ec60015481565b6003546101319067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016100f6565b6000546101319067ffffffffffffffff1681565b6100ec60055481565b6100ec60065481565b6000546101319068010000000000000000900467ffffffffffffffff1681565b6101ab73deaddeaddeaddeaddeaddeaddeaddeaddead000181565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f6565b6100ec60045481565b3373deaddeaddeaddeaddeaddeaddeaddeaddead000114610280576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603b60248201527f4c31426c6f636b3a206f6e6c7920746865206465706f7369746f72206163636f60448201527f756e742063616e20736574204c3120626c6f636b2076616c7565730000000000606482015260840160405180910390fd5b6000805467ffffffffffffffff98891668010000000000000000027fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116998916999099179890981790975560019490945560029290925560038054919094167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009190911617909255600491909155600555600655565b60606103437f00000000000000000000000000000000000000000000000000000000000000006103bb565b61036c7f00000000000000000000000000000000000000000000000000000000000000006103bb565b6103957f00000000000000000000000000000000000000000000000000000000000000006103bb565b6040516020016103a793929190610608565b604051602081830303815290604052905090565b6060816000036103fe57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156104285780610412816106ad565b91506104219050600a83610714565b9150610402565b60008167ffffffffffffffff81111561044357610443610728565b6040519080825280601f01601f19166020018201604052801561046d576020820181803683370190505b5090505b84156104f057610482600183610757565b915061048f600a8661076e565b61049a906030610782565b60f81b8183815181106104af576104af61079a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506104e9600a86610714565b9450610471565b949350505050565b803567ffffffffffffffff8116811461051057600080fd5b919050565b600080600080600080600080610100898b03121561053257600080fd5b61053b896104f8565b975061054960208a016104f8565b9650604089013595506060890135945061056560808a016104f8565b979a969950949793969560a0850135955060c08501359460e001359350915050565b60005b838110156105a257818101518382015260200161058a565b838111156105b1576000848401525b50505050565b60208152600082518060208401526105d6816040850160208701610587565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b6000845161061a818460208901610587565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610656816001850160208a01610587565b60019201918201528351610671816002840160208801610587565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036106de576106de61067e565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610723576107236106e5565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156107695761076961067e565b500390565b60008261077d5761077d6106e5565b500690565b600082198211156107955761079561067e565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1BlockStorageLayoutJSON), L1BlockStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1Block"] = L1BlockStorageLayout
	deployedBytecodes["L1Block"] = L1BlockDeployedBin
}
