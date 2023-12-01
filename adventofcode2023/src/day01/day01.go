package day01

import (
	"adventofcode2023"
	"fmt"
	"regexp"
	"strconv"
)

func GetFileInput() []string {
	return adventofcode2023.ReadData("adventofcode2023/src/day01/input.txt")
}

func RunPartOne() {
	fmt.Printf("Solution: %d", SumOfNumbers(GetFileInput()))
}

func SumOfNumbers(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += CollectNumbersInWord(line)
	}

	return sum
}

func CollectNumbersInWord(word string) int {

	var numbers []int

	// Iteration left
	digitWord := ""
	for i := 0; i < len(word); i++ {
		c := string(word[i])
		digitWord += c
		n := GetNumberFromWord(digitWord, c)
		if n > 0 {
			numbers = append(numbers, n)
			break
		}
	}

	// Iteration right
	digitWord = ""
	for i := len(word) - 1; i >= 0; i-- {
		c := string(word[i])
		digitWord += c
		swapWord := SwapWord(digitWord)
		n := GetNumberFromWord(swapWord, c)
		if n > 0 {
			numbers = append(numbers, n)
			break
		}
	}

	num, _ := strconv.Atoi(strconv.Itoa(numbers[0]) + "" + strconv.Itoa(numbers[len(numbers)-1]))

	return num
}

func SwapWord(wordToSwap string) string {
	word := ""
	for i := len(wordToSwap) - 1; i >= 0; i-- {
		word += string(wordToSwap[i])
	}
	return word
}

func GetNumberFromWord(word string, digit string) int {
	c := digit
	re := regexp.MustCompile("[0-9]")

	if len(word) > 2 {
		n := SpellWord(word)
		if n != 0 {
			c = strconv.Itoa(n)
		}
	}

	if re.MatchString(c) {
		num, err := strconv.Atoi(c)
		if err == nil {
			return num
		}
	}
	
	return 0
}
