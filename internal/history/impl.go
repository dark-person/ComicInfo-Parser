package history

import "gui-comicinfo/internal/database"

// Database value for Genre.
const Genre_Text = "Genre"

// Database value for Publisher.
const Publisher_Text = "Publisher"

// Insert genre value from database.
func InsertGenre(db *database.AppDB, value ...string) error {
	return insertValue(db, Genre_Text, value...)
}

// Get all genre value that from database.
func GetGenreList(db *database.AppDB) ([]string, error) {
	return getHistory(db, Genre_Text)
}

// Insert publisher value from database.
func InsertPublisher(db *database.AppDB, value ...string) error {
	return insertValue(db, Publisher_Text, value...)
}

// Get all publisher value that from database.
func GetPublisherList(db *database.AppDB) ([]string, error) {
	return getHistory(db, Publisher_Text)
}
