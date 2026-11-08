package services

import (
	"io/fs"
	"path/filepath"

	"github.com/dark-person/comicinfo-parser/internal/dircheck"
	"github.com/dark-person/comicinfo-parser/internal/files"
	"github.com/rs/zerolog/log"
)

// AutomationResp will pass automation progress by ID.
type AutomationResp struct {
	jobID        int    // Job ID for tracking automation progress
	errorMessage string // Error message if any
}

// findValidComicFolders will find all valid comic folder in target path
func findValidComicFolders(target string) ([]string, error) {
	comicFolders := make([]string, 0)

	err := filepath.Walk(target, func(path string, info fs.FileInfo, err error) error {
		// Ignore non directory
		if !info.IsDir() {
			return nil
		}

		log.Trace().Str("path", path).Msg("Walking")

		// Check if directory fulfill as comic folder
		isValid, err2 := dircheck.CheckIfComicFolder(path)
		if err2 != nil {
			return err2
		}

		// If it is invalid, then ignore
		if !isValid {
			return nil
		}

		// Marked this directory as it is valid
		comicFolders = append(comicFolders, path)
		return nil
	})

	if err != nil {
		return []string{}, err
	}

	return comicFolders, nil
}

// AutoGenerateXML will identify all comic folder by filepath.Walk,
// and generate ComicInfo.xml inside.
func (s *ComicService) AutoGenerateXML(target string, allowOverwrite bool) AutomationResp {
	// First locate all directory that is valid comic folder
	comicFolders, err := findValidComicFolders(target)

	if err != nil {
		return AutomationResp{jobID: -1, errorMessage: err.Error()}
	}

	// Loop directory
	for _, path := range comicFolders {
		// Check if comicinfo.xml is already exist,
		// if it is exist and overwrite flag is false,
		// then skip current folder
		if !allowOverwrite && files.IsFileExist(filepath.Join(path, "comicinfo.xml")) {
			continue
		}

		// Start generate comicinfo
		cr := s.GetComicInfo(path)
		if cr.ErrorMessage != "" {
			return AutomationResp{jobID: -1, errorMessage: cr.ErrorMessage}
		}

		// Save comicinfo
		err := s.saveComicInfo(path, cr.ComicInfo)

		if err != nil {
			return AutomationResp{jobID: -1, errorMessage: err.Error()}
		}
	}

	return AutomationResp{jobID: -1, errorMessage: ""}
}

func (s *ComicService) GetAutomationProgress(jobID int) {

}
