package services

import (
	"github.com/dark-person/comicinfo-parser/internal/store"
	"github.com/rs/zerolog/log"
)

// Struct that designed for
// send last input record from history module to frontend.
type HistoryResp struct {
	Inputs   []string `json:"Inputs"`
	ErrorMsg string   `json:"ErrorMsg"`
}

// Get all user inputted genre from database.
func (s *ComicService) GetAllGenreInput() HistoryResp {
	return toHistoryResp(store.GetGenreList(s.DB))
}

// Get all user inputted publisher from database.
func (s *ComicService) GetAllPublisherInput() HistoryResp {
	return toHistoryResp(store.GetPublisherList(s.DB))
}

// Get all user inputted tag from database.
func (s *ComicService) GetAllTagInput() HistoryResp {
	return toHistoryResp(store.GetAllTags(s.DB))
}

// Get all word with complete structure that stored in database and will be used in auto fill.
func (s *ComicService) GetAllAutofillWord() []store.AutofillWord {
	words, err := store.GetAllAutofillWord(s.DB)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get autofill word(s)")
		return []store.AutofillWord{}
	}

	return words
}

// toHistoryResp will convert given list and errors to HistoryResp object.
//
// This function is aim to simplify commonly used code.
func toHistoryResp(list []string, err error) HistoryResp {
	if err != nil {
		return HistoryResp{nil, err.Error()}
	}
	return HistoryResp{list, ""}
}
