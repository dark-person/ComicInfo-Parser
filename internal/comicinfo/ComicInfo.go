package comicinfo

import (
	"encoding/xml"
)

// Schema Version of Current ComicInfo Structure
const schemaVersion = "2.1"

// Escape String for XML.
// Prevent appear &#39 .etc
type EscapedString struct {
	InnerXML string `xml:",innerxml" json:"InnerXML"`
}

// The Go Struct Version for ComicInfo.
//
// Developer are suggest to use New() to create a new struct instead of ComicInfo{}
type ComicInfo struct {
	XMLName xml.Name `xml:"ComicInfo"`
	Xsi     string   `xml:"xmlns:xsi,attr,omitempty"`
	Xsd     string   `xml:"xmlns:xsd,attr,omitempty"`

	Title           string        `xml:"Title" json:"Title"`
	Series          string        `xml:"Series" json:"Series"`
	Number          string        `xml:"Number" json:"Number"`
	Count           int           `xml:"Count,omitempty" json:"Count"`
	Volume          int           `xml:"Volume" json:"Volume"`
	AlternateSeries string        `xml:"AlternateSeries" json:"AlternateSeries"`
	AlternateNumber string        `xml:"AlternateNumber" json:"AlternateNumber"`
	AlternateCount  int           `xml:"AlternateCount,omitempty" json:"AlternateCount"`
	StoryArc        string        `xml:"StoryArc" json:"StoryArc"`
	StoryArcNumber  string        `xml:"StoryArcNumber" json:"StoryArcNumber"`
	SeriesGroup     string        `xml:"SeriesGroup" json:"SeriesGroup"`
	Summary         EscapedString `xml:"Summary" json:"Summary"`
	Notes           string        `xml:"Notes" json:"Notes"`
	Year            int           `xml:"Year,omitempty" json:"Year"`
	Month           int           `xml:"Month,omitempty" json:"Month"`
	Day             int           `xml:"Day,omitempty" json:"Day"`
	Writer          string        `xml:"Writer" json:"Writer"`
	Penciller       string        `xml:"Penciller,omitempty" json:"Penciller"`
	Inker           string        `xml:"Inker,omitempty" json:"Inker"`
	Colorist        string        `xml:"Colorist,omitempty" json:"Colorist"`
	Letterer        string        `xml:"Letterer,omitempty" json:"Letterer"`
	CoverArtist     string        `xml:"CoverArtist,omitempty" json:"CoverArtist"`
	Editor          string        `xml:"Editor,omitempty" json:"Editor"`
	Translator      string        `xml:"Translator,omitempty" json:"Translator"`
	Publisher       string        `xml:"Publisher" json:"Publisher"`
	Imprint         string        `xml:"Imprint" json:"Imprint"`
	Genre           string        `xml:"Genre" json:"Genre"`
	Tags            string        `xml:"Tags" json:"Tags"`
	Web             string        `xml:"Web,omitempty" json:"Web"`
	PageCount       int           `xml:"PageCount" json:"PageCount"`
	LanguageISO     string        `xml:"LanguageISO" json:"LanguageISO"`
	Format          string        `xml:"Format" json:"Format"`
	AgeRating       AgeRating     `xml:"AgeRating" json:"AgeRating"`
	BlackAndWhite   YesNo         `xml:"BlackAndWhite,omitempty" json:"BlackAndWhite"`
	Manga           Manga         `xml:"Manga" json:"Manga"`
	Characters      string        `xml:"Characters" json:"Characters"`
	Teams           string        `xml:"Teams" json:"Teams"`
	Locations       string        `xml:"Locations" json:"Locations"`
	ScanInformation string        `xml:"ScanInformation" json:"ScanInformation"`

	Pages               []ComicPageInfo `xml:"Pages>Page" json:"Pages"`
	CommunityRating     float64         `xml:"CommunityRating,omitempty" json:"CommunityRating"`
	MainCharacterOrTeam string          `xml:"MainCharacterOrTeam,omitempty" json:"MainCharacterOrTeam"`
	Review              string          `xml:"Review,omitempty" json:"Review"`
	GTIN                string          `xml:"GTIN,omitempty" json:"GTIN"`
}

// Add Tags to the comic info container.
// This function will handle the comma separation automatically.
func (c *ComicInfo) AddTags(tags ...string) {
	c.Tags = AddValue(c.Tags, tags...)
}

// The Go Struct Version for ComicPageInfo, used to store page information.
type ComicPageInfo struct {
	XMLName xml.Name `xml:"Page"`

	Image       int           `xml:"Image,attr" json:"Image"`
	Type        ComicPageType `xml:"Type,attr,omitempty" json:"Type"`
	DoublePage  bool          `xml:"DoublePage,attr,omitempty" json:"DoublePage"`
	ImageSize   int64         `xml:"ImageSize,attr,omitempty" json:"ImageSize"`
	Key         string        `xml:"Key,attr,omitempty" json:"Key"`
	Bookmark    string        `xml:"Bookmark,attr,omitempty" json:"Bookmark"`
	ImageWidth  string        `xml:"ImageWidth,attr,omitempty" json:"ImageWidth"`
	ImageHeight string        `xml:"ImageHeight,attr,omitempty" json:"ImageHeight"`
}
