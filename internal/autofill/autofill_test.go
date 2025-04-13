package autofill

import (
	"os"
	"slices"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/lazydb"
	"github.com/stretchr/testify/assert"
)

const testingDatabasePath = "./storage.db"

// Create a testing database
func prepareDB() *lazydb.LazyDB {
	db := assets.DefaultDb(testingDatabasePath)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Migrate()
	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}

	// Prepare some dummy data

	// Genre
	db.Exec("INSERT OR IGNORE INTO list_inputted (category, input) VALUES (?, ?)",
		definitions.CategoryGenre,
		"Test-Genre")

	// Publisher
	db.Exec("INSERT OR IGNORE INTO list_inputted (category, input) VALUES (?, ?)",
		definitions.CategoryPublisher,
		"Test-Publisher")

	// Tranlator
	db.Exec("INSERT OR IGNORE INTO list_inputted (category, input) VALUES (?, ?)",
		definitions.CategoryTranslator,
		"Test-Translator")

	// Tags
	db.Exec(`INSERT INTO tags (input) VALUES ('abc'), ('def'), ('ghi')`)

	return db
}

func TestAutoFillRun(t *testing.T) {
	// Prepare database
	db := prepareDB()

	// Init runner
	r := New(db)

	type testResult struct {
		genre      []string
		publisher  []string
		translator []string
		tags       []string
	}

	// Test case
	type testCase struct {
		bookname string
		want     testResult
	}

	// Start test
	tests := []testCase{
		{
			"Some Bookname (abc) [Test-Translator] [Test-Translator2] [ghi]",
			testResult{
				[]string{},
				[]string{},
				[]string{"Test-Translator"},
				[]string{"abc", "ghi"},
			}},
		{
			"(ghi) Another Bookname 2 (abc) [Test-Genre] [Test-Publisher] [Test-TranslatorXTest-Translator2] [invalid] [20240123]",
			testResult{
				[]string{"Test-Genre"},
				[]string{"Test-Publisher"},
				[]string{},
				[]string{"abc", "ghi"},
			}},
	}

	for idx, tt := range tests {
		info, err := r.Run(tt.bookname)
		if err != nil {
			assert.Nilf(t, err, "No error should be generate in case %d", idx)
			continue
		}

		// Sort values for easier testing
		slices.Sort(info.Inputted[definitions.CategoryGenre])
		slices.Sort(info.Inputted[definitions.CategoryPublisher])
		slices.Sort(info.Inputted[definitions.CategoryTranslator])
		slices.Sort(info.Tags)

		// Compare
		assert.EqualValuesf(t, tt.want.genre, info.Inputted[definitions.CategoryGenre], "unmatched genre value in case %d", idx)
		assert.EqualValuesf(t, tt.want.publisher, info.Inputted[definitions.CategoryPublisher], "unmatched publisher value in case %d", idx)
		assert.EqualValuesf(t, tt.want.translator, info.Inputted[definitions.CategoryTranslator], "unmatched translator value in case %d", idx)
		assert.EqualValuesf(t, tt.want.tags, info.Tags, "unmatched tag value in case %d", idx)
	}

	// Remove database after complete
	db.Close()
	os.Remove(testingDatabasePath)
}
