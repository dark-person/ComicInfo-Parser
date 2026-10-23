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
	db := assets.DefaultDB(testingDatabasePath)

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

	// Translator
	db.Exec("INSERT OR IGNORE INTO word_store (category_id, word) VALUES (?, ?)",
		definitions.CategoryTranslator,
		"Test-Translator")

	// Tags
	db.Exec(`INSERT INTO word_store (category_id, word) VALUES (4, 'abc'), (4, 'def'), (4, 'ghi')`)

	// Writers
	db.Exec(`INSERT INTO word_store (category_id, word) VALUES (?, ?)`, definitions.CategoryWriter, "Test-Writer")

	// Triggers
	db.Exec(`INSERT INTO triggers (keyword, word_id) VALUES ('kcs', 4), ('aed', 1)`)

	return db
}

func TestAutoFillRun(t *testing.T) {
	// Prepare database
	db := prepareDB()

	type testResult struct {
		genre      string
		writer     string
		publisher  string
		translator string
		tags       string
	}

	// Test case
	type testCase struct {
		rootFolder string
		bookFolder string
		want       testResult
	}

	// Start test
	tests := []testCase{
		// Plain folder case
		{
			rootFolder: "/c/user/home/input/",
			bookFolder: "/c/user/home/input/[Test-Writer] Some Bookname (abc) [Test-Translator] [Test-Translator2] [def]",
			want:       testResult{"", "Test-Writer", "", "Test-Translator", "abc,def"}},
		{
			rootFolder: "/c/user/home/input/",
			bookFolder: "/c/user/home/input/(ghi) Another Bookname 2 (abc) [Test-Genre] [Test-Publisher] [Test-TranslatorXTest-Translator2] [invalid] [20240123]",
			want:       testResult{"Test-Genre", "", "Test-Publisher", "", "abc,ghi"},
		},
		{
			rootFolder: "/c/user/home/input/",
			bookFolder: "/c/user/home/input/Another Bookname (kcs) (aed) [def]",
			want:       testResult{"Test-Genre", "", "", "", "abc,def"},
		},

		// Intermediate folder case
		{
			rootFolder: "/c/user/home/input/",
			bookFolder: "/c/user/home/input/def/[Test-Writer] Some Bookname (abc) [Test-Translator] [Test-Translator2]",
			want:       testResult{"", "Test-Writer", "", "Test-Translator", "abc,def"}},
		{
			rootFolder: "/c/user/home/input/",
			bookFolder: "/c/user/home/input/kcs/aed/Another Bookname [def]",
			want:       testResult{"Test-Genre", "", "", "", "abc,def"},
		},
		{
			rootFolder: "/c/user/home/input/",
			bookFolder: "/c/user/home/input/kcs/aed/Another Bookname (kcs) [def]",
			want:       testResult{"Test-Genre", "", "", "", "abc,def"},
		},
	}

	var err error
	for idx, tt := range tests {
		// Init runner
		prov := New(db, tt.rootFolder, tt.bookFolder)

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
		assert.EqualValuesf(t, tt.want.writer, info.Writer, "unmatched writer value in case %d", idx)
		assert.EqualValuesf(t, tt.want.publisher, info.Publisher, "unmatched publisher value in case %d", idx)
		assert.EqualValuesf(t, tt.want.translator, info.Translator, "unmatched translator value in case %d", idx)
		assert.EqualValuesf(t, tt.want.tags, info.Tags, "unmatched tag value in case %d", idx)
	}

	// Remove database after complete
	db.Close()
	os.Remove(testingDatabasePath)
}
