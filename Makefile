WEBSITE_REPO=github.com/gefion-tech/gefion.gg
VERSION?="1.0.0"

.PHONY: 
	build

build:
	go build -v ./cmd/gg
	

.DEFAULT_GOAL := build

