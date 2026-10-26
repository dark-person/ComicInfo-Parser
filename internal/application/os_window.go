package application

import (
	"os/exec"
	"runtime"

	"github.com/rs/zerolog/log"
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
			log.Error().Err(err).Msg("Failed to show message dialog.")
		}

		return
	}

	// Run command to open folder
	cmd := exec.Command(`explorer`, path)
	cmd.Run()
}
