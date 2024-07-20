package application

import (
	"context"
	"gui-comicinfo/internal/comicinfo"
	"gui-comicinfo/internal/database"
	"gui-comicinfo/internal/scanner"
	"os"
	"path/filepath"

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
