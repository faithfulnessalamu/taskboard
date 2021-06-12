EXE_DIR := cmd/taskboard
BINARY_NAME := taskboard

.PHONY: build
build:
	go build -o $(BINARY_NAME) $(EXE_DIR)/main.go
