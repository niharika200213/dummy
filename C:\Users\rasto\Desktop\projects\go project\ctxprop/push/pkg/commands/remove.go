package commands

import (
	"os"

	"github.com/Matt-Gleich/fgh/pkg/commands/remove"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "remove <OWNER/REPO>",
	Short:                 "Remove a cloned repo",
	Long:                  longDocStart,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := configuration.GetConfig()
		remove.Remove(args[0], config)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
