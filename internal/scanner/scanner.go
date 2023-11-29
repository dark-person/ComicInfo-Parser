package scanner

import (
	"gui-comicinfo/internal/comicinfo"
	"gui-comicinfo/internal/parser"
	"os"
	"path/filepath"
	"strings"
)

// Scan all image in the Directory. Sorted by filename.
func GetPageInfo(absPath string) []comicinfo.ComicPageInfo {
	entries, err := os.ReadDir(absPath)
	if err != nil {
		return nil
	}

	pageInfos := make([]comicinfo.ComicPageInfo, 0)

	imageIdx := 0
	for _, entry := range entries {
		ext := filepath.Ext(entry.Name())
		ext = strings.ToLower(ext)

		if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
			continue
		}

		info, innerErr := entry.Info()
		if innerErr != nil {
			continue
		}

		page := comicinfo.ComicPageInfo{}
		page.Image = imageIdx
		page.ImageSize = info.Size()

		if imageIdx == 0 {
			page.Type = "FrontCover"
		}

		pageInfos = append(pageInfos, page)
		imageIdx++

		// fmt.Println("Entry ", idx, ":", filepath.Join(absPath, entry.Name()), "Size: ", info.Size())
	}

	return pageInfos
}

// Scan the folderPath as a book/manga, then return comicInfo.
func ScanBooks(folderPath string) comicinfo.ComicInfo {
	folderName := filepath.Base(folderPath)

	// Test XML
	market, author, bookName := parser.Parse(folderName)
	c := comicinfo.New()
	c.Title = bookName
	c.Writer = author
	c.Manga = "Yes"
	if market != "" {
		c.Imprint = market
		c.AddTags(market)
	}

	// Add Special Tags
	tags := parser.GetSpecialTags(folderName)
	c.AddTags(tags...)

	// Get Pages
	pages := GetPageInfo(folderPath)

	c.Pages = pages
	c.PageCount = len(pages)

	return c
}
