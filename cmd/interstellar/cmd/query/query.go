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
	cmds := flags.AddFlags([]*cobra.Command{chainCmd(), txCmd()}, flags.QueryFlags)
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
