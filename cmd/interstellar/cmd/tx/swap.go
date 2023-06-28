package txcli

// import (
// "fmt"
// "strconv"

// "github.com/doggystylez/interstellar/client/query"
// "github.com/doggystylez/interstellar/client/tx"
// "github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
// "github.com/spf13/cobra"
// )

// func swapCmd() (cmd *cobra.Command) {
// 	cmd = &cobra.Command{
// 		Use:   "swap <amount> <denom>",
// 		Short: "Swap on osmosis",
// 		Long:  "Send on osmosis. Key name or private key are required",
// 		Args:  cobra.ExactArgs(3),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			config, err := flags.ProcessGlobalFlags(cmd)
// 			if err != nil {
// 				panic(err)
// 			}
// 			config.TxInfo, err = flags.ProcessTxFlags(cmd)
// 			if err != nil {
// 				panic(err)
// 			}
// 			config.TxInfo.KeyInfo.KeyRing, err = flags.ProcessKeyFlags(cmd)
// 			if err != nil {
// 				panic(err)
// 			}
// 			config.Rpc, err = flags.ProcessQueryFlags(cmd)
// 			if err != nil {
// 				panic(err)
// 			}
// 			err = flags.CheckTxInfo(&config)
// 			if err != nil {
// 				panic(err)
// 			}
// 			msgInfo.Amount, err = strconv.ParseUint(args[1], 10, 64)
// 			if err != nil {
// 				panic(err)
// 			}
// 			msgInfo.From, msgInfo.To, msgInfo.Denom = config.TxInfo.Address, args[0], args[2]
// 			resp, err := tx.AssembleAndBroadcast(msgInfo, config.TxInfo, config.Path, config.Rpc, tx.MakeSwapMsg)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println(query.Jsonify(resp))
// 		},
// 	}
// 	return
// }
