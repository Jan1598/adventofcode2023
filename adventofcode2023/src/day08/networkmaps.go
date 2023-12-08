package day08

import (
	"adventofcode2023/helper/src"
	"strings"
	"sync"
)

func GetInstructionByInput(lines []string) string {
	instruction := ""
	for _, line := range lines {
		if len(line) == 0 {
			return instruction
		}

		instruction += line
	}

	return instruction
}

func CreateNetworkMap(lines []string) map[string][]string {

	networkMap := make(map[string][]string)

	for _, line := range lines {
		output := src.SplitTextByKeyWord(line, "=")

		if len(output) < 2 {
			continue
		}

		keyValue := strings.Trim(output[0], " ")
		valuesOfKey := strings.Replace(output[1], " ", "", -1)

		if strings.Contains(valuesOfKey, "(") {
			valuesOfKey = strings.Trim(strings.Replace(valuesOfKey, "(", "", 1), " ")
		}
		if strings.Contains(valuesOfKey, ")") {
			valuesOfKey = strings.Trim(strings.Replace(valuesOfKey, ")", "", 1), " ")
		}

		networkMap[keyValue] = strings.Split(valuesOfKey, ",")
	}

	return networkMap
}

func ReachValueInMap(networkMap map[string][]string, instruction string) int {
	step := 0

	currentKey := "AAA"
	for currentKey != "ZZZ" {
		for _, c := range instruction {
			step++
			in := string(c)

			if in == "R" {
				currentKey = networkMap[currentKey][1]
			} else if in == "L" {
				currentKey = networkMap[currentKey][0]
			}
		}
	}

	return step
}

func ReachValueInMapSimultaneously(networkMap map[string][]string, instruction string, limit int) int {
	startValues := CollectAllStartPoints(networkMap)
	var wg sync.WaitGroup
	resultChan := make(chan []string, len(startValues))

	for _, startValue := range startValues {
		wg.Add(1)
		go RunValueMap(networkMap, startValue, instruction, limit, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	return CheckList(resultChan)
}

func CheckList(resultChan chan []string) int {
	var collectList [][]string
	i := 0
	for inputList := range resultChan {
		collectList = append(collectList, inputList)
		i++
	}

	for i2, v := range collectList[0] {
		if strings.LastIndex(v, "Z") == 2 {
			zIndex := i2
			end := true
			for i3 := 1; i3 < len(collectList); i3++ {
				inputList := collectList[i3]
				if zIndex > len(inputList) {
					return 0
				}
				if strings.LastIndex(inputList[zIndex], "Z") != 2 {
					end = false
					break
				}
			}
			if end {
				return i2 + 1
			}
		}
	}

	return 0
}

func CheckList2(resultChan chan []string) int {

	step := 0
	for inputList := range resultChan {
		for i, v := range inputList {
			if strings.LastIndex(v, "Z") == 2 {
				step += i + 1
				break
			}
		}
	}

	return step
}

func RunValueMap(networkMap map[string][]string, currentKey, instruction string, limit int, resultChan chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	var valueList []string

	i := 0
	for i < limit {
		for _, c := range instruction {
			in := string(c)

			if in == "R" {
				currentKey = networkMap[currentKey][1]
			} else if in == "L" {
				currentKey = networkMap[currentKey][0]
			}

			valueList = append(valueList, currentKey)
		}
		i++
	}

	resultChan <- valueList
}

func CollectAllStartPoints(networkMap map[string][]string) []string {
	var startValues []string
	for key, _ := range networkMap {
		if strings.LastIndex(key, "A") == 2 {
			startValues = append(startValues, key)
		}
	}
	return startValues
}
