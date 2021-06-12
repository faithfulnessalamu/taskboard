VERSION != git describe --tags --always
OS != uname -s
ARCH != uname -i
VERSION_VAR_PATH = github.com/thealamu/taskboard/internal/command

.PHONY: build
build:
	go build -o taskboard --ldflags="-X $(VERSION_VAR_PATH).version=$(VERSION)-$(OS)-$(ARCH)" cmd/taskboard/main.go
