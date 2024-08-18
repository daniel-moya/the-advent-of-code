package hand

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/daniel-moya/aoc/challenges/day7/algorithm"
	"github.com/daniel-moya/aoc/challenges/day7/card"
)

type Hand struct {
	Index     map[int]card.Card
	Value     int
	Bid       int
	Raw       string
	kinds     []string
	kindCount map[string]int
}

func NewHand() *Hand {
	return &Hand{Index: make(map[int]card.Card), Value: 0}
}

func MarshallHand(line string, h *Hand, algo algorithm.Algorithm) error {
	parsedLine := strings.Split(strings.ReplaceAll(line, " ", "-"), "-")
	kinds := []string{}
	handIndex := make(map[string]int)

	h.Raw = parsedLine[0]
	for index, char := range parsedLine[0] {
		c := &card.Card{}

		err := card.MarshallCard(string(char), c, algo)
		if err != nil {
			return fmt.Errorf("Error parsing hand value, %v \n", err)
		}

		h.Index[index] = *c

		if _, ok := handIndex[c.Label]; !ok {
			kinds = append(kinds, c.Label)
		}
		handIndex[c.Label] += 1
	}

	h.Value = algo.Calculate(kinds, handIndex)
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
		compared := card.CompareCards(a.Index[index], b.Index[index])
		if compared != 0 {
			return compared
		}
	}

	// If all 5 cards are equal in the same order
	return 0
}
