package main

import (
	"fmt"
	"time"

	"github.com/daniel-moya/aoc/challenges/day7"
)

func main() {
	start := time.Now()
	total := day7.Run("./challenges/day7/input-test.txt")
	fmt.Printf("Result: %v in %v\n", total, time.Since(start))
}
