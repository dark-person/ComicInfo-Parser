package services

import (
	"io/fs"
	"path/filepath"

	"github.com/dark-person/comicinfo-parser/internal/dircheck"
	"github.com/dark-person/comicinfo-parser/internal/files"
)

type AutomationResp struct {
	jobId        int    // Job ID for tracking automation progress
	errorMessage string // Error message if any
}

func (s *ComicService) AutoGenerateXML(target string, allowOverwrite bool) AutomationResp {
	comicFolders := make([]string, 0)

	// First locate all directory that is valid comic folder
	err := filepath.Walk(target, func(path string, info fs.FileInfo, err error) error {
		// Ignore non directory
		if !info.IsDir() {
			return nil
		}

		// Check if directory fulfill as comic folder
		isValid, err := dircheck.CheckIfComicFolder(path)
		if err != nil {
			return err
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
		return AutomationResp{jobId: -1, errorMessage: err.Error()}
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
			return AutomationResp{jobId: -1, errorMessage: cr.ErrorMessage}
		}

		// Save comicinfo
		err := s.saveComicInfo(path, cr.ComicInfo)

		if err != nil {
			return AutomationResp{jobId: -1, errorMessage: err.Error()}
		}
	}

	return AutomationResp{jobId: -1, errorMessage: ""}
}

func (s *ComicService) GetAutomationProgress(jobId int) {

}
