package yubihsm

import (
	"github.com/diademnetwork/go-diadem/crypto"
)

// YubiHsmSigner implements signer by YubiHSM secp256k1
type YubiHsmSigner struct {
	PrivateKey crypto.PrivateKey
}
