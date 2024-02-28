package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func Run() {
	fmt.Println("\n Challenge #1")
	day1("./challenges/day1/input-1.txt")
}

func day1(inputPath string) int {
	// Get Calibration Document from File
	var input1 []string = getInputFromFile(inputPath)
	var total int
	for _, line := range input1 {
		// Add calibration number per line
		total += getCalibrationValue(line)
	}

	fmt.Println("Total Calibration Value is %n", total)
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

func getInputFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	check(err)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var inputSlice []string
	for fileScanner.Scan() {
		inputSlice = append(inputSlice, fileScanner.Text())
	}
	file.Close()
	return inputSlice
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
