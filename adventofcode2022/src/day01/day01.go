package day01

import (
	"adventofcode2023"
	"fmt"
	"sort"
	"strconv"
)

func GetFileInput() []string {
	return adventofcode2023.ReadData("adventofcode2022/src/day01/input.txt")
}

func RunPartOne() {

	maxValue := 0
	for _, elf := range CreateElfSumList(GetFileInput()) {
		if elf.Calories > maxValue {
			maxValue = elf.Calories
		}
		fmt.Println(elf.Name + ": " + strconv.Itoa(elf.Calories))
	}

	fmt.Println("Max. Value: " + strconv.Itoa(maxValue))
}

func RunPartTwo() {
	fmt.Println("Top 3 Calories:", SumOfCaloriesOfTop(3, CreateElfSumList(GetFileInput())))
}

func CreateElfSumList(lines []string) ElfList {
	elfList := ElfList{CreateElf("Elf 0")}
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

func SumOfCaloriesOfTop(topValue int, elfList ElfList) int {
	amount := 0
	if amount > len(elfList) {
		return amount
	}
	sort.Sort(elfList)
	for i := 0; i < topValue; i++ {
		amount += elfList[i].Calories
	}
	return amount
}
