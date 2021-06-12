package command

import "github.com/spf13/cobra"

func NewSearchCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "search STRING",
		Aliases: []string{"s"},
		Short:   "Search in tasks",
		Long:    "Display tasks whose descriptions contain the search string",
		Example: `taskboard search "Change the world"`,
		Run:     searchCommand,
	}
}

func searchCommand(cmd *cobra.Command, args []string) {
}
