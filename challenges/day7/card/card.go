package card

import (
	"fmt"
	"strconv"

	"github.com/daniel-moya/aoc/challenges/day7/algorithm"
)

type Card struct {
	Label string
	Value int
}

func CompareCards(a, b Card) int {
	return a.Value - b.Value
}

func MarshallCard(char string, c *Card, algo algorithm.Algorithm) error {
	val := 0

	switch char {
	case "A", "K", "Q", "T", "J":
		val = algo.GetValue(char)
	default:
		num, err := strconv.Atoi(char)
		if err != nil {
			return fmt.Errorf("Error parsing hand value, %v \n", err)
		}
		val = num

	}

	c.Label = string(char)
	c.Value = val
	return nil
}
