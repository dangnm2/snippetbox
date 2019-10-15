PROJECT=$(shell basename "$(PWD)")
BUILD_TARGET="./cmd/web"

all: setup build test clean

setup:
	@mkdir -p bin

build: setup
	go build -o bin/$(PROJECT) $(BUILD_TARGET)

test:
	go test -v -short ./...

clean:
	@rm -rf bin

.PHONY: all
