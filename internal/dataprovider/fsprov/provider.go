// Package to get comicinfo data from file system.
// This package will provider a data provider that get data from folder name, XML & image file.
package fsprov

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider"
	"github.com/dark-person/comicinfo-parser/internal/files"
)

// Comicinfo data provider that use file system to fill comicinfo details.
type FsProvider struct {
	folderPath string // Folder path of comic, in absoulte path
}

var _ dataprovider.DataProvider = (*FsProvider)(nil)

// Create a new file system provider that use file system to fill comicinfo details.
// Folder path SHOULD be absoulte path for precise scan.
func New(folderPath string) *FsProvider {
	return &FsProvider{folderPath: folderPath}
}

// Scan all image in the Directory. Sorted by filename.
func GetPageInfo(absPath string) (pages []comicinfo.ComicPageInfo, err error) {
	entries, err := os.ReadDir(absPath)
	if err != nil {
		return nil, err
	}

	pageInfos := make([]comicinfo.ComicPageInfo, 0)

	// Image must be re-scan due to image contents may changed
	imageIdx := 0
	for _, entry := range entries {
		if !files.IsSupportedImg(entry.Name()) {
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
	}

	return pageInfos, nil
}

func (p *FsProvider) Fill(c *comicinfo.ComicInfo) (*comicinfo.ComicInfo, error) {
	// Prevent Empty path
	if p.folderPath == "" {
		return c, fmt.Errorf("empty folder path")
	}

	// Prevent invalid path
	if !files.IsPathValid(p.folderPath) || !files.IsFileExist(p.folderPath) {
		return c, fmt.Errorf("invalid folder path")
	}

	// Check any comic info file before start parse
	infoPath := filepath.Join(p.folderPath, "ComicInfo.xml")

	// If comicinfo.xml exist, then ignore input comicinfo struct
	if files.IsFileExist(infoPath) {
		// Marshal info to struct
		c, err := comicinfo.Load(infoPath)
		if err != nil {
			return c, err
		}

		// Force Re-scan Pages
		pages, err := GetPageInfo(p.folderPath)
		if err != nil {
			return c, err
		}

		c.Pages = pages
		c.PageCount = len(pages)
		return c, nil
	}

	// Parse Folder to info
	folderName := filepath.Base(p.folderPath)
	market, author, bookName := parseFolder(folderName)
	c.Title = bookName
	c.Writer = author
	c.Manga = "Yes"
	if market != "" {
		c.Imprint = market
		c.AddTags(market)
	}

	// Get Pages
	pages, err := GetPageInfo(p.folderPath)

	c.Pages = pages
	c.PageCount = len(pages)

	return c, err
}
