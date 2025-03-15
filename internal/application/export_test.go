package application

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/archive"
	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/config"
	"github.com/stretchr/testify/assert"
)

// Test `TestQuickExportKomga` function in `app`.
//
// There has some assumptions for this test:
//  1. All file has been copy to `.cbz` correctly
func TestQuickExportKomga(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()

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

func TestExportCbzNoErr(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()

	// Test database
	testDB := filepath.Join(tempFolder, "test_cbz_ok.db")

	// ComicInfo Content
	validInfo := comicinfo.New()
	const expectedFileSize = int64(889)

	// Prepare test case struct
	type testCase struct {
		inputBaseDir  string // base folder for input
		outputBaseDir string // base folder for output
		opt           archive.RenameOption
		expectedPath  string // expected filepath for output cbz
	}

	// Prepare Tests
	tests := []testCase{
		{"input1", "output1", archive.NoWrap(), "output1/input1.cbz"},
		{"input2", "output2", archive.UseDefaultWrap(), "output2/input2/input2.cbz"},
		{"input3", "output3", archive.UseCustomWrap("abc"), "output3/abc/input3.cbz"},
	}

	// Create a new app
	app := NewApp(assets.DefaultDb(testDB))
	app.Startup(context.TODO())

	// Test case for normal situation
	for idx, tt := range tests {
		inputDir := filepath.Join(tempFolder, tt.inputBaseDir)
		outputDir := filepath.Join(tempFolder, tt.outputBaseDir)

		// Prepare content
		createFolderContent(inputDir, false)
		os.MkdirAll(outputDir, 0755)

		// RUN test case
		errMsg := app.exportCbz(inputDir, outputDir, &validInfo, tt.opt)

		// Check if error message is empty
		assert.EqualValuesf(t, "", errMsg, "Case %d, expected empty error message but got non empty", idx)

		// Special Handling for Normal case
		cbzPath := filepath.Join(tempFolder, tt.expectedPath)
		stat, err := os.Stat(cbzPath)

		// Check file is exist & archive size is matched with expected
		assert.EqualValuesf(t, false, os.IsNotExist(err), "Case %d, file is not generated in path=%s", idx, cbzPath)
		assert.EqualValuesf(t, expectedFileSize, stat.Size(), "Case %d, unexpected file size for cbz", idx)
	}

	// Close database file
	app.DB.Close()
}

func TestExportCbzNilInfo(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()

	// Test database
	testDB := filepath.Join(tempFolder, "test_cbz_nil.db")

	// Create a new app
	app := NewApp(assets.DefaultDb(testDB))
	app.Startup(context.TODO())

	// Prepare content
	inputDir := filepath.Join(tempFolder, "input")
	outputDir := filepath.Join(tempFolder, "output")
	createFolderContent(inputDir, false)
	os.MkdirAll(outputDir, 0755)

	// Test for nil comicinfo
	errMsg := app.exportCbz(inputDir, outputDir, nil, archive.NoWrap())

	// Check error message
	assert.EqualValuesf(t, "empty comic info", errMsg, "Unmatched error message")

	// Close database file
	app.DB.Close()
}

func TestExportCbzErrOs(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()

	// Test database
	testDB := filepath.Join(tempFolder, "test_cbz_os.db")

	// Prepare folders
	inputDir := filepath.Join(tempFolder, "input")
	outputDir := filepath.Join(tempFolder, "output")
	invalidDir := filepath.Join(tempFolder, "invalid")

	// Create contents
	createFolderContent(inputDir, false)
	os.MkdirAll(outputDir, 0755)

	// ComicInfo Content
	validInfo := comicinfo.New()

	// Create a new app
	app := NewApp(assets.DefaultDb(testDB))
	app.Startup(context.TODO())

	// Prepare test case struct
	type testCase struct {
		input  string
		output string
		errMsg string
	}

	// Prepare Tests
	tests := []testCase{
		// When input dir is invalid (os.Create fails)
		{invalidDir, outputDir, "input directory does not exist"},
		// When export dir is invalid (os.Create fails)
		{inputDir, invalidDir, "export directory does not exist"},
	}

	// Start Test
	for idx, tt := range tests {
		// RUN test case
		errMsg := app.exportCbz(tt.input, tt.output, &validInfo, archive.NoWrap())

		// Check error message contains
		assert.Containsf(t, errMsg, tt.errMsg, "Case %d, unmatched error message", idx)
	}

	// Close database file
	app.DB.Close()
}

func TestSoftDelete(t *testing.T) {
	// Temp folder creation
	tempFolder := t.TempDir()

	// Test database
	testDB := filepath.Join(tempFolder, "test_soft_delete.db")

	// Create a new app
	app := NewApp(assets.DefaultDb(testDB))
	app.Startup(context.TODO())

	type testcase struct {
		isInputExist bool
		isExported   bool
		hasTrashBin  bool
		wantErr      bool
	}

	tests := []testcase{
		// Normal
		{true, true, true, false},

		// Not exported
		{true, false, true, true},

		// Without trash bin
		{true, true, false, true},

		// Input directory not exist
		{false, true, true, true},
	}

	// Start test
	for idx, tt := range tests {
		app.cfg = config.Default()
		inputDir := filepath.Join(tempFolder, "case-"+strconv.Itoa(idx))

		// Prepare state for test case
		if tt.isInputExist {
			os.MkdirAll(inputDir, 0755)
		}

		if tt.isExported {
			app.lastExportedComic = inputDir
		}

		if tt.hasTrashBin {
			app.cfg.TrashBin = filepath.Join(tempFolder, "trash")
		}

		// Run function
		msg := app.RunSoftDelete()

		if tt.wantErr {
			assert.NotEmpty(t, msg, "Case %d should be not empty", idx)
		} else {
			assert.Empty(t, msg, "Case %d should be empty", idx)
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
