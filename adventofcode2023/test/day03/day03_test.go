package day02

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day03"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestIsSymbolInRange(t *testing.T) {

	input := []string{"...", ".1$", "...", "..."}
	matrix := adventofcode2023.BuildStringMatrix(input)

	startCoordinate := day03.Coordinate{XValue: 0, YValue: 0}
	endCoordinate := day03.Coordinate{XValue: 2, YValue: 1}
	adventofcode2023.AssertEquals(t, true, day03.IsSymbolInRange(matrix, startCoordinate, endCoordinate, 0))

	startCoordinate = day03.Coordinate{XValue: 0, YValue: 2}
	endCoordinate = day03.Coordinate{XValue: 2, YValue: 3}
	adventofcode2023.AssertEquals(t, false, day03.IsSymbolInRange(matrix, startCoordinate, endCoordinate, 0))
}

func TestSumOfSymbolNumbers(t *testing.T) {
	matrix := adventofcode2023.BuildStringMatrix(lines)
	adventofcode2023.AssertEquals(t, 4361, day03.SumOfSymbolNumbers(matrix))

	gearMap := day03.GetGearMap()
	adventofcode2023.AssertEquals(t, 6, len(gearMap))
}

func TestSumOfGears(t *testing.T) {
	result := day03.SumOfGears(lines)
	adventofcode2023.AssertEquals(t, 467835, result)
}
