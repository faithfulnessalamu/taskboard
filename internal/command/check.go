package command

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thealamu/taskboard/internal/data"
	"github.com/thealamu/taskboard/internal/entity"
	"github.com/thealamu/taskboard/internal/ui"
)

func NewCheckCommand(view *ui.UI, db data.Store) *cobra.Command {
	return &cobra.Command{
		Use:     "check TASK_ID",
		Aliases: []string{"c"},
		Short:   "Check/Uncheck a task",
		Args:    cobra.ExactArgs(1),
		Example: `taskboard c 3`,
		Run:     checkCommand(view, db),
	}
}

func checkCommand(view *ui.UI, db data.Store) commandFunc {
	fail := func(err error) {
		view.SetRemark(false, err.Error(), entity.Task{}).Render()
	}

	success := func(msg string, id int) {
		view.SetRemark(true, msg, entity.Task{ID: id}).Render()
	}

	return func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fail(err)
			return
		}
		db.ToggleTask(taskID)
		success("Checked task", taskID)
	}
}
