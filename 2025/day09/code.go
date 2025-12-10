package day09

import (
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

func PartOne() int {
	// INPUT = "7,1\n11,1\n11,7\n9,7\n9,5\n2,5\n2,3\n7,3"
	coords := []*Coord{}
	largestArea := 0

	for _, strPair := range strings.Split(INPUT, "\n") {
		parts := strings.Split(strPair, ",")
		newCoord := &Coord{X: convertToNum(parts[0]), Y: convertToNum(parts[1])}

		for _, existingCoord := range coords {
			x := math.Abs(float64(newCoord.X-existingCoord.X)) + 1
			y := math.Abs(float64(newCoord.Y-existingCoord.Y)) + 1
			area := int(x * y)
			if area > largestArea {
				largestArea = area
			}
		}

		coords = append(coords, newCoord)
	}

	return largestArea
}

func PartTwo() int {
	return 0
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
