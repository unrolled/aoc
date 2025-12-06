package day06

import (
	"strconv"
	"strings"
)

func PartOne() int {
	// INPUT = "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  "
	lines := strings.Split(INPUT, "\n")

	isSpace := func(pos int) bool {
		for _, line := range lines {
			if line[pos] != ' ' {
				return false
			}
		}

		return true
	}

	mathRange := func(min, max int) int {
		total := 0

		op := strings.TrimSpace(lines[len(lines)-1][min:max])
		if op == "*" {
			total = 1
		}

		for _, line := range lines[0 : len(lines)-1] {
			num := convertToNum(strings.TrimSpace(line[min:max]))
			if op == "*" {
				total = total * num
			} else {
				total = total + num
			}

		}

		return total
	}

	total := 0
	lastSpace := -1
	for i := range lines[0] {
		if isSpace(i) || i == len(lines[0])-1 {
			total += mathRange(lastSpace+1, i+1)
			lastSpace = i
		}
	}

	return total
}

func PartTwo() int {
	// INPUT = "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  "
	lines := strings.Split(INPUT, "\n")

	isSpace := func(pos int) bool {
		for _, line := range lines {
			if line[pos] != ' ' {
				return false
			}
		}

		return true
	}

	mathRange := func(min, max int) int {
		total := 0

		op := strings.TrimSpace(lines[len(lines)-1][min:max])
		if op == "*" {
			total = 1
		}

		for i := min; i < max; i++ {
			value := []byte{}

			for _, line := range lines[0 : len(lines)-1] {
				val := line[i]
				if val != ' ' {
					value = append(value, val)
				}
			}

			if len(value) > 0 {
				if op == "*" {
					total = total * convertToNum(string(value))
				} else {
					total = total + convertToNum(string(value))
				}
			}
		}

		return total
	}

	total := 0
	lastSpace := -1
	for i := range lines[0] {
		if isSpace(i) || i == len(lines[0])-1 {
			total += mathRange(lastSpace+1, i+1)
			lastSpace = i
		}
	}

	return total
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
