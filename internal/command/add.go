package command

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/thealamu/taskboard/internal/data"
	"github.com/thealamu/taskboard/internal/entity"
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
	fail := func(task entity.Task, msg string) {
		view.SetRemark(false, msg, task).Render()
	}

	success := func(task entity.Task, msg string) {
		view.SetRemark(true, msg, task).Render()
	}

	return func(cmd *cobra.Command, args []string) {
		task := entity.Task{Description: args[0], Date: time.Now()}
		err := db.AddTask(task)
		if err != nil {
			fail(task, err.Error())
			return
		}
		success(task, "Created task")
	}
}
