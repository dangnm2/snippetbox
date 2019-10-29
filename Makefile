PROJECT=$(shell basename "$(PWD)")
BUILD_TARGET="./cmd/web"

all: setup build test clean

setup:
	@mkdir -p bin

build: setup
	go build -o bin/$(PROJECT) $(BUILD_TARGET)

test:
	CGO_ENABLED=0 go test -v -short ./...

integrate:
	CGO_ENABLED=0 go test -v ./...

clean:
	@rm -rf bin

.PHONY: all
