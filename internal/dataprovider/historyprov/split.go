package historyprov

import (
	"path/filepath"
	"strings"
)

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
func splitKeywords(bookName string) []string {
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

// Parse relative part of folder path to keywords, allow trigger to work.
//
// An example is like this:
//   - root path: /c/user/home/input/
//   - current path: /c/user/home/input/category1/category2/book1
//
// Then this function should return:
//   - result: [category1, category2]
//
// If no word can be loaded from this function, empty string slice will return.
func relativeParentKeywords(rootPath, currentPath string) []string {
	// Do nothing if both root path or current path is empty
	if rootPath == "" || currentPath == "" {
		return []string{}
	}

	root := filepath.Clean(rootPath)
	currentFolder := filepath.Dir(filepath.Clean(currentPath))

	// Try to get relative path from root & current folder
	rel, err := filepath.Rel(root, currentFolder)
	if err != nil || rel == "." || strings.HasPrefix(rel, "..") {
		return []string{}
	}

	// Rel already uses the native separator; Fields collapses any residual whitespace.
	return strings.Fields(strings.ReplaceAll(rel, string(filepath.Separator), " "))
}
