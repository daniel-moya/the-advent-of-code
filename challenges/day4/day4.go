package day4

import (
	"github.com/daniel-moya/aoc/challenges/utils"
	"strings"
)

func Run(path string) int {
	slice := utils.GetInputFromFile(path)
	total := 0
	var cardsMap map[int]int = make(map[int]int)

	for cardNumber := range slice {
		cardsMap[cardNumber] = 1
	}

	for row, line := range slice {
		matches := 0
		card := strings.Split(line, ":")
		numbers := strings.Split(card[1], "|")
		winners := strings.Fields(strings.TrimSpace(numbers[0]))
		assigned := strings.Fields(strings.TrimSpace(numbers[1]))

		for _, current := range assigned {
			for _, winner := range winners {
				if current == winner {
					matches++
					break
				}
			}
		}

		if matches > 0 {
			for index := row + 1; index <= row+matches; index++ {
				cardsMap[index] = cardsMap[index] + cardsMap[row]
			}
		}

	}
	for row := range slice {
		total += cardsMap[row]
	}

	return total
}
