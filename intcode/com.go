package intcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Program represents an Intcode computer program
type Program struct {
	prog   []int
	reader *bufio.Reader
}

// NewProgram creates a new Intcode computer program
func NewProgram(prog []int) Program {
	return Program{prog: cloneArray(prog), reader: bufio.NewReader(os.Stdin)}
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

// Process will run the Intcode program using the inputs
func (p Program) Process() {
	var instrPtr int
	prog := p.prog
	reader := p.reader

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
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			input, _ := strconv.Atoi(text)
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
