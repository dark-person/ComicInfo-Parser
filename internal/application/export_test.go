package application

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/stretchr/testify/assert"
)

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
	app := NewApp(assets.DefaultDb("testing.db"))

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

	// Test database
	testDB := filepath.Join(tempFolder, "test_xml.db")

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
	app := NewApp(assets.DefaultDb(testDB))
	app.Startup(context.TODO())

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

	// Close database file
	app.DB.Close()
}

// Test `ExportCbz` function in `app`, which parameters of `isWrap` is false.
//
// There has some assumptions for this test:
//  1. The content of comicInfo is always correct
//  2. The zip file always content correct image (although this test will check its size)
func TestExportCbzNoWrap(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()
	// tempFolder := "testing"

	// Test database
	testDB := filepath.Join(tempFolder, "test_cbz1.db")

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
	app := NewApp(assets.DefaultDb(testDB))
	app.Startup(context.TODO())

	for idx, tt := range tests {
		errMsg := app.ExportCbzOnly(tt.inputDir, tt.exportDir, tt.comicInfo)

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

	// Close database file
	app.DB.Close()
}

// Test `ExportCbz` function in `app`, which parameters of `isWrap` is true.
//
// There has some assumptions for this test:
//  1. The content of comicInfo is always correct
//  2. The zip file always content correct image (although this test will check its size)
func TestExportCbzWithWrap(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()

	// Test database
	testDB := filepath.Join(tempFolder, "test_cbz2.db")

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
	app := NewApp(assets.DefaultDb(testDB))
	app.Startup(context.TODO())

	for idx, tt := range tests {
		errMsg := app.ExportCbzWithDefaultWrap(tt.inputDir, tt.exportDir, tt.comicInfo)

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

	// Close database file
	app.DB.Close()
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
