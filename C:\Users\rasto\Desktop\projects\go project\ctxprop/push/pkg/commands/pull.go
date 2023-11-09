package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "pull <PULL-NUMBER>",
	Short:                 "Check out the branch from a pull request",
	Long:                  longDocStart,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := configuration.GetConfig()
		pullNumber := repos.ParsePullNumber(args[0])

		pullRequest := repos.PullRequests(config.Owner, config.Repo, []int{pullNumber})
		if len(pullRequest) != 0 {
			repoURL := fmt.Sprintf("https://github.com/%s/%s", config.Owner, config.Repo)
			branchRef := strings.ReplaceAll(pullRequest[0].HeadRef, config.Owner, repoURL)

			utils.CheckoutBranch(branchRef)
		} else {
			fmt.Printf("Cannot find pull request #%v\n", pullNumber)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
