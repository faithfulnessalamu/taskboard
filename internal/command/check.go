package command

import "github.com/spf13/cobra"

func NewCheckCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "check TASK_ID",
		Aliases: []string{"c"},
		Short:   "Check/Uncheck a task",
		Args:    cobra.ExactArgs(1),
		Example: `taskboard c 3`,
		Run:     checkCommand,
	}
}

func checkCommand(cmd *cobra.Command, args []string) {
}
