package query

import (
	"fmt"
	"strconv"

	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
	"github.com/spf13/cobra"
)

func swapCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "swap",
		Short: "Osmosis swap queries",
		Long:  "Osmosis swap queries",
	}
	// cmds := flags.AddFlags([]*cobra.Command{priceCmd(), estimateCmd()}, flags.GlobalFlags, flags.QueryFlags)
	cmds := flags.AddFlags([]*cobra.Command{priceCmd()}, flags.GlobalFlags, flags.QueryFlags)
	cmd.AddCommand(cmds...)
	return
}

func priceCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "spot-price <pool> <denom_in> <denom_out>",
		Short: "Query Osmosis spot price",
		Long:  "Query Osmosis spot price",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			rpc := flags.ProcessQueryFlags(cmd)
			pool, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			price, err := query.SpotPrice(pool, args[1], args[2], rpc)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(price)) //nolint
		},
	}
	return
}

// func estimateCmd() (cmd *cobra.Command) {
// 	cmd = &cobra.Command{
// 		Use:   "estimate <pool> <amount> <denom_in> <denom_out>",
// 		Short: "Estimate Osmosis swap",
// 		Long:  "Estimate Osmosis swap",
// 		Args:  cobra.ExactArgs(4),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			rpc := flags.ProcessQueryFlags(cmd)
// 			pool, err := strconv.Atoi(args[0])
// 			if err != nil {
// 				panic(err)
// 			}
// 			price, err := query.EstimateSwapSinglePool(pool, args[1], args[2], args[3], rpc)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println(query.Jsonify(price)) //nolint
// 		},
// 	}
// 	return
// }
