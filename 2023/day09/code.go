package day09

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne() int {
	// INPUT = "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"
	var result int
	calculation := func(p, c int) int { return p + c }

	for _, line := range strings.Split(INPUT, "\n") {
		sequences := generateSequences(convertLine(line))
		result += nextValue(sequences, calculation)
	}

	return result
}

func PartTwo() int {
	// INPUT = "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"
	var result int
	calculation := func(p, c int) int { return c - p }

	for _, line := range strings.Split(INPUT, "\n") {
		sequences := generateSequences(convertLine(line))

		for _, sequence := range sequences {
			slices.Reverse(sequence)
		}

		result += nextValue(sequences, calculation)
	}

	return result
}

func convertLine(line string) []int {
	result := []int{}
	for _, v := range strings.Split(line, " ") {
		number, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		result = append(result, number)
	}

	return result
}

func generateSequences(input []int) [][]int {
	sequences := [][]int{input}
	nextTotal := -1

	for nextTotal != 0 {
		nextTotal = 0
		currentSequence := sequences[len(sequences)-1]
		nextSequence := []int{}

		for i := 1; i < len(currentSequence); i++ {
			diff := currentSequence[i] - currentSequence[i-1]
			nextSequence = append(nextSequence, diff)
			nextTotal += diff
		}

		sequences = append(sequences, nextSequence)
	}

	return sequences
}

func nextValue(sequences [][]int, mathFn func(int, int) int) int {
	slices.Reverse(sequences)
	sequences[0] = append(sequences[0], 0)

	var nextValue int
	for i := 1; i < len(sequences); i++ {
		nextValue = mathFn(sequences[i-1][len(sequences[i-1])-1], sequences[i][len(sequences[i-1])-1])
		sequences[i] = append(sequences[i], nextValue)
	}

	return nextValue
}
