package day01

import (
	"strconv"
	"strings"
)

const dialStartPos = 50

func PartOne() int {
	// INPUT = "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"

	zeros := 0
	currentPos := dialStartPos
	lines := strings.Split(INPUT, "\n")

	for _, line := range lines {
		var val int

		if num, found := strings.CutPrefix(line, "L"); found {
			val = convertToNum(num)
			val = val * -1
		} else {
			val = convertToNum(strings.TrimPrefix(line, "R"))
		}

		currentPos = currentPos + val

		for currentPos >= 100 {
			currentPos -= 100
		}

		for currentPos < 0 {
			currentPos += 100
		}

		if currentPos == 0 {
			zeros++
		}
	}

	return zeros
}

func PartTwo() int {
	// INPUT = "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"

	zeros := 0
	currentPos := dialStartPos
	lines := strings.Split(INPUT, "\n")

	for _, line := range lines {
		var val int

		if num, found := strings.CutPrefix(line, "L"); found {
			val = convertToNum(num)

			for val != 0 {
				currentPos--
				val--

				if currentPos == 0 {
					zeros++
				}

				if currentPos == -1 {
					currentPos = 99
				}
			}
		} else {
			val = convertToNum(strings.TrimPrefix(line, "R"))

			for val < 0 {
				val += 100
				zeros++
			}

			for val != 0 {
				currentPos++
				val--

				if currentPos == 100 {
					currentPos = 0
					zeros++
				}
			}
		}
	}

	return zeros
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
