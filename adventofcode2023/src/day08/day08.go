package day08

import (
	"adventofcode2023/helper/src"
	"fmt"
)

type Day08 struct {
}

var aocDay = src.AoCDay{
	Name: "day08",
}

func (r Day08) RunPartOne() {
	lines := aocDay.GetFileInput()
	fmt.Printf("Solution: %d", ReachValueInMap(CreateNetworkMap(lines), GetInstructionByInput(lines)))
}

func (r Day08) RunPartTwo() {
	lines := aocDay.GetFileInput()
	fmt.Printf("Solution: %d", ReachValueInMapSimultaneously(CreateNetworkMap(lines), GetInstructionByInput(lines), 100))
}
