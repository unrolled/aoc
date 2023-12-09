package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var cardStrength = map[string]string{"A": "14", "K": "13", "Q": "12", "J": "11", "T": "10", "9": "09", "8": "08", "7": "07", "6": "06", "5": "05", "4": "04", "3": "03", "2": "02", "Z": "01"}

const (
	onePairVal = iota + 1
	twoPairVal
	threeKindVal
	fullHouseVal
	fourKindVal
	fiveKindVal

	Joker = "Z"
)

type Hand struct {
	Cards    string
	Bid      int
	Strength int
}

func PartOne() int {
	// 	INPUT = `32T3K 765
	// T55J5 684
	// KK677 28
	// KTJJT 220
	// QQQJA 483`
	lines := strings.Split(INPUT, "\n")

	hands := []*Hand{}
	for _, line := range lines {
		hands = append(hands, newHand(line))
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Strength < hands[j].Strength
	})

	var result int
	for i, hand := range hands {
		rank := i + 1
		result += (rank * hand.Bid)
	}

	return result
}

func PartTwo() int {
	// 	INPUT = `32T3K 765
	// T55J5 684
	// KK677 28
	// KTJJT 220
	// QQQJA 483`
	INPUT = strings.ReplaceAll(INPUT, "J", Joker)

	return PartOne()
}

func newHand(line string) *Hand {
	cards, bid := "", 0
	fmt.Sscanf(line, "%s %d", &cards, &bid)

	return &Hand{
		Cards:    cards,
		Bid:      bid,
		Strength: determineStrength(cards),
	}
}

func actualCardStrength(cards string) string {
	var numString string

	for _, card := range cards {
		numString += cardStrength[string(card)]
	}

	return numString
}

func makeCardMap(cards string) map[string]int {
	mapper := map[string]int{}
	for _, runeCard := range cards {
		card := string(runeCard)
		if _, exist := mapper[card]; !exist {
			mapper[card] = 0
		}

		mapper[card] = mapper[card] + 1
	}

	return mapper
}

func determineStrength(cards string) int {
	result := "0"

	switch {
	case isFiveKind(makeCardMap(cards)):
		result = strconv.Itoa(fiveKindVal)
	case isFourKind(makeCardMap(cards)):
		result = strconv.Itoa(fourKindVal)
	case isFullHouse(makeCardMap(cards)):
		result = strconv.Itoa(fullHouseVal)
	case isThreeKind(makeCardMap(cards)):
		result = strconv.Itoa(threeKindVal)
	case isTwoPair(makeCardMap(cards)):
		result = strconv.Itoa(twoPairVal)
	case isOnePair(makeCardMap(cards)):
		result = strconv.Itoa(onePairVal)
	}

	result += actualCardStrength(cards)

	return convertToNum(result)
}

func isFiveKind(mapper map[string]int) bool {
	delete(mapper, Joker)

	if len(mapper) <= 1 {
		return true
	}

	return false
}

func isFourKind(mapper map[string]int) bool {
	jokerVal, _ := mapper[Joker]
	delete(mapper, Joker)

	for _, count := range mapper {
		if count+jokerVal == 4 {
			return true
		}
	}

	return false
}

func isFullHouse(mapper map[string]int) bool {
	jokerCount, _ := mapper[Joker]
	delete(mapper, Joker)

	if len(mapper) == 2 || jokerCount >= 2 && len(mapper) == 1 {
		return true
	}

	return false
}

func isThreeKind(mapper map[string]int) bool {
	jokerVal, _ := mapper[Joker]
	delete(mapper, Joker)

	for _, count := range mapper {
		if count+jokerVal == 3 {
			return true
		}
	}

	return false
}

func isTwoPair(mapper map[string]int) bool {
	jokerVal, _ := mapper[Joker]
	delete(mapper, Joker)

	pairCount := 0
	for _, count := range mapper {
		if count == 2 {
			pairCount += 1
		}
		if jokerVal > 0 && count+1 == 2 {
			pairCount += 1
			jokerVal--
		}
	}

	return pairCount == 2
}

func isOnePair(mapper map[string]int) bool {
	_, hasJoker := mapper[Joker]
	delete(mapper, Joker)

	for _, count := range mapper {
		if count == 2 {
			return true
		}
		if hasJoker && count+1 == 2 {
			return true
		}
	}

	return false
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
