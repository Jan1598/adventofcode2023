package day07

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day07"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestCheckKinds(t *testing.T) {
	adventofcode2023.AssertEquals(t, true, day07.CheckKind("QQQJA", 4, 1))
	adventofcode2023.AssertEquals(t, false, day07.CheckKind("AATJT", 4, 1))
	adventofcode2023.AssertEquals(t, false, day07.CheckKind("KKQJQ", 4, 1))

	adventofcode2023.AssertEquals(t, true, day07.CheckKind("3A4JJ", 2, 2))

	adventofcode2023.AssertEquals(t, true, day07.CheckKind("3A42J", 2, 1))

	adventofcode2023.AssertEquals(t, true, day07.CheckKind("AAAA4", 4, 1))
	adventofcode2023.AssertEquals(t, true, day07.CheckKind("BB2B4", 3, 1))
	adventofcode2023.AssertEquals(t, true, day07.CheckKind("2A425", 2, 1))
}

func TestCheckFullHouse(t *testing.T) {
	adventofcode2023.AssertEquals(t, false, day07.CheckFullHouse("4854J"))
	adventofcode2023.AssertEquals(t, false, day07.CheckFullHouse("AA445"))
	adventofcode2023.AssertEquals(t, true, day07.CheckFullHouse("4A44A"))
	adventofcode2023.AssertEquals(t, true, day07.CheckFullHouse("AA55J"))
	adventofcode2023.AssertEquals(t, true, day07.CheckFullHouse("QJ22Q"))
}

func TestCollectHandCards(t *testing.T) {
	result := day07.CollectHandCards(lines)

	adventofcode2023.AssertEquals(t, 3, len(result))
}

func TestRateHandCards(t *testing.T) {
	result := day07.SumRateHandCards(lines)

	adventofcode2023.AssertEquals(t, 5905, result)
}
