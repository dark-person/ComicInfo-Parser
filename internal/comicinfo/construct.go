package comicinfo

// Create a new ComicInfo Instance.
func New() ComicInfo {
	return ComicInfo{
		Xsi: "http://www.w3.org/2001/XMLSchema-instance",
		Xsd: "http://www.w3.org/2001/XMLSchema",
	}
}
