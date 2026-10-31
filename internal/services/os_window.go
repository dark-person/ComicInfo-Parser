package services

import (
	"os/exec"
	"runtime"

	"github.com/rs/zerolog/log"
)

// Open folder in window explorer.
//
// If current os is not window, then this function
// will show error dialog instead of open folder.
func (s *ComicService) OpenFolder(path string) {
	// Reject non-window OS
	if runtime.GOOS != `windows` {
		log.Warn().Msg("Open folder feature is not supported in non-windows OS.")

		if s.app == nil {
			log.Error().Msg("wails application was not provided. No dialog can show")
			return
		}

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
