package comicinfo

import (
	"maps"
	"slices"
	"strings"
)

var Separator string = ","

// Split value by Separator that can be set by comicinfo.Separator.
func splitValue(str string) []string {
	list := strings.Split(str, Separator)

	result := make([]string, 0)

	for _, item := range list {
		str := strings.TrimSpace(item)

		if str == "" {
			continue
		}

		result = append(result, str)
	}

	return result
}

// Remove duplicate item from given string slice.
func removeDuplicates(items []string) []string {
	m := make(map[string]struct{})

	for _, item := range items {
		str := strings.TrimSpace(item)
		m[str] = struct{}{}
	}

	result := make([]string, 0)
	for i := range maps.Keys(m) {
		result = append(result, i)
	}

	// Sort slice to prevent race condition
	slices.Sort(result)
	return result
}

// Add values to given string, by separator.
// This function will also ensure no duplicate item appear in final string.
func AddValue(str string, values ...string) string {
	// Split original string to items
	parsed := splitValue(str)

	// Add values
	parsed = append(parsed, values...)

	// Remove duplicate values
	parsed = removeDuplicates(parsed)

	return strings.Join(parsed, Separator)
}
