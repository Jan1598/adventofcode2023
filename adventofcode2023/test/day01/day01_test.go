package day01

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day01"
	"testing"
)

func TestSwapWord(t *testing.T) {
	adventofcode2023.AssertEquals(t, "test", day01.SwapWord("tset"))
	adventofcode2023.AssertEquals(t, "tone", day01.SwapWord("enot"))
}

func TestCollectNumbersInWord(t *testing.T) {
	adventofcode2023.AssertEquals(t, 23, day01.CollectNumbersInWord("t2asdasd3"))
	adventofcode2023.AssertEquals(t, 27, day01.CollectNumbersInWord("as2asd5asd2opld27"))
	adventofcode2023.AssertEquals(t, 22, day01.CollectNumbersInWord("s2aa"))
	adventofcode2023.AssertEquals(t, 29, day01.CollectNumbersInWord("two1nine"))
	adventofcode2023.AssertEquals(t, 14, day01.CollectNumbersInWord("zoneight234"))
	adventofcode2023.AssertEquals(t, 31, day01.CollectNumbersInWord("three7mxndngcvxbkgxqthreetwoneh"))
}

func TestSumOfNumbers(t *testing.T) {
	result := day01.SumOfNumbers(adventofcode2023.ReadData("input_test.txt"))
	expected := 142

	adventofcode2023.AssertEquals(t, expected, result)
}

func TestSumOfNumbersPart2(t *testing.T) {
	result := day01.SumOfNumbers(adventofcode2023.ReadData("input_test2.txt"))
	expected := 281

	adventofcode2023.AssertEquals(t, expected, result)
}
