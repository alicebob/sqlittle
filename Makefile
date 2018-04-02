.PHONY: all test bench format

all: format test

test:
	go test

bench:
	go test -bench .

format:
	go fmt
