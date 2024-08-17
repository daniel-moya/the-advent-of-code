package day7

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	FIVE_KIND  = 7
	FOUR_KIND  = 6
	FULL_HOUSE = 5
	THREE_KIND = 4
	TWO_PAIR   = 3
	ONE_PAIR   = 2
	HIGH_CARD  = 1
)

type Hand struct {
	Index     map[int]Card
	Value     int
	Bid       int
	Raw       string
	kinds     []string
	kindCount map[string]int
}

func NewHand() *Hand {
	return &Hand{Index: make(map[int]Card), Value: 0}
}

func MarshallHand(line string, h *Hand) error {
	parsedLine := strings.Split(strings.ReplaceAll(line, " ", "-"), "-")
	kinds := []string{}
	handIndex := make(map[string]int)

	h.Raw = parsedLine[0]
	for index, char := range parsedLine[0] {
		card := &Card{}
		err := MarshallCard(string(char), card)
		if err != nil {
			return fmt.Errorf("Error parsing hand value, %v \n", err)
		}
		h.Index[index] = *card

		if _, ok := handIndex[card.Label]; !ok {
			kinds = append(kinds, card.Label)
		}
		handIndex[card.Label] += 1
	}

	h.Value = calculateHandValue(kinds, handIndex)
	h.kinds = kinds
	h.kindCount = handIndex
	bid, err := strconv.Atoi(parsedLine[1])
	if err != nil {
		return fmt.Errorf("Errr parsing hand's bid, %v \n", err)
	}
	h.Bid = bid

	return nil
}

func CompareHands(a, b Hand) int {
	if (a.Value - b.Value) != 0 {
		return a.Value - b.Value
	}
	for index := 0; index < 5; index++ {
		compared := CompareCards(a.Index[index], b.Index[index])
		if compared != 0 {
			return compared
		}
	}

	// If all 5 cards are equal in the same order
	return -1
}

func calculateHandValue(kinds []string, handIndex map[string]int) int {

	fiveKind := false
	fourKind := false
	threeKind := false
	pair := 0
	for _, kind := range kinds {
		switch handIndex[kind] {
		case 5:
			fiveKind = true
			break
		case 4:
			fourKind = true
			break
		case 3:
			threeKind = true
			break
		case 2:
			pair += 1
			break
		default:
		}
	}

	if fiveKind {
		return FIVE_KIND
	}

	if fourKind {
		return FOUR_KIND
	}

	if threeKind && pair == 1 {
		return FULL_HOUSE
	}

	if threeKind {
		return THREE_KIND
	}

	if pair == 2 {
		return TWO_PAIR
	}

	if pair == 1 {
		return ONE_PAIR
	}

	return 1
}
