package commands

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue <ACTION> [NUMBER]",
	Short: "Open the issue in the browser",
	Long:  longDocStart,
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		config := configuration.GetConfig()

		if len(args) == 0 {
			utils.OpenURL(fmt.Sprintf("https://github.com/%v/%v/issues/new/choose", config.Owner, config.Repo))
		} else if len(args) == 1 {
			utils.OpenURL(fmt.Sprintf("https://github.com/%v/%v/issues/%v", config.Owner, config.Repo, args[0]))
		} else {
			cmd.Help()
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)
}ri := range clonedRepos {
		// Remove the root folder from the path
		pathComponents := strings.Split(clonedRepos[i].Path, "/")
		repoPairs = append(repoPairs, fmt.Sprintf("%v/%v", pathComponents[len(pathComponents)-2], pathComponents[len(pathComponents)-1]))
	}

	if len(args) != 0 {
		return cobra.ShellCompDirectiveNoFileComp, []string{}
	}

	return cobra.ShellCompDirectiveNoFileComp, repoPairs
}

// Set the valid args as the repo pairs (OWNER/REPO)
func repoPairsAsValidArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var (
		secrets     = configuration.GetSecrets()
		config      = configuration.GetConfig()
		clonedRepos = reposBasedOffCustomPath(cmd, config)
		repoPairs   = []string{}
	)

	for i := range clonedRepos {
	for i := range clonedRepos {
		// Remove the root folder from the path
		pathComponents := strings.Split(clonedRepos[i].Path, "/")
		repoPairs = append(repoPairs, fmt.Sprintf("%v/%v", pathComponents[len(pathComponents)-2], pathComponents[len(pathComponents)-1]))
	}

	if len(args) != 0 {
		return cobra.ShellCompDirectiveNoFileComp, []string{}
	}

	return cobra.ShellCompDirectiveNoFileComp, repoPairs
}

// Execute the commands
func Execute() {
	for _, cmd := range Factory() {
		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute command", err, 1)
	}
}
