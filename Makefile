.PHONY: all example test

all: example

example:
	go run -race examples/middlewares/main.go

test:
	go test . ./middlewares/...
