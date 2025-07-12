package historyprov

import (
	"os"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
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
	db.Exec("INSERT OR IGNORE INTO word_store (category_id, word) VALUES (?, ?)",
		definitions.CategoryGenre,
		"Test-Genre")

	// Publisher
	db.Exec("INSERT OR IGNORE INTO word_store (category_id, word) VALUES (?, ?)",
		definitions.CategoryPublisher,
		"Test-Publisher")

	// Tranlator
	db.Exec("INSERT OR IGNORE INTO word_store (category_id, word) VALUES (?, ?)",
		definitions.CategoryTranslator,
		"Test-Translator")

	// Tags
	db.Exec(`INSERT INTO word_store (category_id, word) VALUES (4, 'abc'), (4, 'def'), (4, 'ghi')`)

	// Alias tags
	db.Exec(`INSERT INTO triggers (keyword, word_id) VALUES ('kcs', 4)`)

	return db
}

func TestAutoFillRun(t *testing.T) {
	// Prepare database
	db := prepareDB()

	type testResult struct {
		genre      string
		publisher  string
		translator string
		tags       string
	}

	// Test case
	type testCase struct {
		bookname string
		want     testResult
	}

	// Start test
	tests := []testCase{
		{
			"Some Bookname (abc) [Test-Translator] [Test-Translator2] [def]",
			testResult{"", "", "Test-Translator", "abc,def"}},
		{
			"(ghi) Another Bookname 2 (abc) [Test-Genre] [Test-Publisher] [Test-TranslatorXTest-Translator2] [invalid] [20240123]",
			testResult{"Test-Genre", "Test-Publisher", "", "abc,ghi"},
		},
		{
			"Another Bookname (kcs) [def]",
			testResult{"", "", "", "abc,def"},
		},
	}

	var err error
	for idx, tt := range tests {
		// Init runner
		prov := New(db, tt.bookname)

		// Prepare new comicinfo
		temp := comicinfo.New()
		info := &temp

		// Run provider
		info, err = prov.Fill(info)
		if err != nil {
			assert.NoErrorf(t, err, "No error should be generate in case %d", idx)
			continue
		}

		// Compare values
		assert.EqualValuesf(t, tt.want.genre, info.Genre, "unmatched genre value in case %d", idx)
		assert.EqualValuesf(t, tt.want.publisher, info.Publisher, "unmatched publisher value in case %d", idx)
		assert.EqualValuesf(t, tt.want.translator, info.Translator, "unmatched translator value in case %d", idx)
		assert.EqualValuesf(t, tt.want.tags, info.Tags, "unmatched tag value in case %d", idx)
	}

	// Remove database after complete
	db.Close()
	os.Remove(testingDatabasePath)
}
