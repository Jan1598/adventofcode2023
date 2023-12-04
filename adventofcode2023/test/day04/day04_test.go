package day04

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day04"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestIsSymbolInRange(t *testing.T) {

	output := day04.CollectCardSet(lines)

	adventofcode2023.AssertEquals(t, 6, len(output))

	adventofcode2023.AssertEquals(t, 1, output[0].Id)
	adventofcode2023.AssertEquals(t, 5, len(output[0].WinningNumbers))
	adventofcode2023.AssertEquals(t, 8, len(output[0].OwnNumbers))
}

func TestSumOfMatches(t *testing.T) {
	output := day04.SumOfMatches(lines)

	adventofcode2023.AssertEquals(t, 13, output)
}

func TestSumOfScratchCards(t *testing.T) {
	output := day04.SumOfScratchCards(lines)

	adventofcode2023.AssertEquals(t, 30, output)
}
