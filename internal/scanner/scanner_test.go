package scanner

import (
	"fmt"
	"gui-comicinfo/internal/comicinfo"
	"os"
	"path/filepath"
	"testing"
)

// Test GetPageInfo() get correct range of pages and content.
func TestGetPageInfo(t *testing.T) {
	tempDir := t.TempDir()

	// Create Four file, one is not image
	fileNames := []string{"image1.jpg", "image2.png", "image3.jpeg", "test.xml"}
	fileSizes := []int64{1234, 3456, 789, 12}

	for i, filename := range fileNames {
		file, _ := os.Create(filepath.Join(tempDir, filename))
		file.Truncate(fileSizes[i])

		defer file.Close()
	}

	// Start Testing Functions
	pages, err := GetPageInfo(tempDir)

	if err != nil {
		t.Error(err)
	}

	// Check Size, should skip xml
	if len(pages) != 3 {
		t.Error("Wrong number of pages")
	}

	// Check First Page is front page
	if pages[0].Type != comicinfo.ComicPageType_FrontCover {
		t.Error("Wrong Type of first page")
	}

	// Check filesize
	for i, page := range pages {
		if page.ImageSize != fileSizes[i] {
			t.Errorf("Wrong Size of page %d", i)
		}
	}
}

// Test Scan Books can get correct tags, title & author, also correct pages data.
// This test is consider as a integration test.
//
// ScanBooks() is function that combine functions in package parser & comicinfo,
// if the other test is passed, normally this test will pass too.
func TestScanBooks(t *testing.T) {
	tempDir := t.TempDir()

	// Image Folder contents
	imagesFolder := []string{
		"(C99) [author1] title1",
		"[author2] title2",
		"[author3] title3 [DL版]",
		"(C99) [author4] title4 [DL版]",
	}
	tags := []string{"C99", "", "DL版", "C99,DL版"}
	title := []string{"title1", "title2", "title3 [DL版]", "title4 [DL版]"}

	// Create Image folder inside tempDir
	for _, folder := range imagesFolder {
		os.MkdirAll(filepath.Join(tempDir, folder), 0755)
	}

	// Create Four file, one is not image. In first image folder only
	fileNames := []string{"image1.jpg", "image2.png", "image3.jpeg", "test.xml"}
	fileSizes := []int64{1234, 3456, 789, 12}
	folder1 := filepath.Join(tempDir, imagesFolder[0])

	for i, filename := range fileNames {
		file, _ := os.Create(filepath.Join(folder1, filename))
		file.Truncate(fileSizes[i])
		defer file.Close()
	}

	// Run Function and Checks
	for i, folder := range imagesFolder {
		c, err := ScanBooks(filepath.Join(tempDir, folder))

		if err != nil {
			t.Error(err)
		}

		// Check Basic
		if c.Title != title[i] {
			t.Errorf("Error in title %d %s", i, c.Title)
		}

		if c.Writer != fmt.Sprintf("author%d", i+1) {
			t.Errorf("Error in author %d ", i)
		}

		if c.Manga != comicinfo.Manga_Yes {
			t.Errorf("Error in Manga %d ", i)
		}

		// Check Tags
		if c.Tags != tags[i] {
			t.Errorf("Error in Tags %d ", i)
		}

		// Special Check for 1st folder
		if i != 0 {
			continue
		}

		// Check Size, should skip xml
		if c.PageCount != 3 {
			t.Errorf("Wrong number of page count: %d", c.PageCount)
		} else if len(c.Pages) != 3 {
			t.Errorf("Wrong number of pages: %d", len(c.Pages))
		}

		// Check First Page is front page
		if c.Pages[0].Type != comicinfo.ComicPageType_FrontCover {
			t.Error("Wrong Type of first page")
		}

		// Check filesize
		for i, page := range c.Pages {
			if page.ImageSize != fileSizes[i] {
				t.Errorf("Wrong Size of page %d", i)
			}
		}
	}
}

func TestCheckFolder(t *testing.T) {
	// tempDir := t.TempDir()

	// Prepare Test Set
	//  1. Folder contain Image
	//  2. Folder contain another folder
	//  3. Empty Folder

}
