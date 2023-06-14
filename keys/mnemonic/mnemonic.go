package mnemonic

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/go-bip39"
	"github.com/doggystylez/interstellar/types"
)

func KeyFromSeed(key types.KeyRing) (bytes []byte, err error) {
	algo := hd.Secp256k1
	bytes, err = algo.Derive()(key.Mnemonic, "", "m/44'/118'/0'/0/0")
	return
}

func NewKeyWithSeed(key types.KeyRing) (mnemonic string, bytes []byte, err error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return
	}
	mnemonic, err = bip39.NewMnemonic(entropy)
	if err != nil {
		return
	}
	bytes, err = KeyFromSeed(types.KeyRing{Mnemonic: mnemonic})
	return
}
