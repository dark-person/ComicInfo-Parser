package store

import (
	"path/filepath"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/lazydb"
	"github.com/stretchr/testify/assert"
)

// Function to check how many rows in db has given category & value.
func checkHistoryRowCount(a *lazydb.LazyDB, category definitions.CategoryType, value string) (int, error) {
	// Get Inserted rows
	rows, err := a.Query("SELECT COUNT(*) FROM word_store WHERE category_id=? AND word=?", category, value)
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

// Create a opened connection lazydb for testing purposes.
// This database is using almost default migration setting, and have some test data inserted already.
func createTestHistoryDB(path string, withData bool) (*lazydb.LazyDB, error) {
	a := assets.DefaultDB(path)
	err := a.Connect()
	if err != nil {
		return nil, err
	}

	_, err = a.Migrate()
	if err != nil {
		return nil, err
	}

	// Early return if no need to insert data
	if !withData {
		return a, nil
	}

	// Insert data rows
	sql := `INSERT INTO word_store (category_id, word) VALUES (45,'123'), (56, '123'), (56, '456')`
	_, err = a.Exec(sql)
	if err != nil {
		return nil, err
	}

	// Return db object
	return a, nil
}

func getNilDatabase() *lazydb.LazyDB {
	return nil
}

func TestInsertValue(t *testing.T) {
	const db1 = "test1.db"

	// Directory to store db files
	dir := t.TempDir()

	// Test case
	type testCase struct {
		dbPath      string
		category    definitions.CategoryType
		value       []string
		wantErr     bool
		insertedRow []int // Should have same order as `value`
	}

	tests := []testCase{
		// Normal test in same db
		{db1, 45, []string{"123"}, false, []int{1}},
		{db1, 45, []string{"123"}, false, []int{1}},
		{db1, 56, []string{"123"}, false, []int{1}},

		// Duplicate Test
		{"test2.db", 45, []string{"123", "123"}, false, []int{1, 1}},
		{"test2.db", 45, []string{"123", "456"}, false, []int{1, 1}},

		// Empty value
		{"test3.db", 45, []string{}, false, []int{}},
		{"test4.db", 0, []string{"123"}, false, []int{1}},

		// Empty string value
		{"test5.db", 45, []string{"123", ""}, false, []int{1, 0}},

		// Nil database
		{"", 45, []string{"123"}, true, []int{1}},
	}

	// Start testing
	for idx, tt := range tests {
		// Create new database
		var db *lazydb.LazyDB
		var err error

		if tt.dbPath != "" {
			db, err = createTestHistoryDB(filepath.Join(dir, tt.dbPath), false)
			if err != nil {
				t.Errorf("Failed to create database for case %d: %v", idx, err)
			}
			defer db.Close()

		} else {
			// Use nil database if dbPath is empty
			db = getNilDatabase()
		}

		// Perform function
		err = insertValue(db, tt.category, tt.value...)

		// If error expected, check error and contine test as no value need to check
		if tt.wantErr {
			assert.Errorf(t, err, "Case %d: Expected error, but return nil", idx)
			continue
		}

		// Asset no error occur
		assert.NoErrorf(t, err, "Case %d: Unwanted error.", idx)

		// Asset value has inserted
		for i, val := range tt.value {
			count, err := checkHistoryRowCount(db, tt.category, val)
			if err != nil {
				t.Errorf("Failed to check row count: %v", err)
			}

			assert.EqualValuesf(t, tt.insertedRow[i], count, "Case %d: Expected inserted=%d, got %d", idx, tt.insertedRow[i], count)
		}
	}
}

func TestGetHistory(t *testing.T) {
	// Directory to store db files
	dir := t.TempDir()

	// Prepare a database with given data rows
	a, err := createTestHistoryDB(filepath.Join(dir, "t.db"), true)
	if err != nil {
		panic("failed to create database: " + err.Error())
	}
	defer a.Close()

	// Prepare test case
	type testCase struct {
		category definitions.CategoryType
		result   []string
		wantErr  bool
	}

	tests := []testCase{
		{45, []string{"123"}, false},
		{56, []string{"123", "456"}, false},
		{77, []string{}, false},
	}

	// Start Testing
	for idx, tt := range tests {
		results, err := getHistory(a, tt.category)

		// Check error
		if tt.wantErr {
			assert.Errorf(t, err, "Case %d: expected error, but nil returned", idx)
		} else {
			assert.NoErrorf(t, err, "Case %d: Unwanted error", idx)
		}

		// Check values
		assert.EqualValuesf(t, tt.result, results, "unexpected output, expect=%v, got=%v", tt.result, results)
	}
}

func TestGetHistoryNilDB(t *testing.T) {
	// Prepare a database with nil database
	a := getNilDatabase()

	// Prepare test case
	type testCase struct {
		category definitions.CategoryType
		result   []string
		wantErr  bool
	}

	tests := []testCase{
		{45, []string{"123"}, true},
	}

	// Start Testing
	for idx, tt := range tests {
		results, err := getHistory(a, tt.category)

		// If want error, then check error and skip value checking
		if tt.wantErr {
			assert.Errorf(t, err, "Case %d: Expected error, but nil returned.", idx)
			continue
		}

		// Check error
		assert.NoErrorf(t, err, "Case %d: Unwanted error.", idx)

		// Check values
		assert.EqualValuesf(t, tt.result, results, "unexpected output, expect=%v, got=%v", tt.result, results)
	}
}
