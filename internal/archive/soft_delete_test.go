package archive

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoftDeleteComic(t *testing.T) {
	temp := t.TempDir()

	type testcase struct {
		originDir string
		trashBin  string
		wantErr   bool
	}

	tests := []testcase{
		// Error cases
		{"", "", true},
		{filepath.Join(temp, "abc"), "", true},
		{"", filepath.Join(temp, "abc"), true},
		{filepath.Join(temp, "abc"), filepath.Join(temp, "abc"), true},

		// Normal Case
		{filepath.Join(temp, "normal1"), filepath.Join(temp, "trash1"), false},
	}

	for idx, tt := range tests {
		// Prepare folder to test
		if !tt.wantErr {
			os.MkdirAll(tt.originDir, 0755)
		}

		// ============== Test Start =====================
		err := SoftDeleteComic(tt.originDir, tt.trashBin)

		if tt.wantErr {
			assert.Error(t, err, "Case %d : SoftDelete should return an error", idx)
			continue
		}

		// Check no error and directory
		assert.NoError(t, err, "Case %d : SoftDelete should not return an error", idx)

		// Check if directory moved successfully
		expectedDest := filepath.Join(tt.trashBin, filepath.Base(tt.originDir))

		if _, err := os.Stat(expectedDest); os.IsNotExist(err) {
			t.Errorf("Case %d :original directory should not exist", idx)
		}

		if _, err := os.Stat(tt.originDir); !os.IsNotExist(err) {
			t.Errorf("Case %d :original directory should not exist", idx)
		}
	}
}
