package day09

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day09"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestCreateNumberMap(t *testing.T) {
	result := day09.CreateNumberMap(lines, false)

	adventofcode2023.AssertEquals(t, 3, len(result))
}

func TestCreateNextMap(t *testing.T) {
	result := day09.CreateNextMap([][]int{day09.CreateNumberMap(lines, false)[2]})

	adventofcode2023.AssertEquals(t, 5, len(result))
}

func TestGetNextValueFromMap(t *testing.T) {
	result := day09.CreateNextMap([][]int{day09.CreateNumberMap(lines, false)[2]})
	adventofcode2023.AssertEquals(t, 68, day09.GetNextValueFromMap(result))

	result = day09.CreateNextMap([][]int{day09.CreateNumberMap(lines, true)[2]})
	adventofcode2023.AssertEquals(t, 5, day09.GetNextValueFromMap(result))
}

func TestSumNextValuesFromHeapMap(t *testing.T) {
	result := day09.SumNextValuesFromHeapMap(lines, false)
	adventofcode2023.AssertEquals(t, 114, result)

	result = day09.SumNextValuesFromHeapMap(lines, true)
	adventofcode2023.AssertEquals(t, 2, result)
}
