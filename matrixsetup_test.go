package adventofcode2023

import (
	"testing"
)

func TestBuildStringMatrix(t *testing.T) {
	input := []string{"abcd", "defgh", "hijk", "lmnop"}

	output := BuildStringMatrix(input)

	AssertEquals(t, 4, len(output))
}

func TestCreateMatrix(t *testing.T) {
	input := []string{"abcd", "defgh", "hijk", "lmnop"}

	matrix := BuildStringMatrix(input)

	AssertEquals(t, true, IsIndexOnEdge(matrix, 0, 1))
	AssertEquals(t, true, IsIndexOnEdge(matrix, 2, 3))
	AssertEquals(t, false, IsIndexOnEdge(matrix, 2, 2))
	AssertEquals(t, false, IsIndexOnEdge(matrix, 1, 1))
}
