package archive

import (
	"archive/zip"
	"gui-comicinfo/internal/files"
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

// Rename zip file to cbz file.
// If user want to wrap the .cbz file with its filename,
// then put true in isWrap parameter.
//
// The reason for wrap is to designed for komga exports,
// when only one book is available,
// this filepath format would be better for komga:
//
//	{bookName}/{bookName}.cbz
func RenameZip(absPath string, isWrap bool) error {
	originalDir := filepath.Dir(absPath)
	originalFile := filepath.Base(absPath)
	name := files.TrimExt(originalFile)

	// If not wrap, then just rename the file extension to .cbz
	if !isWrap {
		return os.Rename(absPath, filepath.Join(originalDir, name+".cbz"))
	}

	// Create Wrap Folder
	wrappedDir := filepath.Join(originalDir, name)
	err := os.Mkdir(wrappedDir, 0755)
	if err != nil {
		return err
	}

	// Rename
	return os.Rename(absPath, filepath.Join(wrappedDir, name+".cbz"))
}
