package keys

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func FromBytes(key []byte) *secp256k1.PrivKey {
	return &secp256k1.PrivKey{Key: key}
}

// func to convert hrp + hex pubkey to bech
