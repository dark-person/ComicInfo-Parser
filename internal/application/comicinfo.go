package application

import (
	"fmt"
	"path/filepath"

	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider/cfgprov"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider/fsprov"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider/historyprov"
	"github.com/dark-person/comicinfo-parser/internal/dircheck"
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
	isValid, err := dircheck.CheckFolder(absPath, dircheck.DirectoryOpt{SubFolder: dircheck.Reject, Image: dircheck.Allow})
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

	// ------------------- Fill data ---------------------
	var prov dataprovider.DataProvider

	// Prepare empty comicinfo
	temp := comicinfo.New()
	c := &temp

	// Fill comicinfo by file system data provider
	prov = fsprov.New(absPath)
	c, err = prov.Fill(c)
	if err != nil {
		return ComicInfoResponse{
			ComicInfo:    nil,
			ErrorMessage: err.Error(),
		}
	}

	// Autofill by file base name
	prov = historyprov.New(a.DB, filepath.Base(absPath))
	c, err = prov.Fill(c)

	// Consider as acceptable error, log error only
	if err != nil {
		fmt.Println(err)
	}

	// Fill by configuration
	prov = cfgprov.New(a.cfg)
	c, err = prov.Fill(c)

	// Consider as acceptable error, log error only
	if err != nil {
		fmt.Println(err)
	}

	// Return result
	return ComicInfoResponse{
		ComicInfo:    c,
		ErrorMessage: "",
	}
}
