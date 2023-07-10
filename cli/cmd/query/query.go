package query

import (
	"fmt"
	"strconv"

	"github.com/doggystylez/interstellar/cli/cmd/flags"
	"github.com/doggystylez/interstellar/client/query"
	"github.com/spf13/cobra"
)

func QueryCmd() (qyCmd *cobra.Command) {
	qyCmd = &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Query chain via gRPC",
		Long:    "Query chain via gRPC",
	}
	cmds := flags.AddFlags([]*cobra.Command{chainCmd(), blockCmd(), txCmd()}, flags.QueryFlags)
	qyCmd.AddCommand(append(cmds, accountCmd(), swapCmd())...)
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
			chainId, err := query.ChainId(rpc)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(chainId)) //nolint
		},
	}
	return
}

func blockCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "block <height>",
		Short: "Query block by height",
		Long:  "Query block by height. Returns latest height if no height supplied",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var block query.BlockRes
			var height int64
			var err error
			rpc := flags.ProcessQueryFlags(cmd)
			if len(args) > 0 {
				height, err = strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					panic(err)
				}
				block, err = query.Block(height, rpc)
				if err != nil {
					panic(err)
				}
			} else {
				block, err = query.LatestBlock(rpc)
				if err != nil {
					panic(err)
				}
			}
			fmt.Println(query.Jsonify(block)) //nolint
		},
	}
	return
}

func txCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "tx <hash>",
		Short: "Query transaction",
		Long:  "Query transaction by hash",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rpc := flags.ProcessQueryFlags(cmd)
			tx, err := query.Tx(&args[0], rpc)
			if err != nil {
				panic(err)
			}
			fmt.Println(query.Jsonify(tx)) //nolint
		},
	}
	return
}
