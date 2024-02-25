package main

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
