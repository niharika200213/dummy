package commands

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Generate a list of your cloned repos",
	Long:  longDocStart,
	Run: func(cmd *cobra.Command, args []string) {
		config := configuration.GetConfig()
		clonedRepos := reposBasedOffCustomPath(cmd, config)

		for i := range clonedRepos {
			fmt.Println(clonedRepos[i].Path)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
