package flags

import (
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

func LoadKey(config *types.InterstellarConfig) (err error) {
	if len(config.TxInfo.KeyInfo.KeyRing.KeyBytes) == 0 {
		config.TxInfo.KeyInfo.KeyRing.KeyBytes, err = keys.Load(config.TxInfo.KeyInfo.KeyRing.KeyName, config.Path, "")
		if err != nil {
			return
		}
	}
	return
}

func CheckAddress(config *types.InterstellarConfig) (err error) {
	if config.TxInfo.Address == "" {
		err = LoadKey(config)
		if err != nil {
			return
		}
		address, err := query.GetAddressPrefix(config.Rpc)
		if err != nil {
			panic(err)
		}
		config.TxInfo.Address, err = keys.BechAddress(address, config.TxInfo.KeyInfo.KeyRing.KeyBytes)
	}
	return
}

func CheckSigningInfo(config *types.InterstellarConfig) (err error) {
	if config.TxInfo.KeyInfo.AccNum == 0 || config.TxInfo.KeyInfo.SeqNum == 0 {
		account, err := query.GetAccountInfoFromAddress(config.TxInfo.Address, config.Rpc)
		if err != nil {
			panic(err)
		}
		config.TxInfo.KeyInfo.AccNum, config.TxInfo.KeyInfo.SeqNum = account.Account, account.Sequence

	}
	if config.TxInfo.KeyInfo.ChainId == "" {
		chainId, err := query.GetChainId(config.Rpc)
		if err != nil {
			panic(err)
		}
		config.TxInfo.KeyInfo.ChainId = chainId.ChainId
	}
	return
}
