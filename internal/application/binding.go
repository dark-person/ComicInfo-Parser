package application

import "github.com/dark-person/comicinfo-parser/internal/comicinfo"

// The value of `Manga` const to use in wails binding.
var AllMangaValue = []struct {
	Value  comicinfo.Manga
	TSName string
}{
	{comicinfo.MangaUnknown, "Unknown"},
	{comicinfo.MangaNo, "No"},
	{comicinfo.MangaYes, "Yes"},
	{comicinfo.MangaYesAndRightToLeft, "YesAndRightToLeft"},
}

// The value of `AgeRating` const to use in wails binding.
var AllAgeRatingValue = []struct {
	Value  comicinfo.AgeRating
	TSName string
}{
	{comicinfo.AgeRatingUnknown, "Unknown"},
	{comicinfo.AgeRatingAdultsOnly18, "AdultsOnly18"},
	{comicinfo.AgeRatingEarlyChildhood, "EarlyChildhood"},
	{comicinfo.AgeRatingEveryone, "Everyone"},
	{comicinfo.AgeRatingEveryone10Plus, "Everyone10Plus"},
	{comicinfo.AgeRatingG, "G"},
	{comicinfo.AgeRatingKidsToAdults, "KidsToAdults"},
	{comicinfo.AgeRatingM, "M"},
	{comicinfo.AgeRatingMA15Plus, "MA15Plus"},
	{comicinfo.AgeRatingMature17Plus, "Mature17Plus"},
	{comicinfo.AgeRatingPG, "PG"},
	{comicinfo.AgeRatingR18Plus, "R18Plus"},
	{comicinfo.AgeRatingRatingPending, "RatingPending"},
	{comicinfo.AgeRatingTeen, "Teen"},
	{comicinfo.AgeRatingX18Plus, "X18Plus"},
}
