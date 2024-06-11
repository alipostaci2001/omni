package bindings

import (
	_ "embed"
)

const (
	OmniBridgeNativeDeployedBytecode = "0x6080604052600436106100e85760003560e01c80638da5cb5b1161008a578063c3de453d11610059578063c3de453d14610287578063d9caed121461029a578063f2fde38b146102ba578063f35ea557146102da57600080fd5b80638da5cb5b146102135780638fdcb4c914610231578063969b53da146102485780639c5451b01461026857600080fd5b806339acf9f1116100c657806339acf9f1146101725780633abfe55f146101b1578063402914f5146101d1578063715018a6146101fe57600080fd5b806312622e5b146100ed5780631e83409a1461012c57806323b051d91461014e575b600080fd5b3480156100f957600080fd5b5060655461010e9067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b34801561013857600080fd5b5061014c610147366004610c66565b6102fa565b005b34801561015a57600080fd5b5061016460665481565b604051908152602001610123565b34801561017e57600080fd5b5060655461019990600160401b90046001600160a01b031681565b6040516001600160a01b039091168152602001610123565b3480156101bd57600080fd5b506101646101cc366004610c8a565b610598565b3480156101dd57600080fd5b506101646101ec366004610c66565b60686020526000908152604090205481565b34801561020a57600080fd5b5061014c61066a565b34801561021f57600080fd5b506033546001600160a01b0316610199565b34801561023d57600080fd5b5061010e620249f081565b34801561025457600080fd5b50606754610199906001600160a01b031681565b34801561027457600080fd5b506101646a52b7d2dcc80cd2e400000081565b61014c610295366004610c8a565b61067e565b3480156102a657600080fd5b5061014c6102b5366004610cb6565b61068c565b3480156102c657600080fd5b5061014c6102d5366004610c66565b6108ee565b3480156102e657600080fd5b5061014c6102f5366004610d0d565b610967565b60655460408051631799380760e11b81528151600093600160401b90046001600160a01b031692632f32700e92600480820193918290030181865afa158015610347573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036b9190610d58565b606554909150600160401b90046001600160a01b031633146103cc5760405162461bcd60e51b815260206004820152601560248201527413db5b9a509c9a5919d94e881b9bdd081e18d85b1b605a1b60448201526064015b60405180910390fd5b606554815167ffffffffffffffff9081169116146104215760405162461bcd60e51b81526020600482015260126024820152714f6d6e694272696467653a206e6f74204c3160701b60448201526064016103c3565b6020808201516001600160a01b0381166000908152606890925260409091205461048d5760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e694272696467653a206e6f7468696e6720746f207265636c61696d000060448201526064016103c3565b6001600160a01b038181166000908152606860205260408082208054908390559051909286169083908381818185875af1925050503d80600081146104ee576040519150601f19603f3d011682016040523d82523d6000602084013e6104f3565b606091505b50509050806105445760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e694272696467653a207472616e73666572206661696c6564000000000060448201526064016103c3565b846001600160a01b0316836001600160a01b03167ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926838460405161058991815260200190565b60405180910390a35050505050565b606554604080516001600160a01b038581166024830152604480830186905283518084039091018152606490920183526020820180516001600160e01b031663f3fef3a360e01b1790529151632376548f60e21b8152600093600160401b810490931692638dd9523c926106209267ffffffffffffffff90921691620249f090600401610e0b565b602060405180830381865afa15801561063d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106619190610e42565b90505b92915050565b6106726109bf565b61067c6000610a19565b565b6106888282610a6b565b5050565b60655460408051631799380760e11b81528151600093600160401b90046001600160a01b031692632f32700e92600480820193918290030181865afa1580156106d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106fd9190610d58565b606554909150600160401b90046001600160a01b031633146107595760405162461bcd60e51b815260206004820152601560248201527413db5b9a509c9a5919d94e881b9bdd081e18d85b1b605a1b60448201526064016103c3565b60675460208201516001600160a01b039081169116146107b45760405162461bcd60e51b81526020600482015260166024820152754f6d6e694272696467653a206e6f742062726964676560501b60448201526064016103c3565b606554815167ffffffffffffffff9081169116146108095760405162461bcd60e51b81526020600482015260126024820152714f6d6e694272696467653a206e6f74204c3160701b60448201526064016103c3565b816066600082825461081b9190610e71565b90915550506040516000906001600160a01b0385169084908381818185875af1925050503d806000811461086b576040519150601f19603f3d011682016040523d82523d6000602084013e610870565b606091505b50509050806108a7576001600160a01b038516600090815260686020526040812080548592906108a1908490610e71565b90915550505b6040805184815282151560208201526001600160a01b0380871692908816917f2b664ab52fe561d3ace376046aea39744dd736ec1f67d89d504ffd2192825f619101610589565b6108f66109bf565b6001600160a01b03811661095b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103c3565b61096481610a19565b50565b61096f6109bf565b6065805467ffffffffffffffff949094166001600160e01b031990941693909317600160401b6001600160a01b039384160217909255606780546001600160a01b03191692909116919091179055565b6033546001600160a01b0316331461067c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103c3565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008111610abb5760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e694272696467653a20616d6f756e74206d757374206265203e2030000060448201526064016103c3565b606654811115610b0d5760405162461bcd60e51b815260206004820152601860248201527f4f6d6e694272696467653a206e6f206c6971756964697479000000000000000060448201526064016103c3565b6000610b198383610598565b9050610b258183610e71565b3414610b735760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e694272696467653a20696e73756666696369656e742066756e6473000060448201526064016103c3565b8160666000828254610b859190610e84565b9091555050606554606754604080516001600160a01b038781166024830152604480830188905283518084039091018152606490920183526020820180516001600160e01b031663f3fef3a360e01b179052915163c21dda4f60e01b8152600160401b850483169463c21dda4f948794610c1a9467ffffffffffffffff909316936004939190921691620249f0908401610e97565b6000604051808303818588803b158015610c3357600080fd5b505af1158015610c47573d6000803e3d6000fd5b5050505050505050565b6001600160a01b038116811461096457600080fd5b600060208284031215610c7857600080fd5b8135610c8381610c51565b9392505050565b60008060408385031215610c9d57600080fd5b8235610ca881610c51565b946020939093013593505050565b600080600060608486031215610ccb57600080fd5b8335610cd681610c51565b92506020840135610ce681610c51565b929592945050506040919091013590565b67ffffffffffffffff8116811461096457600080fd5b600080600060608486031215610d2257600080fd5b8335610d2d81610cf7565b92506020840135610d3d81610c51565b91506040840135610d4d81610c51565b809150509250925092565b600060408284031215610d6a57600080fd5b6040516040810181811067ffffffffffffffff82111715610d9b57634e487b7160e01b600052604160045260246000fd5b6040528251610da981610cf7565b81526020830151610db981610c51565b60208201529392505050565b6000815180845260005b81811015610deb57602081850181015186830182015201610dcf565b506000602082860101526020601f19601f83011685010191505092915050565b600067ffffffffffffffff808616835260606020840152610e2f6060840186610dc5565b9150808416604084015250949350505050565b600060208284031215610e5457600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561066457610664610e5b565b8181038181111561066457610664610e5b565b600067ffffffffffffffff808816835260ff8716602084015260018060a01b038616604084015260a06060840152610ed260a0840186610dc5565b9150808416608084015250969550505050505056fea26469706673582212203b09f23a62ad3b38b5e1a351a323d98872d94632387f32547f4adc2e0578001e64736f6c63430008180033"
)

//go:embed omnibridgenative_storage_layout.json
var omnibridgenativeStorageLayoutJSON []byte

var OmniBridgeNativeStorageLayout = mustGetStorageLayout(omnibridgenativeStorageLayoutJSON)