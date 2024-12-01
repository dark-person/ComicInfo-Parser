package application

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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

// Attempt to load default output directory.
// If no default directory is set, then return input directory instead.
func (a *App) GetDefaultOutputDirectory(inputDir string) string {
	if a.cfg.DefaultExport == "" {
		return inputDir
	}

	return a.cfg.DefaultExport
}
