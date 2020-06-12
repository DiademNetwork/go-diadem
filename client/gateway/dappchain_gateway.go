// +build evm

package gateway

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	diadem "github.com/diademnetwork/go-diadem"
	tgtypes "github.com/diademnetwork/go-diadem/builtin/types/transfer_gateway"
	"github.com/diademnetwork/go-diadem/client"
	"github.com/diademnetwork/go-diadem/common/evmcompat"
	lptypes "github.com/diademnetwork/go-diadem/plugin/types"
	"github.com/diademnetwork/go-diadem/types"
	ssha "github.com/miguelmota/go-solidity-sha3"
	pubsub "github.com/phonkee/go-pubsub"
	"github.com/pkg/errors"
)

const (
	tokenWithdrawalSignedEventTopic    = "event:TokenWithdrawalSigned"
	contractMappingConfirmedEventTopic = "event:ContractMappingConfirmed"
)

type DAppChainGateway struct {
	contract *client.Contract
	chainID  string

	ws                 *websocket.Conn
	nextMsgID          uint64
	chainEventQuitCh   chan struct{}
	chainEventSubCount int
	chainEventHub      pubsub.Hub

	Address diadem.Address
}

// AddAuthorisedContractMapping creates a bi-directional mapping between a Mainnet & DAppChain
// contract without creating a pending mapping
func (tg *DAppChainGateway) AddAuthorizedContractMapping(
	from common.Address, to diadem.Address, gatewayOwner *client.Identity,
) error {
	fromAddr, err := client.diademAddressFromEthereumAddress(from)
	if err != nil {
		return err
	}
	fmt.Printf("Mapping contract %v to %v\n", fromAddr, to)
	req := &tgtypes.TransferGatewayAddContractMappingRequest{
		ForeignContract: fromAddr.MarshalPB(),
		LocalContract:   to.MarshalPB(),
	}
	_, err = tg.contract.Call("AddAuthorizedContractMapping", req, gatewayOwner.diademSigner, nil)
	return err
}

// AddAuthorizedTronContractMapping same as AddAuthorisedContractMapping but for Tron
func (tg *DAppChainGateway) AddAuthorizedTronContractMapping(
	from common.Address, to diadem.Address, gatewayOwner *client.Identity,
) error {
	fromAddr, err := client.diademAddressFromTronAddress(from)
	if err != nil {
		return err
	}
	fmt.Printf("Mapping contract %v to %v\n", fromAddr, to)
	req := &tgtypes.TransferGatewayAddContractMappingRequest{
		ForeignContract: fromAddr.MarshalPB(),
		LocalContract:   to.MarshalPB(),
	}
	_, err = tg.contract.Call("AddAuthorizedContractMapping", req, gatewayOwner.diademSigner, nil)
	return err
}

// AddAuthorizedBinanceContractMapping same as AddAuthorisedContractMapping but for Binance dex
func (tg *DAppChainGateway) AddAuthorizedBinanceContractMapping(
	from common.Address, to diadem.Address, gatewayOwner *client.Identity,
) error {
	req := &tgtypes.TransferGatewayAddContractMappingRequest{
		ForeignContract: client.diademAddressFromBinanceAddress(from).MarshalPB(),
		LocalContract:   to.MarshalPB(),
	}
	_, err := tg.contract.Call("AddAuthorizedContractMapping", req, gatewayOwner.diademSigner, nil)
	return err
}

// AddContractMapping creates a bi-directional mapping between a Mainnet & DAppChain contract.
// The caller must provide the identity of the creator of the Mainnet contract, along with a Mainnet
// hash of the tx that deployed the contract (which will be used to verify the creator address).
func (tg *DAppChainGateway) AddContractMapping(
	from common.Address, to diadem.Address, creator *client.Identity, contractTxHash string,
) error {
	fromAddr, err := client.diademAddressFromEthereumAddress(from)
	if err != nil {
		return err
	}
	txHash, err := hex.DecodeString(strings.TrimPrefix(contractTxHash, "0x"))
	if err != nil {
		return err
	}
	fmt.Printf("Mapping contract %v to %v\n", fromAddr, to)

	hash := ssha.SoliditySHA3(
		ssha.Address(from),
		ssha.Address(common.BytesToAddress(to.Local)),
	)

	sig, err := evmcompat.GenerateTypedSig(hash, creator.MainnetPrivKey, evmcompat.SignatureType_EIP712)
	if err != nil {
		return err
	}

	req := &tgtypes.TransferGatewayAddContractMappingRequest{
		ForeignContract:           fromAddr.MarshalPB(),
		LocalContract:             to.MarshalPB(),
		ForeignContractCreatorSig: sig,
		ForeignContractTxHash:     txHash,
	}
	_, err = tg.contract.Call("AddContractMapping", req, creator.diademSigner, nil)
	return err
}

// AddTronContractMapping same as AddContractMapping but for Tron
func (tg *DAppChainGateway) AddTronContractMapping(
	from common.Address, to diadem.Address, creator *client.Identity, contractTxHash string,
) error {
	fromAddr, err := client.diademAddressFromTronAddress(from)
	if err != nil {
		return err
	}
	fmt.Printf("Mapping contract %v to %v\n", fromAddr, to)

	hash := ssha.SoliditySHA3(
		ssha.Address(from),
		ssha.Address(common.BytesToAddress(to.Local)),
	)

	sig, err := evmcompat.GenerateTypedSig(
		evmcompat.PrefixHeader(hash, evmcompat.SignatureType_TRON),
		creator.MainnetPrivKey,
		evmcompat.SignatureType_TRON)
	if err != nil {
		return err
	}

	req := &tgtypes.TransferGatewayAddContractMappingRequest{
		ForeignContract:           fromAddr.MarshalPB(),
		LocalContract:             to.MarshalPB(),
		ForeignContractCreatorSig: sig,
	}
	_, err = tg.contract.Call("AddContractMapping", req, creator.diademSigner, nil)
	return err
}

// AddBinanceContractMapping same as AddContractMapping but for Binance
func (tg *DAppChainGateway) AddBinanceContractMapping(
	from common.Address, to diadem.Address, creator *client.Identity,
) error {
	fromAddr := client.diademAddressFromBinanceAddress(from)
	fmt.Printf("Mapping contract %v to %v\n", fromAddr, to)

	hash := evmcompat.GenSHA256(
		ssha.Address(from),
		ssha.Address(common.BytesToAddress(to.Local)),
	)

	sig, err := evmcompat.GenerateTypedSig(hash, creator.MainnetPrivKey, evmcompat.SignatureType_BINANCE)
	if err != nil {
		return err
	}

	req := &tgtypes.TransferGatewayAddContractMappingRequest{
		ForeignContract:           fromAddr.MarshalPB(),
		LocalContract:             to.MarshalPB(),
		ForeignContractCreatorSig: sig,
	}
	_, err = tg.contract.Call("AddContractMapping", req, creator.diademSigner, nil)
	return err
}

func (tg *DAppChainGateway) WithdrawERC721(
	identity *client.Identity, tokenID *big.Int, contract diadem.Address, recipient *common.Address,
) error {
	req := &tgtypes.TransferGatewayWithdrawTokenRequest{
		TokenKind:     tgtypes.TransferGatewayTokenKind_ERC721,
		TokenID:       &types.BigUInt{Value: *diadem.NewBigUInt(tokenID)},
		TokenContract: contract.MarshalPB(),
	}
	if recipient != nil {
		req.Recipient = diadem.Address{
			ChainID: "eth",
			Local:   recipient.Bytes(),
		}.MarshalPB()
	}
	_, err := tg.contract.Call("WithdrawToken", req, identity.diademSigner, nil)
	return err
}

func (tg *DAppChainGateway) WithdrawERC721X(
	identity *client.Identity, tokenID, amount *big.Int, contract diadem.Address, recipient *common.Address,
) error {
	req := &tgtypes.TransferGatewayWithdrawTokenRequest{
		TokenKind:     tgtypes.TransferGatewayTokenKind_ERC721X,
		TokenID:       &types.BigUInt{Value: *diadem.NewBigUInt(tokenID)},
		TokenAmount:   &types.BigUInt{Value: *diadem.NewBigUInt(amount)},
		TokenContract: contract.MarshalPB(),
	}
	if recipient != nil {
		req.Recipient = diadem.Address{
			ChainID: "eth",
			Local:   recipient.Bytes(),
		}.MarshalPB()
	}
	_, err := tg.contract.Call("WithdrawToken", req, identity.diademSigner, nil)
	return err
}

func (tg *DAppChainGateway) WithdrawERC20(identity *client.Identity, amount *big.Int, contract diadem.Address) error {
	req := &tgtypes.TransferGatewayWithdrawTokenRequest{
		TokenKind:     tgtypes.TransferGatewayTokenKind_ERC20,
		TokenAmount:   &types.BigUInt{Value: *diadem.NewBigUInt(amount)},
		TokenContract: contract.MarshalPB(),
	}
	_, err := tg.contract.Call("WithdrawToken", req, identity.diademSigner, nil)
	return err
}

func (tg *DAppChainGateway) WithdrawTRX(identity *client.Identity, amount *big.Int, contract diadem.Address) error {
	req := &tgtypes.TransferGatewayWithdrawTokenRequest{
		TokenKind:     tgtypes.TransferGatewayTokenKind_TRX,
		TokenAmount:   &types.BigUInt{Value: *diadem.NewBigUInt(amount)},
		TokenContract: contract.MarshalPB(),
	}
	_, err := tg.contract.Call("WithdrawToken", req, identity.diademSigner, nil)
	return err
}

func (tg *DAppChainGateway) Withdrawdiadem(identity *client.Identity, amount *big.Int, mainnetdiademCoinAddress common.Address) error {
	req := &tgtypes.TransferGatewayWithdrawdiademCoinRequest{
		TokenContract: diadem.Address{
			ChainID: "eth",
			Local:   mainnetdiademCoinAddress.Bytes(),
		}.MarshalPB(),
		Amount: &types.BigUInt{Value: *diadem.NewBigUInt(amount)},
	}
	_, err := tg.contract.Call("WithdrawdiademCoin", req, identity.diademSigner, nil)
	return err
}

func (tg *DAppChainGateway) WithdrawdiademToBinanceDex(identity *client.Identity, amount *big.Int, mainnetRecipientAddress common.Address) error {
	req := &tgtypes.TransferGatewayWithdrawdiademCoinRequest{
		Recipient: diadem.Address{
			ChainID: "binance",
			Local:   mainnetRecipientAddress.Bytes(),
		}.MarshalPB(),
		Amount: &types.BigUInt{Value: *diadem.NewBigUInt(amount)},
	}
	_, err := tg.contract.Call("WithdrawdiademCoin", req, identity.diademSigner, nil)
	return err
}

func (tg *DAppChainGateway) WithdrawBEP2(identity *client.Identity, amount *big.Int, contract diadem.Address, mainnetRecipientAddress common.Address) error {
	req := &tgtypes.TransferGatewayWithdrawTokenRequest{
		TokenKind:     tgtypes.TransferGatewayTokenKind_BEP2,
		TokenAmount:   &types.BigUInt{Value: *diadem.NewBigUInt(amount)},
		TokenContract: contract.MarshalPB(),
		Recipient:     client.diademAddressFromBinanceAddress(mainnetRecipientAddress).MarshalPB(),
	}
	_, err := tg.contract.Call("WithdrawToken", req, identity.diademSigner, nil)
	return err
}

func (tg *DAppChainGateway) WithdrawETH(identity *client.Identity, amount *big.Int, mainnetGateway common.Address) error {
	req := &tgtypes.TransferGatewayWithdrawETHRequest{
		Amount: &types.BigUInt{Value: *diadem.NewBigUInt(amount)},
		MainnetGateway: diadem.Address{
			ChainID: "eth",
			Local:   mainnetGateway.Bytes(),
		}.MarshalPB(),
	}
	_, err := tg.contract.Call("WithdrawETH", req, identity.diademSigner, nil)
	return err
}

func (tg *DAppChainGateway) WithdrawalReceipt(identity *client.Identity) (*tgtypes.TransferGatewayWithdrawalReceipt, error) {
	req := &tgtypes.TransferGatewayWithdrawalReceiptRequest{
		Owner: identity.diademAddr.MarshalPB(),
	}
	var resp tgtypes.TransferGatewayWithdrawalReceiptResponse
	_, err := tg.contract.StaticCall("WithdrawalReceipt", req, identity.diademAddr, &resp)
	return resp.Receipt, err
}

func (tg *DAppChainGateway) ReclaimDepositorTokens(identity *client.Identity) error {
	req := &tgtypes.TransferGatewayReclaimDepositorTokensRequest{}
	_, err := tg.contract.Call("ReclaimDepositorTokens", req, identity.diademSigner, nil)
	return err
}

func (tg *DAppChainGateway) GetUnclaimedTokens(identity *client.Identity, addr diadem.Address) ([]*tgtypes.TransferGatewayUnclaimedToken, error) {
	req := &tgtypes.TransferGatewayGetUnclaimedTokensRequest{
		Owner: addr.MarshalPB(), // addr is an eth: prefixed ethereum address
	}
	resp := &tgtypes.TransferGatewayGetUnclaimedTokensResponse{}
	_, err := tg.contract.StaticCall("GetUnclaimedTokens", req, identity.diademAddr, resp)
	return resp.UnclaimedTokens, err
}

type EventSub struct {
	subscriber pubsub.Subscriber
	closeFn    func()
}

func (es *EventSub) Close() {
	es.subscriber.Close()
	es.closeFn()
}

func (tg *DAppChainGateway) WatchContractMappingConfirmed(
	sink chan<- *tgtypes.TransferGatewayContractMappingConfirmed) (*EventSub, error) {
	if tg.ws == nil {
		return nil, errors.New("websocket events unavailable")
	}

	if err := tg.subChainEvents(); err != nil {
		return nil, err
	}

	sub := tg.chainEventHub.Subscribe("event")
	sub.Do(func(msg pubsub.Message) {
		ev := lptypes.EventData{}
		if err := proto.Unmarshal(msg.Body(), &ev); err != nil {
			return
		}
		if ev.Topics == nil || ev.Topics[0] != contractMappingConfirmedEventTopic {
			return
		}
		contractAddr := diadem.UnmarshalAddressPB(ev.Address)
		if contractAddr.Compare(tg.Address) != 0 {
			return
		}
		payload := tgtypes.TransferGatewayContractMappingConfirmed{}
		if err := proto.Unmarshal(ev.EncodedBody, &payload); err != nil {
			return
		}
		sink <- &payload
	})

	return &EventSub{
		subscriber: sub,
		closeFn:    tg.unsubChainEvents,
	}, nil
}

func (tg *DAppChainGateway) WatchTokenWithdrawalSigned(
	sink chan<- *tgtypes.TransferGatewayTokenWithdrawalSigned) (*EventSub, error) {
	if tg.ws == nil {
		return nil, errors.New("websocket events unavailable")
	}

	if err := tg.subChainEvents(); err != nil {
		return nil, err
	}

	sub := tg.chainEventHub.Subscribe("event")
	sub.Do(func(msg pubsub.Message) {
		ev := lptypes.EventData{}
		if err := proto.Unmarshal(msg.Body(), &ev); err != nil {
			return
		}
		if ev.Topics == nil || ev.Topics[0] != tokenWithdrawalSignedEventTopic {
			return
		}
		contractAddr := diadem.UnmarshalAddressPB(ev.Address)
		if contractAddr.Compare(tg.Address) != 0 {
			return
		}
		payload := tgtypes.TransferGatewayTokenWithdrawalSigned{}
		if err := proto.Unmarshal(ev.EncodedBody, &payload); err != nil {
			return
		}
		sink <- &payload
	})

	return &EventSub{
		subscriber: sub,
		closeFn:    tg.unsubChainEvents,
	}, nil
}

func (tg *DAppChainGateway) subChainEvents() error {
	tg.chainEventSubCount++
	if tg.chainEventSubCount > 1 {
		return nil // already subscribed
	}

	err := tg.ws.WriteJSON(&client.RPCRequest{
		Version: "2.0",
		Method:  "subevents",
		ID:      strconv.FormatUint(tg.nextMsgID, 10),
	})
	tg.nextMsgID++

	if err != nil {
		return errors.Wrap(err, "failed to subscribe to DAppChain events")
	}

	resp := client.RPCResponse{}
	if err = tg.ws.ReadJSON(&resp); err != nil {
		return errors.Wrap(err, "failed to subscribe to DAppChain events")
	}
	if resp.Error != nil {
		return errors.Wrap(resp.Error, "failed to subscribe to DAppChain events")
	}

	tg.chainEventHub = pubsub.New()
	tg.chainEventQuitCh = make(chan struct{})

	go pumpChainEvents(tg.ws, tg.chainEventHub, tg.chainEventQuitCh)

	return nil
}

func (tg *DAppChainGateway) unsubChainEvents() {
	tg.chainEventSubCount--
	if tg.chainEventSubCount > 0 {
		return // still have subscribers
	}

	close(tg.chainEventQuitCh)

	tg.ws.WriteJSON(&client.RPCRequest{
		Version: "2.0",
		Method:  "unsubevents",
		ID:      strconv.FormatUint(tg.nextMsgID, 10),
	})
	tg.nextMsgID++
}

func pumpChainEvents(ws *websocket.Conn, hub pubsub.Hub, quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			return
		default:
			// TODO: handle ping control messages sent by diadem node to keep the connection alive,
			// otherwise diadem node will close the connection after 30 secs of inactivity.
			resp := client.RPCResponse{}
			if err := ws.ReadJSON(&resp); err != nil {
				panic(err)
			}
			if resp.Error != nil {
				panic(resp.Error)
			}
			unmarshaller := jsonpb.Unmarshaler{}
			reader := bytes.NewBuffer(resp.Result)
			eventData := lptypes.EventData{}
			if err := unmarshaller.Unmarshal(reader, &eventData); err != nil {
				panic(err)
			}
			bytes, err := proto.Marshal(&eventData)
			if err != nil {
				panic(err)
			}
			hub.Publish(pubsub.NewMessage("event", bytes))
		}
	}
}

/**

  Connectors

*/

func ConnectToDAppChaindiademGateway(diademClient *client.DAppChainRPCClient, eventsURI string) (*DAppChainGateway, error) {
	return connectToDAppChainGateway(diademClient, eventsURI, "diademcoin-gateway")
}

func ConnectToDAppChainGateway(diademClient *client.DAppChainRPCClient, eventsURI string) (*DAppChainGateway, error) {
	return connectToDAppChainGateway(diademClient, eventsURI, "gateway")
}

func ConnectToDAppChainTronGateway(diademClient *client.DAppChainRPCClient, eventsURI string) (*DAppChainGateway, error) {
	return connectToDAppChainGateway(diademClient, eventsURI, "tron-gateway")
}

func ConnectToDAppChainBinanceGateway(diademClient *client.DAppChainRPCClient, eventsURI string) (*DAppChainGateway, error) {
	return connectToDAppChainGateway(diademClient, eventsURI, "binance-gateway")
}

func connectToDAppChainGateway(diademClient *client.DAppChainRPCClient, eventsURI string, name string) (*DAppChainGateway, error) {
	gatewayAddr, err := diademClient.Resolve(name)
	if err != nil {
		return nil, err
	}

	fmt.Printf("DAppChain Gateway at %v\n", gatewayAddr)

	var ws *websocket.Conn
	if eventsURI != "" {
		ws, _, err = websocket.DefaultDialer.Dial(eventsURI, nil)
		if err != nil {
			return nil, err
		}
	}

	return &DAppChainGateway{
		contract: client.NewContract(diademClient, gatewayAddr.Local),
		chainID:  diademClient.GetChainID(),
		ws:       ws,
		Address:  gatewayAddr,
	}, nil
}
