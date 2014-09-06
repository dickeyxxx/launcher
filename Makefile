.PHONY: build run

default: build

run: build
	./launcher

build:
	go build
