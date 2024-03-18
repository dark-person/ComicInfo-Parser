package comicinfo

import "encoding/xml"

// Create a new ComicInfo Instance.
func New() ComicInfo {
	return ComicInfo{
		XMLName: xml.Name{Space: "", Local: "ComicInfo"},
		Xsi:     "http://www.w3.org/2001/XMLSchema-instance",
		Xsd:     "http://www.w3.org/2001/XMLSchema",
	}
}
