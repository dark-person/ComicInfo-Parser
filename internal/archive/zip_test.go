package archive

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

const (
	_img1File = "image1.jpg"
	_img2File = "image2.jpg"
	_img3File = "image3.jpg"
	_xmlFile  = "test.xml"
)

// Test Create Zip to destination from input folder.
func TestCreateZipTo(t *testing.T) {
	// Create a temp directory
	tempDir := t.TempDir()

	// Separate two folder
	inputDir := filepath.Join(tempDir, "input")
	os.MkdirAll(inputDir, 0755)

	outputDir := filepath.Join(tempDir, "output")
	os.MkdirAll(outputDir, 0755)

	// Create a set of file
	file1, _ := os.Create(filepath.Join(inputDir, _img1File))
	file2, _ := os.Create(filepath.Join(inputDir, _img2File))
	file3, _ := os.Create(filepath.Join(inputDir, _img3File))
	file4, _ := os.Create(filepath.Join(inputDir, _xmlFile))
	defer file1.Close()
	defer file2.Close()
	defer file3.Close()
	defer file4.Close()

	// Start Testing Functions
	dest, err := CreateZipTo(inputDir, outputDir)
	if err != nil {
		t.Error(err)
	}

	// Check Dest Filename
	destFileName := filepath.Base(inputDir)
	if dest != filepath.Join(outputDir, destFileName+".zip") {
		t.Errorf("Error Destination file: %v", dest)
	}

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

	if !exist1 {
		t.Error("Content 1 missing in zip")
	}

	if !exist2 {
		t.Error("Content 2 missing in zip")
	}

	if !exist3 {
		t.Error("Content 3 missing in zip")
	}

	if !exist4 {
		t.Error("Content 4 missing in zip")
	}
}
