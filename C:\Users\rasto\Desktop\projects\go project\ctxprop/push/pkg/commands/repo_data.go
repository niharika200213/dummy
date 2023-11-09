package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/shurcooL/githubv4"
	"github.com/spf13/cobra"
)

var repoDataCmd = &cobra.Command{
	Use:   "repo-data",
	Short: "Get data about a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		owner, name := parseOwnerAndName(args[0])
		client := api.GenerateClient(token)
		repo, err := api.RepoData(client, owner, name)
		if err != nil {
			log.Fatal(err)
		}
		printRepoData(repo)
	},
}

func printRepoData(repo api.Repo) {
	fmt.Println("Repo Data:")
	fmt.Println("Owner:", repo.Owner)
	fmt.Println("Name:", repo.Name)
	fmt.Println("Main Language:", repo.MainLanguage)
	fmt.Println("Private:", repo.Private)
	fmt.Println("Archived:", repo.Archived)
	fmt.Println("Template:", repo.Template)
	fmt.Println("Disabled:", repo.Disabled)
	fmt.Println("Mirror:", repo.Mirror)
	fmt.Println("Fork:", repo.Fork)
}

func init() {
	rootCmd.AddCommand(repoDataCmd)
}
You should propagate the context wherever required.repository",
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		owner, name := parseOwnerAndName(args[0])
		config := configuration.GetConfig()
		client := api.GenerateClient(config.GithubToken)
		repo, err := api.RepoData(client, owner, name)
		if err != nil {
			log.Fatal(err)
		}
		repoURL := generateRepoURL(repo)
		clone.CloneRepo(repoURL, config.ClonePath)
	},
}

func generateRepoURL(repo api.Repo) string {
	var sb strings.Builder
	sb.WriteString("https://github.com/")
	sb.WriteString(repo.Owner)
	sb.WriteString("/")
	sb.WriteString(repo.Name)
	return sb.String()
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	addClonePathFlag(cloneCmd)
}
You should propagate the context wherever required._, localRepo := range clonedRepos {
		pair := fmt.Sprintf("%s/%s", localRepo.Owner, localRepo.Name)
		repoPairs = append(repoPairs, pair)
	}

	return repoPairs, cobra.ShellCompDirectiveNoFileComp
}

// Parse the owner and name from the clone arg
func parseOwnerAndName(arg string) (string, string) {
	parts := strings.Split(arg, "/")
	owner := parts[0]
	name := parts[1]
	return owner, name
}

var rootCmd = &cobra.Command{
	Use:   "fgh",
	Short: "A CLI to interact with multiple GitHub repositories",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err !}

func init() {
	rootCmd.AddCommand(completionCmd)
}
