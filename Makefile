.PHONY: build test security_scan lint

build:
	go build -v -o bin/day1 ./src/day1

run:
	./bin/day1
