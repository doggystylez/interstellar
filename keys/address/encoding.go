package address

import (
	"encoding/hex"

	"github.com/cosmos/btcutil/bech32"
	"github.com/doggystylez/interstellar/keys/hexkey"
	"github.com/doggystylez/interstellar/keys/keyring"
	"github.com/doggystylez/interstellar/types"
)

func KeyNameToAddress(config *types.InterstellarConfig) (address string, err error) {
	privKey, err := keyring.Load(*config)
	if err != nil {
		return
	}
	config.TxInfo.KeyInfo.KeyRing.HexPriv = hex.EncodeToString(privKey)
	address, err = PrivKeyToAddress(*config)
	return
}

func PrivKeyToAddress(config types.InterstellarConfig) (address string, err error) {
	prefix, err := GetAddressPrefix(config.Rpc)
	if err != nil {
		return
	}
	privKey, err := hexkey.PrivKeyfromString(config.TxInfo.KeyInfo.KeyRing.HexPriv)
	if err != nil {
		return
	}
	address, err = bech32.EncodeFromBase256(prefix, privKey.PubKey().Address())
	return
}
