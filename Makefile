WEBSITE_REPO=github.com/gefion-tech/gefion.gg
VERSION?="1.0.0"

.PHONY: 
	build

build:
	go build -v ./cmd/gg
	rm -rf /home/I0HuKc/NotJob/test
	

.DEFAULT_GOAL := build

