// Package for saving user inputted values to database.
package history

import (
	"gui-comicinfo/internal/database"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Function to check how many rows in db has given category & value.
func checkRowCount(a *database.AppDB, category string, value string) (int, error) {
	// Get Inserted rows
	stmt, err := a.Prepare("SELECT COUNT(*) FROM list_inputted WHERE category=? AND input=?")
	if err != nil {
		return -1, err
	}

	// Execute query
	rows, err := stmt.Query(category, value)
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

func createTestDB(path string) (*database.AppDB, error) {
	a, err := database.NewPathDB(path)
	if err != nil {
		return nil, err
	}

	a.Connect()
	a.StepToLatest()
	defer a.Close()

	// Insert data rows
	sql := `INSERT INTO list_inputted (category, input) VALUES ('abc','123'), ('def', '123'), ('def', '456')`
	stmt, err := a.Prepare(sql)
	if err != nil {
		return nil, err
	}
	stmt.Exec()

	// Return closed db object
	return a, nil
}

func TestInsertValue(t *testing.T) {
	const db1 = "test1.db"

	// Directory to store db files
	dir := t.TempDir()

	// Test case
	type testCase struct {
		dbPath      string
		category    string
		value       []string
		wantErr     bool
		insertedRow []int // Should have same order as `value`
	}

	tests := []testCase{
		// Normal test in  same db
		{db1, "abc", []string{"123"}, false, []int{1}},
		{db1, "abc", []string{"123"}, false, []int{1}},
		{db1, "def", []string{"123"}, false, []int{1}},

		// Duplicate Test
		{"test2.db", "abc", []string{"123", "123"}, false, []int{1, 1}},
		{"test2.db", "abc", []string{"123", "456"}, false, []int{1, 1}},

		// Empty value
		{"test3.db", "abc", []string{}, false, []int{}},
		{"test4.db", "", []string{"123"}, false, []int{1}},
	}

	// Start testing
	for idx, tt := range tests {
		// Create new database
		db, err := database.NewPathDB(filepath.Join(dir, tt.dbPath))
		if err != nil {
			t.Errorf("Failed to create database: %v", err)
		}

		// Connect database
		db.Connect()
		db.StepToLatest()
		defer db.Close()

		// Perform function
		err = insertValue(db, tt.category, tt.value...)

		// Asset no error occur
		assert.EqualValuesf(t, tt.wantErr, err != nil, "Case %d: Expected has error=%t, got %t", idx, tt.wantErr, err != nil)

		// Asset value has inserted
		for i, val := range tt.value {
			count, err := checkRowCount(db, tt.category, val)
			if err != nil {
				t.Errorf("Failed to check row count: %v", err)
			}

			assert.EqualValuesf(t, tt.insertedRow[i], count, "Case %d: Expected inserted=%d, got %d", idx, tt.insertedRow[i], count)
		}
	}
}

func TestGetHistory(t *testing.T) {
	// Prepare a database with given data rows
	a, err := createTestDB("testing/t.db")
	if err != nil {
		panic("failed to create database: " + err.Error())
	}
	a.Connect()
	defer a.Close()

	// Prepare test case
	type testCase struct {
		category string
		result   []string
		wantErr  bool
	}

	tests := []testCase{
		{"abc", []string{"123"}, false},
		{"def", []string{"123", "456"}, false},
		{"kk", []string{}, false},
	}

	// Start Testing
	for _, tt := range tests {
		results, err := getHistory(a, tt.category)

		// Check error
		assert.EqualValuesf(t, tt.wantErr, err != nil, "expected error=%v, but error=%v", tt.wantErr, err)

		// Check values
		assert.EqualValuesf(t, tt.result, results, "unexpected output, expect=%v, got=%v", tt.result, results)
	}
}
