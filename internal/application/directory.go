package application

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Response for get directory method. Designed due to wails API only allow single return.
type DirectoryResp struct {
	SelectedDir string
	ErrMsg      string
}

// Open directory dialog for wails, allow user to select one directory.
//
// If given directory is empty, then default directory will be handled by wails.
// Otherwise, it will open given directory.
//
// If any error occur, the error message will be also returned in 2nd argument.
// Empty string indicate no error occur.
func (a *App) GetDirectory(dir string) DirectoryResp {
	const title = "Select Directory"

	var opt runtime.OpenDialogOptions

	if dir == "" {
		opt = runtime.OpenDialogOptions{Title: title}
	} else {
		opt = runtime.OpenDialogOptions{Title: title, DefaultDirectory: dir}
	}

	directory, err := runtime.OpenDirectoryDialog(a.ctx, opt)
	if err != nil {
		return DirectoryResp{SelectedDir: "", ErrMsg: err.Error()}
	}

	return DirectoryResp{SelectedDir: directory, ErrMsg: ""}
}

// Attempt to load default output directory.
// If no default directory is set, then return input directory instead.
func (a *App) GetDefaultOutputDirectory(inputDir string) string {
	if a.cfg.DefaultExport == "" {
		return inputDir
	}

	return a.cfg.DefaultExport
}
