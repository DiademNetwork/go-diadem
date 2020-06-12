package client

import (
	"github.com/gogo/protobuf/proto"
	"github.com/diademnetwork/go-diadem"
	"github.com/diademnetwork/go-diadem/auth"
	"github.com/diademnetwork/go-diadem/vm"
)

// Contract provides a thin abstraction over DAppChainClient that makes it easier to perform
// read & write operations on a contract running the EVM of a diadem DAppChain.
type EvmContract struct {
	client  *DAppChainRPCClient
	Address diadem.Address
	Name    string
}

func NewEvmContract(client *DAppChainRPCClient, contractAddr diadem.LocalAddress) *EvmContract {
	return &EvmContract{
		client: client,
		Address: diadem.Address{
			ChainID: client.GetChainID(),
			Local:   contractAddr,
		},
	}
}

func DeployContract(client *DAppChainRPCClient, byteCode []byte, signer auth.Signer, name string) (*EvmContract, []byte, error) {
	callerAddr := diadem.Address{
		ChainID: client.GetChainID(),
		Local:   diadem.LocalAddressFromPublicKey(signer.PublicKey()),
	}
	resp, err := client.CommitDeployTx(callerAddr, signer, vm.VMType_EVM, byteCode, name)
	if err != nil {
		return nil, nil, err
	}
	response := vm.DeployResponse{}
	err = proto.Unmarshal(resp, &response)
	if err != nil {
		return nil, nil, err
	}
	data := vm.DeployResponseData{}
	err = proto.Unmarshal(response.Output, &data)
	if err != nil {
		return nil, nil, err
	}
	return &EvmContract{
		client:  client,
		Address: diadem.UnmarshalAddressPB(response.Contract),
		Name:    name,
	}, data.GetTxHash(), nil
}

func (c *EvmContract) Call(input []byte, signer auth.Signer) ([]byte, error) {
	callerAddr := diadem.Address{
		ChainID: c.client.GetChainID(),
		Local:   diadem.LocalAddressFromPublicKey(signer.PublicKey()),
	}
	return c.client.CommitCallTx(callerAddr, c.Address, signer, vm.VMType_EVM, input)
}

func (c *EvmContract) StaticCall(input []byte, caller diadem.Address) ([]byte, error) {
	return c.client.QueryEvm(caller, c.Address.Local, input)
}
