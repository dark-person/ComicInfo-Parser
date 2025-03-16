package archive

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDummyZip(zipPath string) {
	os.MkdirAll(filepath.Dir(zipPath), 0755)

	// Create a zip file
	file1, _ := os.Create(zipPath)
	file1.Close()
}

func TestRenameZip(t *testing.T) {
	// Create a temp directory
	tempDir := t.TempDir()

	type testCase struct {
		zipPath      string
		opt          RenameOption
		expectedPath string
	}

	tests := []testCase{
		{"case1.zip", NoWrap(), "case1.cbz"},
		{"case2.zip", UseDefaultWrap(), "case2/case2.cbz"},
		{"case3.zip", UseCustomWrap("def"), "def/case3.cbz"},

		// Check custom wrap with space
		{"case4.zip", UseCustomWrap("ghi "), "ghi/case4.cbz"},
		{"case5.zip", UseCustomWrap(" jkl"), "jkl/case5.cbz"},
	}

	for _, tt := range tests {
		path := filepath.Join(tempDir, tt.zipPath)

		// create dummy zip
		createDummyZip(path)

		// Run function
		err := RenameZip(path, tt.opt)
		assert.NoError(t, err)

		// Check cbz file exists
		dest, openErr := os.Open(filepath.Join(tempDir, tt.expectedPath))
		assert.NoError(t, openErr)
		defer dest.Close()
	}
}
