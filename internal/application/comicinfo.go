package application

import (
	"gui-comicinfo/internal/comicinfo"
	"gui-comicinfo/internal/scanner"
)

type ComicInfoResponse struct {
	ComicInfo    *comicinfo.ComicInfo `json:"ComicInfo"`
	ErrorMessage string               `json:"ErrorMessage"`
}

// Get the comic info by scan the given folder.
// This function will not create/modify the comicinfo xml.
//
// This function will return a comicInfo struct, with error message in string.
func (a *App) GetComicInfo(folder string) ComicInfoResponse {
	absPath := folder

	// Check Absolute path is empty or not
	if absPath == "" {
		return ComicInfoResponse{
			ComicInfo:    nil,
			ErrorMessage: "folder cannot be empty",
		}
	}

	// Validate the directory
	isValid, err := scanner.CheckFolder(absPath, scanner.ScanOpt{SubFolder: scanner.Reject, Image: scanner.Allow})
	if err != nil {
		return ComicInfoResponse{
			ComicInfo:    nil,
			ErrorMessage: err.Error(),
		}
	} else if !isValid {
		return ComicInfoResponse{
			ComicInfo:    nil,
			ErrorMessage: "folder structure is not correct",
		}
	}

	// Load Abs Path
	c, err := scanner.ScanBooks(absPath)
	if err != nil {
		return ComicInfoResponse{
			ComicInfo:    nil,
			ErrorMessage: err.Error(),
		}
	}

	// Return result
	return ComicInfoResponse{
		ComicInfo:    c,
		ErrorMessage: "",
	}
}
