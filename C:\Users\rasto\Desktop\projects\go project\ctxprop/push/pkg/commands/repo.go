package commands

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var printRepos bool

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.Flags().BoolVarP(&printRepos, "print", "p", false, "Print repos to json")
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Get info about the current repos",
	Long:  "Get info about the current repos",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := utils.GetToken()
		if err != nil {
			statuser.Error(err, 1)
		}

		client := api.GenerateClient(id)

		var repos []*api.Repo
		for _, repo := range repos.CurrentRepos() {
			repoInfo, err := api.RepoData(client, repo.Owner, repo.Name)
			if err != nil {
				statuser.Error(err, 1)
			}
			repos = append(repos, &repoInfo)
			if printRepos {
				statuser.Success(fmt.Sprintf("<<%s>>", repoInfo.Name))
			}
		}
		if printRepos {
			return
		}
		utils.PrintJson(repos)
	},
}
