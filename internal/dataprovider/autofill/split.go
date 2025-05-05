package autofill

import "strings"

// Replace brackets to space from given input.
func replaceBrackets(input string) string {
	// Set of brackets or character to be replace
	bracketsToReplace := []rune{'[', ']', '{', '}', '(', ')'}

	// Create a set of brackets to remove
	brackets := make(map[rune]bool)
	for _, b := range bracketsToReplace {
		brackets[b] = true
	}

	// Replace brackets with space character
	replaceFunc := func(r rune) rune {
		if brackets[r] {
			return ' '
		}
		return r
	}

	// Filter out brackets
	return strings.Map(replaceFunc, input)
}

// Replace bracket and split given bookname to multiple keywords
// by space character.
//
// Split to space character, is to get best performance by minimum effort.
func SplitKeywords(bookName string) []string {
	// Remove brackets
	tmp := replaceBrackets(bookName)

	// Split words
	words := strings.Split(tmp, " ")

	// Loop words to trim space and ignore empty string
	result := make([]string, 0)

	for _, item := range words {
		s := strings.TrimSpace(item)

		// Skip word that is empty string
		if s == "" {
			continue
		}

		result = append(result, s)
	}

	return result
}
