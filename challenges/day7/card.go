package day7

import (
	"fmt"
	"strconv"
)

const (
	A     = 14
	K     = 13
	Q     = 12
	J     = 11
	T     = 10
	NINE  = 9
	EIGHT = 8
	SEVEN = 7
	SIX   = 6
	FIVE  = 5
	FOUR  = 4
	THREE = 3
	TWO   = 2
	ONE   = 1
)

type Card struct {
	Label string
	Value int
}

func CompareCards(a, b Card) int {
	return a.Value - b.Value
}

func (c *Card) Marshall(char string) error {
	val := 0

	switch char {
	case "A":
	case "K":
	case "Q":
	case "J":
	case "T":
		val = getValueFromLetter(char)
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

func getValueFromLetter(letter string) int {
	switch letter {
	case "A":
		return A
	case "K":
		return K
	case "Q":
		return Q
	case "J":
		return J
	}
	return 0
}
