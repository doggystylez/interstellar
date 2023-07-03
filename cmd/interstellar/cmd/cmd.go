package cmd

import (
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/address"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/keys"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/query"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/tx"
	"github.com/spf13/cobra"
)

func RootCmd() {
	rootCmd := &cobra.Command{Use: "interstellar"}
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(tx.TxCmd(), query.QueryCmd(), keys.KeysCmd(), address.AddressCmd())
	_ = rootCmd.Execute()
}
