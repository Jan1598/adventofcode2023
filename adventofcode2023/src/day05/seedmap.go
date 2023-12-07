package day05

import (
	"adventofcode2023/helper/src"
	"fmt"
	"strconv"
	"strings"
)

type SeedMap struct {
	Name  string
	Input []SeedMapInput
}

type SeedMapInput struct {
	DestinationRange int
	SourceRange      int
	RangeLength      int
}

func CollectMaps(lines []string) map[string][]SeedMapInput {

	seedMap := make(map[string][]SeedMapInput)

	currentKey := ""
	for _, line := range lines {
		if strings.Contains(line, "map:") {
			currentKey = strings.Trim(strings.Replace(line, "map:", "", 1), " ")
			continue
		}

		if len(currentKey) <= 0 || len(line) <= 0 {
			continue
		}

		numbers := src.MapNumberStringArrayToIntArray(strings.Split(line, " "))

		seedMapInput := SeedMapInput{
			DestinationRange: numbers[0],
			SourceRange:      numbers[1],
			RangeLength:      numbers[2],
		}

		PersistsValueToSeedInput(seedMap, seedMapInput, currentKey)
	}

	return seedMap
}

func CreateSeeds(line string) []int {
	output := src.SplitTextWithKeyWord(line, "")

	if len(output) < 2 {
		return []int{}
	}

	var seedNumbers []int
	numbers := strings.Split(output[1], " ")
	for _, n := range numbers {

		num, err := strconv.Atoi(n)
		if err == nil {
			seedNumbers = append(seedNumbers, num)
		}
	}

	return seedNumbers
}

func PersistsValueToSeedInput(seedMap map[string][]SeedMapInput, seedMapInput SeedMapInput, key string) {

	value, ok := seedMap[key]

	if ok {
		value = append(value, seedMapInput)
	} else {
		value = []SeedMapInput{seedMapInput}
	}
	seedMap[key] = value
}

func GetNumberFromSeedsByMapInput(seeds []int, input []SeedMapInput) []int {
	var ignoredIndexes []int
	for _, seedInput := range input {
		sourceValue := seedInput.SourceRange
		rangeValue := seedInput.RangeLength
		destinationValue := seedInput.DestinationRange
		for i, seed := range seeds {
			if SliceIncludeValue(ignoredIndexes, i) {
				if len(ignoredIndexes) >= len(seeds) {
					return seeds
				}
				continue
			} else if seed >= sourceValue && seed < sourceValue+rangeValue {
				diff := seed - sourceValue
				seeds[i] = destinationValue + diff
				ignoredIndexes = append(ignoredIndexes, i)
			}
		}
	}

	return seeds
}

func SliceIncludeValue(valueSlice []int, value int) bool {
	for _, v := range valueSlice {
		if v == value {
			return true
		}
	}
	return false
}

// 57451709
var minValue = 0

func GetLowestLocationByMapInput(seeds []int, seedMap map[string][]SeedMapInput) int {
	var ignoredIndexes []int

	input := seedMap[gardenValues[0]]

	minSeedValue := 0
	for _, seedInput := range input {
		sourceValue := seedInput.SourceRange
		rangeValue := seedInput.RangeLength
		destinationValue := seedInput.DestinationRange

		for i := 0; i < len(seeds); i += 2 {

			seedStart := seeds[i]
			seedRange := seeds[i+1]

			if (seedStart+seedRange < sourceValue || seedStart > sourceValue+rangeValue) || SliceIncludeValue(ignoredIndexes, i) {
				if len(ignoredIndexes) >= len(seeds) {
					return minValue
				}
			} else {

				startValue := seedStart

				if seedStart < sourceValue {
					startValue = sourceValue
				}

				if minSeedValue > startValue {
					startValue = minSeedValue
				}

				for i2 := startValue; i2 < seedStart+seedRange; i2++ {

					if i2 < minSeedValue && minSeedValue > 0 {
						break
					}

					if i2 >= sourceValue && i2 < sourceValue+rangeValue {
						seedValue := destinationValue + (i2 - sourceValue)
						for _, key := range gardenValues {
							seedValue = GetNumberFromSeedsByMapInput([]int{seedValue}, seedMap[key])[0]
							if key == gardenValues[len(gardenValues)-1] && (seedValue < minValue || minValue == 0) {
								minValue = seedValue
								minSeedValue = i2
								ignoredIndexes = append(ignoredIndexes, i)
								fmt.Println(minValue)
							}
						}
					}
				}
			}

		}
	}

	return minValue
}
