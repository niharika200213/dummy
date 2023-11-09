package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/status"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var (
	printNoActivity bool
	statusCmd       = &cobra.Command{
		DisableFlagsInUseLine: true,
		Args:                  cobra.ExactArgs(0),
		Use:                   "status",
		Short:                 "Get the status of cloned repositories",
		Long:                  longDocStart + "https://github.com/Matt-Gleich/#-fgh-status",
