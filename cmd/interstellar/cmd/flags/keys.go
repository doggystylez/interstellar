package flags

import (
	"fmt"
	"os"

	"github.com/doggystylez/interstellar/types"
	"github.com/spf13/cobra"
)

func ProcessKeyFlags(cmd *cobra.Command) (keyRing types.KeyRing, err error) {
	keyRing.HexPriv, err = cmd.Flags().GetString("priv")
	if err != nil {
		return
	}
	keyRing.KeyName, err = cmd.Flags().GetString("key")
	if err != nil {
		return
	}
	if keyRing.KeyName == "" && keyRing.HexPriv == "" {
		fmt.Println("address, key name, or private key required")
		os.Exit(1)
	}
	if keyRing.KeyName != "" && keyRing.HexPriv != "" {
		fmt.Println("cannot use both key name and private key")
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
