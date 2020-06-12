// +build evm

package erc20

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	diadem "github.com/diademnetwork/go-diadem"
	"github.com/diademnetwork/go-diadem/auth"
	"github.com/diademnetwork/go-diadem/client"
)

type DAppChainERC20Contract struct {
	*client.MirroredTokenContract
}

func (c *DAppChainERC20Contract) BalanceOf(caller *client.Identity) (*big.Int, error) {
	ownerAddr := common.BytesToAddress(caller.diademAddr.Local)
	var result *big.Int
	if err := c.StaticCallEVM("balanceOf", &result, ownerAddr); err != nil {
		return nil, err
	}
	return result, nil
}

// Approve grants authorization to an entity to transfer the given tokens at a later time
func (c *DAppChainERC20Contract) Approve(identity *client.Identity, to diadem.Address, amount *big.Int) error {
	toAddr := common.BytesToAddress(to.Local)
	return c.CallEVM("approve", identity.diademSigner, toAddr, amount)
}

func (c *DAppChainERC20Contract) Transfer(identity *client.Identity, to diadem.Address, amount *big.Int) error {
	toAddr := common.BytesToAddress(to.Local)
	return c.CallEVM("transfer", identity.diademSigner, toAddr, amount)
}

func (c *DAppChainERC20Contract) MintTo(identity *client.Identity, to diadem.Address, amount *big.Int) error {
	toAddr := common.BytesToAddress(to.Local)
	return c.CallEVM("mintTo", identity.diademSigner, toAddr, amount)
}

/**
  Functions to connect / deploy the ERC20 contracts
*/

func DeployERC20ToDAppChain(diademClient *client.DAppChainRPCClient, contractName string,
	gatewayAddr diadem.Address, creator auth.Signer) (*DAppChainERC20Contract, error) {
	contract, err := client.DeployTokenToDAppChain(diademClient, contractName, "erc20", gatewayAddr, creator)
	if err != nil {
		return nil, err
	}
	return &DAppChainERC20Contract{contract}, nil
}

func ConnectERC20ToDAppChain(
	diademClient *client.DAppChainRPCClient, contractName string,
) (*DAppChainERC20Contract, error) {
	contract, err := client.ConnectToMirroredTokenContract(diademClient, contractName, "erc20")
	if err != nil {
		return nil, err
	}
	return &DAppChainERC20Contract{contract}, nil
}
