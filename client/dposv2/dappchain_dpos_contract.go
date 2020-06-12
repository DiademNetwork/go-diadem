// +build evm

package dposv2

import (
	"github.com/diademnetwork/go-diadem"
	dpostypes "github.com/diademnetwork/go-diadem/builtin/types/dposv2"
	"github.com/diademnetwork/go-diadem/client"
	"github.com/pkg/errors"
	"math/big"
)

// DAppChainDPOSContract is a client-side binding for the builtin coin Go contract.
type DAppChainDPOSContract struct {
	contract *client.Contract
	chainID  string

	Address diadem.Address
}

func ConnectToDAppChainDPOSContract(diademClient *client.DAppChainRPCClient) (*DAppChainDPOSContract, error) {
	contractAddr, err := diademClient.Resolve("dposV2")
	if err != nil {
		return nil, err
	}

	return &DAppChainDPOSContract{
		contract: client.NewContract(diademClient, contractAddr.Local),
		chainID:  diademClient.GetChainID(),
		Address:  contractAddr,
	}, nil
}

func (dpos *DAppChainDPOSContract) CheckDistributions(identity *client.Identity) (*big.Int, error) {
	req := &dpostypes.CheckDistributionRequest{
		Address: identity.diademAddr.MarshalPB(),
	}
	var resp dpostypes.CheckDistributionResponse
	_, err := dpos.contract.StaticCall("CheckDistribution", req, identity.diademAddr, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Amount.Value.Int, err
}

func (dpos *DAppChainDPOSContract) ListCandidates(identity *client.Identity) ([]*dpostypes.CandidateV2, error) {
	req := &dpostypes.ListCandidateRequestV2{}
	var resp dpostypes.ListCandidateResponseV2
	_, err := dpos.contract.StaticCall("ListCandidates", req, identity.diademAddr, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Candidates, err
}

func (dpos *DAppChainDPOSContract) ListValidators(identity *client.Identity) ([]*dpostypes.ValidatorStatisticV2, error) {
	req := &dpostypes.ListValidatorsRequestV2{}
	var resp dpostypes.ListValidatorsResponseV2
	_, err := dpos.contract.StaticCall("ListValidators", req, identity.diademAddr, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Statistics, err
}

func (dpos *DAppChainDPOSContract) ProcessRequestBatch(identity *client.Identity, req *dpostypes.RequestBatchV2) error {
	_, err := dpos.contract.Call("ProcessRequestBatch", req, identity.diademSigner, nil)
	return err
}

func (dpos *DAppChainDPOSContract) GetRequestBatchTally(identity *client.Identity) (*dpostypes.RequestBatchTallyV2, error) {
	req := &dpostypes.GetRequestBatchTallyRequestV2{}
	resp := &dpostypes.RequestBatchTallyV2{}
	if _, err := dpos.contract.StaticCall("GetRequestBatchTally", req, identity.diademAddr, resp); err != nil {
		return nil, errors.Wrap(err, "failed to get request batch tally")
	}

	return resp, nil
}

func (dpos *DAppChainDPOSContract) ChangeFee(identity *client.Identity, candidateFee uint64) error {
	req := &dpostypes.ChangeCandidateFeeRequest{
		Fee: candidateFee,
	}
	_, err := dpos.contract.Call("ChangeFee", req, identity.diademSigner, nil)
	return err
}

func (dpos *DAppChainDPOSContract) ClaimRewards(identity *client.Identity, addr diadem.Address) (*dpostypes.ClaimDistributionResponseV2, error) {
	req := &dpostypes.ClaimDistributionRequestV2{
		WithdrawalAddress: addr.MarshalPB(),
	}
	resp := &dpostypes.ClaimDistributionResponseV2{}

	_, err := dpos.contract.Call("ClaimDistribution", req, identity.diademSigner, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (dpos *DAppChainDPOSContract) RegisterCandidate(identity *client.Identity, pubKey []byte, candidateFee uint64, candidateName string, candidateDescription string, candidateWebsite string) error {
	req := &dpostypes.RegisterCandidateRequestV2{
		PubKey:      pubKey,
		Fee:         candidateFee,
		Name:        candidateName,
		Description: candidateDescription,
		Website:     candidateWebsite,
	}
	_, err := dpos.contract.Call("RegisterCandidate", req, identity.diademSigner, nil)
	return err
}

func (dpos *DAppChainDPOSContract) UnregisterCandidate(identity *client.Identity) error {
	req := &dpostypes.UnregisterCandidateRequestV2{}
	_, err := dpos.contract.Call("UnregisterCandidate", req, identity.diademSigner, nil)
	return err
}
