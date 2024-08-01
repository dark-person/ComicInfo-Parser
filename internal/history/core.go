// Package for saving user inputted values to database.
package history

import (
	"fmt"
	"gui-comicinfo/internal/database"
)

// Error blocks
var (
	// Error when trying to use nil database in this module.
	ErrAppDBNil = fmt.Errorf("AppDB cannot be nil")
)

// Insert value into database. This function is allowed to insert multiple values at once.
func insertValue(db *database.AppDB, category string, value ...string) error {
	// Prevent nil database
	if db == nil {
		return ErrAppDBNil
	}

	// Prepare statement
	stmt, err := db.Prepare("INSERT OR IGNORE INTO list_inputted (category, input) VALUES (?, ?)")
	if err != nil {
		return err
	}

	// Insert multiple value
	for _, item := range value {
		// Skip empty string values
		if item == "" {
			continue
		}

		_, err = stmt.Exec(category, item)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get inputted list from database, by given category.
func getHistory(db *database.AppDB, category string) ([]string, error) {
	// Prevent nil database
	if db == nil {
		return []string{}, ErrAppDBNil
	}

	// Prepare SQL & its args
	query := "SELECT input FROM list_inputted WHERE category = ?"

	var args []any
	args = append(args, category)

	// Prepare query
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	// Execute query
	rows, err := stmt.Query(args...)
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
