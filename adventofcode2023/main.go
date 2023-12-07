package main

import (
	"adventofcode2023/adventofcode2023/src/day05"
	"fmt"
)

func main() {
	currentDay := day05.Day05{}

	currentDay.RunPartOne()
	fmt.Println("")

	// 2117041827 -> too high
	currentDay.RunPartTwo()
}
