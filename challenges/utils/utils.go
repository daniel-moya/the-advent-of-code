package utils

import "os"
import "bufio"

func GetInputFromFile(filePath string) []string {
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
