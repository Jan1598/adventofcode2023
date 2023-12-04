package test

import (
	"adventofcode2023"
	"adventofcode2023/helper/src"
	"testing"
)

func TestSplitTextWithKeyWord(t *testing.T) {
	output := src.SplitTextWithKeyWord("Game 1: test test2 asd", "Game")

	adventofcode2023.AssertEquals(t, "1", output[0])
	adventofcode2023.AssertEquals(t, "test test2 asd", output[1])

	output = src.SplitTextWithKeyWord("asd 1123: asd23fg asd | 23", "asd")

	adventofcode2023.AssertEquals(t, "1123", output[0])
	adventofcode2023.AssertEquals(t, "asd23fg asd | 23", output[1])
}
