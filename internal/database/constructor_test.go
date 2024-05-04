package database

import (
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// Act as an smoke test. It will not check the content of AppDB.
//
// This test will not perform any actions on file,
// to prevent corrupt original database that may be existed in current user $HOME.
func TestNewDB(t *testing.T) {
	_, err := NewDB()
	assert.NoErrorf(t, err, "Failed to create new AppDB (Smoke Test): %v", err)
}

// Test function of AppDB. It will not check content of AppDB,
// but will perform check for nil value.
func Test_new(t *testing.T) {
	type testCase struct {
		path    string // absolute path, non-nil value
		wantNil bool   // Determine *AppDB should not be nil
		wantErr bool   // Determine function should return error
	}

	// Prepare Test Case
	tests := []testCase{
		// Graceful Case
		{filepath.Join(t.TempDir(), "test1.db"), false, false},
		// Database File cannot be created (Not Test every case)
		{"", true, true},
	}

	// Start Test
	for idx, tt := range tests {
		got, err := new(tt.path)

		// Check error
		assert.EqualValuesf(t, tt.wantErr, err != nil, "Case %d: Unexpected error: %v", idx+1, err)

		// Check nil value
		if tt.wantNil {
			assert.Nilf(t, got, "Case %d, unexpected value of non-nil: %v", idx+1, got)
		} else {
			assert.NotNilf(t, got, "Case %d, unexpected value of nil.", idx+1)
		}
	}
}
