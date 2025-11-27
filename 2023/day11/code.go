package day11

import (
	"strings"
)

type Coords struct {
	Y, X int
}

type BlackHoles struct {
	Y, X map[int]int
}

func PartOne() int {
	// INPUT = "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....." // 374@2
	return run(INPUT, 2)
}

func PartTwo() int {
	// INPUT = "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....." // 1030@10
	return run(INPUT, 1_000_000)
}

func run(input string, expansion int) int {
	lines := strings.Split(input, "\n")
	vertical := make([]bool, len(lines[0]))

	galaxies := []Coords{}
	holes := BlackHoles{}
	holes.X = map[int]int{}
	holes.Y = map[int]int{}

	for y, line := range lines {
		hasGalaxy := false
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Coords{Y: y, X: x})
				vertical[x] = true
				hasGalaxy = true
			}
		}

		if !hasGalaxy {
			holes.Y[y] = expansion - 1
		}
	}

	for x, col := range vertical {
		if col == false {
			holes.X[x] = expansion - 1
		}
	}

	total := 0
	for index, currentGalaxy := range galaxies {
		for i := index; i < len(galaxies); i++ {
			total += steps(currentGalaxy.X, galaxies[i].X, holes.X)
			total += steps(currentGalaxy.Y, galaxies[i].Y, holes.Y)
		}
	}

	return total
}

func steps(start, end int, holes map[int]int) int {
	if start > end {
		start, end = end, start
	}

	total := end - start

	for hole, val := range holes {
		if hole > start && hole < end {
			total += val
		}
	}

	return total
}
