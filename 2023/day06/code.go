package day06

import (
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func PartOne() int {
	// INPUT = "Time:      7  15   30\nDistance:  9  40  200"
	data := parseInput(INPUT)

	result := 1
	for _, race := range data {
		result = result * waysToWin(race.Time, race.Distance)
	}

	return result
}

func PartTwo() int {
	// INPUT = "Time:      7  15   30\nDistance:  9  40  200"
	data := parseInput(INPUT)

	var raceTime, raceDistance string
	for _, race := range data {
		raceTime += strconv.Itoa(race.Time)
		raceDistance += strconv.Itoa(race.Distance)
	}

	return waysToWin(convertToNum(raceTime), convertToNum(raceDistance))
}

func parseInput(input string) []*Race {
	lines := strings.Split(input, "\n")
	times := strings.Fields(strings.TrimPrefix(lines[0], "Time:"))
	distances := strings.Fields(strings.TrimPrefix(lines[1], "Distance:"))

	var races []*Race
	for i := 0; i < len(times); i++ {
		races = append(races, &Race{Time: convertToNum(times[i]), Distance: convertToNum(distances[i])})
	}

	return races
}

func waysToWin(raceTime, raceDistance int) int {
	var waysToWin int

	for i := 1; i < raceTime; i++ {
		speed := i
		remainingTime := raceTime - i
		distance := speed * remainingTime

		if distance > raceDistance {
			waysToWin++
		}
	}

	return waysToWin
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
