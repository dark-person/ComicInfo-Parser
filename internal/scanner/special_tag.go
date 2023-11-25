package scanner

import "strings"

// File to Store Special Tags that will added to ComicInfo Tags.
// Structure is quite hardcoded, TODO: Change to config in future.

// The List of Special Tags that may appear in the filename
var specialTags = []string{
	"無修正", "DL版",
}

func GetSpecialTags(filename string) []string {
	finalTags := make([]string, 0)

	for _, tag := range specialTags {
		if strings.Contains(filename, tag) {
			finalTags = append(finalTags, tag)
		}
	}

	return finalTags
}
