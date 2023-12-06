package day05

import (
	"adventofcode2023/helper/src"
	"fmt"
	"sort"
)

type Day05 struct {
}

var aocDay = src.AoCDay{
	Name: "day05",
}

var gardenValues = []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

func (r Day05) RunPartOne() {
	fmt.Printf("Solution: %d", GetLocationValues(aocDay.GetFileInput())[0])
}

func (r Day05) RunPartTwo() {
	fmt.Printf("Solution: %d", 1)
}

func GetLocationValues(lines []string) []int {
	seeds := CreateSeeds(lines[0])
	seedMap := CollectMaps(lines)

	for _, key := range gardenValues {
		seeds = GetNumberFromSeeds(seeds, CreateValueMap(seedMap[key]))
	}

	sort.Ints(seeds)

	return seeds
}
