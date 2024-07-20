package application

import (
	"context"
	"fmt"
	"gui-comicinfo/internal/archive"
	"gui-comicinfo/internal/comicinfo"
	"gui-comicinfo/internal/database"
	"gui-comicinfo/internal/history"
	"gui-comicinfo/internal/scanner"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	DB  *database.AppDB
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	// Init Database
	var err error
	a.DB, err = database.NewDB()
	if err != nil {
		panic(err)
	}

	// Perform connect & migration
	err = a.DB.Connect()
	if err != nil {
		panic(err)
	}

	// Perform migration to database if needed
	err = a.DB.StepToLatest()
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Error",
			Message: "Database version is corrupted, please fix database or remove current database file.",
		})
		os.Exit(1)
	}
}

// Function that is almost same with `startup()`,
// but different on database handling.
//
// This function MUST not used outside test purposes.
func (a *App) StartUpTest(ctx context.Context, dbPath string) {
	a.ctx = ctx

	// Init Database
	var err error
	a.DB, err = database.NewPathDB(dbPath)
	if err != nil {
		panic(err)
	}

	// Perform connect & migration
	err = a.DB.Connect()
	if err != nil {
		panic(err)
	}

	// Perform migration to database if needed
	err = a.DB.StepToLatest()
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Error",
			Message: "Database version is corrupted, please fix database or remove current database file.",
		})
		os.Exit(1)
	}
}

// Open a Dialog for user to select Directory.
//
// If Error is occur, then this function will return an empty string
func (a *App) GetDirectory() string {
	directory, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Directory",
	})

	if err != nil {
		logrus.Warnf("Error when getting directory from user: %v\n", err)
		return ""
	}
	return directory
}

// Open a Dialog for user to select Directory, this dialog will show default directory when open.
//
// If Error is occur, then this function will return an empty string
func (a *App) GetDirectoryWithDefault(defaultDirectory string) string {
	// Try to get parent of default directory
	dir := filepath.Dir(defaultDirectory)

	directory, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Select Directory",
		DefaultDirectory: dir,
	})

	if err != nil {
		logrus.Warnf("Error when getting directory from user: %v\n", err)
		return ""
	}
	return directory
}

type ComicInfoResponse struct {
	ComicInfo    *comicinfo.ComicInfo `json:"ComicInfo"`
	ErrorMessage string               `json:"ErrorMessage"`
}

// Get the comic info by scan the given folder.
// This function will not create/modify the comicinfo xml.
//
// This function will return a comicInfo struct, with error message in string.
func (a *App) GetComicInfo(folder string) ComicInfoResponse {
	absPath := folder

	// Check Absolute path is empty or not
	if absPath == "" {
		return ComicInfoResponse{
			ComicInfo:    nil,
			ErrorMessage: "folder cannot be empty",
		}
	}

	// Validate the directory
	isValid, err := scanner.CheckFolder(absPath, scanner.ScanOpt{SubFolder: scanner.Reject, Image: scanner.Allow})
	if err != nil {
		return ComicInfoResponse{
			ComicInfo:    nil,
			ErrorMessage: err.Error(),
		}
	} else if !isValid {
		return ComicInfoResponse{
			ComicInfo:    nil,
			ErrorMessage: "folder structure is not correct",
		}
	}

	// Load Abs Path
	c, err := scanner.ScanBooks(absPath)
	if err != nil {
		return ComicInfoResponse{
			ComicInfo:    nil,
			ErrorMessage: err.Error(),
		}
	}

	// Return result
	return ComicInfoResponse{
		ComicInfo:    c,
		ErrorMessage: "",
	}
}

// Perform Quick Export Action,
// where ComicInfo.xml file can not be modified before archived.
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
func (a *App) QuickExportKomga(folder string) string {
	absPath := folder

	if absPath == "" {
		return "folder cannot be empty"
	}

	// Validate the directory
	isValid, err := scanner.CheckFolder(absPath, scanner.ScanOpt{SubFolder: scanner.Reject, Image: scanner.Allow})
	if err != nil {
		return err.Error()
	} else if !isValid {
		return "folder structure is not correct"
	}

	// Load Abs Path
	c, err := scanner.ScanBooks(absPath)
	if err != nil {
		return err.Error()
	}

	// Write ComicInfo.xml
	err = comicinfo.Save(c, filepath.Join(absPath, "ComicInfo.xml"))
	if err != nil {
		fmt.Printf("error when saving: %v\n", err)
		return err.Error()
	}

	// Start Archive
	filename, _ := archive.CreateZip(absPath)
	err = archive.RenameZip(filename, true)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err.Error()
	}
	return ""
}

// Save user input to history database.
// All comicinfo handling logic should be inside this function.
func (a *App) saveToHistory(c *comicinfo.ComicInfo) error {

	values := make([]history.HistoryVal, 0)

	//  ------------- Genre ----------------

	// Split the genre into slice of string by comma
	s := strings.Split(c.Genre, ",")
	for _, item := range s {
		values = append(values, history.HistoryVal{
			Category: history.Genre_Text,
			Value:    item,
		})
	}

	// ----------- Publisher ----------------
	// Split the publisher into slice of string by comma
	s = strings.Split(c.Publisher, ",")
	for _, item := range s {
		values = append(values, history.HistoryVal{
			Category: history.Publisher_Text,
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
	err := comicinfo.Save(c, filepath.Join(originalDir, "ComicInfo.xml"))
	if err != nil {
		fmt.Printf("error: %v\n", err)
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
		fmt.Printf("error: %v\n", err)
		return err.Error()
	}

	// Write to database
	err = a.saveToHistory(c)
	if err != nil {
		logrus.Error(err)
	}

	// Start Archive
	filename, _ := archive.CreateZipTo(inputDir, exportDir)
	err = archive.RenameZip(filename, isWrap)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err.Error()
	}
	return ""
}

// -----------------------------------------------

// Struct that designed for
// send last input record from history module to frontend.
type HistoryResp struct {
	Inputs   []string `json:"Inputs"`
	ErrorMsg string   `json:"ErrorMsg"`
}

// Get all user inputted genre from database.
func (a *App) GetAllGenreInput() HistoryResp {
	list, err := history.GetGenreList(a.DB)

	if err != nil {
		return HistoryResp{nil, err.Error()}
	}

	return HistoryResp{list, ""}
}

// Get all user inputted publisher from database.
func (a *App) GetAllPublisherInput() HistoryResp {
	list, err := history.GetPublisherList(a.DB)

	if err != nil {
		return HistoryResp{nil, err.Error()}
	}

	return HistoryResp{list, ""}
}
