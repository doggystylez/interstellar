package tx

import (
	"github.com/doggystylez/interstellar/cli/cmd/flags"
	"github.com/doggystylez/interstellar/client/tx"
	"github.com/spf13/cobra"
)

var msgInfo tx.MsgInfo

func TxCmd() (txCmd *cobra.Command) {
	txCmd = &cobra.Command{
		Use:     "transact",
		Aliases: []string{"tx"},
		Short:   "Send a transaction via gRPC",
		Long:    "Send a transaction via gRPC. Queries chain and account info if not provided",
	}
	cmds := flags.AddFlags([]*cobra.Command{sendCmd(), sendAllCmd(), transferCmd(), transferAllCmd(), swapCmd(), ibcSwapCmd()}, flags.TxFlags, flags.KeySigningFlags, flags.QueryFlags, flags.GlobalFlags)
	txCmd.AddCommand(cmds...)
	return
}
