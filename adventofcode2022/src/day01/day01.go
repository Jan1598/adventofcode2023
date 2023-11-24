package day01

import (
	"adventofcode2023"
	"fmt"
	"strconv"
)

type Elf struct {
	Name     string
	Calories int
}

func RunPartOne() {
	lines := adventofcode2023.ReadData("adventofcode2022/src/day01/input.txt")

	for _, elf := range CreateElfSumList(lines) {
		fmt.Println(elf.Name + ": " + strconv.Itoa(elf.Calories))
	}
}

func CreateElfSumList(lines []string) []Elf {
	elfList := []Elf{CreateElf("Elf 0")}
	for _, line := range lines {
		listSize := len(elfList)
		if len(line) == 0 {
			elfList = append(elfList, CreateElf("Elf "+strconv.Itoa(listSize)))
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Cannot parse string to int", err)
			continue
		}
		currentElf := &elfList[listSize-1]
		currentElf.Calories = currentElf.Calories + num
	}

	return elfList
}

func CreateElf(name string) Elf {
	return Elf{
		Name:     name,
		Calories: 0,
	}
}
