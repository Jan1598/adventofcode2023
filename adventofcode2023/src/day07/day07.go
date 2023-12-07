package day07

import (
	"adventofcode2023/helper/src"
	"fmt"
	"sort"
)

type Day07 struct {
}

var aocDay = src.AoCDay{
	Name: "day07",
}

func (r Day07) RunPartOne() {
	fmt.Printf("Solution: %d", SumRateHandCards(aocDay.GetFileInput()))
}

func SumRateHandCards(lines []string) int {

	t1 := []CardHand{{Hand: "KTJJT", Bid: 220}, {Hand: "KK677", Bid: 18}, {Hand: "T55J5", Bid: 18}, {Hand: "QQQJA", Bid: 18}}

	sortTest(t1)

	cardHandMap := CollectHandCards(lines)

	cardCount := 0
	for i, _ := range cardHandMap {
		cardHandList := cardHandMap[i]
		cardCount += len(cardHandList)
		SortHandCardGroup(cardHandList)
	}

	sum := 0
	for i := 1; i <= 9; i++ {
		cardHandList := cardHandMap[i]
		for _, cardHand := range cardHandList {
			sum += cardHand.Bid * cardCount
			cardCount--
		}
	}

	return sum
}

func sortTest(list CardHandList) {
	sort.Sort(list)
}
