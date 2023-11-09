package repos

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

type Repo struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

// Get all the current repos in the fgh repos file
func CurrentRepos() []*Repo {
	var repos []*Repo
	jsonFile, err := os.Open(utils.REPOS_FILE)
	if err != nil {
		statuser.Error(
			statuser.NewError("Failed to open", err),
			1,
		)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &repos)
	return repos
}GitHub repository",
	Long: `
Clone a repository from GitHub.

The OWNER is the username or organization name that owns the repository.
The REPO is the name of the repository.

The repository will be cloned into the current directory.

For example:

fgh clone Matt-Gleich/fgh
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := args[0]
		config := configuration.GetConfig()
		token, err := configuration.GetToken()
		if err != nil {
			statuser.Error(err, 1)
		}

		clonedRepos := repos.CloneRepos{
			Repos: []repos.Repo{
				{
					Repo: repo,
				},
			},
		}

		clone.CloneRepos(clonedRepos, token, config)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}}

func init() {
	rootCmd.AddCommand(completionCmd)
}b.com/Matt-Gleich/fgh/pkg/commands/remove"
"github.com/Matt-Gleich/fgh/pkg/configuration"
"github.com/Matt-Gleich/fgh/pkg/repos"
"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Use:                   "remove <OWNER/REPO>",
	Short:                 "Remove a cloned GitHub repository",
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-remove",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			username = configuration.GetSecrets().Username
			config   = configuration.GetConfig()
			clonedRepos = reposBasedOffCustomPath(cmd, config)
			repo     = remove.GetRepo(args[0], username, clonedRepos)
		)
		remove.RemoveRepo(repo, config)
	},
	ValidArgsFunction: reposAsValidArgs,
}

func init() {
	rootCmd.AddCommand(removeCmd)
	addCustomPathFlag(removeCmd)
}
