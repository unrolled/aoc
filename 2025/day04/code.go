package day04

import (
	"strings"
)

func PartOne() int {
	// INPUT = "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@."

	accessibleRolls := 0

	grid := map[int]map[int]int{}
	for i, line := range strings.Split(INPUT, "\n") {
		grid[i] = map[int]int{}

		for j, char := range strings.Split(line, "") {
			var num int
			if char == "@" {
				num = 1
			}

			grid[i][j] = num
		}
	}

	fnRollExists := func(i, j int) int {
		if i < 0 || i > len(grid) {
			return 0
		}

		if j < 0 || j > len(grid[0]) {
			return 0
		}

		return grid[i][j]
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			spotCount := 0
			spotCount += fnRollExists(i+1, j)
			spotCount += fnRollExists(i-1, j)
			spotCount += fnRollExists(i+1, j+1)
			spotCount += fnRollExists(i-1, j+1)
			spotCount += fnRollExists(i+1, j-1)
			spotCount += fnRollExists(i-1, j-1)
			spotCount += fnRollExists(i, j-1)
			spotCount += fnRollExists(i, j+1)

			if spotCount < 4 {
				accessibleRolls++
			}
		}
	}

	return accessibleRolls
}

func PartTwo() int {
	// INPUT = "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@."

	totalRemovedRolls := 0

	previous := map[int]map[int]int{}
	next := map[int]map[int]int{}
	for i, line := range strings.Split(INPUT, "\n") {
		previous[i] = map[int]int{}
		next[i] = map[int]int{}

		for j, char := range strings.Split(line, "") {
			var num int
			if char == "@" {
				num = 1
			}

			previous[i][j] = num
			next[i][j] = num
		}
	}

	fnRollExists := func(m map[int]map[int]int, i, j int) int {
		if i < 0 || i > len(m) {
			return 0
		}

		if j < 0 || j > len(m[0]) {
			return 0
		}

		return m[i][j]
	}

	changesMade := -1

	for changesMade != 0 {
		changesMade = 0

		for i := 0; i < len(previous); i++ {
			for j := 0; j < len(previous[0]); j++ {
				if previous[i][j] == 0 {
					continue
				}

				spotCount := 0
				spotCount += fnRollExists(previous, i+1, j)
				spotCount += fnRollExists(previous, i-1, j)
				spotCount += fnRollExists(previous, i+1, j+1)
				spotCount += fnRollExists(previous, i-1, j+1)
				spotCount += fnRollExists(previous, i+1, j-1)
				spotCount += fnRollExists(previous, i-1, j-1)
				spotCount += fnRollExists(previous, i, j-1)
				spotCount += fnRollExists(previous, i, j+1)

				if spotCount < 4 {
					next[i][j] = 0
					changesMade++
				}
			}
		}

		for i := 0; i < len(next); i++ {
			for j := 0; j < len(next[0]); j++ {
				previous[i][j] = next[i][j]
			}
		}

		totalRemovedRolls += changesMade
	}

	return totalRemovedRolls
}
