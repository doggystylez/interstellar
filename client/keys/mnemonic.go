package keys

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/go-bip39"
)

func KeyFromSeed(mnemonic string) (bytes []byte, err error) {
	algo := hd.Secp256k1
	bytes, err = algo.Derive()(mnemonic, "", "m/44'/118'/0'/0/0")
	return
}

func NewKeyWithSeed() (mnemonic string, bytes []byte, err error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return
	}
	mnemonic, err = bip39.NewMnemonic(entropy)
	if err != nil {
		return
	}
	bytes, err = KeyFromSeed(mnemonic)
	return
}
