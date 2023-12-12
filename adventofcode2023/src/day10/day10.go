package day10

import (
	"adventofcode2023/helper/src"
	"fmt"
)

type Day10 struct {
}

var aocDay = src.AoCDay{
	Name: "day10",
}

func (r Day10) RunPartOne() {
	fmt.Printf("Solution: %d", CalcRoadSteps(aocDay.GetFileInput(), "DOWN"))
}
