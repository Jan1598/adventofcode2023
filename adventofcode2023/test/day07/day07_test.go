package day07

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day07"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestCreateSeeds(t *testing.T) {
	result := day07.CollectHandCards(lines)

	adventofcode2023.AssertEquals(t, 5, len(result))
}

func TestCheckKinds(t *testing.T) {
	adventofcode2023.AssertEquals(t, "A", day07.CheckKind("AAAA4", 4, 1)[0])
	adventofcode2023.AssertEquals(t, "B", day07.CheckKind("BB2B4", 3, 1)[0])
	adventofcode2023.AssertEquals(t, "2", day07.CheckKind("2A425", 2, 1)[0])

	result := day07.CheckKind("AA445", 2, 2)
	adventofcode2023.AssertEquals(t, "A", result[0])
	adventofcode2023.AssertEquals(t, "4", result[1])
}

func TestCheckFullHouse(t *testing.T) {
	result := day07.CheckFullHouse("AA445")
	adventofcode2023.AssertEquals(t, 0, len(result))

	result = day07.CheckFullHouse("4A44A")
	adventofcode2023.AssertEquals(t, "4", result[0])
	adventofcode2023.AssertEquals(t, "A", result[1])
}

func TestCollectHandCards(t *testing.T) {
	result := day07.CollectHandCards(lines)

	adventofcode2023.AssertEquals(t, 3, len(result))
}

func TestRateHandCards(t *testing.T) {
	result := day07.SumRateHandCards(lines)

	adventofcode2023.AssertEquals(t, 6440, result)
}
