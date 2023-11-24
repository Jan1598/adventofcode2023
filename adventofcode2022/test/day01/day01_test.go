package day01

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2022/src/day01"
	"adventofcode2023/adventofcode2022/test"
	"testing"
)

func TestAdd(t *testing.T) {
	lines := adventofcode2023.ReadData("input_test.txt")
	result := day01.CreateElfSumList(lines)
	expected := 3

	test.AssertEquals(t, expected, len(result))
	test.AssertEquals(t, 7, result[0].Calories)
	test.AssertEquals(t, 2, result[1].Calories)
	test.AssertEquals(t, 28, result[2].Calories)
}
