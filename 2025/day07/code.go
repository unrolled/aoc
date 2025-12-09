package day07

import (
	"strings"
)

func PartOne() int {
	// INPUT = ".......S.......\n...............\n.......^.......\n...............\n......^.^......\n...............\n.....^.^.^.....\n...............\n....^.^...^....\n...............\n...^.^...^.^...\n...............\n..^...^.....^..\n...............\n.^.^.^.^.^...^.\n..............."
	inputLines := strings.Split(INPUT, "\n")
	beamPaths := map[int]bool{strings.LastIndex(inputLines[0], "S"): true}
	splits := 0

	for _, line := range inputLines[1:] {
		for beamIndex := range beamPaths {
			if line[beamIndex] == '^' {
				delete(beamPaths, beamIndex)

				if before, _ := beamPaths[beamIndex-1]; !before {
					beamPaths[beamIndex-1] = true
				}

				if after, _ := beamPaths[beamIndex+1]; !after {
					beamPaths[beamIndex+1] = true
				}

				splits++
			}
		}
	}

	return splits
}

// var completeBeams int

func PartTwo() int {
	return 0
	// INPUT = ".......S.......\n...............\n.......^.......\n...............\n......^.^......\n...............\n.....^.^.^.....\n...............\n....^.^...^....\n...............\n...^.^...^.^...\n...............\n..^...^.....^..\n...............\n.^.^.^.^.^...^.\n..............."
	// inputLines := strings.Split(INPUT, "\n")
	// beamPaths := map[int]int{strings.LastIndex(inputLines[0], "S"): 1}
	// possibleBeams := 1

	// incrementBeamPath := func(line, index int) {
	// 	key := fmt.Sprintf("%d-%d", line, index)
	// 	if _, ok := beamPaths[key]; ok {

	// 	}
	// }

	// for _, line := range inputLines[1:] {
	// 	d := []rune(line)
	// 	lineBeams := 0
	// 	for beamIndex := range beamPaths {
	// 		switch line[beamIndex] {
	// 		case '^':
	// 			delete(beamPaths, beamIndex)

	// 			if before, _ := beamPaths[beamIndex-1]; !before {
	// 				lineBeams++
	// 				beamPaths[beamIndex-1] = true
	// 				d[beamIndex-1] = '|'
	// 			}

	// 			if after, _ := beamPaths[beamIndex+1]; !after {
	// 				lineBeams++
	// 				beamPaths[beamIndex+1] = true
	// 				d[beamIndex+1] = '|'
	// 			}
	// 		case '|':
	// 			lineBeams++
	// 		case '.':
	// 			d[beamIndex] = '|'
	// 		}
	// 	}

	// 	fmt.Printf("%s -- %d\n", string(d), lineBeams)
	// 	possibleBeams += lineBeams
	// }

	// return possibleBeams
}

// func PartTwo() int {
// 	INPUT = ".......S.......\n...............\n.......^.......\n...............\n......^.^......\n...............\n.....^.^.^.....\n...............\n....^.^...^....\n...............\n...^.^...^.^...\n...............\n..^...^.....^..\n...............\n.^.^.^.^.^...^.\n..............."
// 	inputLines := strings.Split(INPUT, "\n")
// 	startIndex := strings.LastIndex(inputLines[0], "S")
// 	dataLines = inputLines[1:]
// 	beamPaths := map[int]bool{strings.LastIndex(inputLines[0], "S"): true}

// 	findPaths(0, startIndex)

// 	for _, line := range inputLines[1:] {
// 		for beamIndex := range beamPaths {
// 			if line[beamIndex] == '^' {
// 				delete(beamPaths, beamIndex)

// 				if before, _ := beamPaths[beamIndex-1]; !before {
// 					beamPaths[beamIndex-1] = true
// 				}

// 				if after, _ := beamPaths[beamIndex+1]; !after {
// 					beamPaths[beamIndex+1] = true
// 				}

// 				splits++
// 			}
// 		}
// 	}

// 	fmt.Printf("%#v\n", cache)

// 	return len(cache)
// }

// var dataLines []string
// var cache = map[string]int{}

// func findPaths(lineIndex, beamIndex int) {
// 	cacheKey := fmt.Sprintf("%d-%d", lineIndex, beamIndex)
// 	if _, exists := cache[cacheKey]; exists {
// 		return
// 	}

// 	splits := 0
// 	beamPaths := map[int]bool{beamIndex: true}

// 	for _, line := range dataLines[lineIndex:] {
// 		if line[beamIndex] == '^' {
// 			delete(beamPaths, beamIndex)

// 			if before, _ := beamPaths[beamIndex-1]; !before {
// 				beamPaths[beamIndex-1] = true
// 				findPaths(lineIndex+1, beamIndex-1)
// 			}

// 			if after, _ := beamPaths[beamIndex+1]; !after {
// 				beamPaths[beamIndex+1] = true
// 				findPaths(lineIndex+1, beamIndex+1)
// 			}

// 			splits++
// 		}
// 	}

// 	cache[cacheKey] = splits
// }
