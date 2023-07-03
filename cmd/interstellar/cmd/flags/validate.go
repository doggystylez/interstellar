package flags

import (
	"fmt"
	"os"
	"strings"

	"github.com/doggystylez/interstellar/client/keys"
	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/types"
)

func CheckTxInfo(config *types.InterstellarConfig) {
	CheckAddress(config)
	checkAccountInfo(config)
}

func LoadKey(config *types.InterstellarConfig) (err error) {
	if len(config.TxInfo.KeyInfo.KeyRing.KeyBytes) == 0 {
		if !keys.Exists(config.TxInfo.KeyInfo.KeyRing.KeyName, config.Path, config.TxInfo.KeyInfo.KeyRing.Backend) {
			fmt.Println("key named", "`"+config.TxInfo.KeyInfo.KeyRing.KeyName+"`", "does not exist") //nolint
			os.Exit(1)
		}
		config.TxInfo.KeyInfo.KeyRing.KeyBytes, err = keys.Load(config.TxInfo.KeyInfo.KeyRing.KeyName, config.Path, config.TxInfo.KeyInfo.KeyRing.Backend, "")
		if err != nil {
			return
		}
	}
	return
}

func CheckAddress(config *types.InterstellarConfig) {
	checkChainId(config)
	if config.TxInfo.Address == "" {
		config.TxInfo.Address = keys.LoadAddress(config.TxInfo.KeyInfo.KeyRing.KeyName, config.TxInfo.KeyInfo.ChainId, config.Path, true)
		if config.TxInfo.Address == "" {
			err := LoadKey(config)
			if err != nil {
				panic(err)
			}
			address, err := query.GetAddressPrefix(config.Rpc)
			if err != nil {
				panic(err)
			}
			config.TxInfo.Address, err = keys.BechAddress(address, config.TxInfo.KeyInfo.KeyRing.KeyBytes)
			if err != nil {
				panic(err)
			}
		}
	}
}

func checkChainId(config *types.InterstellarConfig) {
	if config.TxInfo.KeyInfo.ChainId == "" {
		chainId, err := query.GetChainId(config.Rpc)
		if err != nil {
			panic(err)
		}
		config.TxInfo.KeyInfo.ChainId = chainId.ChainId
	}
}

func checkAccountInfo(config *types.InterstellarConfig) {
	if config.TxInfo.KeyInfo.AccNum == 0 || config.TxInfo.KeyInfo.SeqNum == 0 {
		account, err := query.GetAccountInfoFromAddress(config.TxInfo.Address, config.Rpc)
		if err != nil {
			if strings.Contains(err.Error(), "NotFound") {
				config.TxInfo.KeyInfo.AccNum, config.TxInfo.KeyInfo.SeqNum = 0, 0
			} else {
				panic(err)
			}
		} else {
			config.TxInfo.KeyInfo.AccNum, config.TxInfo.KeyInfo.SeqNum = account.Account, account.Sequence
		}
	}
}
