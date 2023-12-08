package day06

import (
	"adventofcode2023/helper/src"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func CollectRaces(lines []string) []Race {

	var values [][]int

	for i, line := range lines {
		textArray := src.SplitTextWithKeyWord(line, "Time")
		if i == 1 {
			textArray = src.SplitTextWithKeyWord(line, "Distance")
		}

		if len(textArray) < 2 {
			continue
		}

		lineValues := strings.Split(textArray[1], " ")

		var numbers []int
		for _, v := range lineValues {
			num, err := strconv.Atoi(strings.Trim(v, " "))

			if err == nil {
				numbers = append(numbers, num)
			}
		}
		values = append(values, numbers)
	}

	var result []Race
	for i := 0; i < len(values[0]); i++ {
		race := Race{
			Time:     values[0][i],
			Distance: values[1][i],
		}
		result = append(result, race)
	}

	return result
}

func CalculateWinRatios(race Race) []int {

	var ratios []int

	time := race.Time
	recordDistance := race.Distance

	currentSpeed := 0

	for i := 0; i < time; i++ {
		currentSpeed++
		maxDistanceInTime := currentSpeed * (time - currentSpeed)

		if maxDistanceInTime > recordDistance {
			ratios = append(ratios, currentSpeed)
		}
	}

	return ratios
}
