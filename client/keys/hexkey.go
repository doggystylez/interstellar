package keys

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func PrivKeyfromString(key string) (privKey *secp256k1.PrivKey, err error) {
	bytes, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	privKey = &secp256k1.PrivKey{
		Key: bytes,
	}
	return
}
