package day7

import (
	"github.com/daniel-moya/aoc/challenges/utils"
)

func Run(path string) int {
	return Part1(path)
}

func Part1(path string) int {
	lines := utils.GetInputFromFile(path)
	hands := []Hand{}
	for _, line := range lines {
		hand := NewHand()
		MarshallHand(line, hand)
		hands = append(hands, *hand)
	}

	ranking := NewSortedRanking(hands)
	return ranking.GetBidTotal()
}
