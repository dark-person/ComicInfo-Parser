package database

import (
	"os"
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

// Test Method of Connect().
// This test will NOT consider $HOME directory as a case,
// all tests are using custom path only.
func TestAppDB_Connect(t *testing.T) {
	// Test Case type
	type testCase struct {
		a       *AppDB
		wantErr bool
	}

	// Existing database creation
	existPath := filepath.Join(t.TempDir(), "exist_test.db")
	f, _ := os.Create(existPath)
	f.Close()

	// Prepare Tests
	tests := []testCase{
		// Database that not exist
		{&AppDB{dbPath: filepath.Join(t.TempDir(), "connect_test.db")}, false},
		// Existed database
		{&AppDB{dbPath: existPath}, false},
		// Empty Path
		{&AppDB{dbPath: ""}, true},
	}

	// Run Tests
	for idx, tt := range tests {
		err := tt.a.Connect()

		assert.EqualValuesf(t, tt.wantErr, err != nil,
			"Case %d: Unexpected error result: %v", idx+1, err)

		// Close connection (by sql library but not AppDB)
		if tt.a.db != nil {
			tt.a.db.Close()
		}
	}
}
