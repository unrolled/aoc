package day18

import (
	"bytes"
	"strconv"
	"strings"
)

func PartOne() int {
	// INPUT = "R 6 (#70c710)\nD 5 (#0dc571)\nL 2 (#5713f0)\nD 2 (#d2c081)\nR 2 (#59c680)\nD 2 (#411b91)\nL 5 (#8ceee2)\nU 2 (#caa173)\nL 1 (#1b58a2)\nU 2 (#caa171)\nR 2 (#7807d2)\nU 3 (#a77fa3)\nL 2 (#015232)\nU 2 (#7a21e3)" // 62
	digMap := parse(INPUT)
	m := Map{data: [][]byte{{'#'}}}

	for _, d := range digMap {
		m.ApplyDig(d)
	}

	// m.Print()
	return m.Volume()
}

func PartTwo() int {
	return -1
	// INPUT = "R 6 (#70c710)\nD 5 (#0dc571)\nL 2 (#5713f0)\nD 2 (#d2c081)\nR 2 (#59c680)\nD 2 (#411b91)\nL 5 (#8ceee2)\nU 2 (#caa173)\nL 1 (#1b58a2)\nU 2 (#caa171)\nR 2 (#7807d2)\nU 3 (#a77fa3)\nL 2 (#015232)\nU 2 (#7a21e3)" // 952408144115
	// digMap := parse(INPUT)
	// m := Map{data: [][]byte{{'#'}}}

	// for _, d := range digMap {
	// 	m.ApplyDig(d.Colorize())
	// }

	// // m.Print()
	// return m.Volume()
}

func parse(input string) []Dig {
	input = strings.ReplaceAll(input, "#", "")
	input = strings.ReplaceAll(input, "(", "")
	input = strings.ReplaceAll(input, ")", "")
	digs := []Dig{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		color := parts[2]
		direction := parts[0]
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		digs = append(digs, Dig{Direction: direction[0], Color: color, Value: value})
	}

	return digs
}

type Dig struct {
	Direction byte
	Value     int
	Color     string
}

func (d *Dig) Colorize() Dig {

	value, err := strconv.ParseInt(d.Color[0:len(d.Color)-1], 16, 64)
	if err != nil {
		panic(err)
	}

	// 0 means R, 1 means D, 2 means L, and 3 means U.
	var dir byte
	switch d.Color[len(d.Color)-1] {
	case '0':
		dir = 'R'
	case '1':
		dir = 'D'
	case '2':
		dir = 'L'
	case '3':
		dir = 'U'
	}

	return Dig{Value: int(value), Direction: dir}
}

type Map struct {
	data [][]byte
	Y, X int
}

func (m *Map) Volume() int {
	fillY, fillX := 0, 0
	for y, line := range m.data {
		tmp := []int{}
		for x, char := range line {
			if char == '#' {
				tmp = append(tmp, x)
			}
		}

		if len(tmp) == 2 && tmp[0]+1 != tmp[1] {
			fillY = y
			fillX = tmp[0] + 1
			break
		}
	}

	floodFill(m.data, fillY, fillX)

	total := 0

	for _, line := range m.data {
		for _, char := range line {
			if char == 'F' || char == '#' {
				total++
			}
		}
	}

	return total
}

func (m *Map) Print() {
	orig := m.data[m.Y][m.X]
	m.data[m.Y][m.X] = '$'

	for _, line := range m.data {
		println(string(line))
	}
	m.data[m.Y][m.X] = orig
}

func (m *Map) ApplyDig(d Dig) {
	y, x := 0, 0
	switch d.Direction {
	case 'U':
		y = d.Value * -1
	case 'D':
		y = d.Value
	case 'R':
		x = d.Value
	case 'L':
		x = d.Value * -1
	}

	m.do(y, x)
}

func (m *Map) do(y, x int) {
	currentY := len(m.data)
	currentX := len(m.data[0])

	if y < 0 && m.Y+y < 0 { // Moving north, add rows above
		newRows := [][]byte{}
		for i := 0; i < (m.Y+y)*-1; i++ {
			newRows = append(newRows, bytes.Repeat([]byte("."), currentX))
		}

		m.data = append(newRows, m.data...)
		m.Y += (m.Y + y) * -1
	}

	if y > 0 && m.Y+y >= len(m.data) { // Moving south, add rows below
		for i := len(m.data); i <= m.Y+y; i++ {
			m.data = append(m.data, bytes.Repeat([]byte("."), currentX))
		}
	}

	if x < 0 && m.X+x < 0 { // Moving west, add columns to left
		newChars := bytes.Repeat([]byte("."), (m.X+x)*-1)
		for i := 0; i < currentY; i++ {
			m.data[i] = append(newChars, m.data[i]...)
		}
		m.X += (m.X + x) * -1
	}

	if x > 0 && m.X+x >= len(m.data[0]) { // Moving east, add columns to right
		newChars := bytes.Repeat([]byte("."), x)
		for i := 0; i < currentY; i++ {
			m.data[i] = append(m.data[i], newChars...)
		}
	}

	for x != 0 {
		if x > 0 {
			m.X++
			x--
		} else {
			m.X--
			x++
		}
		m.data[m.Y][m.X] = '#'
	}

	for y != 0 {
		if y > 0 {
			m.Y++
			y--
		} else {
			m.Y--
			y++
		}
		m.data[m.Y][m.X] = '#'
	}
}

func floodFill(data [][]byte, y, x int) [][]byte {
	if y < 0 || y >= len(data) || x < 0 || x >= len(data[0]) {
		return data
	}

	if data[y][x] == 'F' {
		return data
	}

	if data[y][x] != '.' {
		return data
	}

	data[y][x] = 'F'

	data = floodFill(data, y+1, x)
	data = floodFill(data, y-1, x)
	data = floodFill(data, y, x+1)
	data = floodFill(data, y, x-1)

	return data
}
