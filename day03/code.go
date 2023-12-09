package day03

import (
	"fmt"
	"strconv"
	"strings"
)

var gearRation = map[string][]int{}
var lines = []string{}

func PartOne() int {
	// INPUT = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
	lines = strings.Split(INPUT, "\n")
	lines = append([]string{strings.Repeat(".", len(lines[0]))}, lines...)
	lines = append(lines, strings.Repeat(".", len(lines[0])))

	var result int
	for i := range lines {
		if i == 0 || i == len(lines)-1 {
			continue
		}

		lineNumbers := getNumbers(i, i-1, i+1)
		for _, n := range lineNumbers {
			result += n
		}
	}

	return result
}

func PartTwo() int {
	// Run part one if needed.
	if len(gearRation) == 0 {
		PartOne()
	}

	var gearResult int
	for _, numbers := range gearRation {
		if len(numbers) != 2 {
			continue
		}

		gearResult += (numbers[0] * numbers[1])
	}

	return gearResult
}

func getNumbers(current, top, bottom int) (results []int) {
	numStart := -1
	numEnd := -1
	numChars := ""

	for i, char := range lines[current] {
		if isCharNum(char) {
			if numStart == -1 {
				numStart = i
			}
			numChars += string(char)
			numEnd = i
		} else {
			if len(numChars) != 0 {
				value := convertToNum(numChars)
				if hasAdjacentSymbol(current, top, bottom, numStart, numEnd, value) {
					results = append(results, value)
				}

				numStart = -1
				numEnd = -1
				numChars = ""
			}
		}
	}

	if len(numChars) != 0 {
		value := convertToNum(numChars)
		if hasAdjacentSymbol(current, top, bottom, numStart, numEnd, value) {
			results = append(results, value)
		}
	}

	return
}

func hasAdjacentSymbol(current, top, bottom, start, end, value int) bool {
	if start > 0 {
		start -= 1
	}

	if end < len(lines[current])-1 {
		end += 1
	}

	lineResult := false
	for i := start; i <= end; i++ {
		for _, linePos := range []int{top, current, bottom} {
			if isSymbol(lines[linePos][i]) {
				lineResult = true

				key := fmt.Sprintf("%d-%d", linePos, i)
				if _, ok := gearRation[key]; !ok {
					gearRation[key] = []int{}
				}
				gearRation[key] = append(gearRation[key], value)
			}
		}
	}

	return lineResult
}

func isCharNum(c rune) bool {
	_, err := strconv.Atoi(string(c))
	if err != nil {
		return false
	}

	return true
}

func isSymbol(c byte) bool {
	if isCharNum(rune(c)) || c == '.' {
		return false
	}

	return true
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
