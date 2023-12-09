package day09

import (
	"adventofcode2023/helper/src"
	"fmt"
)

type Day09 struct {
}

var aocDay = src.AoCDay{
	Name: "day09",
}

func (r Day09) RunPartOne() {
	fmt.Printf("Solution: %d", SumNextValuesFromHeapMap(aocDay.GetFileInput(), false))
}

func (r Day09) RunPartTwo() {
	fmt.Printf("Solution: %d", SumNextValuesFromHeapMap(aocDay.GetFileInput(), true))
}

func SumNextValuesFromHeapMap(lines []string, reverse bool) int {
	sum := 0
	numberMap := CreateNumberMap(lines, reverse)

	for i := 0; i < len(numberMap); i++ {
		sum += GetNextValueFromMap(CreateNextMap([][]int{numberMap[i]}))
	}

	return sum
}
