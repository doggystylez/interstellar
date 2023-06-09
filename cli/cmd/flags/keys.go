package flags

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/doggystylez/interstellar/client/keys"
	"github.com/spf13/cobra"
)

func ProccesKeyManageFlags(cmd *cobra.Command) (keyRing keys.KeyRing, err error) {
	keyRing.KeyName, err = cmd.Flags().GetString("key")
	if err != nil {
		return
	}
	keyRing.Backend, err = cmd.Flags().GetString("keyring-backend")
	if err != nil {
		return
	}
	if keyRing.Backend != "file" && keyRing.Backend != "test" {
		fmt.Println("invalid keyring backend", keyRing.Backend) //nolint
		os.Exit(1)
	}
	return
}

func ProcessKeySigningFlags(cmd *cobra.Command) (keyRing keys.KeyRing, err error) {
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
	if keyRing.Backend != "file" && keyRing.Backend != "test" {
		fmt.Println("invalid keyring backend", keyRing.Backend) //nolint
		os.Exit(1)
	}
	return
}

func KeyManageFlags(rawCmds ...*cobra.Command) (cmds []*cobra.Command) {
	for _, cmd := range rawCmds {
		cmd.Flags().StringP("key", "k", "", "key name")
		err := cmd.MarkFlagRequired("key")
		if err != nil {
			return
		}
		cmd.Flags().StringP("keyring-backend", "b", "file", "keyring backend")
		cmds = append(cmds, cmd)
	}
	return
}

func KeySigningFlags(rawCmds ...*cobra.Command) (cmds []*cobra.Command) {
	for _, cmd := range rawCmds {
		cmd.Flags().StringP("priv", "v", "", "private key")
		cmd.Flags().StringP("key", "k", "", "key name")
		cmd.Flags().StringP("keyring-backend", "b", "file", "keyring backend")
		cmds = append(cmds, cmd)
	}
	return
}
