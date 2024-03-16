package comicinfo

// This file provides static methods for
//   1. Generate comicinfo struct from file
//   2. Generate ComicInfo.xml from struct

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Load XML file to comicInfo struct in go.
//
// path MUST has base of "ComicInfo.xml", or it will return error for invalid XML.
func Load(path string) (*ComicInfo, error) {
	// Path checking to ensure last element is "ComicInfo.xml"
	if filepath.Base(path) != "ComicInfo.xml" {
		return nil, fmt.Errorf("invalid ComicInfo.xml file")
	}

	// Open XML file
	file, err := os.Open(path)
	if err != nil {
		logrus.Errorf("Error when opening file %s: %v", path, err)
		return nil, err
	}
	defer file.Close()

	// Read byte[] from file
	data, err := io.ReadAll(file)
	if err != nil {
		logrus.Errorf("Error when read file %s: %v", path, err)
		return nil, err
	}

	// Unmarshal XML
	result := New()
	err = xml.Unmarshal(data, &result)
	if err != nil {
		logrus.Errorf("Error when unmarshal to XML file %s: %v", path, err)
		return nil, err
	}

	return &result, nil
}

func Save(info *ComicInfo, path string) error {
	panic("Save is not implemented")
}
