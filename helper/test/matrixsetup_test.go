package test

import (
	"adventofcode2023"
	"adventofcode2023/helper/src"
	"testing"
)

func TestBuildStringMatrix(t *testing.T) {
	input := []string{"abcd", "defgh", "hijk", "lmnop"}

	output := src.BuildStringMatrix(input)

	adventofcode2023.AssertEquals(t, 4, len(output))
}

func TestCreateMatrix(t *testing.T) {
	input := []string{"abcd", "defgh", "hijk", "lmnop"}

	matrix := src.BuildStringMatrix(input)

	adventofcode2023.AssertEquals(t, true, src.IsIndexOnEdge(matrix, 0, 1))
	adventofcode2023.AssertEquals(t, true, src.IsIndexOnEdge(matrix, 2, 3))
	adventofcode2023.AssertEquals(t, false, src.IsIndexOnEdge(matrix, 2, 2))
	adventofcode2023.AssertEquals(t, false, src.IsIndexOnEdge(matrix, 1, 1))
}

func TestMoveCoordinate(t *testing.T) {
	input := []string{"abcd", "defgh", "hijk", "lmnop"}
	matrix := src.BuildStringMatrix(input)
	coordinate := src.Coordinate{XValue: 2, YValue: 3}

	coordinate.MoveCoordinate(matrix, "UP")
	adventofcode2023.AssertEquals(t, 2, coordinate.YValue)

	coordinate.MoveCoordinate(matrix, "DOWN")
	adventofcode2023.AssertEquals(t, 3, coordinate.YValue)

	coordinate.MoveCoordinate(matrix, "RIGHT")
	adventofcode2023.AssertEquals(t, 3, coordinate.XValue)

	coordinate.MoveCoordinate(matrix, "LEFT")
	adventofcode2023.AssertEquals(t, 2, coordinate.XValue)
}

func TestGetValueByCoordinate(t *testing.T) {

	input := []string{"abcd", "defgh", "hijk", "lmnop"}
	matrix := src.BuildStringMatrix(input)

	coordinate := src.Coordinate{XValue: 2, YValue: 3}
	adventofcode2023.AssertEquals(t, "n", coordinate.GetValueByCoordinate(matrix))

	coordinate.MoveCoordinate(matrix, "UP")
	adventofcode2023.AssertEquals(t, "j", coordinate.GetValueByCoordinate(matrix))

	coordinate.MoveCoordinate(matrix, "RIGHT")
	adventofcode2023.AssertEquals(t, "k", coordinate.GetValueByCoordinate(matrix))
}
