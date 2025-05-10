// Package for comicinfo data provider
// that use database record and bookname to fill details of comicinfo.
//
// Bookname will be separate to multiple keyword by space,
// while database store user inputted record.
package historyprov

import (
	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider"
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/lazydb"
)

// Provider that based on user input history.
type HistoryProvider struct {
	db       *lazydb.LazyDB // Database that store tag and other inputted value
	bookname string         // bookname, for split keywords
}

var _ dataprovider.DataProvider = (*HistoryProvider)(nil)

// Create a new HistoryProvider. The database should be connected and initalized.
func New(db *lazydb.LazyDB, bookname string) *HistoryProvider {
	if !db.Connected() {
		panic("database is not connected.")
	}

	return &HistoryProvider{db: db, bookname: bookname}
}

// Parse bookname into multiple keywords, then put into a tempoary table,
// this will help quicker checking on tag/value by SQL.
func (p *HistoryProvider) parseToDB(bookname string) (err error) {
	// Splits words
	words := splitKeywords(bookname)

	// Create temporary table
	_, err = p.db.Exec("CREATE TABLE _tmp_autofill (word text)")
	if err != nil {
		return err
	}

	// Insert words to temporary table
	queries := make([]lazydb.ParamQuery, 0)

	for _, item := range words {
		queries = append(queries, lazydb.Param("INSERT INTO _tmp_autofill (word) VALUES (?)", item))
	}

	_, err = p.db.ExecMultiple(queries)
	if err != nil {
		return err
	}

	return nil
}

// Check if any inputted value match bookname keyword.
func (p *HistoryProvider) matchInputs() (map[definitions.CategoryType][]string, error) {
	var err error

	// Select tags that is matched
	rows, err := p.db.Query("SELECT category, input from _tmp_autofill JOIN list_inputted ON _tmp_autofill.word = list_inputted.input")
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
func (p *HistoryProvider) matchTags() ([]string, error) {
	var err error

	// Select tags that is matched
	rows, err := p.db.Query("SELECT word from _tmp_autofill JOIN tags ON _tmp_autofill.word = tags.input")
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

func (p *HistoryProvider) Fill(c *comicinfo.ComicInfo) (out *comicinfo.ComicInfo, err error) {
	// Prepare bookname into database
	err = p.parseToDB(p.bookname)
	if err != nil {
		return c, err
	}

	// Found Matched Tags
	tags, err := p.matchTags()
	if err != nil {
		return c, err
	}

	// Found matched inputs
	inputted, err := p.matchInputs()
	if err != nil {
		return c, err
	}

	// Drop tempoary table
	_, err = p.db.Exec("DROP TABLE _tmp_autofill")
	if err != nil {
		return c, err
	}

	// Sanitzed input for all category
	for _, c := range definitions.Categories() {
		inputted[c] = sanitized(inputted[c])
	}

	// Fill comicinfo
	c.AddTags(tags...)
	c.AddGenre(inputted[definitions.CategoryGenre]...)
	c.AddPublisher(inputted[definitions.CategoryPublisher]...)
	c.AddTranslator(inputted[definitions.CategoryTranslator]...)

	return c, nil
}
