package flags

import (
	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/keys/address"
	"github.com/doggystylez/interstellar/types"
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
		if config.TxInfo.KeyInfo.KeyRing.HexPriv != "" {
			config.TxInfo.Address, err = address.PrivKeyToAddress(*config)
			return
		} else if config.TxInfo.KeyInfo.KeyRing.KeyName != "" {
			config.TxInfo.Address, err = address.KeyNameToAddress(config)
			if err != nil {
				return
			}
		}
	}
	return
}

func CheckSigningInfo(config *types.InterstellarConfig) (err error) {
	if config.TxInfo.KeyInfo.AccNum == 0 || config.TxInfo.KeyInfo.SeqNum == 0 {
		var account types.AccountRes
		account, err = query.GetAccountInfoFromAddress(*config)
		if err != nil {
			return
		}
		config.TxInfo.KeyInfo.AccNum, config.TxInfo.KeyInfo.SeqNum = account.Account, account.Sequence

	}
	if config.TxInfo.KeyInfo.ChainId == "" {
		var chainId types.ChainIdRes
		chainId, err = query.GetChainId(config.Rpc)
		if err != nil {
			return
		}
		config.TxInfo.KeyInfo.ChainId = chainId.ChainId
	}
	return
}
