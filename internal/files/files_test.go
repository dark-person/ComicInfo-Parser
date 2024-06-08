package files

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFileExist(t *testing.T) {
	// Prepare temp directory
	dir := t.TempDir()

	// Case 1: Exist file
	path1 := filepath.Join(dir, "test1")
	tmp, _ := os.Create(path1)
	defer tmp.Close()

	// Case 2: Not exist file
	path2 := filepath.Join(dir, "test2")

	// Prepare test cases
	tests := []struct {
		path string
		want bool
	}{
		{path1, true},
		{path2, false},
		{"???", false},
	}

	// Start Test
	for _, tt := range tests {
		if got := IsFileExist(tt.path); got != tt.want {
			t.Errorf("IsFileExist() = %v, want %v", got, tt.want)
		}
	}
}

func TestIsPathValid(t *testing.T) {
	// Prepare valid path
	dir := t.TempDir()

	// Struct for perform unit testing
	type testCase struct {
		path string // Path to test
		want bool   // return value
	}

	// Prepare Test
	tests := []testCase{
		// 1. Case for valid & exist path
		{dir, true},
		// 2. Case for valid & not exist path
		{"not exist", true},
		// 3. Case for invalid path
		{"???", false},
	}

	for idx, tt := range tests {
		got := IsPathValid(tt.path)

		// Check errors
		assert.EqualValuesf(t, tt.want, got, "Case %d : result unexpected.", idx)
	}
}

func TestTrimExt(t *testing.T) {
	// Prepare test case struct
	type testCase struct {
		fileName string
		want     string
	}

	// Prepare tests
	tests := []testCase{
		// Filename already trim extension
		{"filename1", "filename1"},
		// Filename not trim extension
		{"filename2.ext", "filename2"},
		// Filename has multiple extension
		{"filename3.tar.gz", "filename3.tar"},
	}

	// Run tests
	for idx, tt := range tests {
		got := TrimExt(tt.fileName)
		assert.EqualValuesf(t, tt.want, got, "Case %d: unexpected result", idx)
	}
}
