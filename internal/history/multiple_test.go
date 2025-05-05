package history

import (
	"path/filepath"
	"testing"

	"github.com/dark-person/lazydb"
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
		{db1, []HistoryVal{{34, "123"}}, false, []int{1}},
		{db1, []HistoryVal{{34, "123"}}, false, []int{1}},
		{db1, []HistoryVal{{45, "123"}}, false, []int{1}},

		// Different values in same time
		{"test2.db", []HistoryVal{{34, "123"}, {45, "123"}}, false, []int{1, 1}},

		// Duplicate Test
		{"test3.db", []HistoryVal{{34, "123"}, {34, "123"}}, false, []int{1, 1}},
		{"test3.db", []HistoryVal{{34, "123"}, {34, "456"}}, false, []int{1, 1}},

		// Empty value
		{"test4.db", []HistoryVal{{34, ""}}, false, []int{0}},

		// Empty string value
		{"test6.db", []HistoryVal{{34, "123"}, {34, ""}}, false, []int{1, 0}},

		// Nil database
		{"", []HistoryVal{{34, "123"}}, true, []int{1}},
	}

	// Start testing
	for idx, tt := range tests {
		// Create new database
		var db *lazydb.LazyDB
		var err error

		if tt.dbPath != "" {
			db, err = createTestDB(filepath.Join(dir, tt.dbPath), false)
			if err != nil {
				t.Errorf("Failed to create database: %v", err)
			}
			defer db.Close()

		} else {
			// Use nil database if dbPath is empty
			db = getNilDatabase()
		}

		// Perform function
		err = InsertMultiple(db, tt.values...)

		// If error expected, check error and contine test as no value need to check
		if tt.wantErr {
			assert.Errorf(t, err, "Case %d: Expected error, but return nil", idx)
			continue
		}

		// Asset no error occur
		assert.NoErrorf(t, err, "Case %d: Unwanted error.", idx)

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
