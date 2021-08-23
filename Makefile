.PHONY: 
	run

build:
	go build -o gg -v ./src/main.go 

run:
	go run -v ./src/main.go

.DEFAULT_GOAL := run

