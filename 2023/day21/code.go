package day21

import (
	"fmt"
	"strings"
)

func PartOne() int {
	// INPUT = "...........\n.....###.#.\n.###.##..#.\n..#.#...#..\n....#.#....\n.##..S####.\n.##..#...#.\n.......##..\n.##.#.####.\n.##..##.##.\n..........." // 16@6
	m := parseMap(INPUT)

	return m.Reach(64)
}

func PartTwo() int {
	return -1
}

func parseMap(input string) *Map {
	m := &Map{data: [][]byte{}, Queue: []*Coord{}, History: map[string]bool{}}

	for y, line := range strings.Split(input, "\n") {
		m.data = append(m.data, []byte(line))

		if x := strings.Index(line, "S"); x != -1 {
			m.Start = &Coord{Y: y, X: x}
		}
	}

	return m
}

type Coord struct {
	Y, X int
}

func (c *Coord) Key() string {
	return fmt.Sprintf("%d|%d", c.Y, c.X)
}

type Map struct {
	data    [][]byte
	Start   *Coord
	Queue   []*Coord
	History map[string]bool
}

var superHistory = map[string]int{}

func (m *Map) Reach(steps int) int {
	m.Queue = append(m.Queue, m.Start)
	m.Step()

	for s := 0; s < steps; s++ {
		m.History = map[string]bool{}
		m.Step()
	}

	return len(m.History)
}

var stepCache = map[string][]*Coord{}

func (m *Map) Step() {
	steps := m.Queue
	m.Queue = []*Coord{}

	for _, step := range steps {
		m.Queue = append(m.Queue, m.findNextSteps(step)...)
	}
}

func (m *Map) findNextSteps(c *Coord) []*Coord {
	result := []*Coord{}

	if !m.canStep(c) {
		return result
	}

	key := c.Key()
	if _, exists := m.History[key]; exists {
		return result
	}

	m.History[key] = true

	result = append(result, &Coord{Y: c.Y + 1, X: c.X})
	result = append(result, &Coord{Y: c.Y - 1, X: c.X})
	result = append(result, &Coord{Y: c.Y, X: c.X + 1})
	result = append(result, &Coord{Y: c.Y, X: c.X - 1})

	return result
}

func (m *Map) canStep(c *Coord) bool {
	if c.Y < 0 || c.Y >= len(m.data) || c.X < 0 || c.X >= len(m.data[0]) {
		return false
	}

	if m.data[c.Y][c.X] == '#' {
		return false
	}

	return true
}
