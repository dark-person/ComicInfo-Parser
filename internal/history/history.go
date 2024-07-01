package history

import (
	"gui-comicinfo/internal/database"
)

func InsertInputted(db *database.AppDB, category string, value string) error {

	stmt, err := db.Prepare("INSERT OR IGNORE INTO list_inputted (category, input) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(category, value)
	if err != nil {
		return err
	}

	return nil
}

// Get inputted list from database, by given part of text.
//
// If text is empty string, then it will select all inputted record from database.
func GetInputtedList(db *database.AppDB, category string, text string) ([]string, error) {
	var query string
	var args []any

	// Determine query by text is empty string or not
	if text == "" {
		query = "SELECT input FROM list_inputted WHERE category = ?"
		args = append(args, category)

	} else {
		query = "SELECT input FROM list_inputted WHERE category = ? AND input LIKE '%?%'"
		args = append(args, category, text)

	}

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
