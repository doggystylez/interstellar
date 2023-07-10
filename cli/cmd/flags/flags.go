package flags

import (
	"os/user"
	"path/filepath"
	"strings"

	"github.com/doggystylez/interstellar/cli/cmd/types"
	"github.com/spf13/cobra"
)

func AddFlags(cmds []*cobra.Command, flags ...func(...*cobra.Command) []*cobra.Command) (flagCmds []*cobra.Command) {
	for _, cmd := range cmds {
		for _, flag := range flags {
			flag(cmd)
		}
		flagCmds = append(flagCmds, cmd)
	}
	return
}

func GlobalFlags(rawCmds ...*cobra.Command) (cmds []*cobra.Command) {
	for _, cmd := range rawCmds {
		cmd.Flags().StringP("path", "p", "~/.interstellar", "path to interstellar config dir")
		cmds = append(cmds, cmd)
	}
	return
}

func ProcessGlobalFlags(cmd *cobra.Command) (config types.InterstellarConfig, err error) {
	config.Path, err = cmd.Flags().GetString("path")
	if err != nil {
		return
	}
	usr, _ := user.Current()
	if err != nil {
		return
	}
	home := usr.HomeDir
	if config.Path == "~" {

		config.Path = home
	} else if strings.HasPrefix(config.Path, "~/") {
		config.Path = filepath.Join(home, config.Path[2:])
	}
	return
}
