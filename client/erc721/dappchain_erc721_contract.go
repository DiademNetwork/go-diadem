// +build evm

package erc721

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/diademnetwork/go-diadem"
	"github.com/diademnetwork/go-diadem/auth"
	"github.com/diademnetwork/go-diadem/client"
)

type DAppChainERC721Contract struct {
	*client.MirroredTokenContract
}

func (c *DAppChainERC721Contract) OwnerOf(tokenID *big.Int) (diadem.LocalAddress, error) {
	var result common.Address
	if err := c.StaticCallEVM("ownerOf", &result, tokenID); err != nil {
		return nil, err
	}
	addr, err := diadem.LocalAddressFromHexString(result.Hex())
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// Approve grants authorization to an entity to transfer the given token at a later time
func (c *DAppChainERC721Contract) Approve(identity *client.Identity, to diadem.Address, tokenID *big.Int) error {
	toAddr := common.BytesToAddress(to.Local)
	return c.CallEVM("approve", identity.diademSigner, toAddr, tokenID)
}

func (c *DAppChainERC721Contract) BalanceOf(identity *client.Identity) (uint64, error) {
	result := new(big.Int)
	addr := common.BytesToAddress(identity.diademAddr.Local)
	if err := c.StaticCallEVM("balanceOf", &result, addr); err != nil {
		return 0, err
	}
	return result.Uint64(), nil
}

func (c *DAppChainERC721Contract) TransferFrom(from *client.Identity, to *client.Identity, tokenID *big.Int) error {
	fromAddr := common.BytesToAddress(from.diademAddr.Local)
	toAddr := common.BytesToAddress(to.diademAddr.Local)
	return c.CallEVM("transferFrom", from.diademSigner, fromAddr, toAddr, tokenID)
}

func (c *DAppChainERC721Contract) MintTo(identity *client.Identity, to diadem.Address, tokenID *big.Int) error {
	toAddr := common.BytesToAddress(to.Local)
	return c.CallEVM("mintTo", identity.diademSigner, toAddr, tokenID)
}

/**
  Connectors
*/

func DeployERC721ToDAppChain(diademClient *client.DAppChainRPCClient, contractName string,
	gatewayAddr diadem.Address, creator auth.Signer) (*DAppChainERC721Contract, error) {
	contract, err := client.DeployTokenToDAppChain(diademClient, contractName, "erc721", gatewayAddr, creator)
	if err != nil {
		return nil, err
	}
	return &DAppChainERC721Contract{contract}, nil
}

func ConnectERC721ToDAppChain(
	diademClient *client.DAppChainRPCClient, contractName string,
) (*DAppChainERC721Contract, error) {
	contract, err := client.ConnectToMirroredTokenContract(diademClient, contractName, "erc721")
	if err != nil {
		return nil, err
	}
	return &DAppChainERC721Contract{contract}, nil
}
