// +build evm

package address_mapper

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/diademnetwork/go-diadem"
	"github.com/diademnetwork/go-diadem/builtin/types/address_mapper"
	"github.com/diademnetwork/go-diadem/client"
	"github.com/diademnetwork/go-diadem/common/evmcompat"
	ssha "github.com/miguelmota/go-solidity-sha3"
)

type DAppChainAddressMapper struct {
	contract *client.Contract
	chainID  string
}

func ConnectToDAppChainAddressMapper(diademClient *client.DAppChainRPCClient) (*DAppChainAddressMapper, error) {
	mapperAddr, err := diademClient.Resolve("addressmapper")
	if err != nil {
		return nil, err
	}
	return &DAppChainAddressMapper{
		contract: client.NewContract(diademClient, mapperAddr.Local),
		chainID:  diademClient.GetChainID(),
	}, nil
}

func (am *DAppChainAddressMapper) GetMappedAccount(account diadem.Address) (diadem.Address, error) {
	req := &address_mapper.AddressMapperGetMappingRequest{
		From: account.MarshalPB(),
	}
	resp := &address_mapper.AddressMapperGetMappingResponse{}
	_, err := am.contract.StaticCall("GetMapping", req, account, resp)
	if err != nil {
		return diadem.Address{}, err
	}
	return diadem.UnmarshalAddressPB(resp.To), nil
}

func (am *DAppChainAddressMapper) HasIdentityMapping(account diadem.Address) (bool, error) {
	req := &address_mapper.AddressMapperHasMappingRequest{
		From: account.MarshalPB(),
	}
	resp := &address_mapper.AddressMapperHasMappingResponse{}
	_, err := am.contract.StaticCall("HasMapping", req, account, resp)
	if err != nil {
		return false, err
	}
	return resp.HasMapping, nil
}

// AddIdentityMapping creates a bi-directional mapping between a Mainnet & DAppChain account.
func (am *DAppChainAddressMapper) AddIdentityMapping(identity *client.Identity) error {
	mainnetAddrBytes, err := diadem.LocalAddressFromHexString(identity.MainnetAddr.Hex())
	if err != nil {
		return err
	}
	from := diadem.Address{
		ChainID: "eth",
		Local:   mainnetAddrBytes,
	}
	to := diadem.Address{
		ChainID: am.chainID,
		Local:   identity.diademAddr.Local,
	}

	mappedAccount, err := am.GetMappedAccount(from)
	if err == nil {
		if mappedAccount.Compare(to) != 0 {
			return fmt.Errorf("Account %v is mapped to %v", from, mappedAccount)
		}
		return nil
	}

	fmt.Printf("Mapping account %v to %v\n", from, to)
	sig, err := signIdentityMapping(from, to, identity.MainnetPrivKey)
	if err != nil {
		return err
	}
	req := &address_mapper.AddressMapperAddIdentityMappingRequest{
		From:      from.MarshalPB(),
		To:        to.MarshalPB(),
		Signature: sig,
	}
	_, err = am.contract.Call("AddIdentityMapping", req, identity.diademSigner, nil)
	return err
}

func signIdentityMapping(from, to diadem.Address, key *ecdsa.PrivateKey) ([]byte, error) {
	hash := ssha.SoliditySHA3(
		ssha.Address(common.BytesToAddress(from.Local)),
		ssha.Address(common.BytesToAddress(to.Local)),
	)
	sig, err := evmcompat.SoliditySign(hash, key)
	if err != nil {
		return nil, err
	}
	// Prefix the sig with a single byte indicating the sig type, in this case EIP712
	return append(make([]byte, 1, 66), sig...), nil
}
