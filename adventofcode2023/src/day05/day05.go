package day05

import (
	"adventofcode2023/helper/src"
	"fmt"
	"sort"
	"sync"
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
	fmt.Printf("Solution: %d", GetLowestLocationValueAsync(aocDay.GetFileInput()))
}

func GetLocationValues(lines []string) []int {
	seeds := CreateSeeds(lines[0])
	seedMap := CollectMaps(lines)

	for _, key := range gardenValues {
		seeds = GetNumberFromSeedsByMapInput(seeds, seedMap[key])
	}

	sort.Ints(seeds)

	return seeds
}

func GetLowestLocationValue(lines []string) int {
	seeds := CreateSeeds(lines[0])
	seedMap := CollectMaps(lines)

	return GetLowestLocationByMapInput(seeds, seedMap)
}

func GetLowestLocationValueAsync(lines []string) int {
	seeds := CreateSeeds(lines[0])
	seedMap := CollectMaps(lines)

	var result []int

	var wg sync.WaitGroup
	resultChan := make(chan int, len(seeds)/2)

	for i := 0; i < len(seeds); i += 2 {

		seedStart := seeds[i]
		seedRange := seeds[i+1]

		wg.Add(1)
		go RunGetLowest([]int{seedStart, seedRange}, seedMap, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	for lowest := range resultChan {
		if lowest > 0 {
			result = append(result, lowest)
		}
	}

	sort.Ints(result)

	return result[0]
}

func RunGetLowest(seeds []int, seedMap map[string][]SeedMapInput, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	resultChan <- GetLowestLocationByMapInput(seeds, seedMap)
}
