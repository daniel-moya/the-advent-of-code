package day1

import (
	"github.com/daniel-moya/aoc/challenges/utils"
	"strconv"
	"unicode/utf8"
)

func Run(inputPath string) int {
	// Get Calibration Document from File
	var input1 []string = utils.GetInputFromFile(inputPath)
	var total int
	for _, line := range input1 {
		// Add calibration number per line
		total += getCalibrationValue(line)
	}

	return total
}

func getCalibrationValue(line string) int {
	var digit string
	counter := 0
	for index, char := range line {
		if _, errA := strconv.Atoi(string(char)); errA == nil {
			digit = string(digit) + string(char)
			counter = index
			break
		}
	}
	for reverseIndex := utf8.RuneCountInString(line) - 1; reverseIndex >= counter; reverseIndex-- {
		if _, errB := strconv.Atoi(string(line[reverseIndex])); errB == nil {
			digit = string(digit) + string(line[reverseIndex])
			break
		}
	}
	intDigit, _ := strconv.Atoi(digit)
	return intDigit
}
