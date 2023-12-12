package day03

import (
	"adventofcode2023/helper/src"
	"fmt"
	"regexp"
	"strconv"
)

type Day03 struct {
}

var aocDay = src.AoCDay{
	Name: "day03",
}

func (r Day03) RunPartOne() {
	matrix := src.BuildStringMatrix(aocDay.GetFileInput())

	fmt.Printf("Solution: %d", SumOfSymbolNumbers(matrix))
}

func (r Day03) RunPartTwo() {
	fmt.Printf("Solution: %d", SumOfGears(aocDay.GetFileInput()))
}

func SumOfSymbolNumbers(matrix [][]string) int {

	CreateGearMap()

	re := regexp.MustCompile("[0-9]")
	sum := 0

	for i := 0; i < len(matrix); i++ {
		number := ""
		for i2 := 0; i2 < len(matrix[i]); i2++ {
			d := matrix[i][i2]
			if re.MatchString(d) {
				number += d
			}
			if len(number) > 0 && (len(matrix[i])-1 == i2) || (!re.MatchString(d)) {
				startCoordinate := src.Coordinate{XValue: (i2 - len(number)) - 1, YValue: i - 1}
				endCoordinate := src.Coordinate{XValue: i2, YValue: i + 1}
				num, err := strconv.Atoi(number)
				if IsSymbolInRange(matrix, startCoordinate, endCoordinate, num) {
					if err == nil {
						sum += num
					}
				}
				number = ""
			}
		}
	}

	return sum
}

func SumOfGears(lines []string) int {
	matrix := src.BuildStringMatrix(lines)
	SumOfSymbolNumbers(matrix)

	gearMap := GetGearMap()

	sum := 0
	for _, value := range gearMap {
		if len(value) == 2 {
			sum += value[0] * value[1]
		}
	}

	return sum
}
