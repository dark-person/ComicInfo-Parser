package application

import "gui-comicinfo/internal/history"

// Struct that designed for
// send last input record from history module to frontend.
type HistoryResp struct {
	Inputs   []string `json:"Inputs"`
	ErrorMsg string   `json:"ErrorMsg"`
}

// Get all user inputted genre from database.
func (a *App) GetAllGenreInput() HistoryResp {
	list, err := history.GetGenreList(a.DB)

	if err != nil {
		return HistoryResp{nil, err.Error()}
	}

	return HistoryResp{list, ""}
}

// Get all user inputted publisher from database.
func (a *App) GetAllPublisherInput() HistoryResp {
	list, err := history.GetPublisherList(a.DB)

	if err != nil {
		return HistoryResp{nil, err.Error()}
	}

	return HistoryResp{list, ""}
}
