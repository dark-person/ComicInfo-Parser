package application

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dark-person/comicinfo-parser/internal/archive"
	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/history"
	"github.com/dark-person/comicinfo-parser/internal/scanner"
	"github.com/dark-person/comicinfo-parser/internal/tagger"
	"github.com/sirupsen/logrus"
)

const comicInfoFile = "ComicInfo.xml"

// Perform Quick Export Action,
// where ComicInfo.xml file can not be modified before archived.
// The input directory MUST be absolute path.
//
// If error is occur, then return a string containing the error message.
// Otherwise, return empty string.
//
// This function will perform these task:
//  1. Scan the directory and create a ComicInfo.xml file
//  2. Archive the directory and xml file as .cbz file
//  3. Wrap the .cbz file with a folder to copy to komga
//
// This function is specific designed for komga folder structure.
func (a *App) QuickExportKomga(inputDir string) string {
	if inputDir == "" {
		return "folder cannot be empty"
	}

	// Validate the directory
	isValid, err := scanner.CheckFolder(inputDir, scanner.ScanOpt{SubFolder: scanner.Reject, Image: scanner.Allow})
	if err != nil {
		return err.Error()
	} else if !isValid {
		return "folder structure is not correct"
	}

	// Load Abs Path
	c, err := scanner.ScanBooks(inputDir)
	if err != nil {
		return err.Error()
	}

	// Write ComicInfo.xml
	err = comicinfo.Save(c, filepath.Join(inputDir, comicInfoFile))
	if err != nil {
		fmt.Printf("error when saving: %v\n", err)
		return err.Error()
	}

	// Start Archive
	filename, _ := archive.CreateZipTo(inputDir, inputDir)
	err = archive.RenameZip(filename, true)
	if err != nil {
		fmt.Printf("error when archive: %v\n", err)
		return err.Error()
	}
	return ""
}

// Save user input to history database.
// All comicinfo handling logic should be inside this function.
func (a *App) saveToHistory(c *comicinfo.ComicInfo) error {
	// ==================== Tags ====================
	// Split the tags into slice of string by comma
	s := strings.Split(c.Tags, ",")
	err := tagger.AddTag(a.DB, s...)

	if err != nil {
		return err
	}

	// ========== For values supported by history pkg ============
	values := make([]history.HistoryVal, 0)

	//  ------------- Genre ----------------
	// Split the genre into slice of string by comma
	s = strings.Split(c.Genre, ",")
	for _, item := range s {
		values = append(values, history.HistoryVal{
			Category: history.CategoryGenre,
			Value:    item,
		})
	}

	// ----------- Publisher ----------------
	// Split the publisher into slice of string by comma
	s = strings.Split(c.Publisher, ",")
	for _, item := range s {
		values = append(values, history.HistoryVal{
			Category: history.CategoryPublisher,
			Value:    item,
		})
	}

	// ----------- INSERT ----------------
	return history.InsertMultiple(a.DB, values...)
}

// Export the ComicInfo struct to XML file.
// This will create/overwrite ComicInfo.xml inside originalDir.
// If the process success, then function will output empty string.
// Otherwise, function will return the reason for error.
//
// The originalDir MUST be absolute path to write it precisely.
func (a *App) ExportXml(originalDir string, c *comicinfo.ComicInfo) (errorMsg string) {
	// Check if comic info is nil value
	if c == nil {
		return "comicinfo is nil value"
	}

	if originalDir == "" {
		return "empty folder path"
	}

	// Save ComicInfo.xml
	err := comicinfo.Save(c, filepath.Join(originalDir, comicInfoFile))
	if err != nil {
		fmt.Printf("error when save xml: %v\n", err)
		return err.Error()
	}

	// Write to database
	err = a.saveToHistory(c)
	if err != nil {
		logrus.Error(err)
	}

	return ""
}

// Export the .cbz file to destination.
// This .cbz file will contain all image in the input directory,
// including newly generated ComicInfo.xml.
//
// This function supports control of using wrap folder.
//
// If the process success, then function will output empty string.
// Otherwise, function will return the reason for error.
//
// Both input directory and output directory MUST be absolute paths.
func (a *App) ExportCbz(inputDir string, exportDir string, c *comicinfo.ComicInfo, isWrap bool) (errMsg string) {
	// Check parameters first
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		return "input directory does not exist"
	}

	if _, err := os.Stat(exportDir); os.IsNotExist(err) {
		return "export directory does not exist"
	}

	if c == nil {
		return "empty comic info"
	}

	// Save ComicInfo.xml
	err := comicinfo.Save(c, filepath.Join(inputDir, "ComicInfo.xml"))
	if err != nil {
		fmt.Printf("error when save: %v\n", err)
		return err.Error()
	}

	// Write to database
	err = a.saveToHistory(c)
	if err != nil {
		logrus.Error(err)
	}

	// Start Archive
	filename, err := archive.CreateZipTo(inputDir, exportDir)
	if err != nil {
		fmt.Printf("error when zip: %v\n", err)
		return err.Error()
	}
	fmt.Printf("Filename: %s\n", filename)

	err = archive.RenameZip(filename, isWrap)
	if err != nil {
		fmt.Printf("error when rename: %v\n", err)
		return err.Error()
	}
	return ""
}
