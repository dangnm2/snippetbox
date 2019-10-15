PROJECT=$(shell basename "$(PWD)")
BUILD_TARGET="./cmd/web"

setup:
	@mkdir -p bin

build: setup
	@go build -o bin/$(PROJECT) $(BUILD_TARGET)

test:
	@go test -short ./...

clean:
	@rm -rf bin
