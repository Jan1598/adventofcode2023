package day04

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day05"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestCreateSeeds(t *testing.T) {
	result := day05.CreateSeeds(lines[0])

	adventofcode2023.AssertEquals(t, 4, len(result))
}

func TestCollectMaps(t *testing.T) {
	result := day05.CollectMaps(lines)

	adventofcode2023.AssertEquals(t, 7, len(result))
}

func TestCreateValueMap(t *testing.T) {
	seedMapInput := []day05.SeedMapInput{{RangeLength: 2, SourceRange: 98, DestinationRange: 50},
		{RangeLength: 48, SourceRange: 50, DestinationRange: 52},
		{RangeLength: 5, SourceRange: 10, DestinationRange: 60}}
	result := day05.CreateValueMap(seedMapInput)

	adventofcode2023.AssertEquals(t, 50, result[98])
	adventofcode2023.AssertEquals(t, 51, result[99])
	adventofcode2023.AssertEquals(t, 52, result[50])
	adventofcode2023.AssertEquals(t, 53, result[51])
	adventofcode2023.AssertEquals(t, 60, result[10])
	adventofcode2023.AssertEquals(t, 64, result[14])
}

func TestGetNumberFromSeeds(t *testing.T) {
	seedMapInput := []day05.SeedMapInput{{RangeLength: 2, SourceRange: 98, DestinationRange: 50},
		{RangeLength: 48, SourceRange: 50, DestinationRange: 52}}
	result := day05.CreateValueMap(seedMapInput)
	seedNumbers := day05.GetNumberFromSeeds([]int{79, 14, 55, 13}, result)

	adventofcode2023.AssertEquals(t, 81, seedNumbers[0])
	adventofcode2023.AssertEquals(t, 14, seedNumbers[1])
	adventofcode2023.AssertEquals(t, 57, seedNumbers[2])
	adventofcode2023.AssertEquals(t, 13, seedNumbers[3])
}

func TestPart(t *testing.T) {
	output := day05.GetLocationValues(lines)

	adventofcode2023.AssertEquals(t, 4, len(output))
	adventofcode2023.AssertEquals(t, 35, output[0])
	adventofcode2023.AssertEquals(t, 43, output[1])
	adventofcode2023.AssertEquals(t, 82, output[2])
	adventofcode2023.AssertEquals(t, 86, output[3])
}
