BINARY_NAME := taskboard

.PHONY: build
build:
	go build -o $(BINARY_NAME) ./...
