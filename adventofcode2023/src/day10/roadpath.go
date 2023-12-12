package day10

import "adventofcode2023/helper/src"

func CalcRoadSteps(lines []string, startAction string) int {
	steps := 0
	matrix := src.BuildStringMatrix(lines)

	startCoordinate := GetStartIndex(matrix)
	endCoordinate := startCoordinate

	startCoordinate.MoveCoordinate(matrix, "DOWN")
	actionBefore := "DOWN"
	for {
		nextRoadPart := startCoordinate.GetValueByCoordinate(matrix)

		action := ""
		switch nextRoadPart {
		case "-":
			action = "RIGHT"
			if actionBefore != "RIGHT" {
				action = "LEFT"
			}
		case "7":
			action = "DOWN"
			if actionBefore == "UP" {
				action = "LEFT"
			}
		case "|":
			action = "DOWN"
			if actionBefore != "DOWN" {
				action = "UP"
			}
		case "J":
			action = "LEFT"
			if actionBefore == "RIGHT" {
				action = "UP"
			}
		case "L":
			action = "UP"
			if actionBefore == "DOWN" {
				action = "RIGHT"
			}
		case "F":
			action = "RIGHT"
			if actionBefore == "LEFT" {
				action = "DOWN"
			}
		}

		startCoordinate.MoveCoordinate(matrix, action)
		actionBefore = action

		if startCoordinate == endCoordinate {
			break
		}
		steps++
	}

	return (steps / 2) + 1
}

func GetStartIndex(matrix [][]string) src.Coordinate {
	for i, _ := range matrix[0] {
		for i2 := 0; i2 < len(matrix[i]); i2++ {
			if matrix[i][i2] == "S" {
				return src.Coordinate{XValue: i2, YValue: i}
			}
		}
	}

	return src.Coordinate{}
}
