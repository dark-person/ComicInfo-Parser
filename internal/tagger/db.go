// Package for manipulating tags in comic info.
package tagger

import (
	"fmt"

	"github.com/dark-person/lazydb"
)

// Error blocks
var (
	// Error when trying to use nil database in this module.
	ErrAppDBNil = fmt.Errorf("AppDB cannot be nil")
)

// Add tags to given AppDB. This function support multiple tags insert at once.
//
// If tag value is empty string, then it will be skipped.
func AddTag(db *lazydb.LazyDB, tags ...string) error {
	// Prevent nil database
	if db == nil {
		return ErrAppDBNil
	}

	// Prepare statement
	prepared := make([]lazydb.ParamQuery, 0)

	for _, item := range tags {
		// Skip empty string values
		if item == "" {
			continue
		}

		prepared = append(prepared, lazydb.Param(
			"INSERT OR IGNORE INTO tags (input) VALUES (?)",
			item,
		))
	}

	// Execute
	_, err := db.ExecMultiple(prepared)
	return err
}

// Get all tags from given AppDB.
func GetAllTags(db *lazydb.LazyDB) ([]string, error) {
	// Prevent nil database
	if db == nil {
		return []string{}, ErrAppDBNil
	}

	// Prepare SQL & its args
	query := "SELECT input FROM tags ORDER BY input"

	// Execute query
	rows, err := db.Query(query)
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
