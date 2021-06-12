package main

import (
	"log"

	"github.com/thealamu/taskboard/internal/command"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	return command.Execute()
}
