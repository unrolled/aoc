package day04

import (
	"strconv"
	"strings"
)

type Card struct {
	ID   int
	Win  []int
	Play []int

	Copies, Points, WinCount int
}

func PartOne() int {
	// 	INPUT = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	cards := []*Card{}

	var result int
	for _, line := range strings.Split(INPUT, "\n") {
		card := parseCard(line)
		cards = append(cards, card)
		result += card.Points
	}

	return result
}

func PartTwo() int {
	// 	INPUT = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	// Setup map of cards.
	cards := map[int]*Card{}
	for _, line := range strings.Split(INPUT, "\n") {
		card := parseCard(line)
		cards[card.ID] = card
	}

	// Determine card count.
	for i := 1; i <= len(cards); i++ {
		card := cards[i]

		for i := 1; i <= card.WinCount; i++ {
			cards[card.ID+i].Copies += +card.Copies
		}
	}

	// Add up card copies.
	var result int
	for _, card := range cards {
		result += card.Copies
	}

	return result
}

func parseCard(line string) *Card {
	result := &Card{Copies: 1}

	// Get ID
	line = strings.TrimPrefix(line, "Card")
	line = strings.TrimSpace(line)
	idParts := strings.Split(line, ":")
	result.ID = convertToNum(idParts[0])
	line = strings.TrimSpace(idParts[1])

	// Split sequences
	numParts := strings.Split(line, "|")
	result.Win = parseNumList(strings.TrimSpace(numParts[0]))
	result.Play = parseNumList(strings.TrimSpace(numParts[1]))

	// Calculate points
	for _, winNum := range result.Win {
		for _, playNum := range result.Play {
			if winNum == playNum {
				if result.Points == 0 {
					result.Points = 1
				} else {
					result.Points = result.Points * 2
				}
			}
		}
	}

	// Calculate win count
	for _, winNum := range result.Win {
		for _, playNum := range result.Play {
			if winNum == playNum {
				result.WinCount++
			}
		}
	}

	return result
}

func parseNumList(input string) []int {
	result := []int{}

	for _, num := range strings.Split(input, " ") {
		num = strings.TrimSpace(num)
		if num == "" {
			continue
		}

		result = append(result, convertToNum(num))
	}

	return result
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
