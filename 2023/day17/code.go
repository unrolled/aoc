package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const maxDirMove = 3

func PartOne() int {
	// INPUT = "2413432311323\n3215453535623\n3255245654254\n3446585845452\n4546657867536\n1438598798454\n4457876987766\n3637877979653\n4654967986887\n4564679986453\n1224686865563\n2546548887735\n4322674655533" // 102
	// m := newMap(INPUT)
	// m.Queue = append(m.Queue, &Path{Map: m, Position: Coord{0, 0}, History: []Coord{}})
	// m.Process()
	// return m.MinCost()

	return -1
}

func PartTwo() int {
	return -1
}

func newMap(input string) *Map {
	grid := [][]int{}

	for _, line := range strings.Split(input, "\n") {
		lineVal := []int{}
		for _, char := range line {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			lineVal = append(lineVal, n)
		}

		grid = append(grid, lineVal)
	}

	return &Map{Grid: grid, Queue: []*Path{}, Completed: []*Path{}}
}

type Map struct {
	Grid      [][]int
	Queue     []*Path
	Completed []*Path
}

func (m *Map) Process() {
	for len(m.Queue) != 0 {
		path := m.Queue[0]
		m.Queue = m.Queue[1:]
		m.Queue = append(m.Queue, path.Step()...)
	}
}

func (m *Map) MinCost() int {
	minCost := math.MaxInt
	for _, path := range m.Completed {
		if path.Cost < minCost {
			minCost = path.Cost
		}
	}
	return minCost
}

type Coord struct {
	Y, X int
}

func (c *Coord) Match(t Coord) bool {
	return c.X == t.X && c.Y == t.Y
}

func (c *Coord) String() string {
	return fmt.Sprintf("%d-%d", c.Y, c.X)
}

type Path struct {
	Map      *Map
	Position Coord
	History  []Coord
	Cost     int
}

func (p *Path) Description() string {
	return fmt.Sprintf("Y:%d|X:%d [%d] %v", p.Position.Y, p.Position.X, p.Cost, p.History)
}

func (p *Path) isAtEnd() bool {
	return p.Position.Y == len(p.Map.Grid)-1 && p.Position.X == len(p.Map.Grid[0])-1
}

func (p *Path) Step() []*Path {
	nextSteps := []*Path{}

	if p.isAtEnd() {
		p.Map.Completed = append(p.Map.Completed, p)
		return nextSteps
	}

	if next := p.move(-1, 0); next != nil {
		nextSteps = append(nextSteps, next)
	}
	if next := p.move(0, 1); next != nil {
		nextSteps = append(nextSteps, next)
	}
	if next := p.move(1, 0); next != nil {
		nextSteps = append(nextSteps, next)
	}
	if next := p.move(0, -1); next != nil {
		nextSteps = append(nextSteps, next)
	}

	return nextSteps
}

func (p *Path) canMoveDirection(y, x int) bool {
	next := Coord{Y: p.Position.Y + y, X: p.Position.X + x}

	// Don't go back to a space we've already been at.
	for _, historyCoord := range p.History {
		if next.Match(historyCoord) {
			return false
		}
	}

	// Make sure the new position is in bounds.
	if next.X == -1 || next.X == len(p.Map.Grid[0]) || next.Y == -1 || next.Y == len(p.Map.Grid) {
		return false
	}

	// Yes if no history.
	if len(p.History) == 0 {
		return true
	}

	// No if matches last position.
	if p.History[len(p.History)-1].Match(next) {
		return false
	}

	change := []Coord{}
	tmpHistory := append(p.History, p.Position, next)
	for i := len(tmpHistory) - 1; i != 0; i-- {
		h0 := tmpHistory[i]
		h1 := tmpHistory[i-1]

		innerX := h0.X - h1.X
		innerY := h0.Y - h1.Y
		change = append(change, Coord{X: innerX, Y: innerY})

		if len(change) == maxDirMove+1 {
			break
		}
	}

	if len(change) <= 2 {
		return true
	}

	firstChange := change[0]
	for _, c := range change[1:] {
		if !firstChange.Match(c) {
			return true
		}
	}

	return false
}

func (p *Path) move(y, x int) *Path {
	if !p.canMoveDirection(y, x) {
		return nil
	}

	nextCoord := Coord{Y: p.Position.Y + y, X: p.Position.X + x}
	newPath := &Path{Map: p.Map, Position: nextCoord}

	hCopy := []Coord{}
	for _, c := range p.History {
		hCopy = append(hCopy, c)
	}
	newPath.History = append(hCopy, p.Position)
	newPath.Cost = p.Cost + p.Map.Grid[nextCoord.Y][nextCoord.X]

	return newPath
}
