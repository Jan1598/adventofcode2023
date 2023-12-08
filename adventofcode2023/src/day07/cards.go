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

	if fiveOfAKind {
		PersistsCardHandValue(cardHandMap, 1, cardHand)
	} else if fourOfAKind {
		PersistsCardHandValue(cardHandMap, 2, cardHand)
	} else if fullHouse {
		PersistsCardHandValue(cardHandMap, 3, cardHand)
	} else if threeOfAKind {
		PersistsCardHandValue(cardHandMap, 4, cardHand)
	} else if twoPair {
		PersistsCardHandValue(cardHandMap, 5, cardHand)
	} else if onePair {
		PersistsCardHandValue(cardHandMap, 6, cardHand)
	} else {
		PersistsCardHandValue(cardHandMap, 7, cardHand)
	}
}

func SortHandCardGroup(handCards CardHandList) {
	sort.Sort(handCards)
}

func CheckKind(hand string, kindAmount, amount int) bool {
	letterMap := CreateLetterMap(hand)

	jokerAmount, _ := letterMap["J"]

	if jokerAmount > 0 {
		if kindAmount == 5 && amount == 1 {
			if len(letterMap) == 2 && jokerAmount > 0 {
				return true
			}
		} else if kindAmount == 4 && amount == 1 {
			if len(letterMap) == 3 && jokerAmount > 1 {
				return true
			} else if len(letterMap) == 3 && jokerAmount > 0 {
				for key, value := range letterMap {
					if key != "J" && value == 3 {
						return true
					}
				}
			}
		} else if kindAmount == 3 && amount == 1 {
			if len(letterMap) == 4 && jokerAmount > 0 {
				return true
			} else if len(letterMap) == 3 && jokerAmount > 1 {
				return true
			}
		} else if kindAmount == 2 && amount == 1 {
			if len(letterMap) == 3 && jokerAmount > 0 {
				return true
			} else if len(letterMap) == 4 && jokerAmount > 0 {
				return true
			} else if len(letterMap) == 5 && jokerAmount > 0 {
				return true
			}
		} else if kindAmount == 2 && amount == 2 {
			if len(letterMap) == 3 && jokerAmount == 2 {
				return true
			} else if len(letterMap) == 4 && jokerAmount > 1 {
				return true
			}
		}

	}

	i := 0
	for _, value := range letterMap {
		if value == kindAmount {
			i++
			if i == amount {
				return true
			}
		}
	}

	return false
}

func CheckFullHouse(hand string) bool {

	letterMap := CreateLetterMap(hand)

	jokerAmount, _ := letterMap["J"]

	if len(letterMap) > 2 && jokerAmount == 0 || len(letterMap) > 3 {
		return false
	}

	threeOfKind := ""
	twoOfKind := ""
	for key, value := range letterMap {
		if value == 3 || (value == 2 && len(threeOfKind) <= 0 && jokerAmount > 0) {
			threeOfKind = key
		} else if value == 2 {
			twoOfKind = key
		}
	}

	if len(threeOfKind) <= 0 || len(twoOfKind) <= 0 {
		return false
	}

	return true
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
