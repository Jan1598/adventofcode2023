package day06

import (
	"adventofcode2023/helper/src"
	"fmt"
	"strconv"
)

type Day06 struct {
}

var aocDay = src.AoCDay{
	Name: "day06",
}

func (r Day06) RunPartOne() {
	fmt.Printf("Solution: %d", CalculateWinRatiosOfRaces(aocDay.GetFileInput()))
}

func (r Day06) RunPartTwo() {
	fmt.Printf("Solution: %d", len(CalculateWinRatios(ParseRacesToOne(CollectRaces(aocDay.GetFileInput())))))
}

func CalculateWinRatiosOfRaces(lines []string) int {

	sum := 1
	races := CollectRaces(lines)

	for _, race := range races {
		sum *= len(CalculateWinRatios(race))
	}

	return sum
}

func ParseRacesToOne(races []Race) Race {

	timeText := "0"
	distanceText := "0"
	for _, race := range races {
		timeText += strconv.Itoa(race.Time)
		distanceText += strconv.Itoa(race.Distance)
	}

	numTime, _ := strconv.Atoi(timeText)
	distanceTime, _ := strconv.Atoi(distanceText)

	return Race{
		Time:     numTime,
		Distance: distanceTime,
	}
}
