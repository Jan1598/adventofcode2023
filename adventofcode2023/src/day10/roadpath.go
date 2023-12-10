package day10

import "adventofcode2023/helper/src"

func BuildMatrixFromInput(lines []string) {
	matrix := src.BuildStringMatrix(lines)

	GetStartIndex(matrix)
}

func GetStartIndex(matrix [][]string) []int {
	for i, _ := range matrix[0] {
		for i2 := 0; i2 < len(matrix[i]); i2++ {
			if matrix[i][i2] == "S" {
				return []int{i, i2}
			}
		}
	}

	return []int{}
}
