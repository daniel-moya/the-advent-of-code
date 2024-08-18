package day7

import (
	"github.com/daniel-moya/aoc/challenges/day7/algorithm"
	"github.com/daniel-moya/aoc/challenges/day7/hand"
	"github.com/daniel-moya/aoc/challenges/day7/ranking"
	"github.com/daniel-moya/aoc/challenges/utils"
)

func Run(path string) int {
	// return Part1(path)
	return Part2(path)
}

func Part1(path string) int {
	simple := algorithm.NewSimpleAlgorithm()
	return ParseCardHands(path, simple)
}

func Part2(path string) int {
	joker := algorithm.NewJokerAlgorithm()
	return ParseCardHands(path, joker)
}

func ParseCardHands(path string, algo algorithm.Algorithm) int {
	lines := utils.GetInputFromFile(path)
	hands := []hand.Hand{}
	for _, line := range lines {
		h := hand.NewHand()
		hand.MarshallHand(line, h, algo)
		hands = append(hands, *h)
	}

	ranking := ranking.NewSortedRanking(hands)
	return ranking.GetBidTotal()
}
