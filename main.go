package main

import (
	"fmt"
	"os"
	"time"

	"github.com/unrolled/adventofcode/day01"
	"github.com/unrolled/adventofcode/day02"
	"github.com/unrolled/adventofcode/day03"
	"github.com/unrolled/adventofcode/day04"
	"github.com/unrolled/adventofcode/day05"
	"github.com/unrolled/adventofcode/day06"
	"github.com/unrolled/adventofcode/day07"
	"github.com/unrolled/adventofcode/day08"
)

type Advent struct {
	Title   string
	PartOne func() int
	PartTwo func() int
}

var adventDays = []Advent{
	{Title: "Day 1: Trebuchet?!", PartOne: day01.PartOne, PartTwo: day01.PartTwo},
	{Title: "Day 2: Cube Conundrum", PartOne: day02.PartOne, PartTwo: day02.PartTwo},
	{Title: "Day 3: Gear Ratios", PartOne: day03.PartOne, PartTwo: day03.PartTwo},
	{Title: "Day 4: Scratchcards", PartOne: day04.PartOne, PartTwo: day04.PartTwo},
	{Title: "Day 5: If You Give A Seed A Fertilizer", PartOne: day05.PartOne, PartTwo: day05.PartTwo},
	{Title: "Day 6: Wait For It", PartOne: day06.PartOne, PartTwo: day06.PartTwo},
	{Title: "Day 7: Camel Cards", PartOne: day07.PartOne, PartTwo: day07.PartTwo},
	{Title: "Day 8: Haunted Wasteland", PartOne: day08.PartOne, PartTwo: day08.PartTwo},
}

func main() {
	var inputDay string
	if len(os.Args) == 2 {
		inputDay = os.Args[1]
	}

	fmt.Printf("\n%s %s %s\n", termCyan("--- Advent of Code:"), termPurple("2023"), termCyan("---"))

	var totalNanoSeconds int64
	for i, adventDay := range adventDays {
		if len(inputDay) == 0 || inputDay == fmt.Sprintf("%d", i+1) {
			totalNanoSeconds += runWithTimer(adventDay)
		}
	}

	if len(inputDay) == 0 {
		fmt.Printf("\n%s %s %s\n\n", termCyan("--- Total run time:"), termPurple("%s", time.Duration(totalNanoSeconds).String()), termCyan("---"))
	}
}

func runWithTimer(adventDay Advent) int64 {
	fmt.Printf("\n%s\n", adventDay.Title)

	partOneStartTime := time.Now()
	partOneResult := adventDay.PartOne()
	partOneDuration := time.Now().Sub(partOneStartTime)
	fmt.Printf("  %s %s %s\n", termBlue("Part 1:"), termGreen("%d", partOneResult), termYellow("(%v)", partOneDuration))

	partTwoStartTime := time.Now()
	partTwoResult := adventDay.PartTwo()
	partTwoDuration := time.Now().Sub(partTwoStartTime)
	fmt.Printf("  %s %s %s\n", termBlue("Part 2:"), termGreen("%d", partTwoResult), termYellow("(%v)", partTwoDuration))

	return partOneDuration.Nanoseconds() + partTwoDuration.Nanoseconds()
}

func termCyan(format string, a ...any) string {
	return "\033[36m" + fmt.Sprintf(format, a...) + "\033[0m"
}
func termGreen(format string, a ...any) string {
	return "\033[32m" + fmt.Sprintf(format, a...) + "\033[0m"
}
func termYellow(format string, a ...any) string {
	return "\033[33m" + fmt.Sprintf(format, a...) + "\033[0m"
}
func termBlue(format string, a ...any) string {
	return "\033[34m" + fmt.Sprintf(format, a...) + "\033[0m"
}
func termPurple(format string, a ...any) string {
	return "\033[35m" + fmt.Sprintf(format, a...) + "\033[0m"
}
