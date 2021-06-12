package ui

import "github.com/thealamu/taskboard/internal/entity"

type Remark struct {
	IsSuccess bool
	Msg       string
	Task      entity.Task
}

type Template struct {
	HasHeader bool
	Remark    Remark
	TaskList  []entity.Task
	HasFooter bool
}
