package flags

import (
	"github.com/cosmos/btcutil/bech32"
	"github.com/doggystylez/interstellar/client/grpc"
	"github.com/doggystylez/interstellar/client/keys"
	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/types"
)

func CheckTxInfo(config *types.InterstellarConfig) (err error) {
	err = CheckAddress(config)
	if err != nil {
		return
	}
	err = CheckSigningInfo(config)
	return
}

func CheckAddress(config *types.InterstellarConfig) (err error) {
	if config.TxInfo.Address == "" {
		if len(config.TxInfo.KeyInfo.KeyRing.KeyBytes) != 0 {
			config.TxInfo.Address, err = privKeyToAddress(config.TxInfo.KeyInfo.KeyRing.KeyBytes, config.Rpc)
			return
		} else if config.TxInfo.KeyInfo.KeyRing.KeyName != "" {
			config.TxInfo.Address, err = keyNameToAddress(&config.TxInfo.KeyInfo.KeyRing, config.Path, config.Rpc)
			if err != nil {
				return
			}
		}
	}
	return
}

func CheckSigningInfo(config *types.InterstellarConfig) (err error) {
	if config.TxInfo.KeyInfo.AccNum == 0 || config.TxInfo.KeyInfo.SeqNum == 0 {
		var account query.AccountRes
		account = query.GetAccountInfoFromAddress(config.TxInfo.Address, config.Rpc)
		config.TxInfo.KeyInfo.AccNum, config.TxInfo.KeyInfo.SeqNum = account.Account, account.Sequence

	}
	if config.TxInfo.KeyInfo.ChainId == "" {
		var chainId query.ChainIdRes
		chainId = query.GetChainId(config.Rpc)
		config.TxInfo.KeyInfo.ChainId = chainId.ChainId
	}
	return
}

func keyNameToAddress(keyRing *keys.KeyRing, path string, rpc grpc.Client) (address string, err error) {
	keyRing.KeyBytes, err = keys.Load(keyRing.KeyName, path)
	if err != nil {
		return
	}
	address, err = privKeyToAddress(keyRing.KeyBytes, rpc)
	return
}

func privKeyToAddress(keyBytes []byte, rpc grpc.Client) (address string, err error) {
	prefix := query.GetAddressPrefix(rpc)
	privKey := keys.FromBytes(keyBytes)
	address, err = bech32.EncodeFromBase256(prefix, privKey.PubKey().Address())
	return
}
