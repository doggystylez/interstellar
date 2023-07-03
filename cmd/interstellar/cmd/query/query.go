package query

import (
	"fmt"

	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
	"github.com/spf13/cobra"
)

func QueryCmd() (qyCmd *cobra.Command) {
	qyCmd = &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Query chain via gRPC",
		Long:    "Query chain via gRPC",
	}
	cmds := flags.AddFlags([]*cobra.Command{accountCmd(), addressCmd(), balanceCmd()}, flags.KeySigningFlags, flags.GlobalFlags)
	cmds = flags.AddFlags(append(cmds, chainCmd()), flags.QueryFlags)
	qyCmd.AddCommand(cmds...)
	return
}

func chainCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "chain-id",
		Short: "Query chain-id",
		Long:  "Query chain-id",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			rpc := flags.ProcessQueryFlags(cmd)
			chainId, err := query.GetChainId(rpc)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(chainId)) //nolint
		},
	}
	return
}

func accountCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "account <address>",
		Short: "Query account info",
		Long:  "Query account info by address, keyname, or privkey",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			config, err := flags.ProcessGlobalFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.Rpc = flags.ProcessQueryFlags(cmd)
			if len(args) == 1 {
				config.TxInfo.Address = args[0]
			} else {
				config.TxInfo.KeyInfo.KeyRing, err = flags.ProcessKeySigningFlags(cmd)
				if err != nil {
					panic(err)
				}
				flags.CheckAddress(&config)
			}
			account, err := query.GetAccountInfoFromAddress(config.TxInfo.Address, config.Rpc)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(account)) //nolint
		},
	}
	return
}

func addressCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "address",
		Short: "Query account address",
		Long:  "Query account address by keyname or privkey",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			config, err := flags.ProcessGlobalFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.Rpc = flags.ProcessQueryFlags(cmd)
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProcessKeySigningFlags(cmd)
			if err != nil {
				panic(err)
			}
			flags.CheckAddress(&config)
			fmt.Println(query.Jsonify(query.AddressRes{Address: config.TxInfo.Address})) //nolint
		},
	}
	return
}

func balanceCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "balance <address>",
		Short: "Query account balance",
		Long:  "Query account balance, with optional denom filter",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				balances interface{}
			)
			config, err := flags.ProcessGlobalFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.Rpc = flags.ProcessQueryFlags(cmd)
			if len(args) == 1 {
				config.TxInfo.Address = args[0]
			} else {
				config.TxInfo.KeyInfo.KeyRing, err = flags.ProcessKeySigningFlags(cmd)
				if err != nil {
					panic(err)
				}
				flags.CheckAddress(&config)
			}
			denom, err := cmd.Flags().GetString("denom")
			if err != nil {
				return
			}
			if denom == "" {
				resp, err := query.GetAllBalances(config.TxInfo.Address, config.Rpc)
				if err != nil {
					panic(err)
				}
				balances = resp.Balances

			} else {
				resp, err := query.GetBalanceByDenom(config.TxInfo.Address, denom, config.Rpc)
				if err != nil {
					panic(err)
				}
				balances = resp
			}
			fmt.Println(query.Jsonify(balances)) //nolint
		},
	}
	cmd.Flags().StringP("denom", "d", "", "denom")
	return
}
