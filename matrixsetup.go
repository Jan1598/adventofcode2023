package adventofcode2023

func BuildStringMatrix(lines []string) [][]string {
	var matrix [][]string
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		var matrixInput []string
		for i2 := 0; i2 < len(line); i2++ {
			matrixInput = append(matrixInput, string(line[i2]))
		}
		matrix = append(matrix, matrixInput)
	}
	return matrix
}

func IsIndexOnEdge(matrix [][]string, indexColumn, indexRow int) bool {
	if indexColumn == 0 || (len(matrix)-1) == indexColumn {
		return true
	}

	return indexRow == 0 || len(matrix[indexColumn])-1 == indexRow
}
