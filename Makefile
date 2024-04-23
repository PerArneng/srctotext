# Simple Makefile for building Go projects

BINARY_NAME=srctotxt

build:
	    mkdir -p build
	    go build -o build/$(BINARY_NAME) main.go

run: build
	    ./$(BINARY_NAME)

clean:
	    go clean
	        rm $(BINARY_NAME)

.PHONY: build run clean

