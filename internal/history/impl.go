package history

import (
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/lazydb"
)

// Insert genre value from database.
func InsertGenre(db *lazydb.LazyDB, value ...string) error {
	return insertValue(db, definitions.CategoryGenre, value...)
}

// Get all genre value that from database.
func GetGenreList(db *lazydb.LazyDB) ([]string, error) {
	return getHistory(db, definitions.CategoryGenre)
}

// Insert publisher value from database.
func InsertPublisher(db *lazydb.LazyDB, value ...string) error {
	return insertValue(db, definitions.CategoryPublisher, value...)
}

// Get all publisher value that from database.
func GetPublisherList(db *lazydb.LazyDB) ([]string, error) {
	return getHistory(db, definitions.CategoryPublisher)
}
