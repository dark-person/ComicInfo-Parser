package fsprov

import (
	"regexp"
	"strings"
)

// Parse the folder name to multiple string for comicInfo.xml.
// The folder name MUST be base folder name, i.e. not include its parent directory
//
// Normally the folder name will be (C102) [Author] Name....
func parseFolder(foldername string) (market, author, bookName string) {
	if foldername[0:1] == "(" {
		// Contains C102 e.g
		re := regexp.MustCompile(`\(([^\)]*)\)?\s?\[([^\]]*)\]{1}(.*)`)
		matches := re.FindStringSubmatch(foldername)

		if len(matches) == 0 {
			// Filename not parse, abort market & author recognize
			return "", "", foldername
		}

		market = strings.TrimSpace(matches[1])
		author = strings.TrimSpace(matches[2])
		bookName = strings.TrimSpace(matches[3])
	} else {
		re := regexp.MustCompile(`\[([^\]]*)\]{1}(.*)`)
		matches := re.FindStringSubmatch(foldername)

		if len(matches) == 0 {
			// Filename not parse, abort market & author recognize
			return "", "", foldername
		}

		market = ""
		author = strings.TrimSpace(matches[1])
		bookName = strings.TrimSpace(matches[2])
	}

	return market, author, bookName
}
