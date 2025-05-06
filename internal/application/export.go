package application

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dark-person/comicinfo-parser/internal/archive"
	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider/scanner"
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/comicinfo-parser/internal/store"
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
	err = a.saveComicInfo(inputDir, c)
	if err != nil {
		return err.Error()
	}

	// Get destination folder
	destDir := a.GetDefaultOutputDirectory(inputDir)

	// Start Archive
	filename, _ := archive.CreateZipTo(inputDir, destDir)
	err = archive.RenameZip(filename, archive.UseDefaultWrap())
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
	err := store.AddTag(a.DB, s...)

	if err != nil {
		return err
	}

	// ========== For values supported by history pkg ============
	values := make([]store.HistoryVal, 0)

	//  ------------- Genre ----------------
	// Split the genre into slice of string by comma
	s = strings.Split(c.Genre, ",")
	for _, item := range s {
		values = append(values, store.HistoryVal{
			Category: definitions.CategoryGenre,
			Value:    item,
		})
	}

	// ----------- Publisher ----------------
	// Split the publisher into slice of string by comma
	s = strings.Split(c.Publisher, ",")
	for _, item := range s {
		values = append(values, store.HistoryVal{
			Category: definitions.CategoryPublisher,
			Value:    item,
		})
	}

	// ----------- Translator ----------------
	// Split the translator into slice of string by comma
	s = strings.Split(c.Translator, ",")
	for _, item := range s {
		values = append(values, store.HistoryVal{
			Category: definitions.CategoryTranslator,
			Value:    item,
		})
	}

	// ----------- INSERT ----------------
	return store.InsertMultiHistory(a.DB, values...)
}

// Export the ComicInfo struct to XML file.
// This will create/overwrite ComicInfo.xml inside originalDir.
//
// OriginalDir & comicinfo MUST be valid, or an error will be returned.
//
// The originalDir MUST be absolute path to write it precisely.
func (a *App) saveComicInfo(originalDir string, c *comicinfo.ComicInfo) error {
	// Check parameter values
	if c == nil {
		return fmt.Errorf("comicinfo is nil value")
	}

	if originalDir == "" {
		return fmt.Errorf("empty folder path")
	}

	// Save ComicInfo.xml
	err := comicinfo.Save(c, filepath.Join(originalDir, comicInfoFile))
	if err != nil {
		return err
	}

	// Write to database
	err = a.saveToHistory(c)
	if err != nil {
		// This is consider as additional part, consider no error here
		logrus.Error(err)
	}

	return nil
}

// API for export comicinfo to original directory.
//
// If the process success, then function will output empty string.
// Otherwise, function will return the reason for error.
func (a *App) ExportXml(originalDir string, c *comicinfo.ComicInfo) (errorMsg string) {
	err := a.saveComicInfo(originalDir, c)
	if err != nil {
		return err.Error()
	}

	return ""
}

// Export the .cbz (contains images & comicInfo) file ONLY to destination.
//
// If the process success, then function will output empty string.
// Otherwise, function will return the reason for error.
//
// Both input directory and output directory MUST be absolute paths.
func (a *App) ExportCbzOnly(inputDir string, exportDir string, c *comicinfo.ComicInfo) (errMsg string) {
	return a.exportCbz(inputDir, exportDir, c, archive.NoWrap())
}

// Export the .cbz (contains images & comicInfo) file to destination,
// wrapped with folder name that same as .cbz base filename.
//
// If the process success, then function will output empty string.
// Otherwise, function will return the reason for error.
//
// Both input directory and output directory MUST be absolute paths.
func (a *App) ExportCbzWithDefaultWrap(inputDir string, exportDir string, c *comicinfo.ComicInfo) (errMsg string) {
	return a.exportCbz(inputDir, exportDir, c, archive.UseDefaultWrap())
}

// Export the .cbz (contains images & comicInfo) file to destination,
// wrapped with folder name specified by user.
//
// If the process success, then function will output empty string.
// Otherwise, function will return the reason for error.
//
// Both input directory and output directory MUST be absolute paths.
func (a *App) ExportCbzWithWrap(inputDir string, exportDir string, wrapFolder string, c *comicinfo.ComicInfo) (errMsg string) {
	return a.exportCbz(inputDir, exportDir, c, archive.UseCustomWrap(wrapFolder))
}

// Core function to export a .cbz file with comicinfo file.
func (a *App) exportCbz(inputDir string, exportDir string, c *comicinfo.ComicInfo, opt archive.RenameOption) (errMsg string) {
	// Reset last export path
	a.lastExportedComic = ""

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
	err := a.saveComicInfo(inputDir, c)
	if err != nil {
		return err.Error()
	}

	// Start Archive
	filename, err := archive.CreateZipTo(inputDir, exportDir)
	if err != nil {
		fmt.Printf("error when zip: %v\n", err)
		return err.Error()
	}
	fmt.Printf("Filename: %s\n", filename)

	// Depend on isWrap value, use different rename option
	err = archive.RenameZip(filename, opt)
	if err != nil {
		fmt.Printf("error when rename: %v\n", err)
		return err.Error()
	}

	// Mark last sucessful export path
	a.lastExportedComic = inputDir
	return ""
}

// Run soft delet process,
// which will move the last exported comic folder to trash bin in config file.
//
// This function only effect when one export function is run successfully.
// If trash bin is not defined, then error will be returned.
func (a *App) RunSoftDelete() (errMsg string) {
	// Ensure config exist
	if a.cfg == nil {
		return "config is corrupted"
	}

	// Ensure last exported path is not empty
	if a.lastExportedComic == "" {
		return "no last exported path"
	}

	// Ensure soft delete is allowed
	trashBin := a.cfg.TrashBin.Path
	if trashBin == "" {
		return "soft delete is not allowed"
	}

	// Run soft delete
	err := archive.SoftDeleteComic(a.lastExportedComic, trashBin)
	if err != nil {
		return err.Error()
	}
	return ""
}
