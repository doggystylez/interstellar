package txcli

import (
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
	"github.com/doggystylez/interstellar/types"
	"github.com/spf13/cobra"
)

var (
	msgInfo types.MsgInfo
	resp    types.TxResponse
)

func TxCmd() (txCmd *cobra.Command) {
	txCmd = &cobra.Command{
		Use:     "transact",
		Aliases: []string{"tx"},
		Short:   "Send a transaction via gRPC",
		Long:    "Send a transaction via gRPC. Queries chain and account info if not provided",
	}
	cmds := flags.AddFlags([]*cobra.Command{sendCmd(), sendAllCmd(), transferCmd(), transferAllCmd()}, flags.TxFlags, flags.KeyFlags, flags.QueryFlags, flags.GlobalFlags)
	txCmd.AddCommand(cmds...)
	return
}
