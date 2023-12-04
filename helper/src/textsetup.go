package src

import (
	"strings"
)

// SplitTextWithKeyWord
// Split a text with format <Text> <ID>: <Content> and returns string array.
// First index is the id.
// Last index the content.
func SplitTextWithKeyWord(input, key string) []string {

	if !strings.Contains(input, key) || !strings.Contains(input, ":") {
		return []string{}
	}

	cubeText := strings.Split(input, ":")
	if len(cubeText) <= 1 {
		return []string{}
	}
	idText := strings.Trim(strings.Replace(cubeText[0], key, "", 1), " ")
	content := strings.Trim(cubeText[1], " ")

	return []string{idText, content}
}
