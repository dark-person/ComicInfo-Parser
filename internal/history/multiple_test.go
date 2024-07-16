package history

import (
	"gui-comicinfo/internal/database"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test function of `InsertMultiple()`.
// Since this function is using core function in `core.go`,
// this test will only act as a simple smoke tests.
func TestInsertMultiple(t *testing.T) {
	const db1 = "test1.db"

	// Directory to store db files
	dir := t.TempDir()

	// Test case
	type testCase struct {
		dbPath      string
		values      []HistoryVal
		wantErr     bool
		insertedRow []int // Should have same order as `value`
	}

	tests := []testCase{
		// Normal test in same db
		{db1, []HistoryVal{{"abc", "123"}}, false, []int{1}},
		{db1, []HistoryVal{{"abc", "123"}}, false, []int{1}},
		{db1, []HistoryVal{{"def", "123"}}, false, []int{1}},

		// Different values in same time
		{"test2.db", []HistoryVal{{"abc", "123"}, {"def", "123"}}, false, []int{1, 1}},

		// Duplicate Test
		{"test3.db", []HistoryVal{{"abc", "123"}, {"abc", "123"}}, false, []int{1, 1}},
		{"test3.db", []HistoryVal{{"abc", "123"}, {"abc", "456"}}, false, []int{1, 1}},

		// Empty value
		{"test4.db", []HistoryVal{{"abc", ""}}, false, []int{0}},
		{"test5.db", []HistoryVal{{"", "123"}}, false, []int{1}},

		// Empty string value
		{"test6.db", []HistoryVal{{"abc", "123"}, {"abc", ""}}, false, []int{1, 0}},

		// Nil database
		{"", []HistoryVal{{"abc", "123"}}, true, []int{1}},
	}

	// Start testing
	for idx, tt := range tests {
		// Create new database
		var db *database.AppDB
		var err error

		if tt.dbPath != "" {
			db, err = database.NewPathDB(filepath.Join(dir, tt.dbPath))
			if err != nil {
				t.Errorf("Failed to create database: %v", err)
			}

			// Connect database
			db.Connect()
			db.StepToLatest()
			defer db.Close()

		} else {
			// Use nil database if dbPath is empty
			db = getNilDatabase()
		}

		// Perform function
		err = InsertMultiple(db, tt.values...)

		// Asset no error occur
		assert.EqualValuesf(t, tt.wantErr, err != nil, "Case %d: Expected has error=%t, got %t", idx, tt.wantErr, err != nil)
		if tt.wantErr {
			continue
		}

		// Asset value has inserted
		for i, val := range tt.values {
			count, err := checkRowCount(db, val.Category, val.Value)
			if err != nil {
				t.Errorf("Failed to check row count: %v", err)
			}

			assert.EqualValuesf(t, tt.insertedRow[i], count, "Case %d: Expected inserted=%d, got %d", idx, tt.insertedRow[i], count)
		}
	}
}
