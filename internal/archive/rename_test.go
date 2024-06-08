package archive

import (
	"os"
	"path/filepath"
	"testing"
)

// Test Rename Zip archive to .cbz archive, with wrap option is enabled.
func TestRenameZip_Wrap(t *testing.T) {
	// Create a temp directory
	tempDir := t.TempDir()

	// Create a folder inside temp directory
	os.MkdirAll(filepath.Join(tempDir, "tmp"), 0755)

	// Create a zip file
	file1, _ := os.Create(filepath.Join(tempDir, "tmp", "hello.zip"))
	file1.Close()

	// Test Function
	err := RenameZip(filepath.Join(tempDir, "tmp", "hello.zip"), true)
	if err != nil {
		t.Error(err)
	}

	// Result Verify
	dest, openErr := os.Open(filepath.Join(tempDir, "tmp", "hello", "hello.cbz"))
	if openErr != nil {
		t.Error(openErr)
	}
	defer dest.Close()
}

// Test Rename Zip archive to .cbz archive, with wrap options is disabled
func TestRenameZip_NoWrap(t *testing.T) {
	// Create a temp directory
	tempDir := t.TempDir()

	// Create a folder inside temp directory
	os.MkdirAll(filepath.Join(tempDir, "tmp"), 0755)

	// Create a zip file
	file1, _ := os.Create(filepath.Join(tempDir, "tmp", "hello.zip"))
	file1.Close()

	// Test Function
	err := RenameZip(filepath.Join(tempDir, "tmp", "hello.zip"), false)
	if err != nil {
		t.Error(err)
	}

	// Result Verify
	dest, openErr := os.Open(filepath.Join(tempDir, "tmp", "hello.cbz"))
	if openErr != nil {
		t.Error(openErr)
	}
	defer dest.Close()
}
