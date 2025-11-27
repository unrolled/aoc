package day10

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	NS    = '|' // | is a vertical pipe connecting north and south.
	EW    = '-' // - is a horizontal pipe connecting east and west.
	NE    = 'L' // L is a 90-degree bend connecting north and east.
	NW    = 'J' // J is a 90-degree bend connecting north and west.
	SW    = '7' // 7 is a 90-degree bend connecting south and west.
	SE    = 'F' // F is a 90-degree bend connecting south and east.
	START = 'S' // S is the starting position of the animal; there is a pipe on this tile.
)

func PartOne() int {
	// INPUT = ".....\n.S-7.\n.|.|.\n.L-J.\n....." // 4
	// INPUT = "7-F7-\n.FJ|7\nSJLL7\n|F--J\nLJ.LJ" // 8
	grid := createGrid(INPUT)
	grid.InitDirection()

	for !grid.IsComplete() {
		grid.Step()
	}

	return grid.steps / 2
}

func PartTwo() int {
	// INPUT = "...........\n.S-------7.\n.|F-----7|.\n.||.....||.\n.||.....||.\n.|L-7.F-J|.\n.|..|.|..|.\n.L--J.L--J.\n..........." // 4
	// INPUT = ".F----7F7F7F7F-7....\n.|F--7||||||||FJ....\n.||.FJ||||||||L7....\nFJL7L7LJLJ||LJ.L-7..\nL--J.L7...LJS7F-7L7.\n....F-J..F7FJ|L7L7L7\n....L7.F7||L7|.L7L7|\n.....|FJLJ|FJ|F7|.LJ\n....FJL-7.||.||||...\n....L---J.LJ.LJLJ..." // 8
	// INPUT = "FF7FSF7F7F7F7F7F---7\nL|LJ||||||||||||F--J\nFL-7LJLJ||||||LJL-77\nF--JF--7||LJLJ7F7FJ-\nL---JF-JLJ.||-FJLJJ7\n|F|F-JF---7F7-L7L|7|\n|FFJF7L7F-JF7|JL---7\n7-L-JL7||F7|L7F-7F7|\nL.L7LFJ|||||FJL7||LJ\nL7JLJL-JLJLJL--JLJ.L" // 10

	grid := createGrid(INPUT)
	grid.InitDirection()

	for !grid.IsComplete() {
		grid.Step()
	}

	return grid.EnclosedTiles()
}

func createGrid(input string) *Grid {
	data := [][]byte{}
	var x, y int

	for i, line := range strings.Split(input, "\n") {
		data = append(data, []byte(line))
		if pos := bytes.Index([]byte(line), []byte{START}); pos != -1 {
			x, y = pos, i
		}
	}

	return &Grid{data: data, cX: x, cY: y, history: map[string]bool{}}
}

type Grid struct {
	data           [][]byte
	cX, cY, pX, pY int
	steps          int
	history        map[string]bool
}

func (g *Grid) MaxX() int {
	return len(g.data[0])
}

func (g *Grid) MaxY() int {
	return len(g.data)
}

func (g *Grid) Step() {
	g.steps++
	g.history[fmt.Sprintf("%d-%d", g.cY, g.cX)] = true
	newX, newY := g.cX, g.cY

	switch g.data[g.cY][g.cX] {
	case NS: // '|'
		if newY+1 == g.pY {
			newY--
		} else {
			newY++
		}
	case EW: // '-'
		if newX+1 == g.pX {
			newX--
		} else {
			newX++
		}
	case NE: // 'L'
		if newY-1 == g.pY {
			newX++
		} else {
			newY--
		}
	case NW: // 'J'
		if newY-1 == g.pY {
			newX--
		} else {
			newY--
		}
	case SW: // '7'
		if newY+1 == g.pY {
			newX--
		} else {
			newY++
		}
	case SE: // 'F'
		if newY+1 == g.pY {
			newX++
		} else {
			newY++
		}
	}

	g.pX, g.pY, g.cX, g.cY = g.cX, g.cY, newX, newY
}

func (g *Grid) InitDirection() {
	g.steps++
	g.history[fmt.Sprintf("%d-%d", g.cY, g.cX)] = true

	newX, newY := g.cX, g.cY

	n, e, s, w := g.data[g.cY-1][g.cX], g.data[g.cY][g.cX+1], g.data[g.cY+1][g.cX], g.data[g.cY][g.cX-1]

	switch {
	case g.cY != 0 && (n == NS || n == SW || n == SE): // look north
		newY = g.cY - 1
	case g.cX+1 != g.MaxX() && (e == EW || e == SW || e == NW): // look east
		newX = g.cX + 1
	case g.cY+1 != g.MaxY() && (s == NS || s == NE || s == NW): // look south
		newY = g.cY + 1
	case g.cX-1 != 0 && (w == EW || w == NE || w == SE): // look west
		newX = g.cX - 1
	}

	g.pX, g.pY, g.cX, g.cY = g.cX, g.cY, newX, newY
}

func (g *Grid) IsComplete() bool {
	return g.data[g.cY][g.cX] == START && g.steps > 0
}

func (g *Grid) EnclosedTiles() int {
	var total int
	var skip byte = '?'

	for y := 0; y < g.MaxY(); y++ {
		var include bool

		for x := 0; x < g.MaxX(); x++ {
			if _, exist := g.history[fmt.Sprintf("%d-%d", y, x)]; exist {
				switch {
				case g.data[y][x] == SE:
					skip = SW
					continue
				case g.data[y][x] == NE:
					skip = NW
					continue
				case g.data[y][x] == EW: // ignore
					continue
				default:
					if g.data[y][x] != skip {
						include = !include
					}
					skip = '?'
				}
			} else if include {
				total++
			}
		}
	}

	return total
}
