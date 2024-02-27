package day3  

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Range struct {
    Start int
    End int
}

func Run() {
    fmt.Println("\n Challenge #3")   
    day3()
}

func day3() {
    var schemaLines []string = getInputFromFile("./challenges/day3/input-3-test.txt")
    for _, line := range schemaLines {
        var digitsSlice = []Range{{-1, -1}}
        for col, char := range line {
            sliceIndex := 0
            currentSlice := &digitsSlice[sliceIndex]
            _, err := strconv.ParseInt(string(char), 10, 32) 
            if err == nil {
                fmt.Println(currentSlice.Start)
                if currentSlice.Start == -1 {
                    fmt.Printf("set start %v \n", col)
                    currentSlice.Start = col 
                } 
            } else if currentSlice.Start > -1 && currentSlice.End == -1 {
                fmt.Printf("set end %v \n", col-1)
                currentSlice.End = col - 1    
                digitsSlice = append(digitsSlice, Range{-1, -1})
                sliceIndex += 1
            }
               
        } 
        fmt.Printf("slices %v \n", digitsSlice)
    }
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
