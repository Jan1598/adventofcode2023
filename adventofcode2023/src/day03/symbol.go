package day03

import "regexp"

type Coordinate struct {
	XValue int
	YValue int
}

var gearMap = make(map[Coordinate][]int)
var symbols = regexp.MustCompile("^[^0-9.]+$")

func IsSymbolInRange(matrix [][]string, startCoordinate, endCoordinate Coordinate, numberValue int) bool {

	startCoordinate.UpdateCoordinate(matrix)
	endCoordinate.UpdateCoordinate(matrix)

	for i := startCoordinate.YValue; i <= endCoordinate.YValue; i++ {
		for i2 := startCoordinate.XValue; i2 <= endCoordinate.XValue; i2++ {
			d := matrix[i][i2]
			if symbols.MatchString(d) {

				symbolCoordinate := Coordinate{XValue: i2, YValue: i}
				PersistsValueToCoordinate(symbolCoordinate, numberValue)

				return true
			}
		}
	}

	return false
}

func (coordinate *Coordinate) UpdateCoordinate(matrix [][]string) {

	if len(matrix) <= coordinate.YValue {
		coordinate.YValue = len(matrix) - 1
	} else if coordinate.YValue < 0 {
		coordinate.YValue = 0
	}

	if len(matrix[0]) <= coordinate.XValue {
		coordinate.XValue = len(matrix[0]) - 1
	} else if coordinate.XValue < 0 {
		coordinate.XValue = 0
	}

}

func PersistsValueToCoordinate(coordinate Coordinate, numberValue int) {

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

func GetGearMap() map[Coordinate][]int {
	return gearMap
}

func CreateGearMap() {
	gearMap = make(map[Coordinate][]int)
}
