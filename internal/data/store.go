package data

import "github.com/thealamu/taskboard/internal/entity"

type Store interface {
	AddTask(entity.Task) error

	// GetTasks returns tasks. If all is true, completed tasks are returned too.
	GetTasks(all bool) ([]entity.Task, error)

	DeleteTask(id int) error

	// ToggleTask checks a task if it is unchecked, unchecks a task if it is checked.
	ToggleTask(id int) error

	UpdateTask(entity.Task) error

	// FindTasks returns tasks with the filter in their description.
	// If all is true, completed tasks are searched too.
	FindTasks(filter string, all bool) ([]entity.Task, error)

	GetLastID() (int, error)
}
