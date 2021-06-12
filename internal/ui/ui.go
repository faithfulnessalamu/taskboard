package ui

import (
	"fmt"

	"github.com/thealamu/taskboard/internal/entity"
)

type Renderer interface {
	Render(Template)
}

// UI builder
type UI struct {
	tmpl *Template
	r    Renderer
}

func New(r Renderer) (*UI, error) {
	if r == nil {
		return nil, fmt.Errorf("got nil renderer")
	}
	return &UI{new(Template), r}, nil
}

func (u *UI) Clear() *UI {
	u.tmpl = new(Template)
	return u
}

func (u *UI) ShowHeader(b bool) *UI {
	u.tmpl.HasHeader = b
	return u
}

func (u *UI) SetRemark(isSuccess bool, msg string, task entity.Task) *UI {
	u.tmpl.Remark = Remark{isSuccess, msg, task}
	return u
}

func (u *UI) SetListResult(tasks []entity.Task) *UI {
	u.tmpl.TaskList = tasks
	return u
}

func (u *UI) ShowFooter(b bool) *UI {
	u.tmpl.HasFooter = b
	return u
}

func (u *UI) Render() {
	u.r.Render(*u.tmpl)
}
