package commands

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Open the repo in the browser",
	Long:  longDocStart,
	Run: func(cmd *cobra.Command, args []string) {
		config := configuration.GetConfig()
		utils.OpenURL(fmt.Sprintf("https://github.com/%v/%v", config.Owner, config.Repo))
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
