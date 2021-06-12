package command

import (
	"github.com/spf13/cobra"
)

func NewAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "add DESCRIPTION",
		Aliases: []string{"a"},
		Short:   "Add a new task",
		Args:    cobra.ExactArgs(1),
		Example: `taskboard add "Change the world"`,
		Run:     addCommand,
	}
}

func addCommand(cmd *cobra.Command, args []string) {
}
