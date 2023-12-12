package day010

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day10"
	"adventofcode2023/helper/src"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")
var lines2 = adventofcode2023.ReadData("input_test2.txt")

func TestGetStartIndex(t *testing.T) {
	result := day10.GetStartIndex(src.BuildStringMatrix(lines))

	adventofcode2023.AssertEquals(t, 1, result.XValue)
	adventofcode2023.AssertEquals(t, 1, result.YValue)
}

func TestCalcRoadSteps(t *testing.T) {
	adventofcode2023.AssertEquals(t, 4, day10.CalcRoadSteps(lines, "RIGHT"))
	adventofcode2023.AssertEquals(t, 8, day10.CalcRoadSteps(lines2, "RIGHT"))
}
