package archive

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Create ZIP File inside folderToAdd.
func CreateZip(folderToAdd string) (dest string, err error) {
	return CreateZipTo(folderToAdd, folderToAdd)
}

// Create ZIP File of inputDir, and output the zip to destDir.
//
// This function is a variant of CreateZip(), purpose to provide flexibility.
func CreateZipTo(inputDir string, destDir string) (dest string, err error) {
	destFileName := filepath.Base(inputDir)

	// Create ZIP File
	destFile, err := os.Create(filepath.Join(destDir, destFileName+".zip"))
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	// Zip Writer
	destZip := zip.NewWriter(destFile)
	defer destZip.Close()

	// Load File Entries inside folderToAdd
	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return "", err
	}

	// Loop File inside folderToAdd
	for _, entry := range entries {
		filename := entry.Name()

		// Skip Zip file
		if filename == destFileName+".zip" {
			continue
		}

		// Create File inside zip
		zipFile, err := destZip.Create(filename)
		if err != nil {
			return "", err
		}

		// Open Actual File
		file, err := os.Open(filepath.Join(inputDir, entry.Name()))
		if err != nil {
			return "", err
		}
		defer file.Close()

		// Copy file content
		_, err = io.Copy(zipFile, file)
		if err != nil {
			return "", err
		}
	}

	return filepath.Join(destDir, destFileName+".zip"), nil
}
