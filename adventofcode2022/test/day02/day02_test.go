package day01

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2022/src/day02"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestCheckPoints(t *testing.T) {
	adventofcode2023.AssertEquals(t, 8, day02.CheckPoints("Y", "A"))
	adventofcode2023.AssertEquals(t, 1, day02.CheckPoints("A", "Y"))
	adventofcode2023.AssertEquals(t, 4, day02.CheckPoints("A", "X"))
}

func TestGetPointsByLetterArray(t *testing.T) {
	result := day02.GetPointsByLetterArray(lines)
	expected := 15

	adventofcode2023.AssertEquals(t, expected, result)
}

func TestGetCalcValueByLetter(t *testing.T) {
	adventofcode2023.AssertEquals(t, "X", day02.GetCalcValueByLetter("Y", "A"))
	adventofcode2023.AssertEquals(t, "Y", day02.GetCalcValueByLetter("Y", "B"))
	adventofcode2023.AssertEquals(t, "Z", day02.GetCalcValueByLetter("Y", "C"))

	adventofcode2023.AssertEquals(t, "X", day02.GetCalcValueByLetter("X", "B"))
	adventofcode2023.AssertEquals(t, "Y", day02.GetCalcValueByLetter("X", "C"))
	adventofcode2023.AssertEquals(t, "Z", day02.GetCalcValueByLetter("X", "A"))

	adventofcode2023.AssertEquals(t, "Y", day02.GetCalcValueByLetter("Z", "A"))
	adventofcode2023.AssertEquals(t, "Z", day02.GetCalcValueByLetter("Z", "B"))
	adventofcode2023.AssertEquals(t, "X", day02.GetCalcValueByLetter("Z", "C"))
}

func TestCheckPointsPartTwo(t *testing.T) {
	adventofcode2023.AssertEquals(t, 4, day02.CheckPoints(day02.GetCalcValueByLetter("Y", "A"), "A"))
	adventofcode2023.AssertEquals(t, 1, day02.CheckPoints(day02.GetCalcValueByLetter("X", "B"), "B"))
	adventofcode2023.AssertEquals(t, 7, day02.CheckPoints(day02.GetCalcValueByLetter("Z", "C"), "C"))
}

func TestGetPointsByLetterArrayPartTwo(t *testing.T) {
	result := day02.GetPointsByLetterArrayPartTwo(lines)
	expected := 12

	adventofcode2023.AssertEquals(t, expected, result)
}
