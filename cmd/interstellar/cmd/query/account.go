package query

import (
	"fmt"

	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
	"github.com/spf13/cobra"
)

func accountCmd() (accountCmd *cobra.Command) {
	accountCmd = &cobra.Command{
		Use:   "account",
		Short: "Account and address queries",
		Long:  "Account and address queries",
	}
	cmds := flags.AddFlags([]*cobra.Command{infoCmd(), addressCmd(), balanceCmd()}, flags.KeySigningFlags, flags.GlobalFlags, flags.QueryFlags)
	accountCmd.AddCommand(cmds...)
	return
}

func infoCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "info <address>",
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
			account, err := query.AccountInfoFromAddress(config.TxInfo.Address, config.Rpc)
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
				resp, err := query.AllBalances(config.TxInfo.Address, config.Rpc)
				if err != nil {
					panic(err)
				}
				balances = resp.Balances

			} else {
				resp, err := query.BalanceByDenom(config.TxInfo.Address, denom, config.Rpc)
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
