// Package for saving user inputted values to database.
package store

import (
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/lazydb"
)

// Insert value into database. This function is allowed to insert multiple values at once.
func insertValue(db *lazydb.LazyDB, category definitions.CategoryType, value ...string) error {
	// Prevent nil database
	if db == nil {
		return ErrDatabaseNil
	}

	// Prepare statement
	prepared := make([]lazydb.ParamQuery, 0)

	for _, item := range value {
		// Skip empty string values
		if item == "" {
			continue
		}

		prepared = append(prepared,
			lazydb.Param(
				"INSERT OR IGNORE INTO list_inputted (category, input) VALUES (?, ?)",
				category, item,
			))
	}

	// Execute
	_, err := db.ExecMultiple(prepared)
	return err
}

// Get inputted list from database, by given category.
func getHistory(db *lazydb.LazyDB, category definitions.CategoryType) ([]string, error) {
	// Prevent nil database
	if db == nil {
		return []string{}, ErrDatabaseNil
	}

	// Execute query
	rows, err := db.Query("SELECT input FROM list_inputted WHERE category = ?", category)
	if err != nil {
		return nil, err
	}

	// Load query result
	list := make([]string, 0)
	for rows.Next() {
		var input string
		rows.Scan(&input)

		list = append(list, input)
	}

	return list, nil
}
