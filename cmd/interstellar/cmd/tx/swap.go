package txcli

import (
	"fmt"
	"strconv"

	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/client/tx"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
	"github.com/spf13/cobra"
)

const swapContract = "osmo18gp2d0p2ye66ek5ejznrpxv5lus0lw6lta45zz2hv0rhz3dd7kwquh8487"

func swapCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "swap <amount> <denom-in> <denom-out>",
		Short: "Swap",
		Long:  "Swap on osmosis. Key name or private key are required",
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
			msgInfo.Amount, err = strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				panic(err)
			}
			msgInfo.From, msgInfo.Denom = config.TxInfo.Address, args[1]

			slippage, err := cmd.Flags().GetString("slippage")
			if err != nil {
				panic(err)
			}
			msgInfo.Contract = swapContract
			msgInfo.ContractMsg = tx.MakeSwapContractMsg(args[0], args[1], args[2], slippage)
			resp, err := tx.AssembleAndBroadcast(msgInfo, config.TxInfo, config.Rpc, tx.MakeWasmMsg)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(resp)) //nolint
		},
	}
	cmd.Flags().StringP("slippage", "i", "2", "slippage %")
	return
}
