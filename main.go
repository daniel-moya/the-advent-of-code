package main

import (
	"fmt"
	"time"

	"github.com/daniel-moya/aoc/challenges/day8"
)

func main() {
	start := time.Now()
	total := day8.Run("./challenges/day8/input.txt")
	fmt.Printf("Result: %v in %v\n", total, time.Since(start))
}
