package src

type Coordinate struct {
	XValue int
	YValue int
}

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

func (coordinate *Coordinate) MoveCoordinate(matrix [][]string, m string) {

	switch m {
	case UP:
		coordinate.YValue = coordinate.YValue - 1
	case RIGHT:
		coordinate.XValue = coordinate.XValue + 1
	case DOWN:
		coordinate.YValue = coordinate.YValue + 1
	case LEFT:
		coordinate.XValue = coordinate.XValue - 1
	}

	coordinate.UpdateCoordinate(matrix)
}

func (coordinate *Coordinate) GetValueByCoordinate(matrix [][]string) string {
	coordinate.UpdateCoordinate(matrix)

	return matrix[coordinate.YValue][coordinate.XValue]
}

const (
	UP    = "UP"
	RIGHT = "RIGHT"
	DOWN  = "DOWN"
	LEFT  = "LEFT"
)
