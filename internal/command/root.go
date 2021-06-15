package command

import (
	"github.com/spf13/cobra"
	"github.com/thealamu/taskboard/internal/data"
	"github.com/thealamu/taskboard/internal/entity"
	"github.com/thealamu/taskboard/internal/ui"
)

func defaultCommand(view *ui.UI, db data.Store) commandFunc {
	fail := func(err error) {
		view.SetRemark(false, err.Error(), entity.Task{})
	}

	success := func(tasks []entity.Task) {
		view.ShowHeader(true).SetListResult(tasks).ShowFooter(true).Render()
	}

	return func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetTasks(false)
		if err != nil {
			fail(err)
			return
		}
		success(tasks)
	}
}
