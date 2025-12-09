package day08

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type JunctionBox struct {
	ID      int
	Circuit int
	X, Y, Z int
}

func (j *JunctionBox) Diff(n *JunctionBox) *Distance {

	x := math.Pow(float64(n.X-j.X), 2)
	y := math.Pow(float64(n.Y-j.Y), 2)
	z := math.Pow(float64(n.Z-j.Z), 2)
	diff := math.Sqrt((x + y + z))

	return &Distance{ID1: j.ID, ID2: n.ID, Value: diff}
}

func (j *JunctionBox) String() string {
	return fmt.Sprintf("%d,%d,%d", j.X, j.Y, j.Z)
}

type Distance struct {
	Value    float64
	ID1, ID2 int
}

type SortableDistances []*Distance

func (s SortableDistances) Len() int {
	return len(s)
}

func (s SortableDistances) Less(i, j int) bool {
	return s[i].Value < s[j].Value
}

func (s SortableDistances) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type SortableCircuits [][]int

func (s SortableCircuits) Len() int {
	return len(s)
}

func (s SortableCircuits) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func (s SortableCircuits) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func PartOne() int {
	// INPUT = "162,817,812\n57,618,57\n906,360,560\n592,479,940\n352,342,300\n466,668,158\n542,29,236\n431,825,988\n739,650,466\n52,470,668\n216,146,977\n819,987,18\n117,168,530\n805,96,715\n346,949,466\n970,615,88\n941,993,340\n862,61,35\n984,92,344\n425,690,689"
	junctionBoxes := []*JunctionBox{}
	distances := []*Distance{}

	for i, line := range strings.Split(INPUT, "\n") {
		parts := strings.Split(line, ",")
		jb := &JunctionBox{ID: i, X: convertToNum(parts[0]), Y: convertToNum(parts[1]), Z: convertToNum(parts[2]), Circuit: -1}

		for _, existingBox := range junctionBoxes {
			distances = append(distances, existingBox.Diff(jb))
		}

		junctionBoxes = append(junctionBoxes, jb)
	}

	sortedDistances := SortableDistances(distances)
	sort.Sort(sortedDistances)

	circuits := [][]int{}
	for _, obj := range sortedDistances[0:1000] {
		jb1 := junctionBoxes[obj.ID1]
		jb2 := junctionBoxes[obj.ID2]

		// If two exists, but one does not... simple.
		if jb1.Circuit == -1 && jb2.Circuit != -1 {
			jb1.Circuit = jb2.Circuit
			circuits[jb2.Circuit] = append(circuits[jb2.Circuit], jb1.ID)
			continue
		}

		// If one exists, but two does not (simple).
		if jb2.Circuit == -1 && jb1.Circuit != -1 {
			jb2.Circuit = jb1.Circuit
			circuits[jb1.Circuit] = append(circuits[jb1.Circuit], jb2.ID)
			continue
		}

		// Both do not have a circuit (create new circuit).
		if jb2.Circuit == -1 && jb1.Circuit == -1 {
			jb1.Circuit = len(circuits)
			jb2.Circuit = len(circuits)
			newCircuit := []int{jb1.ID, jb2.ID}
			circuits = append(circuits, newCircuit)
			continue
		}

		if jb1.Circuit == jb2.Circuit {
			continue
		}

		// Both do have a circuit (join them).
		if jb2.Circuit != -1 && jb1.Circuit != -1 {
			circuits[jb1.Circuit] = append(circuits[jb1.Circuit], circuits[jb2.Circuit]...)
			circuits[jb2.Circuit] = []int{}
			jb2.Circuit = jb1.Circuit
			continue
		}

		panic("please don't make it here...")
	}

	sortedCircuits := SortableCircuits(circuits)
	sort.Sort(sortedCircuits)

	// fmt.Printf("%#v\n", sortedCircuits)

	total := 1
	for i := range 3 {
		total = total * len(sortedCircuits[i])
	}

	return total
}

func PartTwo() int {
	return 0
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
