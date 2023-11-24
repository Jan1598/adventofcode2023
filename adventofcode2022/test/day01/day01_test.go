package day01

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2022/src/day01"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestCreateElfSumList(t *testing.T) {
	result := day01.CreateElfSumList(lines)
	expected := 5

	adventofcode2023.AssertEquals(t, expected, len(result))
	adventofcode2023.AssertEquals(t, 6000, result[0].Calories)
	adventofcode2023.AssertEquals(t, 4000, result[1].Calories)
	adventofcode2023.AssertEquals(t, 11000, result[2].Calories)
	adventofcode2023.AssertEquals(t, 24000, result[3].Calories)
	adventofcode2023.AssertEquals(t, 10000, result[4].Calories)
}

func TestSumOfCaloriesOfTop(t *testing.T) {

	result := day01.SumOfCaloriesOfTop(3, day01.CreateElfSumList(lines))

	adventofcode2023.AssertEquals(t, 45000, result)
}
