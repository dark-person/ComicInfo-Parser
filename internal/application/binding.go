package application

import "gui-comicinfo/internal/comicinfo"

// The value of `Manga` const to use in wails binding.
var AllMangaValue = []struct {
	Value  comicinfo.Manga
	TSName string
}{
	{comicinfo.Manga_Unknown, "Unknown"},
	{comicinfo.Manga_No, "No"},
	{comicinfo.Manga_Yes, "Yes"},
	{comicinfo.Manga_YesAndRightToLeft, "YesAndRightToLeft"},
}

// The value of `AgeRating` const to use in wails binding.
var AllAgeRatingValue = []struct {
	Value  comicinfo.AgeRating
	TSName string
}{
	{comicinfo.AgeRating_Unknown, "Unknown"},
	{comicinfo.AgeRating_AdultsOnly18, "AdultsOnly18"},
	{comicinfo.AgeRating_EarlyChildhood, "EarlyChildhood"},
	{comicinfo.AgeRating_Everyone, "Everyone"},
	{comicinfo.AgeRating_Everyone10Plus, "Everyone10Plus"},
	{comicinfo.AgeRating_G, "G"},
	{comicinfo.AgeRating_KidsToAdults, "KidsToAdults"},
	{comicinfo.AgeRating_M, "M"},
	{comicinfo.AgeRating_MA15Plus, "MA15Plus"},
	{comicinfo.AgeRating_Mature17Plus, "Mature17Plus"},
	{comicinfo.AgeRating_PG, "PG"},
	{comicinfo.AgeRating_R18Plus, "R18Plus"},
	{comicinfo.AgeRating_RatingPending, "RatingPending"},
	{comicinfo.AgeRating_Teen, "Teen"},
	{comicinfo.AgeRating_X18Plus, "X18Plus"},
}
