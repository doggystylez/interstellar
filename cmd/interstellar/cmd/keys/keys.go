package keycli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
	"github.com/doggystylez/interstellar/keys/keyring"
	"github.com/doggystylez/interstellar/keys/mnemonic"
	"github.com/spf13/cobra"
)

func KeysCmd() (keyCmd *cobra.Command) {
	keyCmd = &cobra.Command{
		Use:     "keys",
		Aliases: []string{"k"},
		Short:   "manage keys",
		Long:    "manage keys",
	}
	cmds := flags.AddFlags([]*cobra.Command{newCmd(), restoreCmd()}, flags.KeyFlags, flags.GlobalFlags)
	keyCmd.AddCommand(cmds...)
	return
}

func newCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "new",
		Short: "add a new key",
		Long:  "add a new key to keyring",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			config, err := flags.ProcessGlobalFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProcessKeyFlags(cmd)
			if err != nil {
				panic(err)
			}
			mnemonic, bytes, err := mnemonic.NewKeyWithSeed(config.TxInfo.KeyInfo.KeyRing)
			if err != nil {
				panic(err)
			}
			err = keyring.Save(config, bytes)
			if err != nil {
				panic(err)
			}
			fmt.Println("your new seed phrase is:", mnemonic)
		},
	}
	return
}

func restoreCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "restore",
		Short: "restore an existing key",
		Long:  "restore an existing key",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			config, err := flags.ProcessGlobalFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProcessKeyFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.TxInfo.KeyInfo.KeyRing.Mnemonic, err = input.GetString("enter your mnemonic", bufio.NewReader(os.Stdin))
			if err != nil {
				panic(err)
			}
			bytes, err := mnemonic.KeyFromSeed(config.TxInfo.KeyInfo.KeyRing)
			if err != nil {
				panic(err)
			}
			err = keyring.Save(config, bytes)
			if err != nil {
				panic(err)
			}
		},
	}
	return
}
