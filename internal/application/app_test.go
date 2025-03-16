package application

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/stretchr/testify/assert"
)

// Please note that, the getDirectory method will not be tested
// due to it will call a file selector, instead of input & output that can be control.

// Create folder content that similar to the real environment.
// The folder will contain 3 images, with different image size.
//
// An xml file will be created if `withXml` flag is true.
//
// Please note that this function will try to create all necessary folders.
func createFolderContent(tempDir string, withXml bool) {
	tempDir = strings.TrimSpace(tempDir)

	// Create folder first
	os.MkdirAll(tempDir, 0755)

	// Create Three Image file
	fileNames := []string{"image1.jpg", "image2.png", "image3.jpeg"}
	fileSizes := []int64{1234, 3456, 789}

	// Include XML if flag is true
	if withXml {
		fileNames = append(fileNames, "test.xml")
		fileSizes = append(fileSizes, 12)
	}

	for i, filename := range fileNames {
		file, _ := os.Create(filepath.Join(tempDir, filename))
		file.Truncate(fileSizes[i])
		defer file.Close()
	}
}

// Replace duplicate space to one space.
// This method is to prevent compare two XML string with same contents
// and recognize as different, when only tab indentation is not the same.
func removeExtraSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Test Get ComicInfo function can return expected result.
// This test only perform a brief check on the value returned, i.e. nil value or not
func TestGetComicInfo(t *testing.T) {
	tempFolder := t.TempDir()

	// Test Case Structure
	type testCase struct {
		folder   string // Folder name as parameter
		expected ComicInfoResponse
	}

	// Generate a invalid older path
	invalidPath := filepath.Join(tempFolder, "invalid")

	// Generate a folder with incorrect structure (Contain Subfolder)
	incorrectPath := filepath.Join(tempFolder, "incorrect")
	createFolderContent(incorrectPath, true)
	os.MkdirAll(filepath.Join(incorrectPath, "incorrect"), 0755)

	// Generate a folder with correct structure
	correctPath := filepath.Join(tempFolder, "correct")
	createFolderContent(correctPath, true)

	// Prepare Test Case. Output is defined with special usage,
	// The comicInfo only determine it is nil or not,
	// ErrorMessage only check it contains error message stated,
	// as sometime error may include absolute path which is unpredictable
	tests := []testCase{
		{correctPath, ComicInfoResponse{ComicInfo: &comicinfo.ComicInfo{}, ErrorMessage: ""}},
		{"", ComicInfoResponse{ComicInfo: nil, ErrorMessage: "folder cannot be empty"}},
		{invalidPath, ComicInfoResponse{ComicInfo: nil, ErrorMessage: "The system cannot find the file specified"}},
		{incorrectPath, ComicInfoResponse{ComicInfo: nil, ErrorMessage: "folder structure is not correct"}},
	}

	// Create a new app
	app := NewApp(assets.DefaultDb("test.db"))

	// Start Test
	for idx, tt := range tests {
		temp := app.GetComicInfo(tt.folder)

		// Check comic info is nil or not
		if tt.expected.ComicInfo == nil {
			assert.Nilf(t, temp.ComicInfo, "Case %d: Expect Comicinfo to be nil, got non nil.", idx)
		} else {
			assert.NotNilf(t, temp.ComicInfo, "Case %d: Expect Comicinfo to non nil, got nil.", idx)
		}

		// Check Error Message
		if tt.expected.ErrorMessage == "" {
			assert.EqualValuesf(t, tt.expected.ErrorMessage, temp.ErrorMessage,
				"Case %d: expect error message to be empty, got %v", idx, temp.ErrorMessage)
		} else {
			assert.Containsf(t, temp.ErrorMessage, tt.expected.ErrorMessage,
				"Case %d: expected error message matches, but not.", idx)
		}
	}
}
