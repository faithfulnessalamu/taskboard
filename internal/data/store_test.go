package data

import (
	"testing"
	"time"

	"github.com/thealamu/taskboard/internal/entity"
)

// Test a Store implementation
func testStoreImplementation(t *testing.T, s Store) {
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
	err = s.DeleteTask(1)
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
	err = s.ToggleTask(3)
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
	t0 := testTasks[0]
	t0.Checked = false
	err = s.UpdateTask(t0)
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
	{ID: 0, Checked: true, Description: "Todo 1", Date: time.Now()},
	{ID: 1, Checked: false, Description: "Todo 2", Date: time.Now()},
	{ID: 2, Checked: true, Description: "Todo 3", Date: time.Now()},
	{ID: 3, Checked: false, Description: "Todo 4", Date: time.Now()},
}
