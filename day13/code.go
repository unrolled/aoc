package day13

import (
	"strings"
)

func PartOne() int {
	// INPUT = "#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.\n\n#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#" // 405
	total := 0

	for _, data := range strings.Split(INPUT, "\n\n") {
		p := Pattern{raw: data}
		total += p.Reflections()
	}

	return total
}

func PartTwo() int {
	// INPUT = "#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.\n\n#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#" // 400
	total := 0

	for _, data := range strings.Split(INPUT, "\n\n") {
		p := Pattern{raw: data}
		total += p.Smudges()
	}

	return total
}

type Pattern struct {
	raw string
}

func (p *Pattern) horizontalData() []string {
	return strings.Split(p.raw, "\n")
}

func (p *Pattern) verticalData() []string {
	lines := p.horizontalData()
	vertical := make([]string, len(lines[0]))

	for _, line := range lines {
		for x, val := range line {
			vertical[x] = vertical[x] + string(val)
		}
	}

	return vertical
}

func (p *Pattern) Reflections() int {
	h := p.reflection(p.horizontalData(), -1)
	v := p.reflection(p.verticalData(), -1)

	return v + (h * 100)
}

func (p *Pattern) Smudges() int {
	ogH := p.reflection(p.horizontalData(), -1)
	ogV := p.reflection(p.verticalData(), -1)

	raw := []byte(p.raw)
	for i, char := range raw {
		switch char {
		case '\n':
			continue
		case '.':
			raw[i] = '#'
		case '#':
			raw[i] = '.'
		}

		p.raw = string(raw)

		if newH := p.reflection(p.horizontalData(), ogH); newH > 0 {
			return newH * 100
		}
		if newV := p.reflection(p.verticalData(), ogV); newV > 0 {
			return newV
		}

		raw[i] = char
	}

	return 0
}

func (p *Pattern) reflection(lines []string, old int) int {
	for i := 1; i < len(lines); i++ {
		previous := lines[i-1]
		current := lines[i]

		count := 0
		for previous == current {
			count++

			if i+count >= len(lines) || i-1-count < 0 {
				if i != old {
					return i
				}

				break
			}

			current = current + lines[i+count]
			previous = previous + lines[i-1-count]
		}
	}

	return 0
}
