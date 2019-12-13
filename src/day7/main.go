package main

import (
	"fmt"

	"github.com/nywleswoey/advent_of_code_2019/intcode"
)

// credits from https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func startPuzzle1(data []int, phaseSettings []int, initialInput int) int {
	inChannels := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
	}

	outChannels := []<-chan int{}

	for i := 0; i < 5; i++ {
		amp := intcode.NewProgram(data)
		outChannels = append(outChannels, amp.Process(inChannels[i]))
	}

	for i := 0; i < 5; i++ {
		inChannels[i] <- phaseSettings[i]
	}
	inChannels[0] <- initialInput

	for i := 0; i < 4; i++ {
		go func(index int) {
			for result := range outChannels[index] {
				inChannels[index+1] <- result
			}
		}(i)
	}

	var result int

Done:
	for {
		select {
		case result = <-outChannels[4]:
			break Done
		}
	}

	return result
}

func startPuzzle2(data []int, phaseSettings []int, initialInput int) int {
	inChannels := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
	}

	outChannels := make([]<-chan int, 5)

	ampA := intcode.NewProgram(data)
	outChannels[0] = ampA.Process(inChannels[0])

	for i := 1; i < 5; i++ {
		amp := intcode.NewProgram(data)
		outChannels[i] = amp.Process(inChannels[i])
	}

	for i := 0; i < 5; i++ {
		inChannels[i] <- phaseSettings[i]
	}
	inChannels[0] <- initialInput

	for i := 0; i < 4; i++ {
		go func(index int) {
			for result := range outChannels[index] {
				inChannels[index+1] <- result
			}
		}(i)
	}

	var result int
	for result = range outChannels[4] {
		if !ampA.IsDone {
			inChannels[0] <- result
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testData := [][]int{
		[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
		[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
		[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
	}

	testSequence := [][]int{
		[]int{4, 3, 2, 1, 0},
		[]int{0, 1, 2, 3, 4},
		[]int{1, 0, 4, 3, 2},
	}

	testResult := []string{
		"43210",
		"54321",
		"65210",
	}

	for i, data := range testData {
		fmt.Printf("Input data: %v\n", data)
		fmt.Printf("Phase setting sequence: %v\n", testSequence[i])
		result := startPuzzle1(data, testSequence[i], 0)
		fmt.Printf("Expecting: %v, Received: %v\n", testResult[i], result)
	}
	fmt.Println("")

	puzzleInput := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 42, 67, 88, 101, 114, 195, 276, 357, 438, 99999, 3, 9, 101, 3, 9, 9, 1002, 9, 4, 9, 1001, 9, 5, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 1002, 9, 2, 9, 101, 2, 9, 9, 102, 2, 9, 9, 1001, 9, 5, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 1001, 9, 3, 9, 102, 4, 9, 9, 101, 4, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99}
	permutations := permutation([]int{0, 1, 2, 3, 4})
	highest := 0
	for _, permutation := range permutations {
		result := startPuzzle1(puzzleInput, permutation, 0)
		highest = max(highest, result)
	}

	// puzzle 1 answer
	fmt.Printf("Puzzle 1 answer: %v\n", highest)

	testData = [][]int{
		[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
		[]int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
	}

	testSequence = [][]int{
		[]int{9, 8, 7, 6, 5},
		[]int{9, 7, 8, 5, 6},
	}

	testResult = []string{
		"139629729",
		"18216",
	}

	for i, data := range testData {
		fmt.Printf("Input data: %v\n", data)
		fmt.Printf("Phase setting sequence: %v\n", testSequence[i])
		fmt.Printf("Expecting: %v, Received: %v\n", testResult[i], startPuzzle2(data, testSequence[i], 0))
	}
	fmt.Println("")

	// puzzle 1 answer
	highest = 0
	for _, permutation := range permutations {
		result := startPuzzle2(puzzleInput, permutation, 0)
		highest = max(highest, result)
	}
	fmt.Printf("Puzzle 2 answer: %v\n", highest)
}
