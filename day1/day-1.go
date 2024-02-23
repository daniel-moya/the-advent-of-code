package main

import (
    "strconv"
    "unicode/utf8"
    "bufio"
    "fmt"
    "os"
)
func main() {
    
    day1("./input-1.txt")
}
asdfj;alkfdjadlfkj f sa;ldk fasklfj
func day1(inputPath string) int {
    // Get Calibration Document from File
    var input1 []string = getInputFromFile(inputPath)
    var total int
    for _, line := range input1 {
        // Add calibration number per line    
        fmt.Println(line + ": ")
        total += getCalibrationValue(line) 
    }

    fmt.Println("Total Calibration Value is %n", total)
    return total
}

func getCalibrationValue(line string) int {
    var digit string
    counter := 0
    for index, char := range line  {
        fmt.Println("LEFT: ", string(char))
        if _, errA := strconv.Atoi(string(char)); errA == nil {
            fmt.Println("found Ai ",string(char) )
            digit = string(digit) + string(char)
            counter = index 
            break
        } 
    }
fmt.Println("counter", counter)
    for reverseIndex := utf8.RuneCountInString(line) - 1; reverseIndex >= counter; reverseIndex-- {
        fmt.Println("RIGHT: ", string(line[reverseIndex]))
        if _, errB := strconv.Atoi(string(line[reverseIndex])); errB == nil {
            fmt.Println("FOUND B: ", string(line[reverseIndex]))
            digit = string(digit) + string(line[reverseIndex])
            break
        } 
    }
    fmt.Println("found: ", digit)
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
