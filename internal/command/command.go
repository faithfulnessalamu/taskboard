package command

import (
	"github.com/spf13/cobra"
	"github.com/thealamu/taskboard/internal/data"
	"github.com/thealamu/taskboard/internal/ui"
)

var version = "debug"

type commandFunc = func(cmd *cobra.Command, args []string)

func rootCmd(view *ui.UI, db data.Store) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "taskboard [command]",
		Short:   "📑 Tasks for the command line",
		Version: version,
		Run:     defaultCommand(view, db),
	}

	subCommands := []*cobra.Command{
		NewAddCommand(view, db),
		NewCheckCommand(view, db),
		NewDeleteCommand(view, db),
		NewEditCommand(),
		NewSearchCommand(),
		NewTimelineCommand(),
	}

	rootCmd.AddCommand(subCommands...)
	return rootCmd
}

func Execute(view *ui.UI, db data.Store) error {
	return rootCmd(view, db).Execute()
}
