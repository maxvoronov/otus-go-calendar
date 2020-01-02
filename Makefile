PROJECT?="github.com/maxvoronov/otus-go-calendar"
PROJECTNAME="go-calendar"

COMMIT := $(shell git rev-parse --short HEAD)
VERSION := $(shell git describe --tags --abbrev=0)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

help: Makefile
	@echo "Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## build: Build application
build:
	@GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build \
		-ldflags="-w -s -X ${PROJECT}/internal/version.Commit=${COMMIT} \
		-X ${PROJECT}/internal/version.Version=${VERSION} \
		-X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
		-o ./bin/$(PROJECTNAME)

## test: Start tests
test:
	@go test ./...

## lint: Check source code by linters
lint:
	@echo "Start go vet..." && go vet ./... && echo "Done!\n"
	@echo "Start golint..." && golint ./... && echo "Done!"