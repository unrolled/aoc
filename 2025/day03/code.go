package day03

// Used AI for part 2... I was stomped!

import (
	"strconv"
	"strings"
)

func PartOne() int {
	// INPUT = "987654321111111\n811111111111119\n234234234234278\n818181911112111"

	totalJolts := 0
	banks := strings.Split(INPUT, "\n")

	for _, bank := range banks {
		chars := strings.Split(bank, "")
		first, second, index := 0, 0, -1

		for i := range chars {
			jolts := convertToNum(chars[i])

			if jolts > first && i > index && i < len(chars)-1 {
				first = jolts
				index = i
				second = convertToNum(chars[i+1])
			}

			if jolts > second && i != index {
				second = jolts
			}
		}

		totalJolts += (first * 10) + second
	}

	return totalJolts
}

func PartTwo() int {
	// INPUT = "987654321111111\n811111111111119\n234234234234278\n818181911112111"
	// INPUT = "987654321111111"
	// INPUT = "811111111111119"
	// INPUT = "234234234234278"
	// INPUT = "818181911112111"

	totalJolts := 0
	banks := strings.Split(INPUT, "\n")

	for _, bank := range banks {
		chars := strings.Split(bank, "")
		jolts, start, high, highIndex := 0, 11, 0, -1
		numbers := []int{}

		for i := len(chars) - 1; i >= 0; i-- {
			numbers = append(numbers, convertToNum(chars[i]))
		}

		end := len(numbers)

		for j := start; start > -1 && j < end; j++ {
			if numbers[j] >= high {
				high = numbers[j]
				highIndex = j
			}

			if j == end-1 {
				end = highIndex
				jolts = (jolts * 10) + high
				start += -1
				j = start - 1
				high = 0
				highIndex = -1
			}
		}

		totalJolts += jolts
	}

	return totalJolts
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
