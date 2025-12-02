package day02

import (
	"strconv"
	"strings"
)

func PartOne() int {
	// INPUT = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	invalidSum := 0
	items := strings.Split(INPUT, ",")

	for _, item := range items {
		parts := strings.Split(item, "-")
		currentPos := convertToNum(parts[0])
		end := convertToNum(parts[1])

		for currentPos <= end {
			chars := strconv.Itoa(currentPos)
			if len(chars)%2 == 0 {
				parts := stringParts(chars, len(chars)/2)

				if parts[0] == parts[1] {
					invalidSum += currentPos
				}
			}
			currentPos++
		}
	}

	return invalidSum
}

func PartTwo() int {
	// INPUT = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	invalidSum := 0
	items := strings.Split(INPUT, ",")

	for _, item := range items {
		parts := strings.Split(item, "-")
		currentPos := convertToNum(parts[0])
		end := convertToNum(parts[1])

		for currentPos <= end {
			chars := strconv.Itoa(currentPos)
			charCount := 1

			for charCount <= len(chars)/2 {
				parts := stringParts(chars, charCount)
				if partsEqual(parts) {
					invalidSum += currentPos
					break
				}

				charCount++
			}

			currentPos++
		}
	}

	return invalidSum
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}

func stringParts(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}

	if chunkSize >= len(s) {
		return []string{s}
	}

	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0

	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}

		currentLen++
	}

	chunks = append(chunks, s[currentStart:])

	return chunks
}

func partsEqual(inputs []string) bool {
	if len(inputs) <= 1 {
		return false
	}

	first := inputs[0]
	for _, item := range inputs {
		if first != item {
			return false
		}
	}

	return true
}
