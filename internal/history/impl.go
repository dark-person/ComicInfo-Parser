package history

import "gui-comicinfo/internal/database"

// Database value for Genre.
const Genre_Text = "Genre"

// Insert genre value from database.
func InsertGenre(db *database.AppDB, value ...string) error {
	return insertValue(db, Genre_Text, value...)
}

// Get all genre value that from database.
func GetGenreList(db *database.AppDB) ([]string, error) {
	return getHistory(db, Genre_Text)
}
