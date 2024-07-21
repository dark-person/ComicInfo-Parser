// Package for manipulating tags in comic info.
package tagger

import (
	"gui-comicinfo/internal/database"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Create a new AppDB with database file,
// and ensure that AppDB is connected & update to latest schema.
//
// Developer should ensure returned AppDB will be closed after usage.
//
// If filename is empty string, a nil AppDB will be returned.
func getAppDB(dir, filename string) (*database.AppDB, error) {
	// Return nil if filename is empty string
	if filename == "" {
		return nil, nil
	}

	db, err := database.NewPathDB(filepath.Join(dir, filename))
	if err != nil {
		return nil, err
	}

	// Connect database
	db.Connect()
	db.StepToLatest()

	// Return db
	return db, nil
}

// Create a test database with some value inserted.
//
// // Developer should ensure returned AppDB will be closed after usage.
func createTestDB(path string) (*database.AppDB, error) {
	a, err := database.NewPathDB(path)
	if err != nil {
		return nil, err
	}

	a.Connect()
	a.StepToLatest()

	// Insert data rows
	sql := `INSERT INTO tags (input) VALUES ('abc'), ('def'), ('ghi')`
	stmt, err := a.Prepare(sql)
	if err != nil {
		return nil, err
	}
	stmt.Exec()

	// Return closed db object
	return a, nil
}

// Function to check how many rows in db has given tag.
func checkRowCount(a *database.AppDB, value string) (int, error) {
	// Get Inserted rows
	stmt, err := a.Prepare("SELECT COUNT(*) FROM tags WHERE input=?")
	if err != nil {
		return -1, err
	}

	// Execute query
	rows, err := stmt.Query(value)
	if err != nil {
		return -1, err
	}

	// Load query result
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	return count, nil
}

func TestAddTag(t *testing.T) {
	// Directory to store db files
	dir := t.TempDir()

	// Test case
	type testCase struct {
		dbPath      string
		tags        []string
		wantErr     bool
		insertedRow []int // Should have same order as `tags`
	}

	tests := []testCase{
		// Normal Test
		{"test1.db", []string{"abc"}, false, []int{1}},
		{"test2.db", []string{"abc", "def"}, false, []int{1, 1}},

		// Duplication Test
		{"test-duplicate.db", []string{"abc", "abc"}, false, []int{1, 1}},

		// Empty string
		{"test3.db", []string{""}, false, []int{0}},
		{"test4.db", []string{"abc", ""}, false, []int{1, 0}},

		// Nil Database
		{"", []string{"abc"}, true, []int{1}},
	}

	for idx, tt := range tests {
		// Create new database
		db, err := getAppDB(dir, tt.dbPath)
		if err != nil {
			t.Errorf("Failed to create database: %v", err)
		}

		if db != nil {
			defer db.Close()
		}

		// Perform function
		err = AddTag(db, tt.tags...)

		// Asset no error occur
		assert.EqualValuesf(t, tt.wantErr, err != nil, "Case %d: Expected has error=%t, got %t", idx, tt.wantErr, err != nil)
		if tt.wantErr {
			continue
		}

		// Asset value has inserted
		for i, val := range tt.tags {
			count, err := checkRowCount(db, val)
			if err != nil {
				t.Errorf("Failed to check row count: %v", err)
			}

			assert.EqualValuesf(t, tt.insertedRow[i], count, "Case %d: Expected inserted=%d, got %d", idx, tt.insertedRow[i], count)
		}
	}
}

func TestGetAllTags(t *testing.T) {
	// Directory to store db files
	dir := t.TempDir()

	// A standard test database will be created
	a, err := createTestDB(filepath.Join(dir, "test_get.db"))
	if err != nil {
		t.Errorf("Failed to create db: %v", err)
	}

	if a != nil {
		defer a.Close()
	}

	// Empty database
	empty, err := database.NewPathDB(filepath.Join(dir, "test_get_empty.db"))
	if err != nil {
		t.Errorf("Failed to create db: %v", err)
	}

	if empty != nil {
		empty.Connect()
		empty.StepToLatest()
		defer empty.Close()
	}

	// Test case
	type testCase struct {
		db      *database.AppDB
		results []string
		wantErr bool
	}

	tests := []testCase{
		// Normal Case
		{a, []string{"abc", "def", "ghi"}, false},

		// Empty result
		{empty, []string{}, false},

		// Nil database
		{nil, []string{}, true},
	}

	// Start testing
	for idx, tt := range tests {
		results, err := GetAllTags(tt.db)

		// Check error
		assert.EqualValuesf(t, tt.wantErr, err != nil, "case %d: expected error=%v, but error=%v", idx, tt.wantErr, err)

		// Check values
		assert.EqualValuesf(t, tt.results, results, "case %d: unexpected output, expect=%v, got=%v", idx, tt.results, results)
	}
}
