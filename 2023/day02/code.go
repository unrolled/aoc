package day02

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type LineData struct {
	ID               int
	Power            int
	Red, Blue, Green int
}

func PartOne() int {
	// 	INPUT = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	INPUT = strings.ReplaceAll(INPUT, "Game ", "")

	rawLines := []LineData{}
	var result int
	for _, line := range strings.Split(INPUT, "\n") {
		gameResult := convertLineDate(line)
		rawLines = append(rawLines, gameResult)

		if gameResult.Blue <= maxBlue && gameResult.Green <= maxGreen && gameResult.Red <= maxRed {
			result += gameResult.ID
		}
	}

	return result
}

func PartTwo() int {
	// 	INPUT = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	INPUT = strings.ReplaceAll(INPUT, "Game ", "")

	rawLines := []LineData{}
	var result int
	for _, line := range strings.Split(INPUT, "\n") {
		gameResult := convertLineDate(line)
		rawLines = append(rawLines, gameResult)

		power := gameResult.Blue * gameResult.Red * gameResult.Green
		result += power
	}

	return result
}

func convertLineDate(input string) LineData {
	parts := strings.Split(input, ":")
	input = strings.ReplaceAll(parts[1], ";", ",")
	input = strings.TrimSpace(input)
	data := LineData{ID: convertToNum(parts[0])}

	for _, colors := range strings.Split(input, ",") {
		color, count := "", 0
		fmt.Sscanf(colors, "%d %s", &count, &color)

		switch color {
		case "red":
			if count > data.Red {
				data.Red = count
			}
		case "green":
			if count > data.Green {
				data.Green = count
			}
		case "blue":
			if count > data.Blue {
				data.Blue = count
			}
		}
	}

	return data
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
