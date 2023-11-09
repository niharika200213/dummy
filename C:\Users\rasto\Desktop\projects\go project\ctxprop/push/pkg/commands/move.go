package commands

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/commands/move"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "move <OWNER/REPO> <PATH>",
	Short:                 "Move a repo to a new location",
	Long:                  longDocStart,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		config := configuration.GetConfig()
		move.Move(args[0], args[1], config)
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}
