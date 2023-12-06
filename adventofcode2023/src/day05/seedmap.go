package day05

import (
	"adventofcode2023/helper/src"
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

func FillValueMap(input SeedMapInput, index int, valueMap map[int]int) {

	rangeValue := input.RangeLength
	destinationValue := input.DestinationRange

	a := destinationValue
	for i := index; i < index+rangeValue; i++ {
		valueMap[i] = a
		a++
	}
}

func CreateValueMap(input []SeedMapInput) map[int]int {

	valueMap := make(map[int]int)

	var filledSeedInput []SeedMapInput

	for i := 0; i < input[0].SourceRange+input[0].RangeLength-1; i++ {
		for _, seedInput := range input {
			sourceValue := seedInput.SourceRange
			if i == sourceValue {
				FillValueMap(seedInput, i, valueMap)
				filledSeedInput = append(filledSeedInput, seedInput)
			}
		}
	}

	if len(input) > len(filledSeedInput) {
	outerLoop:
		for _, seedInput := range input {
			for _, filledSeed := range filledSeedInput {
				if seedInput == filledSeed {
					continue outerLoop
				}
			}
			FillValueMap(seedInput, seedInput.SourceRange, valueMap)
		}
	}

	return valueMap
}

func GetNumberFromSeeds(seeds []int, valueMap map[int]int) []int {

	var result []int
	for _, i := range seeds {
		value, ok := valueMap[i]

		if ok {
			result = append(result, value)
		} else {
			result = append(result, i)
		}
	}

	return result
}
