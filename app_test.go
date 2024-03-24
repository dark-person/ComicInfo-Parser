package main

import (
	"gui-comicinfo/internal/comicinfo"
	"os"
	"path/filepath"
	"strings"
	"testing"

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
	app := NewApp()

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

// Test `TestQuickExportKomga` function in `app`.
//
// There has some assumptions for this test:
//  1. All file has been copy to `.cbz` correctly
func TestQuickExportKomga(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()
	// tempFolder := "testing"

	// Prepare Test Case
	type testCase struct {
		folder string
		errMsg string
	}

	// Prepare valid content
	valid := filepath.Join(tempFolder, "valid")
	createFolderContent(valid, false)
	expectedFileSize := int64(956) // File size of valid .cbz file

	// Prepare invalid content
	invalid := filepath.Join(tempFolder, "invalid")
	createFolderContent(invalid, false)
	os.MkdirAll(filepath.Join(invalid, "dummy"), 0755)

	// Prepare list of test case
	tests := []testCase{
		// Normal Case
		{valid, ""},
		// folder is invalid - Contains Subfolder
		{invalid, "folder structure is not correct"},
		// folder is invalid - Failed to ReadDir (Assume in ScanBooks), e.g. path that not exist
		{invalid + "2", "system cannot find the file specified"},
		// folder is empty string
		{"", "folder cannot be empty"},
	}

	// Prepare new app
	app := NewApp()

	// Looping
	for idx, tt := range tests {
		errMsg := app.QuickExportKomga(tt.folder)

		if tt.errMsg == "" {
			// Special Handling for Normal case
			assert.EqualValuesf(t, tt.errMsg, errMsg, "Case %d: expected empty error message, but got non-empty.", idx)

			cbzPath := filepath.Join(valid, "valid", "valid.cbz")
			stat, err := os.Stat(cbzPath)

			// File existence
			assert.EqualValuesf(t, false, os.IsNotExist(err),
				"file is not generated for case %d, path=%s", idx, cbzPath)

			// Archive Size matching
			assert.EqualValuesf(t, expectedFileSize, stat.Size(),
				"Wrong file size for case %d: expected %v, got %v", idx, stat.Size(), expectedFileSize)

		} else {
			assert.Containsf(t, errMsg, tt.errMsg, "Case %d: expected error message matched, but not.", idx)
		}
	}
}

// Test `ExportXml` function in `app`.
//
// There has some assumptions for this test:
//  1. xml.MarshalIndent() not cause any errors
//  2. *os.File sync() not cause any errors
func TestExportXml(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()

	type testCase struct {
		dir       string // Only final part of abs path, i.e. no need to include temp folder value
		c         *comicinfo.ComicInfo
		hasErrMsg bool
	}

	tempInfo := comicinfo.New()

	tests := []testCase{
		// 1. Graceful Test
		{filepath.Join(tempFolder, "case1"), &tempInfo, false},
		// 2. nil value of ComicInfo
		{filepath.Join(tempFolder, "case2"), nil, true},
		// 3. Empty path
		{"", &tempInfo, true},
		// 4. Invalid path (invalid character)
		{filepath.Join(tempFolder, "???"), &tempInfo, true},
	}

	// Prepare dummy App
	app := NewApp()

	// Start test
	for idx, tt := range tests {
		msg := app.ExportXml(tt.dir, tt.c)

		// Check error
		assert.EqualValuesf(t, tt.hasErrMsg, msg != "", "Case %d: unexpected error message: %v", idx, msg)

		// Skip file content compare when tests have expected error message
		if tt.hasErrMsg {
			continue
		}

		// Check exported value
		b, err := os.ReadFile(filepath.Join(tt.dir, "ComicInfo.xml"))
		if err != nil {
			t.Errorf("Reading XML in case id %d : %v", idx, err)
		}

		assert.EqualValuesf(t, removeExtraSpace(expectedXML), removeExtraSpace(string(b)),
			"Case %d: unmatched XML.", idx)
	}
}

// Test `ExportCbz` function in `app`, which parameters of `isWrap` is false.
//
// There has some assumptions for this test:
//  1. The content of comicInfo is always correct
//  2. The zip file always content correct image (although this test will check its size)
func TestExportCbz_NoWrap(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()
	// tempFolder := "testing"

	// Prepare test case struct
	type testCase struct {
		inputDir  string
		exportDir string
		comicInfo *comicinfo.ComicInfo
		errMsg    string
	}

	// Prepare Paths
	invalidPath := filepath.Join(tempFolder, "invalid")
	validInputPath := filepath.Join(tempFolder, "validIn")
	validOutputPath := filepath.Join(tempFolder, "validOut")
	os.MkdirAll(validOutputPath, 0755)

	// Create Content
	createFolderContent(validInputPath, false)
	validInfo := comicinfo.New()
	expectedFileSize := int64(889)

	// Prepare Tests
	tests := []testCase{
		// When input dir is invalid (os.Create fails)
		{invalidPath, validOutputPath, &validInfo, "input directory does not exist"},
		// When export dir is invalid (os.Create fails)
		{validInputPath, invalidPath, &validInfo, "export directory does not exist"},
		// When comic info is nil value
		{validInputPath, validOutputPath, nil, "empty comic info"},
		// Normal value
		{validInputPath, validOutputPath, &validInfo, ""},
	}

	// Create a new app
	app := NewApp()

	for idx, tt := range tests {
		errMsg := app.ExportCbz(tt.inputDir, tt.exportDir, tt.comicInfo, false)

		if tt.errMsg == "" {
			// Check if error message is empty
			assert.EqualValuesf(t, tt.errMsg, errMsg, "Case %d, expected empty error message but got non empty", idx)

			// Special Handling for Normal case
			cbzPath := filepath.Join(tt.exportDir, "validIn.cbz")
			stat, err := os.Stat(cbzPath)

			// Check file is exist & archive size is matched with expected
			assert.EqualValuesf(t, false, os.IsNotExist(err), "Case %d, file is not generated in path=%s", idx, cbzPath)
			assert.EqualValuesf(t, expectedFileSize, stat.Size(), "Case %d, unexpected file size for cbz", idx)
		} else {
			// Check error message contains
			assert.Containsf(t, errMsg, tt.errMsg, "Case %d, unmatched error message", idx)
		}
	}
}

// Test `ExportCbz` function in `app`, which parameters of `isWrap` is true.
//
// There has some assumptions for this test:
//  1. The content of comicInfo is always correct
//  2. The zip file always content correct image (although this test will check its size)
func TestExportCbz_Wrap(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()

	// Prepare Test Case Struct
	type testCase struct {
		inputDir  string
		exportDir string
		comicInfo *comicinfo.ComicInfo
		errMsg    string
	}

	// Prepare Paths
	invalidPath := filepath.Join(tempFolder, "invalid")
	validInputPath := filepath.Join(tempFolder, "validIn")
	validOutputPath := filepath.Join(tempFolder, "validOut")
	os.MkdirAll(validOutputPath, 0755)

	// Create Content
	createFolderContent(validInputPath, false)
	validInfo := comicinfo.New()
	expectedFileSize := int64(889)

	// Prepare Tests
	tests := []testCase{
		// When input dir is invalid (os.Create fails)
		{invalidPath, validOutputPath, &validInfo, "input directory does not exist"},
		// When export dir is invalid (os.Create fails)
		{validInputPath, invalidPath, &validInfo, "export directory does not exist"},
		// When comic info is nil value
		{validInputPath, validOutputPath, nil, "empty comic info"},
		// Normal value
		{validInputPath, validOutputPath, &validInfo, ""},
	}

	// Create a new app
	app := NewApp()

	for idx, tt := range tests {
		errMsg := app.ExportCbz(tt.inputDir, tt.exportDir, tt.comicInfo, true)

		if tt.errMsg == "" {
			// Check if error message is empty
			assert.EqualValuesf(t, tt.errMsg, errMsg, "Case %d, expected empty error message but got non empty", idx)

			// Special Handling for Normal case
			cbzPath := filepath.Join(tt.exportDir, "validIn", "validIn.cbz")
			stat, err := os.Stat(cbzPath)

			// Check file is exist & archive size is matched with expected
			assert.EqualValuesf(t, false, os.IsNotExist(err), "Case %d, file is not generated in path=%s", idx, cbzPath)
			assert.EqualValuesf(t, expectedFileSize, stat.Size(), "Case %d, unexpected file size for cbz", idx)
		} else {
			// Check error message contains
			assert.Containsf(t, errMsg, tt.errMsg, "Case %d, unmatched error message", idx)
		}
	}
}

const expectedXML = `<?xml version="1.0"?>
<ComicInfo xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
	<Title></Title>
	<Series></Series>
	<Number></Number>
	<Volume>0</Volume>
	<AlternateSeries></AlternateSeries>
	<AlternateNumber></AlternateNumber>
	<StoryArc></StoryArc>
	<StoryArcNumber></StoryArcNumber>
	<SeriesGroup></SeriesGroup>
	<Summary></Summary>
	<Notes></Notes>
	<Writer></Writer>
	<Publisher></Publisher>
	<Imprint></Imprint>
	<Genre></Genre>
	<Tags></Tags>
	<PageCount>0</PageCount>
	<LanguageISO></LanguageISO>
	<Format></Format>
	<AgeRating></AgeRating>
	<Manga></Manga>
	<Characters></Characters>
	<Teams></Teams>
	<Locations></Locations>
	<ScanInformation></ScanInformation>
	<Pages></Pages>
</ComicInfo>`
