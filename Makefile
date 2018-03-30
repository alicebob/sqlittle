.PHONY: all test format

all: format test

test:
	go test

format:
	go fmt
