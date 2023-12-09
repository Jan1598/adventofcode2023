package day09

import (
	"strconv"
	"strings"
)

func CreateNumberMap(lines []string, reverse bool) [][]int {

	var numberList [][]int

	for _, line := range lines {
		output := strings.Split(line, " ")

		if len(output) < 2 {
			continue
		}

		var numberInput []int

		if reverse {
			for i := len(output) - 1; i >= 0; i-- {
				num, err := strconv.Atoi(strings.Trim(output[i], " "))
				if err == nil {
					numberInput = append(numberInput, num)
				}
			}
		} else {
			for _, v := range output {
				num, err := strconv.Atoi(strings.Trim(v, " "))
				if err == nil {
					numberInput = append(numberInput, num)
				}
			}
		}

		numberList = append(numberList, numberInput)
	}

	return numberList
}

func CreateNextMap(heapMap [][]int) [][]int {
	numberMap := heapMap[len(heapMap)-1]
	var nextMap []int
	for i, _ := range numberMap {
		nextIndex := i + 1
		if nextIndex < len(numberMap) {
			nextMap = append(nextMap, numberMap[nextIndex]-numberMap[i])
		}
	}

	heapMap = append(heapMap, nextMap)
	if !IsMapEndReached(nextMap) {
		heapMap = CreateNextMap(heapMap)
	}

	return heapMap
}

func IsMapEndReached(numberMap []int) bool {
	for _, v := range numberMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func GetNextValueFromMap(heapMap [][]int) int {
	value := 0

	for i := len(heapMap) - 1; i >= 0; i-- {
		numberMap := heapMap[i]
		value += numberMap[len(numberMap)-1]
	}

	return value
}
