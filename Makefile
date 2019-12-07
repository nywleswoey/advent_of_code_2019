.PHONY: build test security_scan lint

build:
	go build -v -o bin/day1 ./src/day1
	go build -v -o bin/day2 ./src/day2
	go build -v -o bin/day3 ./src/day3
