package day04

import (
	"adventofcode2023/helper/src"
	"strconv"
	"strings"
)

type CardSet struct {
	Id             int
	WinningNumbers []int
	OwnNumbers     []int
	MatchCount     int
}

func CollectCardSet(lines []string) []CardSet {

	var cardSetList []CardSet

	for _, line := range lines {

		textArray := src.SplitTextWithKeyWord(line, "Card")

		if len(textArray) < 2 {
			continue
		}

		num, _ := strconv.Atoi(textArray[0])
		numberSet := strings.Split(textArray[1], "|")

		cardSet := CardSet{
			Id:             num,
			WinningNumbers: make([]int, 0),
			OwnNumbers:     make([]int, 0),
			MatchCount:     1,
		}

		winningNumbers := strings.Split(numberSet[0], " ")
		ownNumbers := strings.Split(numberSet[1], " ")

		cardSet.WinningNumbers = CollectData(cardSet.WinningNumbers, winningNumbers)
		cardSet.OwnNumbers = CollectData(cardSet.OwnNumbers, ownNumbers)

		cardSetList = append(cardSetList, cardSet)
	}

	return cardSetList
}

func CollectData(data []int, values []string) []int {
	for _, w := range values {
		num, err := strconv.Atoi(w)
		if err == nil {
			data = append(data, num)
		}
	}

	return data
}
