package day6

import (
	"log"
	"strconv"
	"strings"

	"github.com/daniel-moya/aoc/challenges/utils"
)

type Millimeters int

type Race struct {
	Record   int
	Duration int
}

func Run(path string) int {
	return Part2(path)
}

func Part1(path string) int {
	lines := utils.GetInputFromFile(path)
	times := parseValues1(lines[0], "Time:")
	distances := parseValues1(lines[1], "Distance:")

	result := 1
	for index := 0; index < len(times); index++ {
		cRace := &Race{Duration: times[index], Record: distances[index]}
		amount := cRace.calculateRaceWinCombinations()
		result *= amount
	}

	return result
}

func Part2(path string) int {
	lines := utils.GetInputFromFile(path)
	time := parseValues2(lines[0], "Time:")
	distance := parseValues2(lines[1], "Distance:")

	println(time)
	println(distance)
	cRace := &Race{Duration: time, Record: distance}
	amount := cRace.calculateRaceWinCombinations()

	return amount
}

func (r *Race) calculateRaceWinCombinations() int {
	combinations := 0
	for hold := 1; hold < r.Duration; hold++ {
		// mm/s
		driving := r.Duration - hold
		distance := driving * hold
		if distance > r.Record {
			combinations += 1
		}
	}

	return combinations
}

func parseValues1(lineValues string, splitTerm string) []int {
	a := strings.Split(lineValues, splitTerm)
	b := strings.ReplaceAll(a[1], " ", "-")
	c := []int{}

	num := ""
	for index, char := range b {
		if index == len(b)-1 {
			if len(num) > 0 {
				num += string(char)
				val, err := strconv.Atoi(string(num))
				if err != nil {
					log.Fatalln("Error converting string to num", err)
				}
				c = append(c, val)
				num = ""

			}
		}
		if string(char) == "-" {
			if len(num) > 0 {
				val, err := strconv.Atoi(string(num))
				if err != nil {
					log.Fatalln("Error converting string to num", err)
				}
				c = append(c, val)
				num = ""
			}
			continue
		}
		num += string(char)
	}

	return c
}

func parseValues2(lineValues string, splitTerm string) int {
	a := strings.Split(lineValues, splitTerm)
	b := strings.ReplaceAll(a[1], " ", "-")

	numStr := ""
	for _, char := range b {
		if string(char) == "-" {
			continue
		}
		numStr += string(char)
	}

	val, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatalln("Error converting string to num", err)
	}

	return val
}
