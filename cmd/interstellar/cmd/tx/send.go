package txcli

import (
	"fmt"
	"strconv"

	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/client/tx"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
	"github.com/spf13/cobra"
)

func sendCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "send <to> <amount> <denom>",
		Short: "Send a specific amount of a denom",
		Long:  "Send a specific amount of a denom. Key name or private key are required",
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
			msgInfo.From, msgInfo.To, msgInfo.Denom = config.TxInfo.Address, args[0], args[2]
			resp, err := tx.AssembleAndBroadcast(msgInfo, config.TxInfo, config.Path, config.Rpc, tx.MakeSendMsg)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(resp))
		},
	}
	return
}

func sendAllCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "send-all <to> <denom>",
		Short: "Send all of a denom",
		Long:  "Send  all of a denom. Key name or private key are required",
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
			amount := query.GetBalanceByDenom(config.TxInfo.Address, args[1], config.Rpc)
			msgInfo.From, msgInfo.To, msgInfo.Denom = config.TxInfo.Address, args[0], args[1]
			if config.TxInfo.FeeDenom == msgInfo.Denom {
				msgInfo.Amount = amount.Amount - config.TxInfo.FeeAmount
			} else {
				msgInfo.Amount = amount.Amount
			}
			resp, err := tx.AssembleAndBroadcast(msgInfo, config.TxInfo, config.Path, config.Rpc, tx.MakeSendMsg)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(resp))
		},
	}
	return
}
