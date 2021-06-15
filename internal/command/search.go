package command

import (
	"github.com/spf13/cobra"
	"github.com/thealamu/taskboard/internal/data"
	"github.com/thealamu/taskboard/internal/entity"
	"github.com/thealamu/taskboard/internal/ui"
)

func NewSearchCommand(view *ui.UI, db data.Store) *cobra.Command {
	return &cobra.Command{
		Use:     "search STRING",
		Aliases: []string{"s"},
		Short:   "Search in tasks",
		Long:    "Display tasks whose descriptions contain the search string",
		Example: `taskboard search "Change the world"`,
		Run:     searchCommand(view, db),
	}
}

func searchCommand(view *ui.UI, db data.Store) commandFunc {
	fail := func(err error) {
		view.SetRemark(false, err.Error(), entity.Task{}).Render()
	}

	success := func(tasks []entity.Task) {
		view.ShowHeader(true).SetListResult(tasks).ShowFooter(true).Render()
	}

	return func(cmd *cobra.Command, args []string) {
		tasks, err := db.FindTasks(args[0], false)
		if err != nil {
			fail(err)
		}
		success(tasks)
	}
}
