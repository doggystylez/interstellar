package cmd

import (
	keycli "github.com/doggystylez/interstellar/cmd/interstellar/cmd/keys"
	querycli "github.com/doggystylez/interstellar/cmd/interstellar/cmd/query"
	txcli "github.com/doggystylez/interstellar/cmd/interstellar/cmd/tx"

	"github.com/spf13/cobra"
)

func RootCmd() {
	rootCmd := &cobra.Command{Use: "interstellar"}
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(txcli.TxCmd(), querycli.QueryCmd(), keycli.KeysCmd())
	_ = rootCmd.Execute()
}
