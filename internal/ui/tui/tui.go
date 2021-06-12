package tui

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/thealamu/taskboard/internal/entity"
	"github.com/thealamu/taskboard/internal/ui"
)

// Terminal UI Renderer
type TUI struct {
	out io.Writer
}

func New(out io.Writer) *TUI {
	return &TUI{out}
}

func (t *TUI) Render(data ui.Template) {
	tuiTemplate, err := template.New("taskboard").
		Funcs(template.FuncMap{
			"tStyleTask":     tStyleTask,
			"tStyleRemark":   tStyleRemark,
			"headerSynopsis": headerSynopsis,
			"underline":      underline,
			"footerSynopsis": footerSynopsis,
		}).
		Parse(templateString)
	if err != nil {
		fmt.Fprintln(t.out, err)
	}
	var buf bytes.Buffer
	tuiTemplate.Execute(&buf, data)
	res := strings.TrimSpace(buf.String())
	fmt.Fprintln(t.out, res)
}

// tStyleRemark styles a remark
func tStyleRemark(r ui.Remark) string {
	green := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	successStatus := green.Render("✔")
	if !r.IsSuccess {
		red := lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
		successStatus = red.Render("✖")
	}

	textStyleFaded := lipgloss.NewStyle().Faint(true)
	id := textStyleFaded.Render(strconv.Itoa(r.Task.ID))

	return fmt.Sprintf("%s  %s: %s", successStatus, r.Msg, id)
}

// tStyleTask styles a task
func tStyleTask(t entity.Task) string {
	textStyleFaded := lipgloss.NewStyle().Faint(true)
	id := textStyleFaded.Render(strconv.Itoa(t.ID) + ".")

	purple := lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	checkedStatus := purple.Render("☐")
	if t.Checked {
		green := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
		checkedStatus = green.Render("✔")
	}

	description := t.Description
	if t.Checked {
		description = textStyleFaded.Render(description)
	}

	daysAgo := ""
	dayDiff := int(time.Since(t.Date).Hours() / 24)
	if dayDiff > 0 {
		daysAgo = strconv.Itoa(dayDiff) + "d"
		daysAgo = textStyleFaded.Render(daysAgo)
	}
	return fmt.Sprintf("%s %s  %s %s", id, checkedStatus, description, daysAgo)
}

func headerSynopsis(tasks []entity.Task) string {
	count := 0
	for _, task := range tasks {
		if task.Checked {
			count++
		}
	}
	synopsis := fmt.Sprintf("[%d/%d]", count, len(tasks))
	textStyleFaded := lipgloss.NewStyle().Faint(true)
	return textStyleFaded.Render(synopsis)
}

func footerSynopsis(tasks []entity.Task) string {
	count := 0
	for _, task := range tasks {
		if task.Checked {
			count++
		}
	}
	pctChecked := (float64(count) / float64(len(tasks))) * 100
	synopsis := fmt.Sprintf("%d%% of all tasks complete.", int(pctChecked))
	textStyleFaded := lipgloss.NewStyle().Faint(true)
	return textStyleFaded.Render(synopsis)
}

func underline(s string) string {
	underline := lipgloss.NewStyle().Underline(true)
	return underline.Render(s)
}

var templateString = `

{{ if .HasHeader }}
{{underline "My Board"}}  {{headerSynopsis .TaskList}}
{{ end }}
{{ range .TaskList }} {{tStyleTask .}}
{{ end }}
{{ with .Remark }} {{tStyleRemark . }} {{ end }}
{{ if .HasFooter }}
{{ footerSynopsis .TaskList}}
{{ end }}
`
