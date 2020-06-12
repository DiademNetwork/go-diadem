// +build evm

package native_coin

import (
	"math/big"

	"github.com/diademnetwork/go-diadem"
	"github.com/diademnetwork/go-diadem/builtin/types/coin"
	"github.com/diademnetwork/go-diadem/client"
	"github.com/diademnetwork/go-diadem/types"
)

// DAppChainNativeCoin is a client-side binding for the builtin coin Go contracts.
type DAppChainNativeCoin struct {
	contract *client.Contract
	chainID  string

	Address diadem.Address
}

func (ec *DAppChainNativeCoin) BalanceOf(identity *client.Identity) (*big.Int, error) {
	ownerAddr := diadem.Address{
		ChainID: ec.chainID,
		Local:   identity.diademAddr.Local,
	}
	req := &coin.BalanceOfRequest{
		Owner: ownerAddr.MarshalPB(),
	}
	var resp coin.BalanceOfResponse
	_, err := ec.contract.StaticCall("BalanceOf", req, ownerAddr, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Balance != nil {
		return resp.Balance.Value.Int, nil
	}
	return nil, nil
}

func (ec *DAppChainNativeCoin) Approve(owner *client.Identity, spender diadem.Address, amount *big.Int) error {
	req := &coin.ApproveRequest{
		Spender: spender.MarshalPB(),
		Amount:  &types.BigUInt{Value: *diadem.NewBigUInt(amount)},
	}
	_, err := ec.contract.Call("Approve", req, owner.diademSigner, nil)
	return err
}

func (ec *DAppChainNativeCoin) Transfer(owner *client.Identity, to diadem.Address, amount *big.Int) error {
	req := &coin.TransferRequest{
		To:     to.MarshalPB(),
		Amount: &types.BigUInt{Value: *diadem.NewBigUInt(amount)},
	}
	_, err := ec.contract.Call("Transfer", req, owner.diademSigner, nil)
	return err
}

/** Connectors */

func ConnectToDAppChaindiademContract(diademClient *client.DAppChainRPCClient) (*DAppChainNativeCoin, error) {
	return connectToDAppChainNativeCoin(diademClient, "coin")
}

func ConnectToDAppChainETHContract(diademClient *client.DAppChainRPCClient) (*DAppChainNativeCoin, error) {
	return connectToDAppChainNativeCoin(diademClient, "ethcoin")
}

func connectToDAppChainNativeCoin(diademClient *client.DAppChainRPCClient, name string) (*DAppChainNativeCoin, error) {
	contractAddr, err := diademClient.Resolve(name)
	if err != nil {
		return nil, err
	}

	return &DAppChainNativeCoin{
		contract: client.NewContract(diademClient, contractAddr.Local),
		chainID:  diademClient.GetChainID(),
		Address:  contractAddr,
	}, nil
}
