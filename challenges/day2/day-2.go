package day2

import (
	"github.com/daniel-moya/aoc/challenges/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type CubeSet struct {
	Red   float64
	Green float64
	Blue  float64
}

func (cubeSet *CubeSet) setProperty(propName string, propValue float64) *CubeSet {
	reflect.ValueOf(cubeSet).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return cubeSet
}

func (cubeSet *CubeSet) getPower() float64 {
	return cubeSet.Red * cubeSet.Green * cubeSet.Blue
}

func Run(filePath string) (int64, float64) {
	var capacitySet = CubeSet{Red: 12, Green: 13, Blue: 14}

	var games []string = utils.GetInputFromFile(filePath)
	var total int64
	var totalOfSetsPower float64

	for _, gameStr := range games {
		var gameSlice []string = strings.Split(gameStr, ":")
		var gameId, _ = strconv.ParseInt(
			strings.Split(gameSlice[0], " ")[1],
			10,
			64,
		)
		id, minCubeSetPower := processGame(gameId, strings.Trim(gameSlice[1], " "), capacitySet)

		total += id
		totalOfSetsPower += minCubeSetPower
	}
	return total, totalOfSetsPower
}

// Returns the game id for games that meet the required game capacity
func processGame(id int64, cubesStr string, capacity CubeSet) (int64, float64) {
	var isGamePossible bool = true
	var minSet CubeSet
	var minCapacitySet CubeSet

	for index, subSet := range strings.Split(cubesStr, ";") {
		currentSet := getSet(strings.Trim(subSet, " "))
		minSet = MergeMaxSet(minSet, currentSet)

		if SetCompare(capacity, minSet) < 0 {
			isGamePossible = false
		}

		if index > 0 {
			minCapacitySet = MergeMaxSetWithoutZero(minCapacitySet, currentSet)
		} else {
			minCapacitySet = currentSet
		}
	}

	if isGamePossible {
		return id, minCapacitySet.getPower()
	}

	return 0, minCapacitySet.getPower()
}

var caser = cases.Title(language.English)

// Parses a cube set string
// Returns a CubeSet struct with the values from the parsed string
func getSet(s string) CubeSet {
	var cubeSet CubeSet
	var cubeString []string = strings.Split(s, ",")

	for _, c := range cubeString {
		var cube string = strings.Trim(c, " ")
		value, _ := strconv.ParseFloat(strings.Split(cube, " ")[0], 64)
		cubeSet.setProperty(
			caser.String(strings.Split(cube, " ")[1]),
			value,
		)
	}

	return cubeSet
}

// Merges 2 sets with the max values from each
// Returns the MaxCubeSet
func MergeMaxSet(a CubeSet, b CubeSet) CubeSet {
	merged := CubeSet{}
	merged.Red = math.Max(a.Red, b.Red)
	merged.Green = math.Max(a.Green, b.Green)
	merged.Blue = math.Max(a.Blue, b.Blue)
	return merged
}

// Merges 2 sets with the min values from each Except 0
// Returns the MinCubeSet
func MergeMaxSetWithoutZero(a CubeSet, b CubeSet) CubeSet {
	merged := CubeSet{}
	merged.Red = maxWithoutZero(a.Red, b.Red)
	merged.Green = maxWithoutZero(a.Green, b.Green)
	merged.Blue = maxWithoutZero(a.Blue, b.Blue)

	return merged
}

func maxWithoutZero(a float64, b float64) float64 {
	if a != 0 && b != 0 {
		return math.Max(a, b)
	} else if a == 0 {
		return b
	}
	return a
}

// Compare 2 CubeSet struct types
// Returns a positive number if any value of A is greater than B
// Returns a negative numebr if any value from B is greater than A
func SetCompare(a CubeSet, b CubeSet) int {
	if a.Red >= b.Red && a.Blue >= b.Blue && a.Green >= b.Green {
		return 1
	}
	return -1
}
