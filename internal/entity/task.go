package entity

import "time"

type Task struct {
	ID          int
	Checked     bool
	Description string
	Date        time.Time
}
