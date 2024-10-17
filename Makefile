.PHONY: build run

build:
	go build -o ./tmp ./...

run:
	go run ./...
