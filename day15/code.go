package day15

import (
	"strconv"
	"strings"
)

func PartOne() int {
	// INPUT = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7" // 1320
	total := 0

	for _, in := range strings.Split(INPUT, ",") {
		total = total + hash(in)
	}

	return total
}

func PartTwo() int {
	// INPUT = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7" // 145
	total := 0

	for i, b := range parseBoxes(INPUT) {
		total += b.Power(i)
	}

	return total
}

type Lens struct {
	Label       string
	FocalLength int
}

type Box struct {
	Lenses []*Lens
}

func (b *Box) Add(newLens *Lens) {
	tmpLens := []*Lens{}
	replaced := false

	for _, l := range b.Lenses {
		if l.Label == newLens.Label {
			tmpLens = append(tmpLens, newLens)
			replaced = true
		} else {
			tmpLens = append(tmpLens, l)
		}
	}

	if !replaced {
		tmpLens = append(tmpLens, newLens)
	}

	b.Lenses = tmpLens
}

func (b *Box) Remove(oldLens *Lens) {
	tmpLens := []*Lens{}

	for _, l := range b.Lenses {
		if l.Label != oldLens.Label {
			tmpLens = append(tmpLens, l)
		}
	}

	b.Lenses = tmpLens
}

func (b *Box) Power(index int) int {
	total := 0

	for i, l := range b.Lenses {
		total += ((index + 1) * (i + 1) * l.FocalLength)
	}

	return total
}

func hash(data string) int {
	var v int

	for _, char := range data {
		v = v + int(char)
		v = v * 17
		v = v % 256
	}

	return v
}

func parseBoxes(input string) []*Box {
	boxes := []*Box{}
	for i := 0; i < 256; i++ {
		boxes = append(boxes, &Box{Lenses: []*Lens{}})
	}

	for _, in := range strings.Split(INPUT, ",") {
		var label, operator string
		var focalLength int
		for _, char := range in {
			switch {
			case char == '=' || char == '-':
				operator = string(char)
			case operator != "":
				focalLength = convertToNum(string(char))
			default:
				label += string(char)
			}
		}

		index := hash(label)
		lens := &Lens{Label: label, FocalLength: focalLength}

		if operator == "=" {
			boxes[index].Add(lens)
		} else {
			boxes[index].Remove(lens)
		}
	}

	return boxes
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
