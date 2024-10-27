package history

import (
	"github.com/dark-person/lazydb"
)

type categoryType int

const (
	CategoryGenre      categoryType = iota + 1 // Database value for Genre.
	CategoryPublisher                          // Database value for Publisher.
	CategoryTranslator                         // Database value for Translator.
)

// Insert genre value from database.
func InsertGenre(db *lazydb.LazyDB, value ...string) error {
	return insertValue(db, CategoryGenre, value...)
}

// Get all genre value that from database.
func GetGenreList(db *lazydb.LazyDB) ([]string, error) {
	return getHistory(db, CategoryGenre)
}

// Insert publisher value from database.
func InsertPublisher(db *lazydb.LazyDB, value ...string) error {
	return insertValue(db, CategoryPublisher, value...)
}

// Get all publisher value that from database.
func GetPublisherList(db *lazydb.LazyDB) ([]string, error) {
	return getHistory(db, CategoryPublisher)
}
