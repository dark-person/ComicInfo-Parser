package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"gui-comicinfo/internal/archive"
	"gui-comicinfo/internal/comicinfo"
	"gui-comicinfo/internal/scanner"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
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

// Get the comic info by scan the given folder.
// This function will not create/modify the comicinfo xml.
//
// This function will return a comicInfo struct, with error message in string.
func (a *App) GetComicInfo(folder string) (*comicinfo.ComicInfo, string) {

	absPath := folder

	if absPath == "" {
		return nil, "folder cannot be empty"
	}

	// Validate the directory
	isValid, err := scanner.CheckFolder(absPath, scanner.ScanOpt{SubFolder: scanner.Reject, Image: scanner.Allow})
	if err != nil {
		return nil, err.Error()
	} else if !isValid {
		return nil, "folder structure is not correct"
	}

	// Load Abs Path
	c, err := scanner.ScanBooks(absPath)
	if err != nil {
		return nil, err.Error()
	}

	return c, ""
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

	output, err := xml.MarshalIndent(c, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err.Error()
	}

	// Open File for reading
	f, err := os.Create(filepath.Join(absPath, "ComicInfo.xml"))
	if err != nil {
		return err.Error()
	}
	defer f.Close()

	// Write XML Content to file
	f.Write([]byte("<?xml version=\"1.0\"?>\n"))
	f.Write(output)

	err = f.Sync()
	if err != nil {
		return err.Error()
	}

	// Start Archive
	filename, _ := archive.CreateZip(absPath)
	err = archive.RenameZip(filename)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err.Error()
	}
	return ""
}
