package scanner

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/stretchr/testify/assert"
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

	// Check file size
	for i, page := range pages {
		if page.ImageSize != fileSizes[i] {
			t.Errorf("Wrong Size of page %d", i)
		}
	}
}

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

func TestScanBooks(t *testing.T) {
	// Prepare testing directory
	dir := t.TempDir()

	// Prepare Test cases
	type testCase struct {
		folderPath string
		want       *comicinfo.ComicInfo
		wantErr    bool
	}

	tests := []testCase{
		// 1. Graceful with no comicInfo
		{filepath.Join(dir, "[author1] title1"), newComicInfo("title1", "author1", ""), false},
		// 2. Graceful with existing comicInfo
		{filepath.Join(dir, "[author2] title2"), newComicInfo("title2", "author2", "tags"), false},
		// 3. Empty path
		{"", nil, true},
		// 4. Invalid path contains invalid characters
		{"???", nil, true},
		// 5. Not existing path
		{filepath.Join(dir, "not exists"), nil, true},
	}

	// Generate needed dummy comicinfo directory
	tests[0].want = dummyComicDir(tests[0].folderPath, tests[0].want)

	tests[1].want = dummyComicDir(tests[1].folderPath, tests[1].want)
	tests[1].want.ScanInformation = "abcd"
	comicinfo.Save(tests[1].want, filepath.Join(tests[1].folderPath, "ComicInfo.xml"))

	// Run Tests
	for idx, tt := range tests {
		got, err := ScanBooks(tt.folderPath)

		// Error Checking
		if tt.wantErr {
			assert.Errorf(t, err, "Case %d: Expected error, but no error return.", idx)
		} else {
			assert.NoErrorf(t, err, "Case %d: Unwanted error.", idx)
		}

		// Value checking
		assert.EqualValuesf(t, tt.want, got, "Case %d: Not equal comicInfo", idx)
	}
}

// Test CheckFolder function works in different options.
func TestCheckFolder(t *testing.T) {
	tempDir := t.TempDir()

	// Prepare Test Set
	//  1. Folder contain Image only
	path1 := filepath.Join(tempDir, "folder1")
	os.MkdirAll(path1, 0755)
	file1, _ := os.Create(filepath.Join(path1, "image1.jpg"))
	defer file1.Close()

	//  2. Folder contain another folder only
	path2 := filepath.Join(tempDir, "folder2")
	os.MkdirAll(filepath.Join(path2, "subfolder2"), 0755)

	//  3. Folder contain both subfolder & Image
	path3 := filepath.Join(tempDir, "folder3")
	os.MkdirAll(filepath.Join(path3, "subfolder3"), 0755)
	file2, _ := os.Create(filepath.Join(path3, "image3.jpg"))
	defer file2.Close()

	//  3. Empty Folder
	path4 := filepath.Join(tempDir, "folder4")
	os.MkdirAll(path4, 0755)

	// Start Image Test
	var tests = []struct {
		path     string
		opt      ScanOpt
		want     bool
		hasError bool
	}{
		// Image Opt Test (1~12)
		{path1, ScanOpt{Image: Unspecific}, true, false},
		{path1, ScanOpt{Image: Allow}, true, false},
		{path1, ScanOpt{Image: AllowOnly}, true, false},
		{path1, ScanOpt{Image: Reject}, false, false},

		{path2, ScanOpt{Image: Unspecific}, true, false},
		{path2, ScanOpt{Image: Allow}, false, false},
		{path2, ScanOpt{Image: AllowOnly}, false, false},
		{path2, ScanOpt{Image: Reject}, true, false},

		{path3, ScanOpt{Image: Unspecific}, true, false},
		{path3, ScanOpt{Image: Allow}, true, false},
		{path3, ScanOpt{Image: AllowOnly}, false, false},
		{path3, ScanOpt{Image: Reject}, false, false},

		{path4, ScanOpt{Image: Unspecific}, true, false},
		{path4, ScanOpt{Image: Allow}, false, false},
		{path4, ScanOpt{Image: AllowOnly}, false, false},
		{path4, ScanOpt{Image: Reject}, true, false},

		// Subfolder Test (13~24)
		{path1, ScanOpt{SubFolder: Unspecific}, true, false},
		{path1, ScanOpt{SubFolder: Allow}, false, false},
		{path1, ScanOpt{SubFolder: AllowOnly}, false, false},
		{path1, ScanOpt{SubFolder: Reject}, true, false},

		{path2, ScanOpt{SubFolder: Unspecific}, true, false},
		{path2, ScanOpt{SubFolder: Allow}, true, false},
		{path2, ScanOpt{SubFolder: AllowOnly}, true, false},
		{path2, ScanOpt{SubFolder: Reject}, false, false},

		{path3, ScanOpt{SubFolder: Unspecific}, true, false},
		{path3, ScanOpt{SubFolder: Allow}, true, false},
		{path3, ScanOpt{SubFolder: AllowOnly}, false, false},
		{path3, ScanOpt{SubFolder: Reject}, false, false},

		{path4, ScanOpt{SubFolder: Unspecific}, true, false},
		{path4, ScanOpt{SubFolder: Allow}, false, false},
		{path4, ScanOpt{SubFolder: AllowOnly}, false, false},
		{path4, ScanOpt{SubFolder: Reject}, true, false},

		// Contradiction Test
		{path3, ScanOpt{SubFolder: AllowOnly, Image: Allow}, false, true},
		{path3, ScanOpt{Image: AllowOnly, SubFolder: Allow}, false, true},
	}

	// Loop the test case and check the result
	for i, testCase := range tests {
		result, err := CheckFolder(testCase.path, testCase.opt)

		// Prevent expected has error, but result is nil
		if testCase.hasError && err == nil {
			t.Errorf("Failed Test Case %d. Expected Error occur", i+1)
			continue
		}

		// Prevent expected not error, but result has error
		if !testCase.hasError && err != nil {
			t.Errorf("Failed Test Case %d. Expected Error Free, got %v", i+1, err)
		}

		// Prevent expected value not matched
		if result != testCase.want {
			t.Errorf("Failed Test Case %d. Expected %v, got %v", i+1, testCase.want, result)
		}
	}

}

func TestIsSupportedImg(t *testing.T) {
	type testCase struct {
		filename string
		want     bool
	}

	tests := []testCase{
		{"foo/image1.jpg", true},
		{"foo/image2.png", true},
		{"foo/image3.jpeg", true},
		{"foo/image4.webp", true},

		// Not Supported
		{"foo/abc/", false},
		{"foo/abc.txt", false},
		{"foo/abc.docx", false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, isSupportedImg(tt.filename))
	}
}
