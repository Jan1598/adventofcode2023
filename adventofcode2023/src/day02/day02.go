package day02

import (
	"adventofcode2023"
	"fmt"
)

func GetFileInput() []string {
	return adventofcode2023.ReadData("adventofcode2023/src/day02/input.txt")
}

func RunPartOne() {
	fmt.Printf("Solution: %d", SumOfPossibleSets(GetFileInput()))
}

func RunPartTwo() {
	fmt.Printf("Solution: %d", SumPowerOfSets(GetFileInput()))
}

func SumOfPossibleSets(lines []string) int {

	sum := 0
	for _, line := range lines {
		gameData := CubeDataFromLine(line)

		valid := true
		for _, gameCube := range gameData.Cubes {
			red := gameCube.RedCubes
			blue := gameCube.BlueCubes
			green := gameCube.GreenCubes

			if red > 12 || green > 13 || blue > 14 {
				valid = false
				break
			}
		}

		if valid {
			sum += gameData.Id
		}
	}

	return sum
}

func SumPowerOfSets(lines []string) int {
	sum := 0
	for _, line := range lines {
		gameData := CubeDataFromLine(line)

		highestRed := 0
		highestGreen := 0
		highestBlue := 0
		for _, gameCube := range gameData.Cubes {
			red := gameCube.RedCubes
			blue := gameCube.BlueCubes
			green := gameCube.GreenCubes

			if highestRed < red {
				highestRed = red
			}
			if highestGreen < green {
				highestGreen = green
			}
			if highestBlue < blue {
				highestBlue = blue
			}
		}
		sum += highestRed * highestGreen * highestBlue
	}
	return sum
}
