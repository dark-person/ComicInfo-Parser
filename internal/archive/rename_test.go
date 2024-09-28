package archive

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	_testZip = "hello.zip"
	_testCbz = "hello.cbz"
)

// Test Rename Zip archive to .cbz archive, with wrap option is enabled.
func TestRenameZipWrap(t *testing.T) {
	// Create a temp directory
	tempDir := t.TempDir()

	// Create a folder inside temp directory
	os.MkdirAll(filepath.Join(tempDir, "tmp"), 0755)

	// Prepare zip path
	zipPath := filepath.Join(tempDir, "tmp", _testZip)

	// Create a zip file
	file1, _ := os.Create(zipPath)
	file1.Close()

	// Test Function
	err := RenameZip(zipPath, true)
	if err != nil {
		t.Error(err)
	}

	// Result Verify
	dest, openErr := os.Open(filepath.Join(tempDir, "tmp", "hello", _testCbz))
	if openErr != nil {
		t.Error(openErr)
	}
	defer dest.Close()
}

// Test Rename Zip archive to .cbz archive, with wrap options is disabled
func TestRenameZipNoWrap(t *testing.T) {
	// Create a temp directory
	tempDir := t.TempDir()

	// Create a folder inside temp directory
	os.MkdirAll(filepath.Join(tempDir, "tmp"), 0755)

	// Prepare zip path
	zipPath := filepath.Join(tempDir, "tmp", _testZip)

	// Create a zip file
	file1, _ := os.Create(zipPath)
	file1.Close()

	// Test Function
	err := RenameZip(zipPath, false)
	if err != nil {
		t.Error(err)
	}

	// Result Verify
	dest, openErr := os.Open(filepath.Join(tempDir, "tmp", _testCbz))
	if openErr != nil {
		t.Error(openErr)
	}
	defer dest.Close()
}
