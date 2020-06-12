// +build evm

package client

import (
	"crypto/ecdsa"
	"encoding/base64"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	diadem "github.com/diademnetwork/go-diadem"
	"github.com/diademnetwork/go-diadem/auth"
)

type Identity struct {
	MainnetPrivKey *ecdsa.PrivateKey
	MainnetAddr    common.Address
	diademSigner     auth.Signer
	diademAddr       diadem.Address
}

func CreateIdentity(mainnetPrivKey *ecdsa.PrivateKey, diademSigner auth.Signer, chainID string) (*Identity, error) {
	var mainnetAddr common.Address
	if mainnetPrivKey != nil {
		mainnetAddr = crypto.PubkeyToAddress(mainnetPrivKey.PublicKey)
	} else {
		mainnetAddr = common.HexToAddress("0x0")
	}
	identity := &Identity{
		MainnetPrivKey: mainnetPrivKey,
		MainnetAddr:    mainnetAddr,
		diademSigner:     diademSigner,
		diademAddr: diadem.Address{
			ChainID: chainID,
			Local:   diadem.LocalAddressFromPublicKey(diademSigner.PublicKey()),
		},
	}
	return identity, nil
}

func CreateIdentityStr(ethKey string, dappchainKey string, chainID string) (*Identity, error) {
	mainnetPrivKey, err := crypto.HexToECDSA(strings.TrimPrefix(ethKey, "0x"))
	if err != nil {
		return nil, err
	}
	keyBytes, err := base64.StdEncoding.DecodeString(dappchainKey)
	if err != nil {
		return nil, err
	}
	diademSigner := auth.NewEd25519Signer(keyBytes)
	identity := &Identity{
		MainnetPrivKey: mainnetPrivKey,
		MainnetAddr:    crypto.PubkeyToAddress(mainnetPrivKey.PublicKey),
		diademSigner:     diademSigner,
		diademAddr: diadem.Address{
			ChainID: chainID,
			Local:   diadem.LocalAddressFromPublicKey(diademSigner.PublicKey()),
		},
	}
	return identity, nil
}

func diademAddressFromEthereumAddress(ethAddr common.Address) (diadem.Address, error) {
	addrBytes, err := diadem.LocalAddressFromHexString(ethAddr.Hex())
	if err != nil {
		return diadem.Address{}, err
	}
	return diadem.Address{
		ChainID: "eth",
		Local:   addrBytes,
	}, nil
}

func diademAddressFromTronAddress(tronAddr common.Address) (diadem.Address, error) {
	addrBytes, err := diadem.LocalAddressFromHexString(tronAddr.Hex())
	if err != nil {
		return diadem.Address{}, err
	}
	return diadem.Address{
		ChainID: "tron",
		Local:   addrBytes,
	}, nil
}

func diademAddressFromBinanceAddress(binanceAddr common.Address) diadem.Address {
	return diadem.Address{ChainID: "binance", Local: binanceAddr.Bytes()}
}
