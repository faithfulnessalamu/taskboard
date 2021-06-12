package command

import "github.com/spf13/cobra"

func NewEditCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "edit TASK_ID",
		Aliases: []string{"e"},
		Short:   "Edit a task",
		Args:    cobra.ExactArgs(1),
		Example: `taskboard e 3`,
		Run:     editCommand,
	}
}

func editCommand(cmd *cobra.Command, args []string) {
}
