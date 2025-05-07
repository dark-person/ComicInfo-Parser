package fsprov

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/stretchr/testify/assert"
)

// Create dummy comicinfo for testing.
func newComicInfo(title, writer, tags string) *comicinfo.ComicInfo {
	c := comicinfo.New()
	c.Title = title
	c.Writer = writer
	c.Tags = tags
	return &c
}

// Create dummy comic directory for testing.
func dummyComicDir(path string, c *comicinfo.ComicInfo) *comicinfo.ComicInfo {
	// Ensure directory exists
	os.MkdirAll(path, 0755)

	// Create files
	fileNames := []string{"image1.jpg", "image2.png", "image3.jpeg"}
	fileSizes := []int64{1234, 3456, 789}

	for i, filename := range fileNames {
		file, _ := os.Create(filepath.Join(path, filename))
		file.Truncate(fileSizes[i])
		defer file.Close()
	}

	c.PageCount = 3
	c.Manga = comicinfo.Manga_Yes
	c.Pages = []comicinfo.ComicPageInfo{
		{Image: 0, Type: comicinfo.ComicPageType_FrontCover, ImageSize: 1234},
		{Image: 1, ImageSize: 3456},
		{Image: 2, ImageSize: 789},
	}
	return c
}

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

	// Check file size
	for i, page := range pages {
		if page.ImageSize != fileSizes[i] {
			t.Errorf("Wrong Size of page %d", i)
		}
	}
}

func TestProviderFillGrace(t *testing.T) {
	// Prepare testing directory
	dir := t.TempDir()

	// Prepare Test cases
	type testCase struct {
		folderPath  string
		isFileExist bool
		want        *comicinfo.ComicInfo
	}

	tests := []testCase{
		// 1. Graceful with no comicInfo
		{filepath.Join(dir, "[author1] title1"), false, newComicInfo("title1", "author1", "")},
		// 2. Graceful with existing comicInfo
		{filepath.Join(dir, "[author2] title2"), true, newComicInfo("title2", "author2", "tags")},
	}

	// Generate needed dummy comicinfo directory
	tests[0].want = dummyComicDir(tests[0].folderPath, tests[0].want)

	tests[1].want = dummyComicDir(tests[1].folderPath, tests[1].want)
	tests[1].want.ScanInformation = "abcd"
	comicinfo.Save(tests[1].want, filepath.Join(tests[1].folderPath, "ComicInfo.xml"))

	// Run Tests
	for idx, tt := range tests {
		// Prepare clean  comicinfo
		temp := comicinfo.New()
		c := &temp

		// Create data provider
		provider := New(tt.folderPath)
		c, err := provider.Fill(c)

		// Error Checking
		assert.NoErrorf(t, err, "Case %d: Unwanted error.", idx)

		// Value checking
		assert.EqualValuesf(t, tt.want, c, "Case %d: Not equal comicInfo", idx)
	}
}

func TestProviderFill(t *testing.T) {
	// Prepare testing directory
	dir := t.TempDir()

	// Prepare Test cases
	type testCase struct {
		folderPath string
	}

	tests := []testCase{
		// Empty path
		{""},
		// Invalid path contains invalid characters
		{"???"},
		// Not existing path
		{filepath.Join(dir, "not exists")},
	}

	// Run Tests
	for idx, tt := range tests {
		// Prepare clean  comicinfo
		temp := comicinfo.New()
		c := &temp

		// Create data provider
		provider := New(tt.folderPath)
		c, err := provider.Fill(c)

		// Check if error occur
		assert.Errorf(t, err, "Case %d: Expected error, but no error return.", idx)

		// Check comicinfo is empty
		assert.Nilf(t, c, "Case %d: Expected nil comicinfo but return non-nil", idx)
	}
}
