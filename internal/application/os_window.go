package application

import (
	"os/exec"
	"runtime"

	"github.com/sirupsen/logrus"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// Open folder in window explorer.
//
// If current os is not window, then this function
// will show error dialog instead of open folder.
func (a *App) OpenFolder(path string) {
	// Reject non-window OS
	if runtime.GOOS != `windows` {
		_, err := wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
			Type:    wailsRuntime.ErrorDialog,
			Title:   "OS not supported",
			Message: "Open folder feature is not supported in non-windows OS.",
		})

		if err != nil {
			logrus.Error(err)
		}

		return
	}

	// Run command to open folder
	cmd := exec.Command(`explorer`, path)
	cmd.Run()
}
