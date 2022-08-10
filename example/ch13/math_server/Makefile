
.DEFAULT_GOAL := build

fmt:
	@go fmt ./...
.PHONY:fmt

lint: fmt
	@golint ./...
.PHONY:lint

vet: fmt
	@go vet ./...
.PHONY:vet

test: vet
	@go test -v ./...
.PHONY:test

test-cover: vet
	@go test -v -cover ./...
.PHONY:test-cover

test-cover-html: vet
	@go test -v -cover -coverprofile=c.out ./...
	@go tool cover -html=c.out
.PHONY:test-cover-html

build: vet
	@go build -ldflags="-w -s"
.PHONY:build

run: build
	@./$(shell basename $(CURDIR))
.PHONY:run

docker-build:
	@docker build -t math_server:latest .
.PHONY:docker-build

docker-run:
	@docker run -p 8080:8080 math_server:latest
.PHONY:docker-run