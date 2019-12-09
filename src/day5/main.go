package main

import (
	"fmt"
)

const (
	add         int = 1
	mult        int = 2
	save        int = 3
	load        int = 4
	jumpOnTrue  int = 5
	jumpOnFalse int = 6
	lessThan    int = 7
	equals      int = 8
	end         int = 99
)

func cloneArray(a []int) []int {
	return append(a[:0:0], a...)
}

func extractMode(modes int) (int, int) {
	// returns the right most mode and the remaining unused modes
	return modes % 10, modes / 10
}

func getValue(prog []int, param, mode int) int {
	if mode == 1 {
		return param
	}
	return prog[param]
}

func executeAdd(prog []int, param1, param2, param3, modes int) int {
	mode, modes := extractMode(modes)
	value1 := getValue(prog, param1, mode)
	mode, modes = extractMode(modes)
	value2 := getValue(prog, param2, mode)
	mode, modes = extractMode(modes)
	if mode != 0 {
		// do some error handling
	}
	prog[param3] = value1 + value2
	return 4
}

func executeMult(prog []int, param1, param2, param3, modes int) int {
	mode, modes := extractMode(modes)
	value1 := getValue(prog, param1, mode)
	mode, modes = extractMode(modes)
	value2 := getValue(prog, param2, mode)
	mode, modes = extractMode(modes)
	if mode != 0 {
		// do some error handling
	}
	prog[param3] = value1 * value2
	return 4
}

func executeSave(prog []int, param1, input int) int {
	prog[param1] = input
	return 2
}

func executeLoad(prog []int, param1, modes int) int {
	mode, modes := extractMode(modes)
	value1 := getValue(prog, param1, mode)
	fmt.Printf("output: %v\n", value1)
	return 2
}

func executeJump(onTrue bool, prog []int, param1, param2, modes, currPtr int) int {
	mode, modes := extractMode(modes)
	value1 := getValue(prog, param1, mode)
	mode, modes = extractMode(modes)
	value2 := getValue(prog, param2, mode)
	if (onTrue && value1 != 0) || (!onTrue && value1 == 0) {
		return value2
	}
	return currPtr + 3
}

func executeLessThan(prog []int, param1, param2, param3, modes int) int {
	mode, modes := extractMode(modes)
	value1 := getValue(prog, param1, mode)
	mode, modes = extractMode(modes)
	value2 := getValue(prog, param2, mode)
	mode, modes = extractMode(modes)
	if mode != 0 {
		// do some error handling
	}

	if value1 < value2 {
		prog[param3] = 1
	} else {
		prog[param3] = 0
	}
	return 4
}

func executeEquals(prog []int, param1, param2, param3, modes int) int {
	mode, modes := extractMode(modes)
	value1 := getValue(prog, param1, mode)
	mode, modes = extractMode(modes)
	value2 := getValue(prog, param2, mode)
	mode, modes = extractMode(modes)
	if mode != 0 {
		// do some error handling
	}

	if value1 == value2 {
		prog[param3] = 1
	} else {
		prog[param3] = 0
	}
	return 4
}

func process(prog []int, input int) {
	var instrPtr int

Done:
	for instrPtr < len(prog) {
		opCode := prog[instrPtr] % 100
		modes := prog[instrPtr] / 100

		switch opCode {
		case add:
			instrPtr += executeAdd(prog, prog[instrPtr+1], prog[instrPtr+2], prog[instrPtr+3], modes)
		case mult:
			instrPtr += executeMult(prog, prog[instrPtr+1], prog[instrPtr+2], prog[instrPtr+3], modes)
		case save:
			instrPtr += executeSave(prog, prog[instrPtr+1], input)
		case load:
			// for this puzzle we'll just print to console
			instrPtr += executeLoad(prog, prog[instrPtr+1], modes)
		case jumpOnTrue:
			instrPtr = executeJump(true, prog, prog[instrPtr+1], prog[instrPtr+2], modes, instrPtr)
		case jumpOnFalse:
			instrPtr = executeJump(false, prog, prog[instrPtr+1], prog[instrPtr+2], modes, instrPtr)
		case lessThan:
			instrPtr += executeLessThan(prog, prog[instrPtr+1], prog[instrPtr+2], prog[instrPtr+3], modes)
		case equals:
			instrPtr += executeEquals(prog, prog[instrPtr+1], prog[instrPtr+2], prog[instrPtr+3], modes)
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
		[]int{1002, 4, 3, 4, 33},
		[]int{1101, 100, -1, 4, 0},
	}

	for _, v := range testData {
		fmt.Printf("Received value: %v, ", v)
		process(v, 1)
		fmt.Printf("Processed value: %v\n\n", v)
	}

	puzzleInput := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1101, 72, 36, 225, 1101, 87, 26, 225, 2, 144, 13, 224, 101, -1872, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1102, 66, 61, 225, 1102, 25, 49, 224, 101, -1225, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 5, 224, 1, 223, 224, 223, 1101, 35, 77, 224, 101, -112, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1002, 195, 30, 224, 1001, 224, -2550, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1102, 30, 44, 225, 1102, 24, 21, 225, 1, 170, 117, 224, 101, -46, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 224, 223, 223, 1102, 63, 26, 225, 102, 74, 114, 224, 1001, 224, -3256, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 224, 223, 223, 1101, 58, 22, 225, 101, 13, 17, 224, 101, -100, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 224, 223, 223, 1101, 85, 18, 225, 1001, 44, 7, 224, 101, -68, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 329, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 344, 1001, 223, 1, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 359, 1001, 223, 1, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 374, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 389, 101, 1, 223, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 404, 101, 1, 223, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 419, 1001, 223, 1, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 434, 101, 1, 223, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 449, 101, 1, 223, 223, 1108, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 464, 101, 1, 223, 223, 1007, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 494, 101, 1, 223, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 524, 101, 1, 223, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 539, 1001, 223, 1, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 554, 101, 1, 223, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 569, 101, 1, 223, 223, 8, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 584, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 599, 1001, 223, 1, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 614, 1001, 223, 1, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 629, 101, 1, 223, 223, 7, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 108, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 659, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}

	// puzzle 1 answer
	process(cloneArray(puzzleInput), 1)

	testData = [][]int{
		[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
		[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
	}

	testInput := [][]int{
		[]int{8, 9},
		[]int{3, 8},
		[]int{8, 12},
		[]int{7, 8},
		[]int{1, 0},
		[]int{-1, 0},
		[]int{7, 8, 9},
	}

	for i, data := range testData {
		for _, input := range testInput[i] {
			process(cloneArray(data), input)
		}
	}

	// puzzle 2 answer
	process(cloneArray(puzzleInput), 5)
}
