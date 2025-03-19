package archive

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/dark-person/comicinfo-parser/internal/files"
)

// Create ZIP File of inputDir, and output the zip to destDir.
//
// This function is a variant of CreateZip(), purpose to provide flexibility.
func CreateZipTo(inputDir string, destDir string) (dest string, err error) {
	// Prevent space include in paths
	inputDir = strings.TrimSpace(inputDir)
	destDir = strings.TrimSpace(destDir)

	// Get destination zip filename
	destFileName := strings.TrimSpace(filepath.Base(inputDir))

	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "comicinfo-zip-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create ZIP File
	tmpFile, err := createArchive(inputDir, tmpDir, destFileName)
	if err != nil {
		return tmpFile, err
	}

	// Move zip in temp folder to dest directory
	newPath := filepath.Join(destDir, destFileName+".zip")

	// Return dest file path
	return newPath, files.MoveFile(tmpFile, newPath)
}

// Create archive to temporary directory by given input directory & its content.
// Destination file will be named with "{destFileName}.zip".
//
// Please note that, "zipFilename" is not include file extension, e.g. "zip".
//
// This function will return path of created zip file.
// If any error occur, this function will return that error directly.
func createArchive(inputDir, destDir, zipFilename string) (createdZip string, err error) {
	// Ensure all input not contains spaces
	inputDir = strings.TrimSpace(inputDir)
	destDir = strings.TrimSpace(destDir)
	zipFilename = strings.TrimSpace(zipFilename)

	// Create ZIP File
	f, err := os.Create(filepath.Join(destDir, zipFilename+".zip"))
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Zip Writer
	w := zip.NewWriter(f)
	defer w.Close()

	// Load File Entries inside folderToAdd
	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return "", err
	}

	// Loop File inside folderToAdd
	for _, entry := range entries {
		filename := entry.Name()

		// Skip Zip file
		if filename == zipFilename+".zip" {
			continue
		}

		// Create File inside zip
		fileInside, err := w.Create(filename)
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
		_, err = io.Copy(fileInside, file)
		if err != nil {
			return "", err
		}
	}

	return f.Name(), nil
}
