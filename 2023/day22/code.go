package day22

import (
	"fmt"
	"sort"
	"strings"
)

func PartOne() int {
	// INPUT = "1,0,1~1,2,1\n0,0,2~2,0,2\n0,2,3~2,2,3\n0,0,4~0,2,4\n2,0,5~2,2,5\n0,1,6~2,1,6\n1,1,8~1,1,9" // 5
	stack := newStack(INPUT)
	stack.Drop()

	return stack.Disintegratable()
}

func PartTwo() int {
	return -1
}

func newStack(input string) *Stack {
	brickMap := map[int]*Brick{}
	bricks := []*Brick{}

	var mX, mY, mZ int

	for i, line := range strings.Split(input, "\n") {
		b := &Brick{ID: i + 1, Supporting: map[int]bool{}, SupportedBy: map[int]bool{}}
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &b.x1, &b.y1, &b.z1, &b.x2, &b.y2, &b.z2)
		bricks = append(bricks, b)

		switch {
		case b.x1 > b.x2 && b.x1 > mX:
			mX = b.x1
		case b.x2 > mX:
			mX = b.x2
		case b.y1 > b.y2 && b.y1 > mY:
			mY = b.y1
		case b.y2 > mY:
			mY = b.y2
		case b.z1 > b.z2 && b.z1 > mZ:
			mZ = b.z1
		case b.z2 > mZ:
			mZ = b.z2
		}
	}

	mX++
	mY++
	mZ++

	sort.Slice(bricks, func(i, j int) bool {
		first := bricks[i].z1
		if first > bricks[i].z2 {
			first = bricks[i].z2
		}

		second := bricks[j].z1
		if second > bricks[j].z2 {
			second = bricks[j].z2
		}

		return first < second
	})

	xAxis := make([][]int, mZ)
	for i := 0; i < mZ; i++ {
		xAxis[i] = make([]int, mX)
	}

	yAxis := make([][]int, mZ)
	for i := 0; i < mZ; i++ {
		yAxis[i] = make([]int, mY)
	}

	for i, brick := range bricks {
		brickMap[i+1] = brick
		for z := brick.z1; z <= brick.z2; z++ {
			for x := brick.x1; x <= brick.x2; x++ {
				xAxis[z][x] = brick.ID
			}
			for y := brick.y1; y <= brick.y2; y++ {
				yAxis[z][y] = brick.ID
			}
		}
	}

	return &Stack{bricks: brickMap, xAxis: xAxis, yAxis: yAxis, movedBricks: map[int]bool{}}
}

type Brick struct {
	ID, x1, y1, z1, x2, y2, z2 int
	Supporting, SupportedBy    map[int]bool
}

func (b *Brick) LowestZ() int {
	if b.z1 < b.z2 {
		return b.z1
	}

	return b.z2
}

func (b *Brick) HighestZ() int {
	if b.z1 > b.z2 {
		return b.z1
	}

	return b.z2
}

type Stack struct {
	bricks       map[int]*Brick
	xAxis, yAxis [][]int
	movedBricks  map[int]bool
}

func (s *Stack) PrintLetters() {
	var xPart, yPart string
	for z := len(s.xAxis) - 1; z > 0; z-- {
		xPart, yPart = "", ""
		for x := 0; x < len(s.xAxis[0]); x++ {
			if s.xAxis[z][x] == 0 {
				xPart += "."
			} else {
				xPart += string(byte(64 + s.xAxis[z][x]))
			}

			if s.yAxis[z][x] == 0 {
				yPart += "."
			} else {
				yPart += string(byte(64 + s.yAxis[z][x]))
			}
		}
		fmt.Printf("%s    %s\n", xPart, yPart)
	}
	fmt.Printf("%s____%s\n", strings.Repeat("-", len(xPart)), strings.Repeat("-", len(yPart)))
}

func (s *Stack) Print() {
	var xPart, yPart string
	for z := len(s.xAxis) - 1; z > 0; z-- {
		xPart, yPart = "", ""
		for x := 0; x < len(s.xAxis[0]); x++ {
			if s.xAxis[z][x] == 0 {
				xPart += "."
			} else {
				xPart += "#"
			}

			if s.yAxis[z][x] == 0 {
				yPart += "."
			} else {
				yPart += "#"
			}
		}
		fmt.Printf("%s    %s\n", xPart, yPart)
	}
	fmt.Printf("%s____%s\n", strings.Repeat("-", len(xPart)), strings.Repeat("-", len(yPart)))
}

func (s *Stack) Drop() {
	for i := 1; i <= len(s.bricks); i++ {
		brick := s.bricks[i]
		for !s.movedBricks[brick.ID] {
			lowestZ := brick.z1
			if lowestZ > brick.z2 {
				lowestZ = brick.z2
			}

			nextZ := lowestZ - 1

			if nextZ == 0 {
				s.movedBricks[brick.ID] = true
				continue
			}

			var xBlocked, yBlocked bool

			for x := brick.x1; x <= brick.x2; x++ {
				if s.xAxis[nextZ][x] != 0 {
					xBlocked = true
				}
			}
			for y := brick.y1; y <= brick.y2; y++ {
				if s.yAxis[nextZ][y] != 0 {
					yBlocked = true
				}
			}

			if xBlocked && yBlocked {
				s.movedBricks[brick.ID] = true
				continue
			}

			for z := brick.z1; z <= brick.z2; z++ {
				for x := brick.x1; x <= brick.x2; x++ {
					s.xAxis[z][x] = 0
				}
				for y := brick.y1; y <= brick.y2; y++ {
					s.yAxis[z][y] = 0
				}
			}

			brick.z1--
			brick.z2--

			for z := brick.z1; z <= brick.z2; z++ {
				for x := brick.x1; x <= brick.x2; x++ {
					s.xAxis[z][x] = brick.ID
				}
				for y := brick.y1; y <= brick.y2; y++ {
					s.yAxis[z][y] = brick.ID
				}
			}
		}
	}
}

func (s *Stack) Disintegratable() int {
	for i := 1; i <= len(s.bricks); i++ {

		currentBrick := s.bricks[i]
		for j := 1; j <= len(s.bricks); j++ {
			nextBrick := s.bricks[j]
			if currentBrick.ID == nextBrick.ID {
				continue
			}

			if currentBrick.HighestZ()+1 == nextBrick.LowestZ() {
				if nextBrick.x1 <= currentBrick.x1 && nextBrick.x2 >= currentBrick.x1 {
					currentBrick.Supporting[nextBrick.ID] = true
					nextBrick.SupportedBy[currentBrick.ID] = true
				}
				if nextBrick.x1 <= currentBrick.x2 && nextBrick.x2 >= currentBrick.x2 {
					currentBrick.Supporting[nextBrick.ID] = true
					nextBrick.SupportedBy[currentBrick.ID] = true
				}
				if nextBrick.y1 <= currentBrick.y1 && nextBrick.y2 >= currentBrick.y1 {
					currentBrick.Supporting[nextBrick.ID] = true
					nextBrick.SupportedBy[currentBrick.ID] = true
				}
				if nextBrick.y1 <= currentBrick.y2 && nextBrick.y2 >= currentBrick.y2 {
					currentBrick.Supporting[nextBrick.ID] = true
					nextBrick.SupportedBy[currentBrick.ID] = true
				}
			}
		}
	}

	var count int

	for j := 1; j <= len(s.bricks); j++ {
		brick := s.bricks[j]
		canDisintegrate := true
		for k := range brick.Supporting {
			if len(s.bricks[k].SupportedBy) == 1 {
				canDisintegrate = false
			}
		}

		if canDisintegrate {
			count++
		}
	}

	return count
}
