package main

import (
	"fmt"
	"time"

	"github.com/daniel-moya/aoc/challenges/day6"
)

func main() {
	start := time.Now()
	total := day6.Run("./challenges/day6/input.txt")
	fmt.Printf("Result: %v in %v\n", total, time.Since(start))
}
