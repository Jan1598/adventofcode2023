package day04

import (
	"adventofcode2023/helper/src"
	"fmt"
	"math"
)

type Day04 struct {
}

var aocDay = src.AoCDay{
	Name: "day04",
}

func (r Day04) RunPartOne() {
	fmt.Printf("Solution: %d", SumOfMatches(aocDay.GetFileInput()))
}

func (r Day04) RunPartTwo() {
	fmt.Printf("Solution: %d", SumOfScratchCards(aocDay.GetFileInput()))
}

func SumOfMatches(lines []string) int {
	sum := 0

	cardSet := CollectCardSet(lines)

	for _, card := range cardSet {

		winningNumbersAmount := 0
		winSum := 0
		for _, w := range card.WinningNumbers {
			for _, o := range card.OwnNumbers {
				if w == o {
					winSum = int(math.Pow(2, float64(winningNumbersAmount)))
					winningNumbersAmount++
					break
				}
			}
		}
		sum += winSum
	}

	return sum
}

func SumOfScratchCards(lines []string) int {
	sum := 0

	cardSet := CollectCardSet(lines)

	for i, card := range cardSet {

		winningNumbersAmount := 0
		for _, w := range card.WinningNumbers {
			for _, o := range card.OwnNumbers {
				if w == o {
					winningNumbersAmount++
					break
				}
			}
		}

		currentCount := card.MatchCount
		for w := 1; w <= winningNumbersAmount; w++ {
			cardSet[i+w].MatchCount += currentCount
		}

	}

	for _, card := range cardSet {
		sum += card.MatchCount
	}

	return sum
}
