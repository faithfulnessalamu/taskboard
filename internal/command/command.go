package command

import "github.com/spf13/cobra"

var version = "debug"

var commands = []*cobra.Command{
	NewAddCommand(),
	NewTimelineCommand(),
	NewDeleteCommand(),
	NewCheckCommand(),
	NewEditCommand(),
	NewSearchCommand(),
}

var rootCmd = &cobra.Command{
	Use:     "taskboard [command]",
	Short:   "📑 Tasks for the command line",
	Version: version,
	Run:     rootCommand,
}

func init() {
	rootCmd.AddCommand(commands...)
}

func Execute() error {
	return rootCmd.Execute()
}
