all: run

build:
	go build -o cam-example example/main.go 

run:
	go run -race example/main.go
