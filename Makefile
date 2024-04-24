# Simple Makefile for building Go projects

BINARY_NAME=srctotxt

build:
	    mkdir -p build
	    go build -o build/$(BINARY_NAME) main.go

run: build
	    ./build/$(BINARY_NAME)

clean:
	    go clean
	    rm -rf ./build

.PHONY: build run clean

