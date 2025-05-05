package store

import (
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/lazydb"
)

// Insert genre value from database.
func InsertGenre(db *lazydb.LazyDB, value ...string) error {
	return insertValue(db, definitions.CategoryGenre, value...)
}

// Get all genre value that from database.
func GetGenreList(db *lazydb.LazyDB) ([]string, error) {
	return getHistory(db, definitions.CategoryGenre)
}

// Insert publisher value from database.
func InsertPublisher(db *lazydb.LazyDB, value ...string) error {
	return insertValue(db, definitions.CategoryPublisher, value...)
}

// Get all publisher value that from database.
func GetPublisherList(db *lazydb.LazyDB) ([]string, error) {
	return getHistory(db, definitions.CategoryPublisher)
}

// ================================================================================

// Container for history module,
// which wraps all necessary value will be used to insert value.
//
// This type is designed for insert value with different category at a time.
type HistoryVal struct {
	Category definitions.CategoryType // category to be inserted
	Value    string                   // value to be inserted
}

// Insert values to database, which can include multiple record with different categories.
//
// The string values & category values should be wrapped in `HistoryObj`.
//
// If any error is occur during insert, the insert process will stop and return its error.
func InsertMultiple(db *lazydb.LazyDB, values ...HistoryVal) error {
	// Prevent nil database
	if db == nil {
		return ErrDatabaseNil
	}

	// Loop value item with core function
	for _, val := range values {
		err := insertValue(db, val.Category, val.Value)

		// Stop loop if any error occurs
		if err != nil {
			return err
		}
	}

	return nil
}
