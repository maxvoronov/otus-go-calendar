PROJECT?="github.com/maxvoronov/otus-go-calendar"
PROJECTNAME="go-calendar"
DOCKER=docker-compose -f deployments/docker-compose.yml

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

## start: Start application in docker containers with hot reload
start:
	@$(DOCKER) up -d --build

## stop: Stop application and remove docker containers
stop:
	@$(DOCKER) down --remove-orphans

## genproto: Generate gRPC interfaces by proto files
genproto:
	@protoc grpc/proto/*.proto --go_out=plugins=grpc:.

## migrate: Apply all migrations
migrate:
	@$(DOCKER) run --rm migrations up

## migrate-down: Rollback migrations
migrate-down:
	@$(DOCKER) run --rm migrations down

## migrate-status: Rollback migrations
migrate-status:
	@$(DOCKER) run --rm migrations status

## test: Start tests
test:
	@go test ./...

## lint: Check source code by linters
lint:
	@echo "Checking go vet..." && go vet ./... && echo "Done!\n"
	@echo "Checking golint..." && golint ./... && echo "Done!\n"
	@echo "Checking golangci-lint..." && golangci-lint run ./... && echo "Done!"
