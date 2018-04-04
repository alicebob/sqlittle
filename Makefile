.PHONY: all test bench format fuzz

all: format test

test:
	go test

bench:
	go test -bench .

format:
	go fmt

sqlittle-fuzz.zip:
	go-fuzz-build github.com/alicebob/sqlittle

fuzz: sqlittle-fuzz.zip
	mkdir workdir
	cp -r corpus workdir
	go-fuzz -bin=sqlittle-fuzz.zip -workdir=workdir
