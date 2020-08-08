.PHONY: all example test

all: example

example:
	go run -race examples/simple/main.go 

test:
	go test . ./middlewares/...
