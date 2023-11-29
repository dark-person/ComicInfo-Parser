package comicinfo

import (
	"encoding/xml"
	"strings"
)

// Schema Version of Current ComicInfo Structure
const schemaVersion = "2.1"

// Escape String for XML.
// Prevent appear &#39 .etc
type EscapedString struct {
	InnerXML string `xml:",innerxml"`
}

// The Go Struct Version for ComicInfo.
//
// Developer are suggest to use New() to create a new struct instead of ComicInfo{}
type ComicInfo struct {
	XMLName xml.Name `xml:"ComicInfo"`
	Xsi     string   `xml:"xmlns:xsi,attr,omitempty"`
	Xsd     string   `xml:"xmlns:xsd,attr,omitempty"`

	Title           string        `xml:"Title"`
	Series          string        `xml:"Series"`
	Number          string        `xml:"Number"`
	Count           int           `xml:"Count,omitempty"`
	Volume          int           `xml:"Volume"`
	AlternateSeries string        `xml:"AlternateSeries"`
	AlternateNumber string        `xml:"AlternateNumber"`
	AlternateCount  int           `xml:"AlternateCount,omitempty"`
	StoryArc        string        `xml:"StoryArc"`
	StoryArcNumber  string        `xml:"StoryArcNumber"`
	SeriesGroup     string        `xml:"SeriesGroup"`
	Summary         EscapedString `xml:"Summary"`
	Notes           string        `xml:"Notes"`
	Year            int           `xml:"Year,omitempty"`
	Month           int           `xml:"Month,omitempty"`
	Day             int           `xml:"Day,omitempty"`
	Writer          string        `xml:"Writer"`
	Penciller       string        `xml:"Penciller,omitempty"`
	Inker           string        `xml:"Inker,omitempty"`
	Colorist        string        `xml:"Colorist,omitempty"`
	Letterer        string        `xml:"Letterer,omitempty"`
	CoverArtist     string        `xml:"CoverArtist,omitempty"`
	Editor          string        `xml:"Editor,omitempty"`
	Translator      string        `xml:"Translator,omitempty"`
	Publisher       string        `xml:"Publisher"`
	Imprint         string        `xml:"Imprint"`
	Genre           string        `xml:"Genre"`
	Tags            string        `xml:"Tags"`
	Web             string        `xml:"Web,omitempty"`
	PageCount       int           `xml:"PageCount"`
	LanguageISO     string        `xml:"LanguageISO"`
	Format          string        `xml:"Format"`
	AgeRating       AgeRating     `xml:"AgeRating"`
	BlackAndWhite   YesNo         `xml:"BlackAndWhite,omitempty"`
	Manga           Manga         `xml:"Manga"`
	Characters      string        `xml:"Characters"`
	Teams           string        `xml:"Teams"`
	Locations       string        `xml:"Locations"`
	ScanInformation string        `xml:"ScanInformation"`

	Pages               []ComicPageInfo `xml:"Pages>Page"`
	CommunityRating     float64         `xml:"CommunityRating,omitempty"`
	MainCharacterOrTeam string          `xml:"MainCharacterOrTeam,omitempty"`
	Review              string          `xml:"Review,omitempty"`
	GTIN                string          `xml:"GTIN,omitempty"`
}

// Add Tags to the comic info container.
// This function will handle the comma separation automatically.
func (c *ComicInfo) AddTags(tags ...string) {
	original := strings.Split(c.Tags, ",")
	new := append(original, tags...)

	temp := make([]string, 0)
	for _, tag := range new {
		// Prevent Empty Strings
		if strings.TrimSpace(tag) == "" {
			continue
		}
		temp = append(temp, tag)
	}

	c.Tags = strings.Join(temp, ",")
}

type ComicPageInfo struct {
	XMLName xml.Name `xml:"Page"`

	Image       int           `xml:"Image,attr"`
	Type        ComicPageType `xml:"Type,attr,omitempty"`
	DoublePage  bool          `xml:"DoublePage,attr,omitempty"`
	ImageSize   int64         `xml:"ImageSize,attr,omitempty"`
	Key         string        `xml:"Key,attr,omitempty"`
	Bookmark    string        `xml:"Bookmark,attr,omitempty"`
	ImageWidth  string        `xml:"ImageWidth,attr,omitempty"`
	ImageHeight string        `xml:"ImageHeight,attr,omitempty"`
}
