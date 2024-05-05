package database

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createFile(t *testing.T) {
	prepareTest()

	// Test Case Type
	type testCase struct {
		dbPath  string // Database path, as this function only accept this
		wantErr bool
	}

	// Prepare a existed database
	existPath := filepath.Join(t.TempDir(), "existed.db")
	f, _ := os.Create(existPath)
	f.Close()

	// Prepare Test Case, with direct declare with `&AppDB{}`
	tests := []testCase{
		// Valid Path, not exist database
		{filepath.Join(t.TempDir(), "new.db"), false},
		// Valid Path, existed database
		{existPath, false},
		// Invalid path, with directory not exist
		{filepath.Join(t.TempDir(), "not_exist", "new.db"), true},
		// Empty filepath
		{"", true},
	}

	for idx, tt := range tests {
		// Run Create
		err := createFile(tt.dbPath)
		assert.EqualValuesf(t, tt.wantErr, err != nil, "Case %d - Unexpected result error: %v", idx+1, err)
	}
}
