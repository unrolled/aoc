package day16

import (
	"fmt"
	"strings"
)

const (
	dirRight = iota
	dirLeft
	dirDown
	dirUp
)

func PartOne() int {
	// INPUT = ".|...\\....\n|.-.\\.....\n.....|-...\n........|.\n..........\n.........\\\n..../.\\\\..\n.-.-/..|..\n.|....-|.\\\n..//.|...." // 46
	runner := &Runner{
		Grid:      parseGrid(INPUT),
		Energized: map[string]bool{},
		HasRun:    map[string]bool{},
	}
	runner.Add(dirRight, 0, -1)

	return runner.Run()
}

func PartTwo() int {
	// INPUT = ".|...\\....\n|.-.\\.....\n.....|-...\n........|.\n..........\n.........\\\n..../.\\\\..\n.-.-/..|..\n.|....-|.\\\n..//.|...." // 51
	mostEnergized := 0
	runner := &Runner{
		Grid:      parseGrid(INPUT),
		Energized: map[string]bool{},
		HasRun:    map[string]bool{},
	}
	maxY, maxX := len(runner.Grid), len(runner.Grid[0])

	for x := 0; x < maxX; x++ {
		runner.Clear()
		runner.Add(dirDown, -1, x)
		if result := runner.Run(); result > mostEnergized {
			mostEnergized = result
		}

		runner.Clear()
		runner.Add(dirUp, maxY, x)
		if result := runner.Run(); result > mostEnergized {
			mostEnergized = result
		}
	}

	for y := 0; y < maxY; y++ {
		runner.Clear()
		runner.Add(dirRight, y, -1)
		if result := runner.Run(); result > mostEnergized {
			mostEnergized = result
		}

		runner.Clear()
		runner.Add(dirLeft, y, maxX)
		if result := runner.Run(); result > mostEnergized {
			mostEnergized = result
		}
	}

	return mostEnergized
}

func parseGrid(input string) [][]byte {
	grid := [][]byte{}

	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []byte(line))
	}

	return grid
}

type Runner struct {
	Grid         [][]byte
	Energized    map[string]bool
	HasRun       map[string]bool
	Contraptions []*Contraption
}

func (r *Runner) Add(d, y, x int) {
	r.Contraptions = append(r.Contraptions, &Contraption{
		Runner: r,
		Dir:    d, Y: y, X: x,
		Key: fmt.Sprintf("%d-%d-%d", d, y, x),
	})
}

func (r *Runner) Run() int {
	for len(r.Contraptions) != 0 {
		current := r.Contraptions[0]
		current.Run()
		r.Contraptions = r.Contraptions[1:]
	}

	return len(r.Energized)
}

func (r *Runner) Clear() {
	clear(r.Energized)
	clear(r.HasRun)
	clear(r.Contraptions)
}

type Contraption struct {
	Dir, Y, X int
	Key       string
	Done      bool
	Runner    *Runner
}

func (c *Contraption) Run() {
	if _, exists := c.Runner.HasRun[c.Key]; exists {
		return
	}

	for !c.Done {
		c.Step()
	}

	c.Runner.HasRun[c.Key] = true
}

func (c *Contraption) Step() {
	switch c.Dir {
	case dirRight:
		c.X++
	case dirLeft:
		c.X--
	case dirDown:
		c.Y++
	case dirUp:
		c.Y--
	}

	if c.X < 0 || c.X == len(c.Runner.Grid[0]) || c.Y < 0 || c.Y == len(c.Runner.Grid) {
		c.Done = true
		return
	}

	c.Runner.Energized[fmt.Sprintf("%d-%d", c.Y, c.X)] = true

	char := c.Runner.Grid[c.Y][c.X]
	switch {
	case char == '/' && c.Dir == dirRight:
		c.Dir = dirUp
	case char == '/' && c.Dir == dirLeft:
		c.Dir = dirDown
	case char == '/' && c.Dir == dirDown:
		c.Dir = dirLeft
	case char == '/' && c.Dir == dirUp:
		c.Dir = dirRight

	case char == '\\' && c.Dir == dirRight:
		c.Dir = dirDown
	case char == '\\' && c.Dir == dirLeft:
		c.Dir = dirUp
	case char == '\\' && c.Dir == dirDown:
		c.Dir = dirRight
	case char == '\\' && c.Dir == dirUp:
		c.Dir = dirLeft

	case char == '|' && (c.Dir == dirLeft || c.Dir == dirRight):
		c.Done = true
		c.Runner.Add(dirUp, c.Y, c.X)
		c.Runner.Add(dirDown, c.Y, c.X)

	case char == '-' && (c.Dir == dirDown || c.Dir == dirUp):
		c.Done = true
		c.Runner.Add(dirRight, c.Y, c.X)
		c.Runner.Add(dirLeft, c.Y, c.X)
	}
}
