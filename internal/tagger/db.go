// Package for manipulating tags in comic info.
package tagger

import (
	"fmt"
	"gui-comicinfo/internal/database"
)

// Error blocks
var (
	// Error when trying to use nil database in this module.
	ErrAppDBNil = fmt.Errorf("AppDB cannot be nil")
)

// Add tags to given AppDB. This function support multiple tags insert at once.
//
// If tag value is empty string, then it will be skipped.
func AddTag(db *database.AppDB, tags ...string) error {
	// Prevent nil database
	if db == nil {
		return ErrAppDBNil
	}

	// Prepare statement
	stmt, err := db.Prepare("INSERT OR IGNORE INTO tags (input) VALUES (?)")
	if err != nil {
		return err
	}

	// Insert multiple value
	for _, item := range tags {
		// Skip empty string values
		if item == "" {
			continue
		}

		_, err = stmt.Exec(item)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get all tags from given AppDB.
func GetAllTags(db *database.AppDB) ([]string, error) {
	// Prevent nil database
	if db == nil {
		return []string{}, ErrAppDBNil
	}

	// Prepare SQL & its args
	query := "SELECT input FROM tags ORDER BY input"

	// Prepare query
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	// Execute query
	rows, err := stmt.Query()
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
