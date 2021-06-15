package command

import (
	"github.com/spf13/cobra"
	"github.com/thealamu/taskboard/internal/data"
	"github.com/thealamu/taskboard/internal/ui"
)

func NewAddCommand(view *ui.UI, db data.Store) *cobra.Command {
	return &cobra.Command{
		Use:     "add DESCRIPTION",
		Aliases: []string{"a"},
		Short:   "Add a new task",
		Args:    cobra.ExactArgs(1),
		Example: `taskboard add "Change the world"`,
		Run:     addCommand(view, db),
	}
}

func addCommand(view *ui.UI, db data.Store) commandFunc {
	return func(cmd *cobra.Command, args []string) {
	}
}
