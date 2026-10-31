package services

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
func (s *ComicService) GetDirectory(dir string) DirectoryResp {
	dialog := s.app.Dialog.OpenFile().
		SetTitle("Select Directory").
		CanChooseDirectories(true).
		CanChooseFiles(false)

	if dir != "" {
		dialog.SetDirectory(dir)
	}

	directory, err := dialog.PromptForSingleSelection()
	if err != nil {
		return DirectoryResp{SelectedDir: "", ErrMsg: err.Error()}
	}

	return DirectoryResp{SelectedDir: directory, ErrMsg: ""}
}

// Attempt to load default output directory.
// If no default directory is set, then return input directory instead.
func (s *ComicService) GetDefaultOutputDirectory(inputDir string) string {
	if s.cfg.Preference.ExportDir == "" {
		return inputDir
	}

	return s.cfg.Preference.ExportDir
}

// Attempt to load default comic folder.
// If config is not set, then return a empty string instead.
func (s *ComicService) GetComicFolder() string {
	if s.cfg == nil {
		return ""
	}

	return s.cfg.Preference.ComicDir
}
