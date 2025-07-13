package comicinfo

// This File Store the validation related for ComicInfo

// The ComicPageType Type for comicInfo
type ComicPageType string

const (
	ComicPageTypeFrontCover    ComicPageType = "FrontCover"
	ComicPageTypeInnerCover    ComicPageType = "InnerCover"
	ComicPageTypeRoundup       ComicPageType = "Roundup"
	ComicPageTypeStory         ComicPageType = "Story"
	ComicPageTypeAdvertisement ComicPageType = "Advertisement"
	ComicPageTypeEditorial     ComicPageType = "Editorial"
	ComicPageTypeLetters       ComicPageType = "Letters"
	ComicPageTypePreview       ComicPageType = "Preview"
	ComicPageTypeBackCover     ComicPageType = "BackCover"
	ComicPageTypeOther         ComicPageType = "Other"
	ComicPageTypeDeleted       ComicPageType = "Deleted"
)

// The AgeRating Type for comicInfo
type AgeRating string

const (
	AgeRatingUnknown        AgeRating = "Unknown"
	AgeRatingAdultsOnly18   AgeRating = "Adults Only 18+"
	AgeRatingEarlyChildhood AgeRating = "Early Childhood"
	AgeRatingEveryone       AgeRating = "Everyone"
	AgeRatingEveryone10Plus AgeRating = "Everyone 10+"
	AgeRatingG              AgeRating = "G"
	AgeRatingKidsToAdults   AgeRating = "Kids to Adults"
	AgeRatingM              AgeRating = "M"
	AgeRatingMA15Plus       AgeRating = "MA15+"
	AgeRatingMature17Plus   AgeRating = "Mature 17+"
	AgeRatingPG             AgeRating = "PG"
	AgeRatingR18Plus        AgeRating = "R18+"
	AgeRatingRatingPending  AgeRating = "Rating Pending"
	AgeRatingTeen           AgeRating = "Teen"
	AgeRatingX18Plus        AgeRating = "X18+"
)

// The YesNo Type for comicInfo
type YesNo string

const (
	YesNoValUnknown YesNo = "Unknown"
	YesNoValNo      YesNo = "No"
	YesNoValYes     YesNo = "Yes"
)

// The Manga Type for comicInfo
type Manga string

const (
	MangaUnknown           Manga = "Unknown"
	MangaNo                Manga = "No"
	MangaYes               Manga = "Yes"
	MangaYesAndRightToLeft Manga = "YesAndRightToLeft"
)
