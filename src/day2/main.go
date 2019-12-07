package main

import (
	"fmt"
)

const (
	add  int = 1
	mult int = 2
	end  int = 99
)

func executeAdd(prog []int, param1, param2, param3 int) int {
	prog[param3] = prog[param1] + prog[param2]
	return 4
}

func executeMult(prog []int, param1, param2, param3 int) int {
	prog[param3] = prog[param1] * prog[param2]
	return 4
}

func process(prog []int) {
	var instrPtr int

Done:
	for instrPtr < len(prog) {
		opCode := prog[instrPtr]

		switch opCode {
		case add:
			instrPtr += executeAdd(prog, prog[instrPtr+1], prog[instrPtr+2], prog[instrPtr+3])
		case mult:
			instrPtr += executeMult(prog, prog[instrPtr+1], prog[instrPtr+2], prog[instrPtr+3])
		case end:
			instrPtr++
			break Done
		default:
			fmt.Printf("encountered an invalid op code: %v at address %v\n", prog[instrPtr], instrPtr)
			break Done
		}
	}
}

func main() {
	// test input
	testData := [][]int{
		[]int{1, 0, 0, 0, 99},
		[]int{2, 3, 0, 3, 99},
		[]int{2, 4, 4, 5, 99, 0},
		[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
	}

	for _, v := range testData {
		fmt.Printf("Received value: %v, ", v)
		process(v)
		fmt.Printf("Processed value: %v\n\n", v)
	}

	puzzleInput := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 9, 1, 19, 1, 19, 5, 23, 1, 9, 23, 27, 2, 27, 6, 31, 1, 5, 31, 35, 2, 9, 35, 39, 2, 6, 39, 43, 2, 43, 13, 47, 2, 13, 47, 51, 1, 10, 51, 55, 1, 9, 55, 59, 1, 6, 59, 63, 2, 63, 9, 67, 1, 67, 6, 71, 1, 71, 13, 75, 1, 6, 75, 79, 1, 9, 79, 83, 2, 9, 83, 87, 1, 87, 6, 91, 1, 91, 13, 95, 2, 6, 95, 99, 1, 10, 99, 103, 2, 103, 9, 107, 1, 6, 107, 111, 1, 10, 111, 115, 2, 6, 115, 119, 1, 5, 119, 123, 1, 123, 13, 127, 1, 127, 5, 131, 1, 6, 131, 135, 2, 135, 13, 139, 1, 139, 2, 143, 1, 143, 10, 0, 99, 2, 0, 14, 0}

	// puzzle 1 answer
	input := append(puzzleInput[:0:0], puzzleInput...)
	input[1] = 12
	input[2] = 2
	process(input)
	fmt.Printf("Puzzle 1 answer: %v\n\n", input[0])

	// puzzle 2 answer
Done:
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			input = append(puzzleInput[:0:0], puzzleInput...)
			input[1] = noun
			input[2] = verb
			process(input)
			if input[0] == 19690720 {
				break Done
			}
		}
	}
	fmt.Printf("Puzzle 2 answer: %v%v\n", input[1], input[2])
}
