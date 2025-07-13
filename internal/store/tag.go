package store

import "github.com/dark-person/lazydb"

// Add tags to given LazyDB. This function support multiple tags insert at once.
//
// If tag value is empty string, then it will be skipped.
func AddTag(db *lazydb.LazyDB, tags ...string) error {
	// Prevent nil database
	if db == nil {
		return ErrDatabaseNil
	}

	// Prepare statement
	prepared := make([]lazydb.ParamQuery, 0)

	for _, item := range tags {
		// Skip empty string values
		if item == "" {
			continue
		}

		prepared = append(prepared, lazydb.Param(
			"INSERT OR IGNORE INTO word_store (word, category_id) VALUES (?, 4)",
			item,
		))
	}

	// Execute
	_, err := db.ExecMultiple(prepared)
	return err
}

// Get all tags from given LazyDB.
func GetAllTags(db *lazydb.LazyDB) ([]string, error) {
	// Prevent nil database
	if db == nil {
		return []string{}, ErrDatabaseNil
	}

	// Prepare SQL & its args
	query := "SELECT word FROM word_store WHERE category_id = 4 ORDER BY word"

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
