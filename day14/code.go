package day14

import (
	"strings"
)

const (
	roundRock  = 'O'
	emptySpace = '.'
)

func PartOne() int {
	// INPUT = "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#...." // 136
	data := parse(INPUT)
	data.TiltNorth()
	return data.TotalLoad()
}

func PartTwo() int {
	// INPUT = "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#...." // 64
	data := parse(INPUT)
	data.Spin(1_000_000_000)
	return data.TotalLoad()
}

type Data struct {
	raw         [][]byte
	spinHistory map[string]int
}

func (d *Data) String() string {
	var s string
	for _, line := range d.raw {
		s += string(line)
	}
	return s
}

func (d *Data) Spin(times int) {
	for i := 0; i < times; i++ {
		key := d.String()
		if val, exists := d.spinHistory[key]; exists {
			step := i - val - 1
			if i+step < times {
				i += step
				continue
			}
		}

		d.TiltNorth()
		d.TiltWest()
		d.TiltSouth()
		d.TiltEast()

		d.spinHistory[key] = i
	}
}

func (d *Data) TiltNorth() {
	d.tilt()
}

func (d *Data) TiltSouth() {
	d.invert()
	d.tilt()
	d.invert()
}

func (d *Data) TiltWest() {
	d.transpose()
	d.tilt()
	d.transpose()
}

func (d *Data) TiltEast() {
	d.transpose()
	d.invert()
	d.tilt()
	d.invert()
	d.transpose()
}

func (d *Data) transpose() {
	rowCount := len(d.raw)
	charCount := len(d.raw[0])
	newData := make([][]byte, charCount)
	for i := 0; i < charCount; i++ {
		newData[i] = make([]byte, rowCount)
	}

	for i := 0; i < rowCount; i++ {
		for j := 0; j < charCount; j++ {
			newData[j][i] = d.raw[i][j]
		}
	}
	d.raw = newData
}

func (d *Data) invert() {
	newData := [][]byte{}
	for i := len(d.raw) - 1; i >= 0; i-- {
		newData = append(newData, d.raw[i])
	}
	d.raw = newData
}

func (d *Data) tilt() {
	var changed bool

	for i := len(d.raw) - 1; i > 0; i-- {
		for j := 0; j < len(d.raw[i]); j++ {
			if d.raw[i][j] == roundRock && d.raw[i-1][j] == emptySpace {
				d.raw[i][j], d.raw[i-1][j] = d.raw[i-1][j], d.raw[i][j]
				changed = true
			}
		}
	}

	if changed {
		d.tilt()
	}
}

func (d *Data) TotalLoad() int {
	var total int

	for i := 0; i < len(d.raw); i++ {
		points := len(d.raw) - i
		for _, char := range d.raw[i] {
			if char == roundRock {
				total += points
			}
		}
	}

	return total
}

func parse(input string) Data {
	result := [][]byte{}
	for _, line := range strings.Split(input, "\n") {
		result = append(result, []byte(line))
	}
	return Data{raw: result, spinHistory: map[string]int{}}
}
