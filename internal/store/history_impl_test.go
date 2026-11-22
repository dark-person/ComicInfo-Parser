package store

import (
	"path/filepath"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/lazydb"
	"github.com/stretchr/testify/assert"
)

func TestSplitToHistoryValSeparator(t *testing.T) {
	const (
		str      = "A,B C"
		category = definitions.CategoryGenre
	)

	type testCase struct {
		name string
		sep  string
		want []HistoryVal
	}

	tests := []testCase{
		{
			name: "comma separator",
			sep:  ",",
			want: []HistoryVal{
				{Category: category, Value: "A"},
				{Category: category, Value: "B C"},
			},
		},
		{
			name: "space separator",
			sep:  " ",
			want: []HistoryVal{
				{Category: category, Value: "A,B"},
				{Category: category, Value: "C"},
			},
		},
		{
			name: "separator not found",
			sep:  "|",
			want: []HistoryVal{
				{Category: category, Value: str},
			},
		},
		{
			name: "empty separator",
			sep:  "",
			want: []HistoryVal{
				{Category: category, Value: "A"},
				{Category: category, Value: ","},
				{Category: category, Value: "B"},
				{Category: category, Value: " "},
				{Category: category, Value: "C"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SplitToHistoryVal(str, tt.sep, category))
		})
	}
}

func TestSplitToHistoryValCategory(t *testing.T) {
	const (
		str = "Action,Adventure"
		sep = ","
	)

	type testCase struct {
		name     string
		category definitions.CategoryType
	}

	tests := []testCase{
		{name: "genre", category: definitions.CategoryGenre},
		{name: "publisher", category: definitions.CategoryPublisher},
		{name: "translator", category: definitions.CategoryTranslator},
		{name: "tag", category: definitions.CategoryTag},
		{name: "writer", category: definitions.CategoryWriter},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := []HistoryVal{
				{Category: tt.category, Value: "Action"},
				{Category: tt.category, Value: "Adventure"},
			}

			assert.Equal(t, want, SplitToHistoryVal(str, sep, tt.category))
		})
	}
}

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
			db, err = createTestHistoryDB(filepath.Join(dir, tt.dbPath), false)
			if err != nil {
				t.Errorf("Failed to create database: %v", err)
			}
			defer db.Close()

		} else {
			// Use nil database if dbPath is empty
			db = getNilDatabase()
		}

		// Perform function
		err = InsertMultiHistory(db, tt.values...)

		// If error expected, check error and contine test as no value need to check
		if tt.wantErr {
			assert.Errorf(t, err, "Case %d: Expected error, but return nil", idx)
			continue
		}

		// Asset no error occur
		assert.NoErrorf(t, err, "Case %d: Unwanted error.", idx)

		// Asset value has inserted
		for i, val := range tt.values {
			count, err := checkHistoryRowCount(db, val.Category, val.Value)
			if err != nil {
				t.Errorf("Failed to check row count: %v", err)
			}

			assert.EqualValuesf(t, tt.insertedRow[i], count, "Case %d: Expected inserted=%d, got %d", idx, tt.insertedRow[i], count)
		}
	}
}
