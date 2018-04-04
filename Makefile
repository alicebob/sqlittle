.PHONY: all test bench format fuzz

all: format test

test:
	go test

bench:
	go test -bench .

format:
	go fmt

fuzz:
	go get -v github.com/dvyukov/go-fuzz/...

	rm sqlittle-fuzz.zip
	go-fuzz-build github.com/alicebob/sqlittle
	mkdir -p workdir
	cp -r corpus workdir
	go-fuzz -bin=sqlittle-fuzz.zip -workdir=workdir
