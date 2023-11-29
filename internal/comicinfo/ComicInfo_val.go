package comicinfo

// This File Store the validation related for ComicInfo

// The ComicPageType Type for comicInfo
type ComicPageType string

const (
	ComicPageType_FrontCover    ComicPageType = "FrontCover"
	ComicPageType_InnerCover    ComicPageType = "InnerCover"
	ComicPageType_Roundup       ComicPageType = "Roundup"
	ComicPageType_Story         ComicPageType = "Story"
	ComicPageType_Advertisement ComicPageType = "Advertisement"
	ComicPageType_Editorial     ComicPageType = "Editorial"
	ComicPageType_Letters       ComicPageType = "Letters"
	ComicPageType_Preview       ComicPageType = "Preview"
	ComicPageType_BackCover     ComicPageType = "BackCover"
	ComicPageType_Other         ComicPageType = "Other"
	ComicPageType_Deleted       ComicPageType = "Deleted"
)

// The AgeRating Type for comicInfo
type AgeRating string

const (
	AgeRating_Unknown        AgeRating = "Unknown"
	AgeRating_AdultsOnly18   AgeRating = "Adults Only 18+"
	AgeRating_EarlyChildhood AgeRating = "Early Childhood"
	AgeRating_Everyone       AgeRating = "Everyone"
	AgeRating_Everyone10Plus AgeRating = "Everyone 10+"
	AgeRating_G              AgeRating = "G"
	AgeRating_KidsToAdults   AgeRating = "Kids to Adults"
	AgeRating_M              AgeRating = "M"
	AgeRating_MA15Plus       AgeRating = "MA15+"
	AgeRating_Mature17Plus   AgeRating = "Mature 17+"
	AgeRating_PG             AgeRating = "PG"
	AgeRating_R18Plus        AgeRating = "R18+"
	AgeRating_RatingPending  AgeRating = "Rating Pending"
	AgeRating_Teen           AgeRating = "Teen"
	AgeRating_X18Plus        AgeRating = "X18+"
)

// The YesNo Type for comicInfo
type YesNo string

const (
	YesNo_Unknown YesNo = "Unknown"
	YesNo_No      YesNo = "No"
	YesNo_Yes     YesNo = "Yes"
)

// The Manga Type for comicInfo
type Manga string

const (
	Manga_Unknown           Manga = "Unknown"
	Manga_No                Manga = "No"
	Manga_Yes               Manga = "Yes"
	Manga_YesAndRightToLeft Manga = "YesAndRightToLeft"
)
