# Simple Makefile for Go project

BINARY_NAME=go-boilerplate.exe
MAIN_PATH=./cmd/main.go

.PHONY: all build run test clean

all: build

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

test:
	go test ./...

clean:
	if exist $(BINARY_NAME) del $(BINARY_NAME)
	go clean

tidy:
	go mod tidy
