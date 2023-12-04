package day02

import (
	"adventofcode2023"
	"adventofcode2023/adventofcode2023/src/day02"
	"testing"
)

var lines = adventofcode2023.ReadData("input_test.txt")

func TestCubeCollectionFromLine(t *testing.T) {

	gameData := day02.CubeDataFromLine("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

	adventofcode2023.AssertEquals(t, 1, gameData.Id)

	adventofcode2023.AssertEquals(t, 3, gameData.Cubes[0].BlueCubes)
	adventofcode2023.AssertEquals(t, 4, gameData.Cubes[0].RedCubes)
	adventofcode2023.AssertEquals(t, 0, gameData.Cubes[0].GreenCubes)

	adventofcode2023.AssertEquals(t, 6, gameData.Cubes[1].BlueCubes)
	adventofcode2023.AssertEquals(t, 1, gameData.Cubes[1].RedCubes)
	adventofcode2023.AssertEquals(t, 2, gameData.Cubes[1].GreenCubes)

	adventofcode2023.AssertEquals(t, 0, gameData.Cubes[2].BlueCubes)
	adventofcode2023.AssertEquals(t, 0, gameData.Cubes[2].RedCubes)
	adventofcode2023.AssertEquals(t, 2, gameData.Cubes[2].GreenCubes)
}

func TestCubeCollectionFromLineWithMoreSets(t *testing.T) {

	gameData := day02.CubeDataFromLine("Game 90: 16 blue, 7 red, 4 green; 4 green, 6 red, 11 blue; 2 red, 8 blue, 2 green; 5 green, 8 red, 10 blue; 4 red, 2 green, 7 blue; 4 green, 5 blue, 5 red")

	adventofcode2023.AssertEquals(t, 90, gameData.Id)
	adventofcode2023.AssertEquals(t, 6, len(gameData.Cubes))
}

func TestSumOfPossibleSets(t *testing.T) {
	result := day02.SumOfPossibleSets(lines)

	adventofcode2023.AssertEquals(t, 8, result)
}

func TestSumPowerOfSets(t *testing.T) {
	result := day02.SumPowerOfSets(lines)

	adventofcode2023.AssertEquals(t, 2286, result)
}
