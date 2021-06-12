package command

import "github.com/spf13/cobra"

func NewDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "delete TASK_ID",
		Aliases: []string{"d"},
		Short:   "Delete a task",
		Args:    cobra.ExactArgs(1),
		Example: `taskboard delete 3`,
		Run:     deleteCommand,
	}
}

func deleteCommand(cmd *cobra.Command, args []string) {
}
