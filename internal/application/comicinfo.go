package application

import (
	"fmt"
	"path/filepath"

	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider/autofill"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider/scanner"
	"github.com/dark-person/comicinfo-parser/internal/definitions"
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

	// Autofill by file base name
	r := autofill.New(a.DB)
	result, err := r.Run(filepath.Base(absPath))

	// Consider as acceptable error
	if err != nil {
		fmt.Println(err)

		return ComicInfoResponse{
			ComicInfo:    c,
			ErrorMessage: "",
		}
	}

	// Use autofill result
	c.AddTags(result.Tags...)
	c.AddGenre(result.Inputted[definitions.CategoryGenre]...)
	c.AddPublisher(result.Inputted[definitions.CategoryPublisher]...)
	c.AddTranslator(result.Inputted[definitions.CategoryTranslator]...)

	// Return result
	return ComicInfoResponse{
		ComicInfo:    c,
		ErrorMessage: "",
	}
}
