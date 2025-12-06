package day05

import (
	"strconv"
	"strings"
)

type FreshRange struct {
	Low, High int
}

func PartOne() int {
	// INPUT = "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32"
	inputParts := strings.Split(INPUT, "\n\n")

	freshRanges := []*FreshRange{}

	for _, freshRange := range strings.Split(inputParts[0], "\n") {
		parts := strings.Split(freshRange, "-")
		freshRanges = append(freshRanges, &FreshRange{Low: convertToNum(parts[0]), High: convertToNum(parts[1])})
	}

	results := map[int]bool{}
	for _, ingredient := range strings.Split(inputParts[1], "\n") {
		id := convertToNum(ingredient)
		for _, freshRange := range freshRanges {
			if id >= freshRange.Low && id <= freshRange.High {
				results[id] = true
			}
		}
	}

	return len(results)
}

func PartTwo() int {
	// INPUT = "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32"
	inputParts := strings.Split(INPUT, "\n\n")
	total := 0
	pairs := [][]int{}

	for _, freshRange := range strings.Split(inputParts[0], "\n") {
		parts := strings.Split(freshRange, "-")
		low := convertToNum(parts[0])
		high := convertToNum(parts[1])

		for _, existing := range pairs {
			existingLow := existing[0]
			existingHigh := existing[1]

			if low >= existingLow && low <= existingHigh {
				low = existingHigh + 1
			}
			if high <= existingHigh && high >= existingLow {
				high = existingLow - 1
			}
		}

		if low <= high {
			total += (high - low) + 1
			pairs = append(pairs, []int{low, high})
		}
	}

	return total - 1
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
