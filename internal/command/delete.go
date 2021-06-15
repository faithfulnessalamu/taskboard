package command

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thealamu/taskboard/internal/data"
	"github.com/thealamu/taskboard/internal/entity"
	"github.com/thealamu/taskboard/internal/ui"
)

func NewDeleteCommand(view *ui.UI, db data.Store) *cobra.Command {
	return &cobra.Command{
		Use:     "delete TASK_ID",
		Aliases: []string{"d"},
		Short:   "Delete a task",
		Args:    cobra.ExactArgs(1),
		Example: `taskboard delete 3`,
		Run:     deleteCommand(view, db),
	}
}

func deleteCommand(view *ui.UI, db data.Store) commandFunc {
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
		err = db.DeleteTask(taskID)
		if err != nil {
			fail(err)
			return
		}
		success("Deleted item", taskID)
	}
}
