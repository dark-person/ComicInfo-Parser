// Package to control autofill behavior for comicinfo, based on database record and bookname.
//
// Bookname will be separate to multiple keyword by space, while database store user inputted record.
package autofill

import (
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/lazydb"
)

// Runner for auto fill process.
type AutoFillRunner struct {
	db *lazydb.LazyDB // Database that store tag and other inputted value
}

// Create a new autofill runner. The database should be connected and initalized.
func New(db *lazydb.LazyDB) *AutoFillRunner {
	if !db.Connected() {
		panic("database is not connected.")
	}

	return &AutoFillRunner{
		db: db,
	}
}

// Parse bookname into multiple keywords, then put into a tempoary table,
// this will help quicker checking on tag/value by SQL.
func (r *AutoFillRunner) parseToDB(bookname string) (err error) {
	// Splits words
	words := splitKeywords(bookname)

	// Create temporary table
	_, err = r.db.Exec("CREATE TABLE _tmp_autofill (word text)")
	if err != nil {
		return err
	}

	// Insert words to temporary table
	queries := make([]lazydb.ParamQuery, 0)

	for _, item := range words {
		queries = append(queries, lazydb.Param("INSERT INTO _tmp_autofill (word) VALUES (?)", item))
	}

	_, err = r.db.ExecMultiple(queries)
	if err != nil {
		return err
	}

	return nil
}

// Check if any inputted value match bookname keyword.
func (r *AutoFillRunner) matchInputs() (map[definitions.CategoryType][]string, error) {
	var err error

	// Select tags that is matched
	rows, err := r.db.Query("SELECT category, input from _tmp_autofill JOIN list_inputted ON _tmp_autofill.word = list_inputted.input")
	if err != nil {
		return nil, err
	}

	result := make(map[definitions.CategoryType][]string, 0)

	for rows.Next() {
		var category definitions.CategoryType
		var word string

		err = rows.Scan(&category, &word)
		if err != nil {
			return nil, err
		}

		_, exist := result[category]
		if !exist {
			result[category] = make([]string, 0)
		}

		result[category] = append(result[category], word)
	}

	return result, nil
}

// Check if any tag match bookname keyword.
func (r *AutoFillRunner) matchTags() ([]string, error) {
	var err error

	// Select tags that is matched
	rows, err := r.db.Query("SELECT word from _tmp_autofill JOIN tags ON _tmp_autofill.word = tags.input")
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	for rows.Next() {
		var word string

		err = rows.Scan(&word)
		if err != nil {
			return nil, err
		}

		result = append(result, word)
	}

	return result, nil
}

// Ensure input slice must be a empty slice instead of nil.
func sanitized(input []string) []string {
	if input == nil {
		return make([]string, 0)
	}

	return input
}

// Result for auto fill runner.
type AutoFillResult struct {
	Tags     []string                              // Tag that is matched with bookname keywords.
	Inputted map[definitions.CategoryType][]string // Map of Category and Inputted value that is matched with bookname keyword.
}

// Start the auto fill runner.
// Bookname will separate to multiple keyword and insert to tempoary table,
// then find matched value by SQL.
func (r *AutoFillRunner) Run(bookname string) (*AutoFillResult, error) {
	var err error
	result := &AutoFillResult{}

	// Prepare bookname into database
	err = r.parseToDB(bookname)
	if err != nil {
		return nil, err
	}

	// Found Matched Tags
	result.Tags, err = r.matchTags()
	if err != nil {
		return nil, err
	}

	// Found matched inputs
	result.Inputted, err = r.matchInputs()
	if err != nil {
		return nil, err
	}

	// Drop tempoary table
	_, err = r.db.Exec("DROP TABLE _tmp_autofill")
	if err != nil {
		return nil, err
	}

	// Sanitzed input for all category
	for _, c := range definitions.Categories() {
		result.Inputted[c] = sanitized(result.Inputted[c])
	}

	return result, nil
}
