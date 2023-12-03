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
