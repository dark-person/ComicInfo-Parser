package files

import (
	"os"
	"path/filepath"
	"testing"
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
	}

	// Start Test
	for _, tt := range tests {
		if got := IsFileExist(tt.path); got != tt.want {
			t.Errorf("IsFileExist() = %v, want %v", got, tt.want)
		}
	}
}
