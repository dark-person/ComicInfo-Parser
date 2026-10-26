package application

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
func (a *App) GetAllGenreInput() HistoryResp {
	list, err := store.GetGenreList(a.DB)

	if err != nil {
		return HistoryResp{nil, err.Error()}
	}

	return HistoryResp{list, ""}
}

// Get all user inputted publisher from database.
func (a *App) GetAllPublisherInput() HistoryResp {
	list, err := store.GetPublisherList(a.DB)

	if err != nil {
		return HistoryResp{nil, err.Error()}
	}

	return HistoryResp{list, ""}
}

// Get all user inputted tag from database.
func (a *App) GetAllTagInput() HistoryResp {
	list, err := store.GetAllTags(a.DB)

	if err != nil {
		return HistoryResp{nil, err.Error()}
	}

	return HistoryResp{list, ""}
}

// Get all word with complete structure that stored in database and will be used in auto fill.
func (a *App) GetAllAutofillWord() []store.AutofillWord {
	words, err := store.GetAllAutofillWord(a.DB)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get autofill word(s)")
		return []store.AutofillWord{}
	}

	return words
}
