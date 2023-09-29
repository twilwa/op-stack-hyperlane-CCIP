// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const HyperlaneValidatorAnnounceStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/hyperlane/HyperlaneValidatorAnnounce.sol:HyperlaneValidatorAnnounce\",\"label\":\"validators\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_struct(AddressSet)1003_storage\"},{\"astId\":1001,\"contract\":\"contracts/hyperlane/HyperlaneValidatorAnnounce.sol:HyperlaneValidatorAnnounce\",\"label\":\"storageLocations\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_array(t_string_storage)dyn_storage)\"},{\"astId\":1002,\"contract\":\"contracts/hyperlane/HyperlaneValidatorAnnounce.sol:HyperlaneValidatorAnnounce\",\"label\":\"replayProtection\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_mapping(t_bytes32,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_bytes32)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"bytes32[]\",\"numberOfBytes\":\"32\",\"base\":\"t_bytes32\"},\"t_array(t_string_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"string[]\",\"numberOfBytes\":\"32\",\"base\":\"t_string_storage\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_array(t_string_storage)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e string[])\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_array(t_string_storage)dyn_storage\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_struct(AddressSet)1003_storage\":{\"encoding\":\"inplace\",\"label\":\"struct EnumerableSet.AddressSet\",\"numberOfBytes\":\"64\"},\"t_struct(Set)1004_storage\":{\"encoding\":\"inplace\",\"label\":\"struct EnumerableSet.Set\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var HyperlaneValidatorAnnounceStorageLayout = new(solc.StorageLayout)

var HyperlaneValidatorAnnounceDeployedBin = "0x608060405234801561001057600080fd5b50600436106100675760003560e01c8063690cb78611610050578063690cb786146100b45780638d3638f4146100c9578063d5438eae1461010557600080fd5b806321f717811461006c57806351abe7cc14610094575b600080fd5b61007f61007a366004610cbf565b610151565b60405190151581526020015b60405180910390f35b6100a76100a2366004610d40565b61046c565b60405161008b9190610de5565b6100bc610639565b60405161008b9190610ee6565b6100f07f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161008b565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161008b565b60008086868660405160200161016993929190610f34565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915060ff161561021e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f7265706c6179000000000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b600081815260036020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558051601f89018390048302810183019091528781526102d5917f0000000000000000000000000000000000000000000000000000000000000000917f0000000000000000000000000000000000000000000000000000000000000000918b908b908190840183828082843760009201919091525061064a92505050565b905060006103198287878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061077892505050565b90508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146103b0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f217369676e6174757265000000000000000000000000000000000000000000006044820152606401610215565b6103bb60008a610794565b6103cc576103ca60008a6107c6565b505b73ffffffffffffffffffffffffffffffffffffffff891660009081526002602090815260408220805460018101825590835291200161040c888a83611044565b508873ffffffffffffffffffffffffffffffffffffffff167f78066d8adb677a1353d1fc8be28cf03e2a8de7157bbab979953587d78076c11e898960405161045592919061115f565b60405180910390a250600198975050505050505050565b606060008267ffffffffffffffff81111561048957610489610f73565b6040519080825280602002602001820160405280156104bc57816020015b60608152602001906001900390816104a75790505b50905060005b8381101561062f57600260008686848181106104e0576104e06111ac565b90506020020160208101906104f591906111db565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156105fb57838290600052602060002001805461056e90610fa2565b80601f016020809104026020016040519081016040528092919081815260200182805461059a90610fa2565b80156105e75780601f106105bc576101008083540402835291602001916105e7565b820191906000526020600020905b8154815290600101906020018083116105ca57829003601f168201915b50505050508152602001906001019061054f565b50505050828281518110610611576106116111ac565b60200260200101819052508080610627906111f6565b9150506104c2565b5090505b92915050565b606061064560006107e8565b905090565b6000808373ffffffffffffffffffffffffffffffffffffffff861660405160e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016602083015260248201527f48595045524c414e455f414e4e4f554e43454d454e54000000000000000000006044820152605a0160405160208183030381529060405280519060200120905061076f81846040516020016106f1929190611255565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201207f19457468657265756d205369676e6564204d6573736167653a0a33320000000084830152603c8085019190915282518085039091018152605c909301909152815191012090565b95945050505050565b600080600061078785856107f5565b9150915061062f8161083a565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415155b9392505050565b60006107bf8373ffffffffffffffffffffffffffffffffffffffff8416610a91565b606060006107bf83610ae0565b600080825160410361082b5760208301516040840151606085015160001a61081f87828585610b3c565b94509450505050610833565b506000905060025b9250929050565b600081600481111561084e5761084e61127b565b036108565750565b600181600481111561086a5761086a61127b565b036108d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610215565b60028160048111156108e5576108e561127b565b0361094c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610215565b60038160048111156109605761096061127b565b036109ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610215565b6004816004811115610a0157610a0161127b565b03610a8e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610215565b50565b6000818152600183016020526040812054610ad857508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610633565b506000610633565b606081600001805480602002602001604051908101604052809291908181526020018280548015610b3057602002820191906000526020600020905b815481526020019060010190808311610b1c575b50505050509050919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610b735750600090506003610c4b565b8460ff16601b14158015610b8b57508460ff16601c14155b15610b9c5750600090506004610c4b565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610bf0573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116610c4457600060019250925050610c4b565b9150600090505b94509492505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c7857600080fd5b919050565b60008083601f840112610c8f57600080fd5b50813567ffffffffffffffff811115610ca757600080fd5b60208301915083602082850101111561083357600080fd5b600080600080600060608688031215610cd757600080fd5b610ce086610c54565b9450602086013567ffffffffffffffff80821115610cfd57600080fd5b610d0989838a01610c7d565b90965094506040880135915080821115610d2257600080fd5b50610d2f88828901610c7d565b969995985093965092949392505050565b60008060208385031215610d5357600080fd5b823567ffffffffffffffff80821115610d6b57600080fd5b818501915085601f830112610d7f57600080fd5b813581811115610d8e57600080fd5b8660208260051b8501011115610da357600080fd5b60209290920196919550909350505050565b60005b83811015610dd0578181015183820152602001610db8565b83811115610ddf576000848401525b50505050565b60208152600060208201835180825260408401915060408160051b8501016020860160005b83811015610eda578683037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018552815180518085526020918201918086019190600582901b87010160005b82811015610ec0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08089840301855285518051808552610e9e816020870160208501610db5565b60209788019796870196601f9190910190921693909301019150600101610e56565b506020988901989096509490940193505050600101610e0a565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b81811015610eda57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101610f02565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008460601b168152818360148301376000910160140190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c90821680610fb657607f821691505b602082108103610fef577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f82111561103f57600081815260208120601f850160051c8101602086101561101c5750805b601f850160051c820191505b8181101561103b57828155600101611028565b5050505b505050565b67ffffffffffffffff83111561105c5761105c610f73565b6110708361106a8354610fa2565b83610ff5565b6000601f8411600181146110c2576000851561108c5750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355611158565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b8281101561111157868501358255602094850194600190920191016110f1565b508682101561114c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156111ed57600080fd5b6107bf82610c54565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361124e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b8281526000825161126d816020850160208701610db5565b919091016020019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(HyperlaneValidatorAnnounceStorageLayoutJSON), HyperlaneValidatorAnnounceStorageLayout); err != nil {
		panic(err)
	}

	layouts["HyperlaneValidatorAnnounce"] = HyperlaneValidatorAnnounceStorageLayout
	deployedBytecodes["HyperlaneValidatorAnnounce"] = HyperlaneValidatorAnnounceDeployedBin
}
