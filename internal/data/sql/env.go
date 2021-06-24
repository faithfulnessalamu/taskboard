package sql

import (
	"os"
	"path"
)

var (
	TASKBOARD_DB_PATH = "TASKBOARD_DB_PATH"

	// $HOME/.taskboard
	defaultDBPath = path.Join(os.Getenv("HOME"), ".taskboard")
)

func getDBPath() (string, error) {
	dbPath := os.Getenv(TASKBOARD_DB_PATH)
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	if err := os.MkdirAll(dbPath, 0755); err != nil {
		return "", err
	}
	return dbPath, nil
}
