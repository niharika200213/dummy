package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clone"
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/configuration"
	"github.com/Matt-Gleich/fgh/pkg/commands/issue"
	"github.com/Matt-Gleich/fgh/pkg/commands/list"
	"github.com/Matt-Gleich/fgh/pkg/commands/move"
	"github.com/Matt-Gleich/fgh/pkg/commands/pr"
	"github.com/Matt-Gleich/fgh/pkg/commands/pull"
	"github.com/Matt-Gleich/fgh/pkg/commands/remove"
	"github.com/Matt-Gleich/fgh/pkg/commands/repo"
	"github.com/Matt-Gleich/fgh/pkg/commands/view"
	"github.com/spf13/cobra"
)

func Factory() []*cobra.Command {
	return []*cobra.Command{
		cloneCmd,
		listCmd,
		moveCmd,
		prCmd,
		pullCmd,
		issueCmd,
		removeCmd,
		cleanCmd,
		configurationCmd,
		repoCmd,
		viewCmd,
	}
}
