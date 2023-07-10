package flags

import (
	"github.com/doggystylez/interstellar/client/grpc"
	"github.com/spf13/cobra"
)

func ProcessQueryFlags(cmd *cobra.Command) grpc.Client {
	endpoint, err := cmd.Flags().GetString("node")
	if err != nil {
		panic(err)
	}
	timeout, err := cmd.Flags().GetInt("timeout")
	if err != nil {
		panic(err)
	}
	retries, err := cmd.Flags().GetInt("retries")
	if err != nil {
		panic(err)
	}
	return grpc.New(endpoint, timeout, retries, 2)
}

func QueryFlags(rawCmds ...*cobra.Command) (cmds []*cobra.Command) {
	for _, cmd := range rawCmds {
		cmd.Flags().StringP("node", "n", "localhost:9090", "gRPC server address")
		cmd.Flags().IntP("retries", "r", 0, "gRPC retries")
		cmd.Flags().IntP("timeout", "t", 1, "gRPC timeout in seconds")
		cmds = append(cmds, cmd)
	}
	return
}
