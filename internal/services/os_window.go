package services

import (
	"os/exec"
	"runtime"
)

// Open folder in window explorer.
//
// If current os is not window, then this function
// will show error dialog instead of open folder.
func (s *ComicService) OpenFolder(path string) {
	// Reject non-window OS
	if runtime.GOOS != `windows` {
		s.app.Dialog.
			Error().
			SetTitle("OS not supported").
			SetMessage("Open folder feature is not supported in non-windows OS.").
			Show()

		return
	}

	// Run command to open folder
	cmd := exec.Command(`explorer`, path)
	cmd.Run()
}
