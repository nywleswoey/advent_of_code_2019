.PHONY: build test security_scan lint

build:
	go build ./intcode
	go build -v -o bin/day1 ./src/day1
	go build -v -o bin/day2 ./src/day2
	go build -v -o bin/day3 ./src/day3
	go build -v -o bin/day4 ./src/day4
	go build -v -o bin/day5 ./src/day5
	go build -v -o bin/day6 ./src/day6
