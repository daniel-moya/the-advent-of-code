package day4

import (
	"github.com/daniel-moya/aoc/challenges/utils"
	"strings"
)

func Run(path string) int {
	slice := utils.GetInputFromFile(path)
	total := 0
	for _, line := range slice {
		points := 0
		card := strings.Split(line, ":")
		numbers := strings.Split(card[1], "|")
		winners := strings.Fields(strings.TrimSpace(numbers[0]))
		assigned := strings.Fields(strings.TrimSpace(numbers[1]))

		for _, current := range assigned {
			for _, winner := range winners {
				if current == winner {
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
					break
				}
			}
		}
		total += points
	}

	return total
}
