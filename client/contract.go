package client

import (
	"errors"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/diademnetwork/go-diadem"
	"github.com/diademnetwork/go-diadem/auth"
	"github.com/diademnetwork/go-diadem/plugin"
	ptypes "github.com/diademnetwork/go-diadem/plugin/types"
	"github.com/diademnetwork/go-diadem/types"
	"github.com/diademnetwork/go-diadem/vm"
)

var (
	typeOfPBMessage = reflect.TypeOf((*proto.Message)(nil)).Elem()
)

// Contract provides a thin abstraction over DAppChainClient that makes it easier to perform
// read & write operations on a contract running on a diadem DAppChain.
type Contract struct {
	client  *DAppChainRPCClient
	Address diadem.Address
	Name    string
}

func NewContract(client *DAppChainRPCClient, contractAddr diadem.LocalAddress) *Contract {
	return &Contract{
		client: client,
		Address: diadem.Address{
			ChainID: client.GetChainID(),
			Local:   contractAddr,
		},
	}
}

func (c *Contract) Call(method string, args proto.Message, signer auth.Signer, result interface{}) (interface{}, error) {
	if result != nil && !reflect.TypeOf(result).Implements(typeOfPBMessage) {
		return nil, errors.New("Contract.Call result parameter must be a protobuf")
	}

	argsBytes, err := proto.Marshal(args)
	if err != nil {
		return nil, err
	}
	methodCallBytes, err := proto.Marshal(&plugin.ContractMethodCall{
		Method: method,
		Args:   argsBytes,
	})
	if err != nil {
		return nil, err
	}
	requestBytes, err := proto.Marshal(&plugin.Request{
		ContentType: plugin.EncodingType_PROTOBUF3,
		Accept:      plugin.EncodingType_PROTOBUF3,
		Body:        methodCallBytes,
	})
	if err != nil {
		return nil, err
	}
	callTxBytes, err := proto.Marshal(&vm.CallTx{
		VmType: vm.VMType_PLUGIN,
		Input:  requestBytes,
	})
	if err != nil {
		return nil, err
	}
	callerAddr := diadem.Address{
		ChainID: c.client.GetChainID(),
		Local:   diadem.LocalAddressFromPublicKey(signer.PublicKey()),
	}
	msgTxBytes, err := proto.Marshal(&vm.MessageTx{
		From: callerAddr.MarshalPB(),
		To:   c.Address.MarshalPB(),
		Data: callTxBytes,
	})
	if err != nil {
		return nil, err
	}
	resultBytes, err := c.client.CommitTx(signer, &types.Transaction{
		Id:   2,
		Data: msgTxBytes,
	})
	if err != nil {
		return nil, err
	}
	if result != nil && len(resultBytes) > 0 {
		response := &ptypes.Response{}
		err = proto.Unmarshal(resultBytes, response)
		if err != nil {
			return nil, nil
		}
		if err := proto.Unmarshal(response.Body, result.(proto.Message)); err != nil {
			return result, err
		}
	}
	return nil, nil
}

func (c *Contract) StaticCall(method string, args proto.Message, caller diadem.Address, result interface{}) (interface{}, error) {
	if result == nil || !reflect.TypeOf(result).Implements(typeOfPBMessage) {
		return nil, errors.New("Contract.StaticCall result parameter must be a protobuf")
	}
	argsBytes, err := proto.Marshal(args)
	if err != nil {
		return nil, err
	}
	methodCall := &plugin.ContractMethodCall{
		Method: method,
		Args:   argsBytes,
	}
	resultBytes, err := c.client.Query(caller, c.Address.Local, methodCall)
	if err != nil {
		return nil, err
	}
	if len(resultBytes) > 0 {
		if err := proto.Unmarshal(resultBytes, result.(proto.Message)); err != nil {
			return result, err
		}
	}
	return nil, nil
}
