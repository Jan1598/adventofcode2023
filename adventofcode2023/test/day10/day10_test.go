package day010

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day10"
	"adventofcode2023/helper/src"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestGetStartIndex(t *testing.T) {
	result := day10.GetStartIndex(src.BuildStringMatrix(lines))

	adventofcode2023.AssertEquals(t, 1, result[0])
	adventofcode2023.AssertEquals(t, 1, result[1])
}
