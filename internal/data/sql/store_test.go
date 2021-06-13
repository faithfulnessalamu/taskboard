package sql

import (
	"testing"
	"time"

	"github.com/thealamu/taskboard/internal/data"
	"github.com/thealamu/taskboard/internal/entity"
)

func TestSQLStore(t *testing.T) {
	// use test dsn
	_DSN = "file:taskboard-test.db?cache=shared&mode=memory"

	store, err := NewStore()
	if err != nil {
		t.Errorf("could not create new store, got %v", err)
	}
	DoStoreImplementationTest(t, store)
}

// Test a Store implementation
func DoStoreImplementationTest(t *testing.T, s data.Store) {
	// Add Tasks
	for _, testTask := range testTasks {
		err := s.AddTask(testTask)
		if err != nil {
			t.Errorf("could not add new task, got error '%v'", err)
		}
	}

	// Get Tasks
	// GetTasks(false) because we only want uncompleted tasks i.e where 'checked' is false
	gotTestTasks, err := s.GetTasks(false)
	if err != nil {
		t.Errorf("could not retrieve tasks, got error '%v'", err)
	}
	if len(gotTestTasks) != 2 {
		t.Errorf("expected %d uncompleted tasks, got %d", 2, len(gotTestTasks))
	}

	// Delete a Task
	err = s.DeleteTask(2)
	if err != nil {
		t.Errorf("could not delete task, got error '%v'", err)
	}
	// validate deletion
	// GetTasks(false) because we only want uncompleted tasks i.e where 'checked' is false
	gotTestTasks, err = s.GetTasks(false)
	if err != nil {
		t.Errorf("could not retrieve tasks, got error '%v'", err)
	}
	if len(gotTestTasks) != 1 {
		t.Errorf("expected %d uncompleted tasks, got %d", 1, len(gotTestTasks))
	}

	// Toggle Task
	err = s.ToggleTask(4)
	if err != nil {
		t.Errorf("could not toggle task, got error '%v'", err)
	}

	// validate toggle
	// GetTasks(false) because we only want uncompleted tasks i.e where 'checked' is false
	gotTestTasks, err = s.GetTasks(false)
	if err != nil {
		t.Errorf("could not retrieve tasks, got error '%v'", err)
	}
	if len(gotTestTasks) != 0 {
		t.Errorf("expected %d uncompleted tasks, got %d", 0, len(gotTestTasks))
	}

	// Update Task
	task1 := entity.Task{ID: 1, Checked: false, Description: "Todo 1", Date: time.Now()}
	err = s.UpdateTask(task1)
	if err != nil {
		t.Errorf("could not toggle task, got error '%v'", err)
	}
	// validate update
	// GetTasks(false) because we only want uncompleted tasks i.e where 'checked' is false
	gotTestTasks, err = s.GetTasks(false)
	if err != nil {
		t.Errorf("could not retrieve tasks, got error '%v'", err)
	}
	if len(gotTestTasks) != 1 {
		t.Errorf("expected %d uncompleted tasks, got %d", 1, len(gotTestTasks))
	}

	// Find Tasks
	// FindTasks('filter', true) because we want to include checked tasks
	gotTestTasks, err = s.FindTasks("Todo", true)
	if err != nil {
		t.Errorf("could not toggle task, got error '%v'", err)
	}
	if len(gotTestTasks) != 3 {
		t.Errorf("expected %d tasks, got %d", 3, len(gotTestTasks))
	}
}

var testTasks = []entity.Task{
	{Checked: true, Description: "Todo 1", Date: time.Now()},
	{Checked: false, Description: "Todo 2", Date: time.Now()},
	{Checked: true, Description: "Todo 3", Date: time.Now()},
	{Checked: false, Description: "Todo 4", Date: time.Now()},
}
