package store

import "github.com/dark-person/lazydb"

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
