package day01

import "strings"

func SpellWord(word string) int {

	switch {
	case strings.Contains(word, "one"):
		return 1
	case strings.Contains(word, "two"):
		return 2
	case strings.Contains(word, "three"):
		return 3
	case strings.Contains(word, "four"):
		return 4
	case strings.Contains(word, "five"):
		return 5
	case strings.Contains(word, "six"):
		return 6
	case strings.Contains(word, "seven"):
		return 7
	case strings.Contains(word, "eight"):
		return 8
	case strings.Contains(word, "nine"):
		return 9
	}

	return 0
}
