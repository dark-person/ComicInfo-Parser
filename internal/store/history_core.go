package store

import (
	"strings"

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
				"INSERT OR IGNORE INTO word_store (category_id, word) VALUES (?, ?)",
				category, strings.TrimSpace(item),
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
	rows, err := db.Query("SELECT word FROM word_store WHERE category_id = ?", category)
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

// Word that will be used in auto fill function.
type AutofillWord struct {
	ID       int    `json:"id"`       // ID of word
	Word     string `json:"word"`     // word that user has inputted once and will be re-used when autofill
	Category string `json:"category"` // Category of the word, in human readable format
}

// Get all word that stored in database and will be used when autofill.
func GetAllAutofillWord(db *lazydb.LazyDB) ([]AutofillWord, error) {
	// Prevent nil database
	if db == nil {
		return []AutofillWord{}, ErrDatabaseNil
	}

	// Execute query
	rows, err := db.Query(`
		SELECT 
			w.word_id, w.word, c.category_name 
		FROM 
			word_store AS w 
		LEFT JOIN 
			category AS c ON w.category_id = c.category_id 
		ORDER BY w.word`)

	if err != nil {
		return nil, err
	}

	// Load query result
	list := make([]AutofillWord, 0)
	for rows.Next() {
		var word AutofillWord
		rows.Scan(&word.ID, &word.Word, &word.Category)

		list = append(list, word)
	}

	return list, nil
}
