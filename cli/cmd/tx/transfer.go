package tx

import (
	"fmt"
	"strconv"

	"github.com/doggystylez/interstellar/cli/cmd/flags"
	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/client/tx"
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
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProcessKeySigningFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.Rpc = flags.ProcessQueryFlags(cmd)
			flags.CheckTxInfo(&config)
			msgInfo.Channel, err = cmd.Flags().GetString("channel-id")
			if err != nil {
				panic(err)
			}
			msgInfo.From, msgInfo.To, msgInfo.Amount, msgInfo.Denom = config.TxInfo.Address, args[0], args[1], args[2]
			msgInfo.Maker = tx.MakeTransferMsg
			resp, err := tx.AssembleAndBroadcast([]tx.MsgInfo{msgInfo}, config.TxInfo, config.Rpc)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(resp)) //nolint
		},
	}
	cmd.Flags().StringP("channel-id", "i", "", "ibc channel")
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
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProcessKeySigningFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.Rpc = flags.ProcessQueryFlags(cmd)
			flags.CheckTxInfo(&config)
			msgInfo.Channel, err = cmd.Flags().GetString("channel-id")
			if err != nil {
				panic(err)
			}
			token, err := query.BalanceByDenom(config.TxInfo.Address, args[1], config.Rpc)
			if err != nil {
				panic(err)
			}
			msgInfo.From, msgInfo.To, msgInfo.Denom = config.TxInfo.Address, args[0], args[1]
			if config.TxInfo.FeeDenom == msgInfo.Denom {
				var amount uint64
				amount, err = strconv.ParseUint(token.Amount, 10, 64)
				if err != nil {
					panic(err)
				}
				msgInfo.Amount = strconv.FormatUint(amount-config.TxInfo.FeeAmount, 10)
			} else {
				msgInfo.Amount = token.Amount
			}
			msgInfo.Maker = tx.MakeTransferMsg
			resp, err := tx.AssembleAndBroadcast([]tx.MsgInfo{msgInfo}, config.TxInfo, config.Rpc)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(resp)) //nolint
		},
	}
	cmd.Flags().StringP("channel-id", "i", "", "ibc channel")
	return
}
