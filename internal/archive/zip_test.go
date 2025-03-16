package archive

import (
	"archive/zip"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_img1File = "image1.jpg"
	_img2File = "image2.jpg"
	_img3File = "image3.jpg"
	_xmlFile  = "test.xml"
)

func createDummyFolder(inputDir string) {
	inputDir = strings.TrimSpace(inputDir)

	// Create folder
	os.MkdirAll(inputDir, 0755)

	// Create files
	file1, _ := os.Create(filepath.Join(inputDir, _img1File))
	file2, _ := os.Create(filepath.Join(inputDir, _img2File))
	file3, _ := os.Create(filepath.Join(inputDir, _img3File))
	file4, _ := os.Create(filepath.Join(inputDir, _xmlFile))
	defer file1.Close()
	defer file2.Close()
	defer file3.Close()
	defer file4.Close()
}

// Test Create Zip to destination from input folder.
func TestCreateZipTo(t *testing.T) {
	// Create a temp directory
	tempDir := t.TempDir()

	// Prepare test case
	type testCase struct {
		inputDir  string
		outputDir string
		destZip   string
	}

	tests := []testCase{
		{"input1", "output1", "output1/input1.zip"},
	}

	for _, tt := range tests {
		// Create input
		inputDir := filepath.Join(tempDir, tt.inputDir)
		createDummyFolder(inputDir)

		// Create destination
		outputDir := filepath.Join(tempDir, tt.outputDir)
		os.MkdirAll(outputDir, 0755)

		// Start Testing Functions
		dest, err := CreateZipTo(inputDir, outputDir)
		assert.Nil(t, err, "Unexpected error")

		// Check Dest Filename
		assert.EqualValuesf(t, dest, filepath.Join(tempDir, tt.destZip), "Error destination file.")

		// Check Zip Content
		reader, err := zip.OpenReader(dest)
		if err != nil {
			t.Error(err)
		}
		defer reader.Close()

		list := make(map[string]int, 0)
		for _, f := range reader.File {
			list[f.Name] = 1
		}

		_, exist1 := list[_xmlFile]
		_, exist2 := list[_img1File]
		_, exist3 := list[_img2File]
		_, exist4 := list[_img3File]

		assert.True(t, exist1, "XML missing in zip")
		assert.True(t, exist2, "content 2 missing in zip")
		assert.True(t, exist3, "content 3 missing in zip")
		assert.True(t, exist4, "content 4 missing in zip")
	}
}
