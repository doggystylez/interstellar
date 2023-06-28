package txcli

import (
	"fmt"
	"strconv"

	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/client/tx"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
	"github.com/spf13/cobra"
)

func transferCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "transfer <to> <amount> <denom>",
		Short: "Transfer a specific amount of a denom via IBC",
		Long:  "Transfer a specific amount of a denom via IBC. Key name or private key are required",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			config, err := flags.ProcessGlobalFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.TxInfo, err = flags.ProcessTxFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProcessKeyFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.Rpc, err = flags.ProcessQueryFlags(cmd)
			if err != nil {
				panic(err)
			}
			err = flags.CheckTxInfo(&config)
			if err != nil {
				panic(err)
			}
			msgInfo.Amount, err = strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				panic(err)
			}
			msgInfo.Channel, err = cmd.Flags().GetString("channel-id")
			if err != nil {
				panic(err)
			}
			msgInfo.From, msgInfo.To, msgInfo.Denom = config.TxInfo.Address, args[0], args[2]
			resp, err := tx.AssembleAndBroadcast(msgInfo, config.TxInfo, config.Rpc, tx.MakeTransferMsg)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(resp))
		},
	}
	return
}

func transferAllCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "transfer-all <to> <denom>",
		Short: "Transfer all of a specific amount of a denom via IBC",
		Long:  "Transfer all of a specific amount of a denom via IBC. Key name or private key are required",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			config, err := flags.ProcessGlobalFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.TxInfo, err = flags.ProcessTxFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProcessKeyFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.Rpc, err = flags.ProcessQueryFlags(cmd)
			if err != nil {
				panic(err)
			}
			err = flags.CheckTxInfo(&config)
			if err != nil {
				panic(err)
			}
			msgInfo.Channel, err = cmd.Flags().GetString("channel-id")
			if err != nil {
				panic(err)
			}
			amount := query.GetBalanceByDenom(config.TxInfo.Address, args[1], config.Rpc)
			msgInfo.From, msgInfo.To, msgInfo.Denom = config.TxInfo.Address, args[0], args[1]
			if config.TxInfo.FeeDenom == msgInfo.Denom {
				msgInfo.Amount = amount.Amount - config.TxInfo.FeeAmount
			} else {
				msgInfo.Amount = amount.Amount
			}
			resp, err := tx.AssembleAndBroadcast(msgInfo, config.TxInfo, config.Rpc, tx.MakeTransferMsg)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(resp))
		},
	}
	return
}
