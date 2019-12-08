package main

import (
	"fmt"
	"strconv"
)

func hasDuplicateDigit(p string) string {
	for i := 1; i < len(p); i++ {
		if p[i-1:i] == p[i:i+1] {
			return ""
		}
	}
	return "no double"
}

func hasDuplicateDigitV2(p string) string {
	// also checks that the adjacent digits are
	// not part of a larger group of matching digits
	if len(p) < 1 {
		return "no double"
	}

	result := "no double"
	consecutiveCount := 1
	matchedDigit := p[:1]
	for i := 1; i < len(p); i++ {
		if p[i:i+1] == matchedDigit {
			consecutiveCount++
			result = ""

			if consecutiveCount > 2 {
				result = fmt.Sprintf("the repeated %v%v is part of a larger group of %v%v%v", matchedDigit, matchedDigit, matchedDigit, matchedDigit, matchedDigit)
			}
		} else {
			if consecutiveCount == 2 {
				return ""
			}

			matchedDigit = p[i : i+1]
			consecutiveCount = 1
		}
	}
	return result
}

func hasNoDecreasingPair(p string) string {
	for i := 0; i < len(p)-1; i++ {
		digit1, err := strconv.Atoi(p[i : i+1])
		if err != nil {
			return "received invalid digit"
		}
		digit2, err := strconv.Atoi(p[i+1 : i+2])
		if err != nil {
			return "received invalid digit"
		}
		if digit1 > digit2 {
			return fmt.Sprintf("decreasing pair of digits %v%v", digit1, digit2)
		}
	}
	return ""
}

func checkValidityV2(p string) string {
	result := hasNoDecreasingPair(p)
	if len(result) > 0 {
		return result
	}

	return hasDuplicateDigitV2(p)
}

func checkValidity(p string) string {
	decreasingPairResult := hasNoDecreasingPair(p)
	if len(decreasingPairResult) > 0 {
		return decreasingPairResult
	}

	return hasDuplicateDigit(p)

}

func main() {
	// test input
	testData := []string{
		"111111",
		"223450",
		"123789",
	}

	for _, v := range testData {
		status := checkValidity(v)
		fmt.Printf("Error for %v: %v\n", v, status)
	}

	passedCount := 0
	for candidate := 265275; candidate <= 781584; candidate++ {
		puzzleInput := strconv.Itoa(candidate)
		status := checkValidity(puzzleInput)
		if len(status) == 0 {
			passedCount++
		}
	}

	// puzzle 1 answer

	fmt.Printf("Puzzle 1 answer: %v\n\n", passedCount)

	// puzzle 2 test
	testData = []string{
		"112233",
		"123444",
		"111122",
	}

	for _, v := range testData {
		status := checkValidityV2(v)
		fmt.Printf("Error for %v: %v\n", v, status)
	}

	passedCount = 0
	for candidate := 265275; candidate <= 781584; candidate++ {
		puzzleInput := strconv.Itoa(candidate)
		status := checkValidityV2(puzzleInput)
		if len(status) == 0 {
			passedCount++
		}
	}

	// puzzle 2 answer

	fmt.Printf("Puzzle 2 answer: %v\n\n", passedCount)
}
