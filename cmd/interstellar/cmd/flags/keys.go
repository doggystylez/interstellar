package flags

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/doggystylez/interstellar/client/keys"
	"github.com/spf13/cobra"
)

func ProcessKeyFlags(cmd *cobra.Command) (keyRing keys.KeyRing, err error) {
	keyRing.KeyName, err = cmd.Flags().GetString("key")
	if err != nil {
		return
	}
	hexPriv, err := cmd.Flags().GetString("priv")
	if err != nil {
		return
	}
	if hexPriv != "" {
		if keyRing.KeyName != "" {
			fmt.Println("cannot use both key name and private key") //nolint
			os.Exit(1)
		}
		keyRing.KeyBytes, err = hex.DecodeString(hexPriv)
		if err != nil {
			return
		}
	} else if keyRing.KeyName == "" {
		fmt.Println("address, key name, or private key required") //nolint
		os.Exit(1)
	}
	keyRing.Backend, err = cmd.Flags().GetString("keyring-backend")
	if err != nil {
		return
	}
	return
}

func KeyFlags(rawCmds ...*cobra.Command) (cmds []*cobra.Command) {
	for _, cmd := range rawCmds {
		cmd.Flags().StringP("priv", "v", "", "private key")
		cmd.Flags().StringP("key", "k", "", "key name")
		cmd.Flags().StringP("keyring-backend", "b", "test", "keyring backend")
		cmds = append(cmds, cmd)
	}
	return
}
