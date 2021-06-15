package command

import (
	"github.com/spf13/cobra"
	"github.com/thealamu/taskboard/internal/data"
	"github.com/thealamu/taskboard/internal/ui"
)

var version = "debug"

type commandFunc = func(cmd *cobra.Command, args []string)

var rootCmd = &cobra.Command{
	Use:     "taskboard [command]",
	Short:   "📑 Tasks for the command line",
	Version: version,
	Run:     defaultCommand,
}

func getSubCommands(view *ui.UI, db data.Store) []*cobra.Command {
	return []*cobra.Command{
		NewAddCommand(view, db),
		NewTimelineCommand(),
		NewDeleteCommand(),
		NewCheckCommand(),
		NewEditCommand(),
		NewSearchCommand(),
	}
}

func Execute(view *ui.UI, db data.Store) error {
	commands := getSubCommands(view, db)
	rootCmd.AddCommand(commands...)
	return rootCmd.Execute()
}
