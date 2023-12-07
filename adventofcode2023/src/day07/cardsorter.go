package day07

func CardWeight(card string) int {
	weights := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}

	return weights[card]
}

type CardHandList []CardHand

func (s CardHandList) Len() int {
	return len(s)
}

func (s CardHandList) Less(i, j int) bool {
	return CompareWords(s[i].Hand, s[j].Hand)
}

func (s CardHandList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func CompareWords(word1, word2 string) bool {
	for i := 0; i < len(word1) && i < len(word2); i++ {
		weight1 := CardWeight(string(word1[i]))
		weight2 := CardWeight(string(word2[i]))

		if weight1 != weight2 {
			return weight1 > weight2
		}
	}

	return len(word1) > len(word2)
}
