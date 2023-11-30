package scanner

import (
	"fmt"
	"gui-comicinfo/internal/comicinfo"
	"gui-comicinfo/internal/parser"
	"os"
	"path/filepath"
	"strings"
)

// Scan all image in the Directory. Sorted by filename.
func GetPageInfo(absPath string) (pages []comicinfo.ComicPageInfo, err error) {
	entries, err := os.ReadDir(absPath)
	if err != nil {
		return nil, err
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

	return pageInfos, nil
}

// Scan the folderPath as a book/manga, then return comicInfo.
func ScanBooks(folderPath string) (comicinfo.ComicInfo, error) {
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
	pages, err := GetPageInfo(folderPath)

	c.Pages = pages
	c.PageCount = len(pages)

	return c, err
}

// Check the folder fulfill requirement of the given Scanner Options
func CheckFolder(folderPath string, opt ScanOpt) (bool, error) {
	if !opt.Valid() {
		return false, fmt.Errorf("invalid scan options")
	}

	// Get all file/folder in given path
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return false, err
	}

	// Prepare variable
	subfolderCount := 0
	imageCount := 0
	totalCount := 0

	// Loop the entries
	for _, entry := range entries {
		totalCount++

		// Directory Check
		if entry.IsDir() {
			subfolderCount++
			continue
		}

		// Image Extension check
		ext := filepath.Ext(entry.Name())
		ext = strings.ToLower(ext)
		if ext == ".jpg" || ext == ".png" || ext == ".jpeg" {
			imageCount++
			continue
		}
	}

	// Check Contain Only Option
	if opt.Image == AllowOnly && (totalCount != imageCount) {
		return false, nil
	}

	if opt.SubFolder == AllowOnly && (totalCount != subfolderCount) {
		return false, nil
	}

	// Check Reject Option
	if opt.Image == Reject && imageCount > 0 {
		return false, nil
	}

	if opt.SubFolder == Reject && subfolderCount > 0 {
		return false, nil
	}

	// Check Contain Option
	if (opt.Image == Allow || opt.Image == AllowOnly) && imageCount <= 0 {
		return false, nil
	}

	if (opt.SubFolder == Allow || opt.SubFolder == AllowOnly) && subfolderCount <= 0 {
		return false, nil
	}

	// All Checking Passed
	return true, nil
}
