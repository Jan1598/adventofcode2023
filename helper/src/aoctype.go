package src

import "adventofcode2023"

type AoCDayInterface interface {
	RunPartOne()
	RunPartTwo()
}

type AoCDay struct {
	Name string
}

func (r AoCDay) GetFileInput() []string {
	return adventofcode2023.ReadData("adventofcode2023/src/" + r.Name + "/input.txt")
}
