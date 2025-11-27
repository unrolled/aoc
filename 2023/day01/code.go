package day01

import (
	"fmt"
	"strconv"
	"strings"
)

var numberWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func PartOne() int {
	// 	INPUT = `1abc2
	// pqr3stu8vwx
	// a1b2c3d4e5f
	// treb7uchet`

	var result int
	for _, line := range strings.Split(INPUT, "\n") {
		result += getNum(line)
	}

	return result
}

func PartTwo() int {
	// 	INPUT = `two1nine
	// eightwothree
	// abcone2threexyz
	// xtwone3four
	// 4nineeightseven2
	// zoneight234
	// 7pqrstsixteen`

	for index, word := range numberWords {
		INPUT = strings.ReplaceAll(INPUT, word, fmt.Sprintf("%s%d%s", word, index+1, word))
	}

	return PartOne()
}

func getNum(input string) int {
	numbers := []int{}
	for _, char := range input {
		num, err := strconv.Atoi(string(char))
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	strResult := fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1])
	return convertToNum(strResult)
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
