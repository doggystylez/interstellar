package keys

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

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
	cmds := flags.AddFlags([]*cobra.Command{newCmd(), restoreCmd()}, flags.KeyManageFlags, flags.GlobalFlags)
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
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProccesKeyManageFlags(cmd)
			if err != nil {
				panic(err)
			}
			mnemonic, bytes, err := keys.NewKeyWithSeed()
			if err != nil {
				panic(err)
			}
			err = keys.Save(bytes, config.TxInfo.KeyInfo.KeyRing.KeyName, config.Path, config.TxInfo.KeyInfo.KeyRing.Backend, "")
			if err != nil {
				panic(err)
			}
			fmt.Println("your new seed phrase is:", mnemonic) //nolint
		},
	}
	return
}

func restoreCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "restore",
		Short: "restore an existing key",
		Long:  "restore an existing key from mnemonic or hex",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			var bytes []byte
			config, err := flags.ProcessGlobalFlags(cmd)
			if err != nil {
				panic(err)
			}
			config.TxInfo.KeyInfo.KeyRing, err = flags.ProccesKeyManageFlags(cmd)
			if err != nil {
				panic(err)
			}
			isHex, err := cmd.Flags().GetBool("hex")
			if err != nil {
				panic(err)
			}
			if !isHex {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("enter your mnemonic: ") //nolint
				mnemonic, err := reader.ReadString('\n')
				if err != nil {
					panic(err)
				}
				bytes, err = keys.KeyFromSeed(mnemonic)
				if err != nil {
					panic(err)
				}
			} else {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("enter your hex key: ") //nolint
				privKey, err := reader.ReadString('\n')
				if err != nil {
					panic(err)
				}
				bytes, err = hex.DecodeString(strings.TrimSuffix(privKey, string('\n')))
				if err != nil {
					panic(err)
				}
			}
			err = keys.Save(bytes, config.TxInfo.KeyInfo.KeyRing.KeyName, config.Path, config.TxInfo.KeyInfo.KeyRing.Backend, "")
			if err != nil {
				panic(err)
			} else {
				fmt.Println("saved key to keyring") //nolint
			}
		},
	}
	cmd.Flags().Bool("hex", false, "restore from hex encoded private key")
	return
}
