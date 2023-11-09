package commands

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr <PR-NUMBER>",
	Short: "Open the pull request in the browser",
	Long:  longDocStart,
	Run: func(cmd *cobra.Command, args []string) {
		config := configuration.GetConfig()

		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}

		pulls := repos.PullRequests(config.Owner, config.Repo, []int{repos.ParsePrNumber(args[0])})
		if len(pulls) == 0 {
			utils.OpenURL(fmt.Sprintf("https://github.com/%v/%v/pulls/%v", config.Owner, config.Repo, args[0]))
		} else {
			utils.OpenURL(fmt.Sprintf("https://github.com/%v/%v/pull/%v", config.Owner, config.Repo, args[0]))
		}
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
}
