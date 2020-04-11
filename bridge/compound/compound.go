// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compound

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

// CompoundABI is the input ABI used to generate the binding from.
const CompoundABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_compoundAgentLogic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"UpdateAgentLogic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"UpdateVaultCompound\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"agents\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compoundAgentLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"srcTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"executeMulti\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"getAgentAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"sigToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_compoundAgentLogic\",\"type\":\"address\"}],\"name\":\"updateAgentLogic\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"name\":\"updateVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CompoundBin is the compiled bytecode used for deploying new contracts.
var CompoundBin = "0x608060405234801561001057600080fd5b5060405161263c38038061263c8339818101604052606081101561003357600080fd5b81019080805190602001909291908051906020019092919080519060200190929190505050826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506125108061012c6000396000f3fe6080604052600436106100c25760003560e01c8063b42a644b1161007f578063ea1efaf411610059578063ea1efaf414610a14578063f851a44014610a65578063fd66091e14610abc578063fee8a5c114610b4d576100c2565b8063b42a644b14610919578063e4bd707414610970578063e7563f3f146109c3576100c2565b80631ea1940e146100c457806324d536e31461011757806327d54a92146104c55780633fec6b401461055657806351f120371461066857806372e94bf6146108c2575b005b3480156100d057600080fd5b506100fd600480360360208110156100e757600080fd5b8101908080359060200190929190505050610ba4565b604051808215151515815260200191505060405180910390f35b610426600480360360a081101561012d57600080fd5b810190808035906020019064010000000081111561014a57600080fd5b82018360208201111561015c57600080fd5b8035906020019184602083028401116401000000008311171561017e57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156101de57600080fd5b8201836020820111156101f057600080fd5b8035906020019184602083028401116401000000008311171561021257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561027257600080fd5b82018360208201111561028457600080fd5b803590602001918460018302840111640100000000831117156102a657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561030957600080fd5b82018360208201111561031b57600080fd5b8035906020019184600183028401116401000000008311171561033d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156103a057600080fd5b8201836020820111156103b257600080fd5b803590602001918460018302840111640100000000831117156103d457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610bc4565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561046d578082015181840152602081019050610452565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156104af578082015181840152602081019050610494565b5050505090500194505050505060405180910390f35b3480156104d157600080fd5b50610514600480360360208110156104e857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061112b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561056257600080fd5b506106266004803603604081101561057957600080fd5b810190808035906020019064010000000081111561059657600080fd5b8201836020820111156105a857600080fd5b803590602001918460018302840111640100000000831117156105ca57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050611194565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610879600480360360a081101561067e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156106c557600080fd5b8201836020820111156106d757600080fd5b803590602001918460018302840111640100000000831117156106f957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561075c57600080fd5b82018360208201111561076e57600080fd5b8035906020019184600183028401116401000000008311171561079057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156107f357600080fd5b82018360208201111561080557600080fd5b8035906020019184600183028401116401000000008311171561082757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061123a565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b3480156108ce57600080fd5b506108d7611662565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561092557600080fd5b5061092e611667565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561097c57600080fd5b506109a96004803603602081101561099357600080fd5b810190808035906020019092919050505061168c565b604051808215151515815260200191505060405180910390f35b3480156109cf57600080fd5b50610a12600480360360208110156109e657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116c7565b005b348015610a2057600080fd5b50610a6360048036036020811015610a3757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611813565b005b348015610a7157600080fd5b50610a7a611960565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610ac857600080fd5b50610b0b60048036036020811015610adf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611986565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610b5957600080fd5b50610b626119b9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60016020528060005260406000206000915054906101000a900460ff1681565b6060806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c2057600080fd5b8551875114610c2e57600080fd5b6000610cf786866040516020018083805190602001908083835b60208310610c6b5780518252602082019150602081019050602083039250610c48565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b60208310610cbc5780518252602082019150602081019050602083039250610c99565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052856119df565b90506000610d0482611a43565b905060008090505b8951811015610e3857600073ffffffffffffffffffffffffffffffffffffffff168a8281518110610d3957fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614610e2b57898181518110610d6857fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb838b8481518110610d9857fe5b60200260200101516040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610e0957600080fd5b505af1158015610e1d573d6000803e3d6000fd5b50505050610e29611c73565b505b8080600101915050610d0c565b50600060608273ffffffffffffffffffffffffffffffffffffffff16348a6040518082805190602001908083835b60208310610e895780518252602082019150602081019050602083039250610e66565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114610eeb576040519150601f19603f3d011682016040523d82523d6000602084013e610ef0565b606091505b509150915081610eff57600080fd5b6060818060200190516020811015610f1657600080fd5b8101908080516040519392919084640100000000821115610f3657600080fd5b83820191506020820185811115610f4c57600080fd5b8251866001820283011164010000000082111715610f6957600080fd5b8083526020830192505050908051906020019080838360005b83811015610f9d578082015181840152602081019050610f82565b50505050905090810190601f168015610fca5780820380516001836020036101000a031916815260200191505b506040525050509050808060200190516040811015610fe857600080fd5b810190808051604051939291908464010000000082111561100857600080fd5b8382019150602082018581111561101e57600080fd5b825186602082028301116401000000008211171561103b57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015611072578082015181840152602081019050611057565b505050509050016040526020018051604051939291908464010000000082111561109b57600080fd5b838201915060208201858111156110b157600080fd5b82518660208202830111640100000000821117156110ce57600080fd5b8083526020830192505050908051906020019060200280838360005b838110156111055780820151818401526020810190506110ea565b505050509050016040525050508191508090509650965050505050509550959350505050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000806000806020860151915060408601519250601b866040815181106111b757fe5b602001015160f81c60f81b60f81c01905060018582848660405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611225573d6000803e3d6000fd5b50505060206040510351935050505092915050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461129657600080fd5b6000611396888787604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140183805190602001908083835b6020831061130957805182526020820191506020810190506020830392506112e6565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b6020831061135a5780518252602082019150602081019050602083039250611337565b6001836020036101000a0380198251168184511680821785525050505050509050019350505050604051602081830303815290604052856119df565b905060006113a382611a43565b9050600073ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614611482578873ffffffffffffffffffffffffffffffffffffffff1663a9059cbb828a6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561146057600080fd5b505af1158015611474573d6000803e3d6000fd5b50505050611480611c73565b505b600060608273ffffffffffffffffffffffffffffffffffffffff16348a6040518082805190602001908083835b602083106114d257805182526020820191506020810190506020830392506114af565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611534576040519150601f19603f3d011682016040523d82523d6000602084013e611539565b606091505b50915091508161154857600080fd5b606081806020019051602081101561155f57600080fd5b810190808051604051939291908464010000000082111561157f57600080fd5b8382019150602082018581111561159557600080fd5b82518660018202830111640100000000821117156115b257600080fd5b8083526020830192505050908051906020019080838360005b838110156115e65780820151818401526020810190506115cb565b50505050905090810190601f1680156116135780820380516001836020036101000a031916815260200191505b50604052505050905080806020019051604081101561163157600080fd5b8101908080519060200190929190805190602001909291905050508191509650965050505050509550959350505050565b600081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006001600083815260200190815260200160002060009054906101000a900460ff16156116bd57600190506116c2565b600090505b919050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461176d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124ba6022913960400191505060405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f856bb47bb5dc3e7d1f5c5d594120747b26fb9b49f9a530679772ac961d7c58a481604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146118b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806124ba6022913960400191505060405180910390fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f1d7094459f506b06afa7f2c3bec49b9073c63236bb41bcb97513be10f0ff632781604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080838051906020012090506119f58161168c565b156119ff57600080fd5b6000611a0b8483611194565b9050600180600084815260200190815260200160002060006101000a81548160ff021916908315150217905550809250505092915050565b60008073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611c0c5730600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051611b0890611cb1565b808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050604051809103906000f080158015611b8d573d6000803e3d6000fd5b50600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600080600090503d60008114611c905760208114611c9957611ca5565b60019150611ca5565b60206000803e60005191505b50600081141591505090565b6107fb80611cbf8339019056fe608060405234801561001057600080fd5b506040516107fb3803806107fb8339818101604052604081101561003357600080fd5b810190808051906020019092919080519060200190929190505050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061071b806100e06000396000f3fe60806040526004361061004a5760003560e01c8063370f81fe1461004c5780639e955149146100a3578063da67bcc41461013b578063ea1efaf41461022d578063fee8a5c1146102c5575b005b34801561005857600080fd5b5061006161031c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100af57600080fd5b506100f2600480360360208110156100c657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610341565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b6101b26004803603602081101561015157600080fd5b810190808035906020019064010000000081111561016e57600080fd5b82018360208201111561018057600080fd5b803590602001918460018302840111640100000000831117156101a257600080fd5b909192939192939050505061043b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101f25780820151818401526020810190506101d7565b50505050905090810190601f16801561021f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561023957600080fd5b5061027c6004803603602081101561025057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610596565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b3480156102d157600080fd5b506102da610691565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001806106b8602f913960400191505060405180910390fd5b826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008081915080905091509150915091565b60606000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001806106b8602f913960400191505060405180910390fd5b60006060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168585604051808383808284378083019250505092505050600060405180830381855af49150503d8060008114610571576040519150601f19603f3d011682016040523d82523d6000602084013e610576565b606091505b5091509150600082141561058b573d60208201fd5b809250505092915050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461063e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001806106b8602f913960400191505060405180910390fd5b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008081915080905091509150915091565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fe4f6e6c792070726f787920636f6d706f756e6420616c6c6f77656420746f2063616c6c2074686973206167656e7421a265627a7a723158208ca4ba3daf721213ab2951c45a2192f970b8539dd4d2ae8d2e984744a49a2dc564736f6c634300051000324f6e6c792061646d696e2063616e2063616c6c20746869732066756e6374696f6e21a265627a7a723158201f35028da8a61e47c1f0b32bb83560a9f44ced1f421ae587cae2aa7f0200074564736f6c63430005100032"

// DeployCompound deploys a new Ethereum contract, binding an instance of Compound to it.
func DeployCompound(auth *bind.TransactOpts, backend bind.ContractBackend, _incognitoSmartContract common.Address, _compoundAgentLogic common.Address, _admin common.Address) (common.Address, *types.Transaction, *Compound, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompoundBin), backend, _incognitoSmartContract, _compoundAgentLogic, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Compound{CompoundCaller: CompoundCaller{contract: contract}, CompoundTransactor: CompoundTransactor{contract: contract}, CompoundFilterer: CompoundFilterer{contract: contract}}, nil
}

// Compound is an auto generated Go binding around an Ethereum contract.
type Compound struct {
	CompoundCaller     // Read-only binding to the contract
	CompoundTransactor // Write-only binding to the contract
	CompoundFilterer   // Log filterer for contract events
}

// CompoundCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundSession struct {
	Contract     *Compound         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompoundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundCallerSession struct {
	Contract *CompoundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CompoundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundTransactorSession struct {
	Contract     *CompoundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CompoundRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundRaw struct {
	Contract *Compound // Generic contract binding to access the raw methods on
}

// CompoundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundCallerRaw struct {
	Contract *CompoundCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundTransactorRaw struct {
	Contract *CompoundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompound creates a new instance of Compound, bound to a specific deployed contract.
func NewCompound(address common.Address, backend bind.ContractBackend) (*Compound, error) {
	contract, err := bindCompound(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compound{CompoundCaller: CompoundCaller{contract: contract}, CompoundTransactor: CompoundTransactor{contract: contract}, CompoundFilterer: CompoundFilterer{contract: contract}}, nil
}

// NewCompoundCaller creates a new read-only instance of Compound, bound to a specific deployed contract.
func NewCompoundCaller(address common.Address, caller bind.ContractCaller) (*CompoundCaller, error) {
	contract, err := bindCompound(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundCaller{contract: contract}, nil
}

// NewCompoundTransactor creates a new write-only instance of Compound, bound to a specific deployed contract.
func NewCompoundTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundTransactor, error) {
	contract, err := bindCompound(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundTransactor{contract: contract}, nil
}

// NewCompoundFilterer creates a new log filterer instance of Compound, bound to a specific deployed contract.
func NewCompoundFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundFilterer, error) {
	contract, err := bindCompound(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundFilterer{contract: contract}, nil
}

// bindCompound binds a generic wrapper to an already deployed contract.
func bindCompound(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compound *CompoundRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Compound.Contract.CompoundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compound *CompoundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compound.Contract.CompoundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compound *CompoundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compound.Contract.CompoundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compound *CompoundCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Compound.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compound *CompoundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compound.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compound *CompoundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compound.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Compound *CompoundCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Compound *CompoundSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Compound.Contract.ETHCONTRACTADDRESS(&_Compound.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Compound *CompoundCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Compound.Contract.ETHCONTRACTADDRESS(&_Compound.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Compound *CompoundCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Compound *CompoundSession) Admin() (common.Address, error) {
	return _Compound.Contract.Admin(&_Compound.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Compound *CompoundCallerSession) Admin() (common.Address, error) {
	return _Compound.Contract.Admin(&_Compound.CallOpts)
}

// Agents is a free data retrieval call binding the contract method 0xfd66091e.
//
// Solidity: function agents(address ) constant returns(address)
func (_Compound *CompoundCaller) Agents(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "agents", arg0)
	return *ret0, err
}

// Agents is a free data retrieval call binding the contract method 0xfd66091e.
//
// Solidity: function agents(address ) constant returns(address)
func (_Compound *CompoundSession) Agents(arg0 common.Address) (common.Address, error) {
	return _Compound.Contract.Agents(&_Compound.CallOpts, arg0)
}

// Agents is a free data retrieval call binding the contract method 0xfd66091e.
//
// Solidity: function agents(address ) constant returns(address)
func (_Compound *CompoundCallerSession) Agents(arg0 common.Address) (common.Address, error) {
	return _Compound.Contract.Agents(&_Compound.CallOpts, arg0)
}

// CompoundAgentLogic is a free data retrieval call binding the contract method 0xfee8a5c1.
//
// Solidity: function compoundAgentLogic() constant returns(address)
func (_Compound *CompoundCaller) CompoundAgentLogic(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "compoundAgentLogic")
	return *ret0, err
}

// CompoundAgentLogic is a free data retrieval call binding the contract method 0xfee8a5c1.
//
// Solidity: function compoundAgentLogic() constant returns(address)
func (_Compound *CompoundSession) CompoundAgentLogic() (common.Address, error) {
	return _Compound.Contract.CompoundAgentLogic(&_Compound.CallOpts)
}

// CompoundAgentLogic is a free data retrieval call binding the contract method 0xfee8a5c1.
//
// Solidity: function compoundAgentLogic() constant returns(address)
func (_Compound *CompoundCallerSession) CompoundAgentLogic() (common.Address, error) {
	return _Compound.Contract.CompoundAgentLogic(&_Compound.CallOpts)
}

// GetAgentAddress is a free data retrieval call binding the contract method 0x27d54a92.
//
// Solidity: function getAgentAddress(address caller) constant returns(address)
func (_Compound *CompoundCaller) GetAgentAddress(opts *bind.CallOpts, caller common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "getAgentAddress", caller)
	return *ret0, err
}

// GetAgentAddress is a free data retrieval call binding the contract method 0x27d54a92.
//
// Solidity: function getAgentAddress(address caller) constant returns(address)
func (_Compound *CompoundSession) GetAgentAddress(caller common.Address) (common.Address, error) {
	return _Compound.Contract.GetAgentAddress(&_Compound.CallOpts, caller)
}

// GetAgentAddress is a free data retrieval call binding the contract method 0x27d54a92.
//
// Solidity: function getAgentAddress(address caller) constant returns(address)
func (_Compound *CompoundCallerSession) GetAgentAddress(caller common.Address) (common.Address, error) {
	return _Compound.Contract.GetAgentAddress(&_Compound.CallOpts, caller)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Compound *CompoundCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Compound *CompoundSession) IncognitoSmartContract() (common.Address, error) {
	return _Compound.Contract.IncognitoSmartContract(&_Compound.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Compound *CompoundCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _Compound.Contract.IncognitoSmartContract(&_Compound.CallOpts)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) constant returns(bool)
func (_Compound *CompoundCaller) IsSigDataUsed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "isSigDataUsed", hash)
	return *ret0, err
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) constant returns(bool)
func (_Compound *CompoundSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Compound.Contract.IsSigDataUsed(&_Compound.CallOpts, hash)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) constant returns(bool)
func (_Compound *CompoundCallerSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Compound.Contract.IsSigDataUsed(&_Compound.CallOpts, hash)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) constant returns(bool)
func (_Compound *CompoundCaller) SigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "sigDataUsed", arg0)
	return *ret0, err
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) constant returns(bool)
func (_Compound *CompoundSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Compound.Contract.SigDataUsed(&_Compound.CallOpts, arg0)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) constant returns(bool)
func (_Compound *CompoundCallerSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Compound.Contract.SigDataUsed(&_Compound.CallOpts, arg0)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) constant returns(address)
func (_Compound *CompoundCaller) SigToAddress(opts *bind.CallOpts, signData []byte, hash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "sigToAddress", signData, hash)
	return *ret0, err
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) constant returns(address)
func (_Compound *CompoundSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Compound.Contract.SigToAddress(&_Compound.CallOpts, signData, hash)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) constant returns(address)
func (_Compound *CompoundCallerSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Compound.Contract.SigToAddress(&_Compound.CallOpts, signData, hash)
}

// Execute is a paid mutator transaction binding the contract method 0x51f12037.
//
// Solidity: function execute(address srcToken, uint256 amount, bytes callData, bytes timestamp, bytes signData) returns(address, uint256)
func (_Compound *CompoundTransactor) Execute(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "execute", srcToken, amount, callData, timestamp, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x51f12037.
//
// Solidity: function execute(address srcToken, uint256 amount, bytes callData, bytes timestamp, bytes signData) returns(address, uint256)
func (_Compound *CompoundSession) Execute(srcToken common.Address, amount *big.Int, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Compound.Contract.Execute(&_Compound.TransactOpts, srcToken, amount, callData, timestamp, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x51f12037.
//
// Solidity: function execute(address srcToken, uint256 amount, bytes callData, bytes timestamp, bytes signData) returns(address, uint256)
func (_Compound *CompoundTransactorSession) Execute(srcToken common.Address, amount *big.Int, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Compound.Contract.Execute(&_Compound.TransactOpts, srcToken, amount, callData, timestamp, signData)
}

// ExecuteMulti is a paid mutator transaction binding the contract method 0x24d536e3.
//
// Solidity: function executeMulti(address[] srcTokens, uint256[] amounts, bytes callData, bytes timestamp, bytes signData) returns(address[], uint256[])
func (_Compound *CompoundTransactor) ExecuteMulti(opts *bind.TransactOpts, srcTokens []common.Address, amounts []*big.Int, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "executeMulti", srcTokens, amounts, callData, timestamp, signData)
}

// ExecuteMulti is a paid mutator transaction binding the contract method 0x24d536e3.
//
// Solidity: function executeMulti(address[] srcTokens, uint256[] amounts, bytes callData, bytes timestamp, bytes signData) returns(address[], uint256[])
func (_Compound *CompoundSession) ExecuteMulti(srcTokens []common.Address, amounts []*big.Int, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Compound.Contract.ExecuteMulti(&_Compound.TransactOpts, srcTokens, amounts, callData, timestamp, signData)
}

// ExecuteMulti is a paid mutator transaction binding the contract method 0x24d536e3.
//
// Solidity: function executeMulti(address[] srcTokens, uint256[] amounts, bytes callData, bytes timestamp, bytes signData) returns(address[], uint256[])
func (_Compound *CompoundTransactorSession) ExecuteMulti(srcTokens []common.Address, amounts []*big.Int, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Compound.Contract.ExecuteMulti(&_Compound.TransactOpts, srcTokens, amounts, callData, timestamp, signData)
}

// UpdateAgentLogic is a paid mutator transaction binding the contract method 0xea1efaf4.
//
// Solidity: function updateAgentLogic(address _compoundAgentLogic) returns()
func (_Compound *CompoundTransactor) UpdateAgentLogic(opts *bind.TransactOpts, _compoundAgentLogic common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "updateAgentLogic", _compoundAgentLogic)
}

// UpdateAgentLogic is a paid mutator transaction binding the contract method 0xea1efaf4.
//
// Solidity: function updateAgentLogic(address _compoundAgentLogic) returns()
func (_Compound *CompoundSession) UpdateAgentLogic(_compoundAgentLogic common.Address) (*types.Transaction, error) {
	return _Compound.Contract.UpdateAgentLogic(&_Compound.TransactOpts, _compoundAgentLogic)
}

// UpdateAgentLogic is a paid mutator transaction binding the contract method 0xea1efaf4.
//
// Solidity: function updateAgentLogic(address _compoundAgentLogic) returns()
func (_Compound *CompoundTransactorSession) UpdateAgentLogic(_compoundAgentLogic common.Address) (*types.Transaction, error) {
	return _Compound.Contract.UpdateAgentLogic(&_Compound.TransactOpts, _compoundAgentLogic)
}

// UpdateVault is a paid mutator transaction binding the contract method 0xe7563f3f.
//
// Solidity: function updateVault(address _incognitoSmartContract) returns()
func (_Compound *CompoundTransactor) UpdateVault(opts *bind.TransactOpts, _incognitoSmartContract common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "updateVault", _incognitoSmartContract)
}

// UpdateVault is a paid mutator transaction binding the contract method 0xe7563f3f.
//
// Solidity: function updateVault(address _incognitoSmartContract) returns()
func (_Compound *CompoundSession) UpdateVault(_incognitoSmartContract common.Address) (*types.Transaction, error) {
	return _Compound.Contract.UpdateVault(&_Compound.TransactOpts, _incognitoSmartContract)
}

// UpdateVault is a paid mutator transaction binding the contract method 0xe7563f3f.
//
// Solidity: function updateVault(address _incognitoSmartContract) returns()
func (_Compound *CompoundTransactorSession) UpdateVault(_incognitoSmartContract common.Address) (*types.Transaction, error) {
	return _Compound.Contract.UpdateVault(&_Compound.TransactOpts, _incognitoSmartContract)
}

// CompoundUpdateAgentLogicIterator is returned from FilterUpdateAgentLogic and is used to iterate over the raw logs and unpacked data for UpdateAgentLogic events raised by the Compound contract.
type CompoundUpdateAgentLogicIterator struct {
	Event *CompoundUpdateAgentLogic // Event containing the contract specifics and raw log

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
func (it *CompoundUpdateAgentLogicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundUpdateAgentLogic)
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
		it.Event = new(CompoundUpdateAgentLogic)
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
func (it *CompoundUpdateAgentLogicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundUpdateAgentLogicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundUpdateAgentLogic represents a UpdateAgentLogic event raised by the Compound contract.
type CompoundUpdateAgentLogic struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAgentLogic is a free log retrieval operation binding the contract event 0x1d7094459f506b06afa7f2c3bec49b9073c63236bb41bcb97513be10f0ff6327.
//
// Solidity: event UpdateAgentLogic(address )
func (_Compound *CompoundFilterer) FilterUpdateAgentLogic(opts *bind.FilterOpts) (*CompoundUpdateAgentLogicIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "UpdateAgentLogic")
	if err != nil {
		return nil, err
	}
	return &CompoundUpdateAgentLogicIterator{contract: _Compound.contract, event: "UpdateAgentLogic", logs: logs, sub: sub}, nil
}

// WatchUpdateAgentLogic is a free log subscription operation binding the contract event 0x1d7094459f506b06afa7f2c3bec49b9073c63236bb41bcb97513be10f0ff6327.
//
// Solidity: event UpdateAgentLogic(address )
func (_Compound *CompoundFilterer) WatchUpdateAgentLogic(opts *bind.WatchOpts, sink chan<- *CompoundUpdateAgentLogic) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "UpdateAgentLogic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundUpdateAgentLogic)
				if err := _Compound.contract.UnpackLog(event, "UpdateAgentLogic", log); err != nil {
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

// ParseUpdateAgentLogic is a log parse operation binding the contract event 0x1d7094459f506b06afa7f2c3bec49b9073c63236bb41bcb97513be10f0ff6327.
//
// Solidity: event UpdateAgentLogic(address )
func (_Compound *CompoundFilterer) ParseUpdateAgentLogic(log types.Log) (*CompoundUpdateAgentLogic, error) {
	event := new(CompoundUpdateAgentLogic)
	if err := _Compound.contract.UnpackLog(event, "UpdateAgentLogic", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CompoundUpdateVaultCompoundIterator is returned from FilterUpdateVaultCompound and is used to iterate over the raw logs and unpacked data for UpdateVaultCompound events raised by the Compound contract.
type CompoundUpdateVaultCompoundIterator struct {
	Event *CompoundUpdateVaultCompound // Event containing the contract specifics and raw log

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
func (it *CompoundUpdateVaultCompoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundUpdateVaultCompound)
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
		it.Event = new(CompoundUpdateVaultCompound)
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
func (it *CompoundUpdateVaultCompoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundUpdateVaultCompoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundUpdateVaultCompound represents a UpdateVaultCompound event raised by the Compound contract.
type CompoundUpdateVaultCompound struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdateVaultCompound is a free log retrieval operation binding the contract event 0x856bb47bb5dc3e7d1f5c5d594120747b26fb9b49f9a530679772ac961d7c58a4.
//
// Solidity: event UpdateVaultCompound(address )
func (_Compound *CompoundFilterer) FilterUpdateVaultCompound(opts *bind.FilterOpts) (*CompoundUpdateVaultCompoundIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "UpdateVaultCompound")
	if err != nil {
		return nil, err
	}
	return &CompoundUpdateVaultCompoundIterator{contract: _Compound.contract, event: "UpdateVaultCompound", logs: logs, sub: sub}, nil
}

// WatchUpdateVaultCompound is a free log subscription operation binding the contract event 0x856bb47bb5dc3e7d1f5c5d594120747b26fb9b49f9a530679772ac961d7c58a4.
//
// Solidity: event UpdateVaultCompound(address )
func (_Compound *CompoundFilterer) WatchUpdateVaultCompound(opts *bind.WatchOpts, sink chan<- *CompoundUpdateVaultCompound) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "UpdateVaultCompound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundUpdateVaultCompound)
				if err := _Compound.contract.UnpackLog(event, "UpdateVaultCompound", log); err != nil {
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

// ParseUpdateVaultCompound is a log parse operation binding the contract event 0x856bb47bb5dc3e7d1f5c5d594120747b26fb9b49f9a530679772ac961d7c58a4.
//
// Solidity: event UpdateVaultCompound(address )
func (_Compound *CompoundFilterer) ParseUpdateVaultCompound(log types.Log) (*CompoundUpdateVaultCompound, error) {
	event := new(CompoundUpdateVaultCompound)
	if err := _Compound.contract.UnpackLog(event, "UpdateVaultCompound", log); err != nil {
		return nil, err
	}
	return event, nil
}
