package main

import (
	"fmt"
	"gui-comicinfo/internal/comicinfo"
	"os"
	"path/filepath"
	"strings"
	"testing"
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

	// Generate a invalid older path
	invalidPath := filepath.Join(tempFolder, "invalid")

	// Generate a folder with incorrect structure (Contain Subfolder)
	incorrectPath := filepath.Join(tempFolder, "incorrect")
	createFolderContent(incorrectPath, true)
	os.MkdirAll(filepath.Join(incorrectPath, "incorrect"), 0755)

	// Generate a folder with correct structure
	correctPath := filepath.Join(tempFolder, "correct")
	createFolderContent(correctPath, true)

	// Case of Input
	caseInput := []string{"", invalidPath, incorrectPath, correctPath}

	// Case of Output. Defined with special usage,
	// The comicInfo only determine it is nil or not,
	// ErrorMessage only check it contains error message stated,
	// as sometime error may include absolute path which is unpredictable
	caseOutput := []ComicInfoResponse{
		{ComicInfo: nil, ErrorMessage: "folder cannot be empty"},
		{ComicInfo: nil, ErrorMessage: "The system cannot find the file specified"},
		{ComicInfo: nil, ErrorMessage: "folder structure is not correct"},
		{ComicInfo: &comicinfo.ComicInfo{}, ErrorMessage: ""},
	}

	// Create a new app
	app := NewApp()
	for i := 0; i < len(caseInput); i++ {
		temp := app.GetComicInfo(caseInput[i])

		fmt.Println("Result: ", temp)
		fmt.Println("Expected: ", caseOutput[i])

		// Check comic info is nil or not
		if caseOutput[i].ComicInfo == nil && temp.ComicInfo != nil {
			t.Errorf("Error when running case %d: ComicInfo expected nil, but got non nil\n", i)
			continue
		} else if caseOutput[i].ComicInfo != nil && temp.ComicInfo == nil {
			t.Errorf("Error when running case %d: ComicInfo expected non nil, but got nil\n", i)
			continue
		}

		// Check error message empty/not empty matches
		if caseOutput[i].ErrorMessage == "" && temp.ErrorMessage != "" {
			t.Errorf("Error when running case %d: ErrorMessage expected empty, but got %s\n", i, temp.ErrorMessage)
			continue
		} else if caseOutput[i].ErrorMessage != "" && temp.ErrorMessage == "" {
			t.Errorf("Error when running case %d: ErrorMessage expected non-empty, but got empty string\n", i)
			continue
		}

		// Check error message similarity
		if caseOutput[i].ErrorMessage != "" && !strings.Contains(temp.ErrorMessage, caseOutput[i].ErrorMessage) {
			t.Errorf("Error when running case %d: ErrorMessage expected contain %s, but got %s\n",
				i, caseOutput[i].ErrorMessage, temp.ErrorMessage)
			continue
		}
	}
}

func TestQuickExportKomga(t *testing.T) {
	// TODO: Implementation
}

// Test `ExportXml` function in `app`.
//
// There has some assumptions for this test:
//  1. xml.MarshalIndent() not cause any errors
//  2. *os.File sync() not cause any errors
func TestExportXml(t *testing.T) {
	// Prepare list of test case
	dirInput := make([]string, 0)
	infoInput := make([]*comicinfo.ComicInfo, 0)
	textOutput := make([]string, 0)

	// Temp folder creation
	tempFolder := t.TempDir()
	// tempFolder := "testing"

	// Check if input comicinfo is nil value
	dirInput = append(dirInput, tempFolder)
	infoInput = append(infoInput, nil)
	textOutput = append(textOutput, "comicinfo is nil value")

	// Demo os.Create() error (target directory doesn't exist)
	dirInput = append(dirInput, filepath.Join(tempFolder, "invalid"))
	infoInput = append(infoInput, &comicinfo.ComicInfo{})
	textOutput = append(textOutput, "The system cannot find the path specified")

	// Check output xml
	c := comicinfo.New()

	dirInput = append(dirInput, tempFolder)
	infoInput = append(infoInput, &c)
	textOutput = append(textOutput, "")

	// Create a new app
	app := NewApp()

	for i := 0; i < len(textOutput); i++ {
		errMsg := app.ExportXml(dirInput[i], infoInput[i])

		// Check error message
		if !strings.Contains(errMsg, textOutput[i]) {
			t.Errorf("Case id = %d: Expected %v, got %v", i, textOutput[i], errMsg)
		}

		// Early Return for error cases
		if textOutput[i] != "" {
			continue
		}

		// Check output xml equals to expected
		b, err := os.ReadFile(filepath.Join(tempFolder, "ComicInfo.xml"))
		if err != nil {
			t.Errorf("Reading XML in case id %d : %v", i, err)
		} else if removeExtraSpace(string(b)) != removeExtraSpace(expectedXML) {
			t.Errorf("Unmatched XML in case id %d: %s vs %s", i, string(b), expectedXML)
		}
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

	// Prepare Paths
	invalidPath := filepath.Join(tempFolder, "invalid")
	validInputPath := filepath.Join(tempFolder, "validIn")
	validOutputPath := filepath.Join(tempFolder, "validOut")
	os.MkdirAll(validOutputPath, 0755)

	// Create Content
	createFolderContent(validInputPath, false)
	validInfo := comicinfo.New()

	// Prepare list for test case
	inputDirList := make([]string, 0)
	exportDirList := make([]string, 0)
	comicInfoList := make([]*comicinfo.ComicInfo, 0)
	errMsgList := make([]string, 0)

	// When input dir is invalid (os.Create fails)
	inputDirList = append(inputDirList, invalidPath)
	exportDirList = append(exportDirList, validOutputPath)
	comicInfoList = append(comicInfoList, &validInfo)
	errMsgList = append(errMsgList, "input directory does not exist")

	// When export dir is invalid (os.Create fails)
	inputDirList = append(inputDirList, validInputPath)
	exportDirList = append(exportDirList, invalidPath)
	comicInfoList = append(comicInfoList, &validInfo)
	errMsgList = append(errMsgList, "export directory does not exist")

	// When comic info is nil value
	inputDirList = append(inputDirList, validInputPath)
	exportDirList = append(exportDirList, validOutputPath)
	comicInfoList = append(comicInfoList, nil)
	errMsgList = append(errMsgList, "empty comic info")

	// Normal value
	inputDirList = append(inputDirList, validInputPath)
	exportDirList = append(exportDirList, validOutputPath)
	comicInfoList = append(comicInfoList, &validInfo)
	errMsgList = append(errMsgList, "")
	expectedFileSize := int64(894)

	// Create a new app
	app := NewApp()

	for i := 0; i < len(errMsgList); i++ {
		errMsg := app.ExportCbz(inputDirList[i], exportDirList[i], comicInfoList[i], false)

		if errMsg == "" && errMsg == errMsgList[i] {
			// Special Handling for Normal case
			cbzPath := filepath.Join(exportDirList[i], "validIn.cbz")

			stat, err := os.Stat(cbzPath)

			// Check file is exist & archive size is matched with expected
			if os.IsNotExist(err) {
				t.Errorf("file is not generated for case %d, path=%s", i, cbzPath)
			} else if stat.Size() != expectedFileSize {
				t.Errorf("Wrong file size for case %d: expected %v, got %v", i, stat.Size(), expectedFileSize)
			}

			continue
		} else if strings.Contains(errMsg, errMsgList[i]) {
			// Pass when error message is highly matched
			continue
		}

		// Error Message not expected
		t.Errorf("Wrong error message for case %d: expected %v, got %v", i, errMsgList[i], errMsg)
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
	// tempFolder := "testing"

	// Prepare Paths
	invalidPath := filepath.Join(tempFolder, "invalid")
	validInputPath := filepath.Join(tempFolder, "validIn")
	validOutputPath := filepath.Join(tempFolder, "validOut")
	os.MkdirAll(validOutputPath, 0755)

	// Create Content
	createFolderContent(validInputPath, false)
	validInfo := comicinfo.New()

	// Prepare list for test case
	inputDirList := make([]string, 0)
	exportDirList := make([]string, 0)
	comicInfoList := make([]*comicinfo.ComicInfo, 0)
	errMsgList := make([]string, 0)

	// When input dir is invalid (os.Create fails)
	inputDirList = append(inputDirList, invalidPath)
	exportDirList = append(exportDirList, validOutputPath)
	comicInfoList = append(comicInfoList, &validInfo)
	errMsgList = append(errMsgList, "input directory does not exist")

	// When export dir is invalid (os.Create fails)
	inputDirList = append(inputDirList, validInputPath)
	exportDirList = append(exportDirList, invalidPath)
	comicInfoList = append(comicInfoList, &validInfo)
	errMsgList = append(errMsgList, "export directory does not exist")

	// When comic info is nil value
	inputDirList = append(inputDirList, validInputPath)
	exportDirList = append(exportDirList, validOutputPath)
	comicInfoList = append(comicInfoList, nil)
	errMsgList = append(errMsgList, "empty comic info")

	// Normal value
	inputDirList = append(inputDirList, validInputPath)
	exportDirList = append(exportDirList, validOutputPath)
	comicInfoList = append(comicInfoList, &validInfo)
	errMsgList = append(errMsgList, "")
	expectedFileSize := int64(894)

	// Create a new app
	app := NewApp()

	for i := 0; i < len(errMsgList); i++ {
		errMsg := app.ExportCbz(inputDirList[i], exportDirList[i], comicInfoList[i], true)

		if errMsg == "" && errMsg == errMsgList[i] {
			// Special Handling for Normal case
			cbzPath := filepath.Join(exportDirList[i], "validIn", "validIn.cbz")

			stat, err := os.Stat(cbzPath)

			// Check file is exist & archive size is matched with expected
			if os.IsNotExist(err) {
				t.Errorf("file is not generated for case %d, path=%s", i, cbzPath)
			} else if stat.Size() != expectedFileSize {
				t.Errorf("Wrong file size for case %d: expected %v, got %v", i, stat.Size(), expectedFileSize)
			}

			continue
		} else if strings.Contains(errMsg, errMsgList[i]) {
			// Pass when error message is highly matched
			continue
		}

		// Error Message not expected
		t.Errorf("Wrong error message for case %d: expected %v, got %v", i, errMsgList[i], errMsg)
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
