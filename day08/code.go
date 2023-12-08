package day08

import (
	"fmt"
	"strings"
)

var (
	nodes map[string]*Node
	steps []byte
)

type Node struct {
	ID, Left, Right string
}

func PartOne() int {
	// INPUT = "RL\n\nAAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)\nDDD = (DDD, DDD)\nEEE = (EEE, EEE)\nGGG = (GGG, GGG)\nZZZ = (ZZZ, ZZZ)"
	parts := strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(INPUT, "= (", ""), ",", ""), ")", ""), "\n\n")
	steps = []byte(parts[0])
	nodes = map[string]*Node{}

	for _, input := range strings.Split(parts[1], "\n") {
		node := parseNode(input)
		nodes[node.ID] = node
	}

	return traverseNodes("AAA", "ZZZ")
}

func PartTwo() int {
	// INPUT = "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)"
	parts := strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(INPUT, "= (", ""), ",", ""), ")", ""), "\n\n")
	steps = []byte(parts[0])
	nodes = map[string]*Node{}

	paths := []string{}

	for _, input := range strings.Split(parts[1], "\n") {
		node := parseNode(input)
		nodes[node.ID] = node

		if strings.HasSuffix(node.ID, "A") {
			paths = append(paths, node.ID)
		}
	}

	stepCounts := []int{}
	for _, path := range paths {
		stepCounts = append(stepCounts, traverseNodes(path, "Z"))
	}

	return leastCommonMultiple(stepCounts)
}

func parseNode(input string) *Node {
	var i, l, r string
	fmt.Sscanf(input, "%s %s %s", &i, &l, &r)

	return &Node{ID: i, Left: l, Right: r}
}

func traverseNodes(currentNodeID string, endSuffix string) int {
	var directionIndex, stepCount int

	for {
		stepCount++

		direction := steps[directionIndex]
		directionIndex++
		if directionIndex == len(steps) {
			directionIndex = 0
		}

		if direction == 'R' {
			currentNodeID = nodes[currentNodeID].Right
		} else {
			currentNodeID = nodes[currentNodeID].Left
		}

		if strings.HasSuffix(nodes[currentNodeID].ID, endSuffix) {
			break
		}
	}

	return stepCount
}

func leastCommonMultiple(input []int) int {
	if len(input) == 1 {
		return input[0]
	}

	firstNum := input[0]
	secondNum := input[1]

	for secondNum != 0 {
		secondNum, firstNum = firstNum%secondNum, secondNum
	}

	multiple := []int{input[0] * input[1] / firstNum}
	newInput := append(multiple, input[2:]...)

	return leastCommonMultiple(newInput)
}
