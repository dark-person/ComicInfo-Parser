package archive

import (
	"archive/zip"
	"gui-comicinfo/internal/parser"
	"io"
	"os"
	"path/filepath"
)

// Create ZIP File inside folderToAdd.
func CreateZip(folderToAdd string) (dest string, err error) {
	destFileName := filepath.Base(folderToAdd)

	// Create ZIP File
	destFile, err := os.Create(filepath.Join(folderToAdd, destFileName+".zip"))
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	// Zip Writer
	destZip := zip.NewWriter(destFile)
	defer destZip.Close()

	// Load File Entries inside folderToAdd
	entries, err := os.ReadDir(folderToAdd)
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
		file, err := os.Open(filepath.Join(folderToAdd, entry.Name()))
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

	return filepath.Join(folderToAdd, destFileName+".zip"), nil
}

// Rename zip file to cbz file, and wrap it with special folder.
func RenameZip(absPath string) error {
	dir := filepath.Dir(absPath)
	originalFile := filepath.Base(absPath)
	name := parser.FilenameWithoutExt(originalFile)

	// Create Wrap Folder
	wrap := filepath.Join(dir, name)
	err := os.Mkdir(wrap, 0755)
	if err != nil {
		return err
	}

	// Rename
	return os.Rename(absPath, filepath.Join(wrap, name+".cbz"))
}
