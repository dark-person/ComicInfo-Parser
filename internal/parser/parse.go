package parser

import (
	"path/filepath"
	"regexp"
	"strings"
)

// Get the filename, without the file extension.
//
// The Filename should be Base, not containing any directory inside.
func FilenameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// Parse the zip filename to multiple string for comicInfo.xml
// Normally the filename will be (C102) [Author] Name....  .zip
func Parse(filename string) (market, author, bookName string) {
	// Get Filename without extension

	name := FilenameWithoutExt(filename)
	if name[0:1] == "(" {
		// Contains C102 e.g
		re := regexp.MustCompile(`\(([^\)]*)\)?\s?\[([^\]]*)\]{1}(.*)`)
		matches := re.FindStringSubmatch(filename)

		market = strings.TrimSpace(matches[1])
		author = strings.TrimSpace(matches[2])
		bookName = strings.TrimSpace(matches[3])
	} else {
		re := regexp.MustCompile(`\[([^\]]*)\]{1}(.*)`)
		matches := re.FindStringSubmatch(name)

		market = ""
		author = strings.TrimSpace(matches[1])
		bookName = strings.TrimSpace(matches[2])
	}

	return market, author, bookName
}
