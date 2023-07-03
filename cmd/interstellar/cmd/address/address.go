package address

import (
	"fmt"

	"github.com/doggystylez/interstellar/client/keys"
	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
	"github.com/spf13/cobra"
)

func AddressCmd() (addrCmd *cobra.Command) {
	addrCmd = &cobra.Command{
		Use:     "address",
		Aliases: []string{"a"},
		Short:   "manage addresses",
		Long:    "manage address book",
	}
	cmds := flags.AddFlags([]*cobra.Command{fetchCmd(), saveCmd()}, flags.QueryFlags, flags.KeyManageFlags, flags.GlobalFlags)
	addrCmd.AddCommand(cmds...)
	return
}

func fetchCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "fetch",
		Short: "fetch address from chain",
		Long:  "fetch address from chain and save to address book. Key name required",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			config, err := flags.ProcessGlobalFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.Rpc = flags.ProcessQueryFlags(cmd)
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProccesKeyManageFlags(cmd)
			if err != nil {
				panic(err)
			}
			flags.CheckTxInfo(&config)
			if keys.AddressExists(config.TxInfo.KeyInfo.KeyRing.KeyName, config.TxInfo.KeyInfo.ChainId, config.Path, true) {
				fmt.Println("address for key", config.TxInfo.KeyInfo.KeyRing.KeyName, "already exists for", config.TxInfo.KeyInfo.ChainId) //nolint
				return
			}
			err = keys.SaveAccountInfo(config.TxInfo.KeyInfo.KeyRing.KeyName, config.TxInfo.Address,
				config.TxInfo.KeyInfo.ChainId, config.TxInfo.KeyInfo.AccNum, config.TxInfo.KeyInfo.SeqNum, config.Path)
			if err != nil {
				panic(err)
			} else {
				fmt.Println("address", config.TxInfo.Address, "saved to address book") //nolint
			}
		},
	}
	return
}
func saveCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "save <address> <chain-id>",
		Short: "Save address to address book",
		Long:  "Save an external address to address book. Key name required, chain id optional, will be queried if blank",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var chainId string
			address := args[0]
			config, err := flags.ProcessGlobalFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.Rpc = flags.ProcessQueryFlags(cmd)
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProccesKeyManageFlags(cmd)
			if err != nil {
				panic(err)
			}
			if len(args) > 1 {
				chainId = args[1]
			} else {
				var res query.ChainIdRes
				res, err = query.GetChainId(config.Rpc)
				if err != nil {
					panic(err)
				}
				chainId = res.ChainId
			}
			err = keys.SaveToAddrbook(config.TxInfo.KeyInfo.KeyRing.KeyName, address, chainId, config.Path, false)
			if err != nil {
				panic(err)
			} else {
				fmt.Println("address", address, "saved to address book") //nolint
			}
		},
	}
	return
}
