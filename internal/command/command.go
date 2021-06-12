package command

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "taskboard [command]",
	Short: "📑 Tasks for the command line",
}

func Execute() error {
	return rootCmd.Execute()
}
