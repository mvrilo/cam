.PHONY: all example test

all: example

example:
	go run -race examples/classifier/main.go examples/classifier/haarcascade_eye.xml

test:
	go test . ./middlewares/...
