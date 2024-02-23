package main

import (
    "strings"
    "golang.org/x/text/language"
    "golang.org/x/text/cases"
    "reflect"
    _ "os"
    "math"
    "fmt"
    _ "bufio"
)    

type CubeSet struct {
    Red int32
    Green int32
    Blue int32
}

func main (){
    day2()
}

func day2() {
    var capacitySet = CubeSet{Red:12, Green: 13, Blue: 14}
    var testStringSubSet string = "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
    var minSet CubeSet 

    var isGamePossible bool = true

    for _, subSet := range strings.Split(testStringSubSet, ";") {
        minSet = MergeMaxSet(minSet, getSet(subSet))    
        if SetCompare(capacitySet, minSet) < 0 {
            isGamePossible = false
            break
        }
    }
    fmt.Println("Game Possible: %t", isGamePossible)
}

func (i *CubeSet) setProperty(propName string, propValue int) *CubeSet {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

func getSet(s string) CubeSet {
    var chunk []string = strings.Split(s, ",")
    var cubeSet CubeSet
    for _, cube := range chunk {
        var cubeValues []string = strings.Split(cube, " ")
        c := cases.Title(language.English)
        color := c.String(cubeValues[1])
        cubeSet = cubeSet.setProperty(color, cubeValues[1])
    }
    return cubeSet
}

func MergeMaxSet(a CubeSet, b CubeSet) CubeSet {
    merged := CubeSet{}
    merged.Red = math.Max(a.Red, b.Red)
    merged.Green =  math.Max(a.Green, b.Green)
    merged.Blue =  math.Max(a.Blue, b.Blue)
    return merged
}

func SetCompare(a CubeSet, b CubeSet) int {
    if a.Red >= b.Red || a.Blue >= b.Blue || a.Green >= b.Green {
        return 1
    }
    return -1
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}
