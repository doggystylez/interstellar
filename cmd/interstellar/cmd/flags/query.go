package flags

import (
	"github.com/doggystylez/interstellar/types"
	"github.com/spf13/cobra"
)

func ProcessQueryFlags(cmd *cobra.Command) (rpc types.Client, err error) {
	rpc.Endpoint, err = cmd.Flags().GetString("node")
	if err != nil {
		return
	}
	rpc.Timeout, err = cmd.Flags().GetInt("timeout")
	if err != nil {
		return
	}
	return
}

func QueryFlags(rawCmds ...*cobra.Command) (cmds []*cobra.Command) {
	for _, cmd := range rawCmds {
		cmd.Flags().StringP("node", "n", "localhost:9090", "gRPC server address")
		cmd.Flags().IntP("timeout", "t", 1, "gRPC timeout")
		cmds = append(cmds, cmd)
	}
	return
}
