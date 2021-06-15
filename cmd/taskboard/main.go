package main

import (
	"log"
	"os"

	"github.com/thealamu/taskboard/internal/command"
	"github.com/thealamu/taskboard/internal/data/sql"
	"github.com/thealamu/taskboard/internal/ui"
	"github.com/thealamu/taskboard/internal/ui/tui"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := sql.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	terminalRenderer := tui.New(os.Stdout)
	view, err := ui.New(terminalRenderer)
	if err != nil {
		log.Fatal(err)
	}

	return command.Execute(view, db)
}
