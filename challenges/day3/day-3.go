package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PartNumberBit struct {
	index         int
	isLeftCorner  bool
	isMiddle      bool
	isRightCorner bool
}

func Run(path string) int {
	var schemaLines []string = getInputFromFile(path)
	for _, line := range schemaLines {
		var chunks [][]PartNumberBit
		var currChunk []PartNumberBit
		var pushed bool
		for index, char := range line {
			_, err := strconv.ParseInt(string(char), 10, 64)
			if pushed {
				currChunk = []PartNumberBit{}
			}

			currPartNumberBit := PartNumberBit{index: index}
			if err == nil {
				if len(currChunk) == 0 {
					currPartNumberBit.isLeftCorner = true
				} else {
					currPartNumberBit.isMiddle = true
				}
				currChunk = append(currChunk, currPartNumberBit)
				pushed = false
			} else {
				if !pushed && len(currChunk) > 0 {
					currChunk[len(currChunk)-1].isRightCorner = true
					currChunk[len(currChunk)-1].isMiddle = false
					if len(currChunk) == 0 {
						currChunk[0].isMiddle = true
						currChunk[0].isRightCorner = true
					}

					chunks = append(chunks, currChunk)
					pushed = true
					// TODO: scan chunk after pushed
				}
			}
		}
		fmt.Println(chunks)
	}
	return 0
}

// Util to parse file into string input
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
