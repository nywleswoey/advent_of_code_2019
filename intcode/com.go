package intcode

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

// Program represents an Intcode computer program
type Program struct {
	Instructions []int
	instrPtr     int
	IsDone       bool
}

// NewProgram creates a new Intcode computer program
func NewProgram(prog []int) Program {
	return Program{Instructions: cloneArray(prog), instrPtr: 0}
}

func (p Program) getValue(param, mode int) int {
	if mode == 1 {
		return param
	}
	return p.Instructions[param]
}

func (p *Program) executeAdd(modes int) {
	param1 := p.Instructions[p.instrPtr+1]
	param2 := p.Instructions[p.instrPtr+2]
	param3 := p.Instructions[p.instrPtr+3]

	mode, modes := extractMode(modes)
	value1 := p.getValue(param1, mode)
	mode, modes = extractMode(modes)
	value2 := p.getValue(param2, mode)
	mode, modes = extractMode(modes)
	if mode != 0 {
		// do some error handling
	}
	p.Instructions[param3] = value1 + value2
	p.instrPtr += 4
}

func (p *Program) executeMult(modes int) {
	param1 := p.Instructions[p.instrPtr+1]
	param2 := p.Instructions[p.instrPtr+2]
	param3 := p.Instructions[p.instrPtr+3]

	mode, modes := extractMode(modes)
	value1 := p.getValue(param1, mode)
	mode, modes = extractMode(modes)
	value2 := p.getValue(param2, mode)
	mode, modes = extractMode(modes)
	if mode != 0 {
		// do some error handling
	}
	p.Instructions[param3] = value1 * value2
	p.instrPtr += 4
}

func (p *Program) executeSave(input int) {
	param1 := p.Instructions[p.instrPtr+1]

	p.Instructions[param1] = input
	p.instrPtr += 2
}

func (p *Program) executeLoad(modes int) int {
	param1 := p.Instructions[p.instrPtr+1]

	mode, modes := extractMode(modes)
	value := p.getValue(param1, mode)
	p.instrPtr += 2
	return value
}

func (p *Program) executeJump(onTrue bool, modes int) {
	param1 := p.Instructions[p.instrPtr+1]
	param2 := p.Instructions[p.instrPtr+2]

	mode, modes := extractMode(modes)
	value1 := p.getValue(param1, mode)
	mode, modes = extractMode(modes)
	value2 := p.getValue(param2, mode)
	if (onTrue && value1 != 0) || (!onTrue && value1 == 0) {
		p.instrPtr = value2
	} else {
		p.instrPtr += 3
	}
}

func (p *Program) executeLessThan(modes int) {
	param1 := p.Instructions[p.instrPtr+1]
	param2 := p.Instructions[p.instrPtr+2]
	param3 := p.Instructions[p.instrPtr+3]

	mode, modes := extractMode(modes)
	value1 := p.getValue(param1, mode)
	mode, modes = extractMode(modes)
	value2 := p.getValue(param2, mode)
	mode, modes = extractMode(modes)
	if mode != 0 {
		// do some error handling
	}

	if value1 < value2 {
		p.Instructions[param3] = 1
	} else {
		p.Instructions[param3] = 0
	}
	p.instrPtr += 4
}

func (p *Program) executeEquals(modes int) {
	param1 := p.Instructions[p.instrPtr+1]
	param2 := p.Instructions[p.instrPtr+2]
	param3 := p.Instructions[p.instrPtr+3]

	mode, modes := extractMode(modes)
	value1 := p.getValue(param1, mode)
	mode, modes = extractMode(modes)
	value2 := p.getValue(param2, mode)
	mode, modes = extractMode(modes)
	if mode != 0 {
		// do some error handling
	}

	if value1 == value2 {
		p.Instructions[param3] = 1
	} else {
		p.Instructions[param3] = 0
	}
	p.instrPtr += 4
}

// Process will run the Intcode program using the inputs
func (p *Program) Process(input <-chan int) <-chan int {
	prog := p.Instructions
	output := make(chan int)

	go func() {
	Done:
		for p.instrPtr < len(prog) {
			opCode := prog[p.instrPtr] % 100
			modes := prog[p.instrPtr] / 100

			switch opCode {
			case add:
				p.executeAdd(modes)
			case mult:
				p.executeMult(modes)
			case save:
				isWaiting := true
				for isWaiting {
					select {
					case result := <-input:
						isWaiting = false
						p.executeSave(result)
					}
				}
			case load:
				output <- p.executeLoad(modes)
			case jumpOnTrue:
				p.executeJump(true, modes)
			case jumpOnFalse:
				p.executeJump(false, modes)
			case lessThan:
				p.executeLessThan(modes)
			case equals:
				p.executeEquals(modes)
			case end:
				p.instrPtr++
				break Done
			default:
				fmt.Printf("encountered an invalid op code: %v at address %v\n", p.Instructions[p.instrPtr], p.instrPtr)
				break Done
			}
		}

		p.IsDone = true
		close(output)
	}()

	return output
}
