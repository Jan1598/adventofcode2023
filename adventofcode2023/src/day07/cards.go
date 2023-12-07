package day07

import (
	"sort"
	"strconv"
	"strings"
)

type CardHand struct {
	Hand string
	Bid  int
}

func CollectHandCards(lines []string) map[int][]CardHand {

	cardHandMap := make(map[int][]CardHand)

	for _, line := range lines {
		output := strings.Split(line, " ")

		if len(output) < 2 {
			return cardHandMap
		}

		hand := strings.Trim(output[0], " ")
		num, err := strconv.Atoi(strings.Trim(output[1], " "))
		if err != nil {
			continue
		}
		cardHand := CardHand{
			Hand: hand,
			Bid:  num,
		}

		PersistHandCard(cardHandMap, cardHand)
	}

	return cardHandMap
}

func PersistHandCard(cardHandMap map[int][]CardHand, cardHand CardHand) {

	hand := cardHand.Hand

	fiveOfAKind := CheckKind(hand, 5, 1)
	fourOfAKind := CheckKind(hand, 4, 1)
	fullHouse := CheckFullHouse(hand)
	threeOfAKind := CheckKind(hand, 3, 1)
	twoPair := CheckKind(hand, 2, 2)
	onePair := CheckKind(hand, 2, 1)

	if len(fiveOfAKind) > 0 {
		PersistsCardHandValue(cardHandMap, 1, cardHand)
	} else if len(fourOfAKind) > 0 {
		PersistsCardHandValue(cardHandMap, 2, cardHand)
	} else if len(fullHouse) > 0 {
		PersistsCardHandValue(cardHandMap, 3, cardHand)
	} else if len(threeOfAKind) > 0 {
		PersistsCardHandValue(cardHandMap, 4, cardHand)
	} else if len(twoPair) > 0 {
		PersistsCardHandValue(cardHandMap, 5, cardHand)
	} else if len(onePair) > 0 {
		PersistsCardHandValue(cardHandMap, 6, cardHand)
	} else {
		PersistsCardHandValue(cardHandMap, 7, cardHand)
	}
}

func SortHandCardGroup(handCards CardHandList) {
	sort.Sort(handCards)
}

func CheckKind(hand string, kindAmount, amount int) []string {
	letterMap := CreateLetterMap(hand)

	var result []string
	i := 0
	for key, value := range letterMap {
		if value == kindAmount {
			i++
			result = append(result, key)
			if i == amount {
				return result
			}
		}
	}

	return []string{}
}

func CheckFullHouse(hand string) []string {

	letterMap := CreateLetterMap(hand)

	if len(letterMap) > 2 {
		return []string{}
	}

	threeOfKind := ""
	twoOfKind := ""
	for key, value := range letterMap {
		if value == 3 {
			threeOfKind = key
		} else if value == 2 {
			twoOfKind = key
		}
	}

	if len(threeOfKind) <= 0 {
		return []string{}
	}

	return []string{threeOfKind, twoOfKind}
}

func CreateLetterMap(word string) map[string]int {
	letterMap := make(map[string]int)

	for _, c := range word {
		cc := string(c)
		value, ok := letterMap[cc]

		if ok {
			value += 1
		} else {
			value = 1
		}
		letterMap[cc] = value
	}
	return letterMap
}

func PersistsCardHandValue(groupMap map[int][]CardHand, key int, input CardHand) {

	value, ok := groupMap[key]

	if ok {
		value = append(value, input)
	} else {
		value = []CardHand{input}
	}
	groupMap[key] = value
}
