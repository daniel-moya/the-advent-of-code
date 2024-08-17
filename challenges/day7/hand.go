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
	Index map[int]Card
	Value int
	Bid   int
}

func NewHand() *Hand {
	return &Hand{Index: make(map[int]Card), Value: 0}
}

func (h *Hand) Marshall(line string) error {
	parsedLine := strings.Split(strings.ReplaceAll(line, " ", "-"), "-")
	kinds := []string{}
	handIndex := make(map[string]int)

	for index, char := range parsedLine[0] {
		card := &Card{}
		err := card.Marshall(string(char))
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
	bid, err := strconv.Atoi(parsedLine[1])
	if err != nil {
		return fmt.Errorf("Errr parsing hand's bid, %v \n", err)
	}
	h.Bid = bid

	return nil
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

	return 0
}
