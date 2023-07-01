package keys

import (
	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func FromBytes(key []byte) *secp256k1.PrivKey {
	return &secp256k1.PrivKey{Key: key}
}

func BechAddress(prefix string, keyBytes []byte) (string, error) {
	return bech32.EncodeFromBase256(prefix, FromBytes(keyBytes).PubKey().Address())
}

func DecodeBech(address string) (string, []byte, error) {
	return bech32.DecodeToBase256(address)
}
