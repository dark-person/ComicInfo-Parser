package services

import (
	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider/cfgprov"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider/fsprov"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider/historyprov"
	"github.com/dark-person/comicinfo-parser/internal/dircheck"
	"github.com/rs/zerolog/log"
)

type ComicInfoResponse struct {
	ComicInfo    *comicinfo.ComicInfo `json:"ComicInfo"`
	ErrorMessage string               `json:"ErrorMessage"`
}

// Get the comic info by scan the given folder.
// This function will not create/modify the comicinfo xml.
//
// This function will return a comicInfo struct, with error message in string.
func (s *ComicService) GetComicInfo(folder string) ComicInfoResponse {
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

	// Autofill by folder path (intermediate folder + book folder name)
	prov = historyprov.New(s.DB, s.cfg.Preference.ComicDir, absPath)
	c, err = prov.Fill(c)

	// Consider as acceptable error, log error only
	if err != nil {
		log.Debug().Err(err).Msg("acceptable error caught")
	}

	// Fill by configuration
	prov = cfgprov.New(s.cfg, absPath)
	c, err = prov.Fill(c)

	// Consider as acceptable error, log error only
	if err != nil {
		log.Debug().Err(err).Msg("acceptable error caught")
	}

	// Return result
	return ComicInfoResponse{
		ComicInfo:    c,
		ErrorMessage: "",
	}
}
