package cli

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	diadem "github.com/diademnetwork/go-diadem"
	amtypes "github.com/diademnetwork/go-diadem/builtin/types/address_mapper"
	"github.com/diademnetwork/go-diadem/client"
	"github.com/pkg/errors"
)

func ParseBytes(s string) ([]byte, error) {
	if strings.HasPrefix(s, "0x") {
		return hex.DecodeString(s[2:])
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		b, err = base64.StdEncoding.DecodeString(s)
	}

	return b, err
}

// ParseAddress attempts to parse the given string into an address, if the resulting address doesn't
// have a chain ID the given chain ID will be used instead.
func ParseAddress(s, chainID string) (diadem.Address, error) {
	addr, err := diadem.ParseAddress(s)
	if err == nil {
		return addr, nil
	}

	b, err := ParseBytes(s)
	if err != nil {
		return diadem.Address{}, err
	}
	if len(b) != 20 {
		return diadem.Address{}, diadem.ErrInvalidAddress
	}

	return diadem.Address{ChainID: chainID, Local: diadem.LocalAddress(b)}, nil
}

// ResolveAddress attempts to parse the given string into an address, if that fails it assumes the
// string corresponds to a contract name and attempts to obtain the corresponding contract address.
func ResolveAddress(s, chainID, URI string) (diadem.Address, error) {
	rpcClient := client.NewDAppChainRPCClient(chainID, URI+"/rpc", URI+"/query")
	contractAddr, err := ParseAddress(s, chainID)
	if err != nil {
		// if address invalid, try to resolve it using registry
		contractAddr, err = rpcClient.Resolve(s)
		if err != nil {
			return diadem.Address{}, err
		}
	}

	return contractAddr, nil
}

func sciNot(m, n int64) *diadem.BigUInt {
	ret := diadem.NewBigUIntFromInt(10)
	ret.Exp(ret, diadem.NewBigUIntFromInt(n), nil)
	ret.Mul(ret, diadem.NewBigUIntFromInt(m))
	return ret
}

func ParseAmount(s string) (*diadem.BigUInt, error) {
	// TODO: allow more precision
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return sciNot(val, 18), nil
}

func getMappedAccount(mapper *client.Contract, account diadem.Address) (diadem.Address, error) {
	req := &amtypes.AddressMapperGetMappingRequest{
		From: account.MarshalPB(),
	}
	resp := &amtypes.AddressMapperGetMappingResponse{}
	_, err := mapper.StaticCall("GetMapping", req, account, resp)
	if err != nil {
		return diadem.Address{}, err
	}
	return diadem.UnmarshalAddressPB(resp.To), nil
}

// ResolveAccountAddress attempts to parse the given string into the address of a user account.
// If the chain ID on the parsed address doesn't match the chain ID specified in chainFlags then
// the address is resolved to an on-chain address via the address mapper.
func ResolveAccountAddress(address string, chainFlags *ContractCallFlags) (diadem.Address, error) {
	addr, err := ParseAddress(address, chainFlags.ChainID)
	if err != nil {
		return addr, errors.Wrap(err, "failed to parse address")
	}
	// Resolve address if chainID doesn't match
	if addr.ChainID != chainFlags.ChainID {
		rpcClient := client.NewDAppChainRPCClient(chainFlags.ChainID, chainFlags.URI+"/rpc", chainFlags.URI+"/query")
		mapperAddr, err := rpcClient.Resolve("addressmapper")
		if err != nil {
			return addr, errors.Wrap(err, "failed to resolve DAppChain Address Mapper address")
		}
		mapper := client.NewContract(rpcClient, mapperAddr.Local)
		mappedAccount, err := getMappedAccount(mapper, addr)
		if err != nil {
			return addr, fmt.Errorf("No account information found for %v", addr)
		}
		addr = mappedAccount
	}
	return addr, nil
}
