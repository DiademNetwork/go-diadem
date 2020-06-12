// +build evm

package erc721x

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/diademnetwork/go-diadem"
	"github.com/diademnetwork/go-diadem/auth"
	"github.com/diademnetwork/go-diadem/client"
)

type DAppChainERC721XContract struct {
	*client.MirroredTokenContract
}

func (c *DAppChainERC721XContract) OwnerOf(tokenID *big.Int) (diadem.LocalAddress, error) {
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

// SetApprovalForAll grants or revokes authorization to an entity to transfer any tokens from the
// owner account at a later time.
func (c *DAppChainERC721XContract) SetApprovalForAll(owner *client.Identity, operator diadem.Address, isApproved bool) error {
	operatorAddr := common.BytesToAddress(operator.Local)
	return c.CallEVM("setApprovalForAll", owner.diademSigner, operatorAddr, isApproved)
}

func (c *DAppChainERC721XContract) BalanceOf(identity *client.Identity, tokenID *big.Int) (*big.Int, error) {
	result := new(big.Int)
	addr := common.BytesToAddress(identity.diademAddr.Local)
	if err := c.StaticCallEVM("balanceOfToken", &result, addr, tokenID); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *DAppChainERC721XContract) TransferFrom(from *client.Identity, to *client.Identity, tokenID, amount *big.Int) error {
	fromAddr := common.BytesToAddress(from.diademAddr.Local)
	toAddr := common.BytesToAddress(to.diademAddr.Local)
	return c.CallEVM("transferFrom", from.diademSigner, fromAddr, toAddr, tokenID, amount)
}

func DeployERC721XToDAppChain(diademClient *client.DAppChainRPCClient, contractName string,
	gatewayAddr diadem.Address, creator auth.Signer) (*DAppChainERC721XContract, error) {
	contract, err := client.DeployTokenToDAppChain(diademClient, contractName, "erc721x", gatewayAddr, creator)
	if err != nil {
		return nil, err
	}
	return &DAppChainERC721XContract{contract}, nil
}

func ConnectERC721XToDAppChain(
	diademClient *client.DAppChainRPCClient, contractName string,
) (*DAppChainERC721XContract, error) {
	contract, err := client.ConnectToMirroredTokenContract(diademClient, contractName, "erc721x")
	if err != nil {
		return nil, err
	}
	return &DAppChainERC721XContract{contract}, nil
}
