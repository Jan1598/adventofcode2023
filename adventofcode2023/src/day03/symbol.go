package day03

import (
	"adventofcode2023/helper/src"
	"regexp"
)

var gearMap = make(map[src.Coordinate][]int)
var symbolsRegex = regexp.MustCompile("^[^0-9.]+$")

func IsSymbolInRange(matrix [][]string, startCoordinate, endCoordinate src.Coordinate, numberValue int) bool {

	startCoordinate.UpdateCoordinate(matrix)
	endCoordinate.UpdateCoordinate(matrix)

	for i := startCoordinate.YValue; i <= endCoordinate.YValue; i++ {
		for i2 := startCoordinate.XValue; i2 <= endCoordinate.XValue; i2++ {
			d := matrix[i][i2]
			if symbolsRegex.MatchString(d) {

				symbolCoordinate := src.Coordinate{XValue: i2, YValue: i}
				PersistsValueToCoordinate(symbolCoordinate, numberValue)

				return true
			}
		}
	}

	return false
}

func PersistsValueToCoordinate(coordinate src.Coordinate, numberValue int) {

	if numberValue <= 0 {
		return
	}

	value, ok := gearMap[coordinate]

	if ok {
		value = append(value, numberValue)
	} else {
		value = []int{numberValue}
	}
	gearMap[coordinate] = value
}

func GetGearMap() map[src.Coordinate][]int {
	return gearMap
}

func CreateGearMap() {
	gearMap = make(map[src.Coordinate][]int)
}
