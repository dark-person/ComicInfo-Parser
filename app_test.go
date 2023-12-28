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
// The folder will contain 3 images, and 1 xml file inside.
// Four file will have different file size.
func createFolderContent(tempDir string) {
	// Create Four file, one is not image
	fileNames := []string{"image1.jpg", "image2.png", "image3.jpeg", "test.xml"}
	fileSizes := []int64{1234, 3456, 789, 12}

	for i, filename := range fileNames {
		file, _ := os.Create(filepath.Join(tempDir, filename))
		file.Truncate(fileSizes[i])

		defer file.Close()
	}
}

// Test Get ComicInfo function can return expected result.
// This test only perform a brief check on the value returned, i.e. nil value or not
func TestGetComicInfo(t *testing.T) {
	tempFolder := t.TempDir()

	// Generate a invalid older path
	invalidPath := filepath.Join(tempFolder, "invalid")

	// Generate a folder with incorrect structure (Contain Subfolder)
	incorrectPath := filepath.Join(tempFolder, "incorrect")
	os.MkdirAll(incorrectPath, 0755)
	createFolderContent(incorrectPath)
	os.MkdirAll(filepath.Join(incorrectPath, "incorrect"), 0755)

	// Generate a folder with correct structure
	correctPath := filepath.Join(tempFolder, "correct")
	os.MkdirAll(correctPath, 0755)
	createFolderContent(correctPath)

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

func TestExportXml(t *testing.T) {
	// TODO: Implementation
}

func TestExportCbz_NoWrap(t *testing.T) {
	// TODO: Implementation
}

func TestExportCbz_Wrap(t *testing.T) {
	// TODO: Implementation
}
