package commands

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/spf13/cobra"
)

var configurationCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Print the configuration used by fgh",
	Long:  longDocStart,
	Run: func(cmd *cobra.Command, args []string) {
		config := configuration.GetConfig()
		fmt.Fprint(os.Stdout, config.ToString())
	},
}

func init() {
	rootCmd.AddCommand(configurationCmd)
}
