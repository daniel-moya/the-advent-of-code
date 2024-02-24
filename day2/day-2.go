package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "reflect"
    "strconv"
    "strings"
    "golang.org/x/text/cases"
    "golang.org/x/text/language"
)    

type CubeSet struct {
    Red float64 
    Green float64 
    Blue float64 
}

func (cubeSet *CubeSet) setProperty(propName string, propValue float64) *CubeSet {
    reflect.ValueOf(cubeSet).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
    return cubeSet
}

func main (){
    day2()
}

func day2() {
    var capacitySet = CubeSet{Red:12, Green: 13, Blue: 14}
    
    var games []string = getInputFromFile("input-2.txt")
    var total int64

    for _, gameStr := range games {
	var gameSlice []string = strings.Split(gameStr, ":")
	var gameId, _  = strconv.ParseInt(
	    strings.Split(gameSlice[0], " ")[1],
	    10,
	    64,
	)
	total += processGame(gameId, strings.Trim(gameSlice[1], " "), capacitySet)
    }
    fmt.Printf("Total is %d \n", total)
}

// Returns the game id for games that meet the required game capacity
func processGame(id int64, cubesStr string, capacity CubeSet) int64 {
    var isGamePossible bool = true
    var minSet CubeSet 

    for _, subSet := range strings.Split(cubesStr, ";") {
	if id == 1 {
	    fmt.Println(minSet, subSet)
	    fmt.Printf("%v \n", MergeMaxSet(minSet, getSet(strings.Trim(subSet, " "))))
	}
        minSet = MergeMaxSet(minSet, getSet(strings.Trim(subSet, " ")))    
        if SetCompare(capacity, minSet) < 0 {
            isGamePossible = false
            break
        }
    }

    if isGamePossible {
	return id
    }

    return 0
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
    merged.Green =  math.Max(a.Green, b.Green)
    merged.Blue =  math.Max(a.Blue, b.Blue)
    return merged
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
