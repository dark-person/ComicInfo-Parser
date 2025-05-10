package fsprov

import (
	"fmt"
	"os"

	"github.com/dark-person/comicinfo-parser/internal/files"
)

// Check the folder fulfill requirement of the given Scanner Options
func CheckFolder(folderPath string, opt ScanOpt) (bool, error) {
	if !opt.Valid() {
		return false, fmt.Errorf("invalid scan options")
	}

	// Get all file/folder in given path
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return false, err
	}

	// Prepare variable
	subfolderCount := 0
	imageCount := 0
	totalCount := 0

	// Loop the entries
	for _, entry := range entries {
		totalCount++

		// Directory Check
		if entry.IsDir() {
			subfolderCount++
			continue
		}

		// Image Extension check
		if files.IsSupportedImg(entry.Name()) {
			imageCount++
			continue
		}
	}

	// Check Contain Only Option
	if opt.Image == AllowOnly && (totalCount != imageCount) {
		return false, nil
	}

	if opt.SubFolder == AllowOnly && (totalCount != subfolderCount) {
		return false, nil
	}

	// Check Reject Option
	if opt.Image == Reject && imageCount > 0 {
		return false, nil
	}

	if opt.SubFolder == Reject && subfolderCount > 0 {
		return false, nil
	}

	// Check Contain Option
	if (opt.Image == Allow || opt.Image == AllowOnly) && imageCount <= 0 {
		return false, nil
	}

	if (opt.SubFolder == Allow || opt.SubFolder == AllowOnly) && subfolderCount <= 0 {
		return false, nil
	}

	// All Checking Passed
	return true, nil
}
