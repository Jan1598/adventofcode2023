package day02

import (
	"adventofcode2023"
	"fmt"
	"strings"
)

func GetFileInput() []string {
	return adventofcode2023.ReadData("adventofcode2022/src/day02/input.txt")
}

func RunPartOne() {
	fmt.Println(GetPointsByLetterArray(GetFileInput()))
}

func GetPointsByLetterArray(letters []string) int {
	sum := 0
	for _, line := range letters {
		shapeLetters := GetArrayByString(line)
		sum += CheckPoints(shapeLetters[1], shapeLetters[0])
	}

	return sum
}

func GetArrayByString(line string) []string {
	return strings.Split(line, " ")
}
