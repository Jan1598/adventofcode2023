package day06

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day06"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestCollectHandCards(t *testing.T) {
	result := day06.CollectRaces(lines)

	adventofcode2023.AssertEquals(t, 3, len(result))
}

func TestCalculateWinRatios(t *testing.T) {
	race := day06.Race{Distance: 9, Time: 7}
	result := day06.CalculateWinRatios(race)

	adventofcode2023.AssertEquals(t, 4, len(result))
	adventofcode2023.AssertEquals(t, 2, result[0])
	adventofcode2023.AssertEquals(t, 3, result[1])
	adventofcode2023.AssertEquals(t, 4, result[2])
	adventofcode2023.AssertEquals(t, 5, result[3])

}

func TestCalculateWinRatios2(t *testing.T) {
	race := day06.Race{Distance: 40, Time: 15}
	result := day06.CalculateWinRatios(race)

	adventofcode2023.AssertEquals(t, 8, len(result))
	adventofcode2023.AssertEquals(t, 4, result[0])
	adventofcode2023.AssertEquals(t, 5, result[1])
	adventofcode2023.AssertEquals(t, 6, result[2])
	adventofcode2023.AssertEquals(t, 7, result[3])
	adventofcode2023.AssertEquals(t, 8, result[4])
	adventofcode2023.AssertEquals(t, 9, result[5])
	adventofcode2023.AssertEquals(t, 10, result[6])
	adventofcode2023.AssertEquals(t, 11, result[7])
}

func TestCalculateWinRatios3(t *testing.T) {
	race := day06.Race{Distance: 200, Time: 30}
	result := day06.CalculateWinRatios(race)

	adventofcode2023.AssertEquals(t, 9, len(result))
	adventofcode2023.AssertEquals(t, 11, result[0])
	adventofcode2023.AssertEquals(t, 12, result[1])
	adventofcode2023.AssertEquals(t, 13, result[2])
	adventofcode2023.AssertEquals(t, 14, result[3])
	adventofcode2023.AssertEquals(t, 15, result[4])
	adventofcode2023.AssertEquals(t, 16, result[5])
	adventofcode2023.AssertEquals(t, 17, result[6])
	adventofcode2023.AssertEquals(t, 18, result[7])
	adventofcode2023.AssertEquals(t, 19, result[8])
}

func TestCalculateWinRatios4(t *testing.T) {
	race := day06.Race{Distance: 940200, Time: 71530}
	result := day06.CalculateWinRatios(race)

	adventofcode2023.AssertEquals(t, 71503, len(result))
}

func TestCalculateWinRatiosOfRaces(t *testing.T) {
	result := day06.CalculateWinRatiosOfRaces(lines)

	adventofcode2023.AssertEquals(t, 288, result)
}

func TestParseRacesToOne(t *testing.T) {
	result := day06.ParseRacesToOne(day06.CollectRaces(lines))

	adventofcode2023.AssertEquals(t, 940200, result.Distance)
	adventofcode2023.AssertEquals(t, 71530, result.Time)
}
