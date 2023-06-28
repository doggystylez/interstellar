package keycli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/doggystylez/interstellar/client/keys"
	"github.com/doggystylez/interstellar/cmd/interstellar/cmd/flags"
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
			if keys.Exists(config.TxInfo.KeyInfo.KeyRing.KeyName, config.Path) {
				fmt.Println("key named", "`"+config.TxInfo.KeyInfo.KeyRing.KeyName+"`", "already exists")
				return
			}
			mnemonic, bytes, err := keys.NewKeyWithSeed()
			if err != nil {
				panic(err)
			}
			err = keys.Save(config.TxInfo.KeyInfo.KeyRing.KeyName, config.Path, bytes, "")
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
			if keys.Exists(config.TxInfo.KeyInfo.KeyRing.KeyName, config.Path) {
				fmt.Println("key named", "`"+config.TxInfo.KeyInfo.KeyRing.KeyName+"`", "already exists")
				return
			}
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("enter your mnemonic")
			config.TxInfo.KeyInfo.KeyRing.Mnemonic, err = reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			bytes, err := keys.KeyFromSeed(config.TxInfo.KeyInfo.KeyRing.Mnemonic)
			if err != nil {
				panic(err)
			}
			err = keys.Save(config.TxInfo.KeyInfo.KeyRing.KeyName, config.Path, bytes, "")
			if err != nil {
				panic(err)
			} else {
				fmt.Println("saved key to keyring")
			}
		},
	}
	return
}
