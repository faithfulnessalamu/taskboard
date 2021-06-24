package sql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/thealamu/taskboard/internal/entity"
)

type SQLStore struct {
	db *sql.DB
}

func NewStore() (*SQLStore, error) {
	// Build DSN
	dbPath, err := getDBPath()
	if err != nil {
		return nil, err
	}

	DSN := fmt.Sprintf(`file:%s/taskboard.db`, dbPath)
	return newStore(DSN)
}

func newStore(DSN string) (*SQLStore, error) {
	db, err := newSqliteDB(DSN)
	if err != nil {
		return nil, err
	}
	createStoreTable(db)
	return &SQLStore{db}, nil
}

func (s *SQLStore) AddTask(t entity.Task) error {
	qry := `INSERT INTO tasks(checked, description, date) VALUES (?, ?, ?)`
	_, err := s.db.Exec(qry, t.Checked, t.Description, t.Date.Unix())
	return err
}

// GetTasks returns tasks. If all is true, completed tasks are returned too.
func (s *SQLStore) GetTasks(all bool) ([]entity.Task, error) {
	qry := `SELECT * FROM tasks`
	if !all {
		qry += ` WHERE checked=false`
	}
	rows, err := s.db.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readTasksFromRows(rows)
}

func (s *SQLStore) GetLastID() (int, error) {
	qry := `SELECT * FROM tasks ORDER BY date DESC LIMIT 1`
	rows, err := s.db.Query(qry)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	t, err := readTasksFromRows(rows)
	if err != nil {
		return -1, err
	}

	if len(t) < 1 {
		return 0, nil
	}

	return t[0].ID, nil
}

func (s *SQLStore) DeleteTask(id int) error {
	qry := `DELETE FROM tasks WHERE id=?`
	_, err := s.db.Exec(qry, id)
	return err
}

// ToggleTask checks a task if it is unchecked, unchecks a task if it is checked.
func (s *SQLStore) ToggleTask(id int) error {
	qry := `UPDATE tasks SET checked= ((checked | 1) - (checked & 1)) WHERE id=?`
	_, err := s.db.Exec(qry, id)
	return err
}

func (s *SQLStore) UpdateTask(t entity.Task) error {
	qry := `UPDATE tasks SET checked=?, description=?, date=? WHERE id=?`
	_, err := s.db.Exec(qry, t.Checked, t.Description, t.Date.Unix(), t.ID)
	return err
}

// FindTasks returns tasks with the filter in their description.
// If all is true, completed tasks are searched too.
func (s *SQLStore) FindTasks(filter string, all bool) ([]entity.Task, error) {
	qry := `SELECT * FROM tasks WHERE description LIKE '%' || ? || '%'`
	if !all {
		qry += ` AND checked=false`
	}
	rows, err := s.db.Query(qry, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readTasksFromRows(rows)
}

func readTasksFromRows(rows *sql.Rows) ([]entity.Task, error) {
	var tasks []entity.Task

	// loop through result rows, inserting a task into tasks each iteration
	for rows.Next() {
		t := entity.Task{}
		var timeInt int64

		err := rows.Scan(&t.ID, &t.Checked, &t.Description, &timeInt)
		if err != nil {
			return nil, err
		}

		t.Date = time.Unix(timeInt, 0)
		tasks = append(tasks, t)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func newSqliteDB(DSN string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DSN)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "could not connect to sqlite database")
	}
	return db, nil
}

func createStoreTable(db *sql.DB) error {
	schema := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY,
			checked INTEGER DEFAULT 0,
			description TEXT NOT NULL,
			date INTEGER 
		)
	`
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
