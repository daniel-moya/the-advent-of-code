package day3

import (
	"fmt"
	"github.com/daniel-moya/aoc/challenges/utils"
	"strconv"
)

type PartNumberBit struct {
	row           int
	index         int
	isLeftCorner  bool
	isMiddle      bool
	isRightCorner bool
	isValid       bool
}

func Run(path string) int {
	var schemaLines []string = utils.GetInputFromFile(path)
	for row, line := range schemaLines {
		processLine(schemaLines, line, row)
	}
	return 0
}

func processLine(schema []string, line string, row int) {
	var chunks [][]PartNumberBit
	var currChunk []PartNumberBit
	var pushed bool
	for col, char := range line {
		_, err := strconv.ParseInt(string(char), 10, 64)
		if pushed {
			currChunk = []PartNumberBit{}
		}

		currPartNumberBit := PartNumberBit{row: row, index: col}
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
				scanChunk(schema, row, col, line, currChunk)
			}
		}
	}
}

func scanChunk(
	schema []string,
	row int,
	col int,
	line string,
	chunk []PartNumberBit,
) {
	fmt.Println(line)
	for _, partNumberBit := range chunk {
		fmt.Println(string(line[partNumberBit.index]))
		// if partNumberBit.isLeftCorner {
		// 	partNumberBit.isValid = scanLeftCorner(schema, line, partNumberBit)
		// }
		//
		if partNumberBit.isMiddle {
			partNumberBit.isValid = scanMiddle(schema, line, partNumberBit)
		}

		// if partNumberBit.isRightCorner {
		// 	partNumberBit.isValid = scanRightCorner(schema, line, partNumberBit)
		// }
	}
}

// func scanLeftCorner(schema []string, line string, part PartNumberBit) bool {
// 	left, topLeftDiag, bottomLeftDiag := false, false, false
// 	middle := scanMiddle(schema, line, part)
//
// 	if part.row > 0 {
// 		topLeftDiag := isSymbol(schema[part.row-1][part.index-1])
// 	}
//
// 	if part.row < len(schema) - 1 {
// 		bottomLeftDiag := isSymbol(schema[part.row+1][part.index-1])
// 	}
//
// 	if part.index > 0 {
// 		left := isSymbol(schema[part.row][part.index-1])
// 	}
//
//
// 	return middle || topLeftDiag || left || bottomLeftDiag
// }

func scanMiddle(schema []string, line string, part PartNumberBit) bool {
	var top, bottom bool

	if part.row > 0 {
		top = isSymbol(schema[part.row-1][part.index])
	}

	if part.row < len(schema)-1 {
		bottom = isSymbol(schema[part.row+1][part.index])
	}

	return top || bottom
}

// func scanRightCorner(schema []string, line string, part PartNumberBit) bool {
// 	middle := scanMiddle(schema, line, part)
//
// 	return middle
// }

func isSymbol(char byte) bool {
	_, err := strconv.ParseInt(string(char), 10, 64)
	return err == nil && char != '.'
}
