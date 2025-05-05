package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFolder(t *testing.T) {
	// Prepare Struct
	type testCase struct {
		foldername   string // Folder name, as Input
		wantBookName string // Output Book name
		wantAuthor   string // Output Author
		wantMarket   string // Output Market
	}

	// Prepare List
	//cSpell:disable
	tests := []testCase{
		// Test Case that in real world example
		{`(C01) [Author (Alias)] bookname (parody)[translator]`, "bookname (parody)[translator]", "Author (Alias)", "C01"},
		{`[Author (Alias)] Series Name 1-3`, `Series Name 1-3`, "Author (Alias)", ""},
		{`[Author (Alias)] bookname [Translated]`, "bookname [Translated]", "Author (Alias)", ""},
		{`[Author (Alias)] bookname (parody) [Translated] [DL]`, "bookname (parody) [Translated] [DL]", "Author (Alias)", ""},
		{`[Author (Alias)] series name 4.5 (Side Story)`, "series name 4.5 (Side Story)", "Author (Alias)", ""},

		// Test case that is created
		{"([Author Bookname", "([Author Bookname", "", ""},
		{"[Author Bookname", "[Author Bookname", "", ""},
	}
	//cSpell:enable

	// Run Test Case
	for idx, tt := range tests {
		gotMarket, gotAuthor, gotBookName := ParseFolder(tt.foldername)

		// Compare value
		assert.EqualValuesf(t, tt.wantBookName, gotBookName, "Case %d : Book name is not equal.", idx)
		assert.EqualValuesf(t, tt.wantAuthor, gotAuthor, "Case %d : Author is not equal.", idx)
		assert.EqualValuesf(t, tt.wantMarket, gotMarket, "Case %d : Market is not equal.", idx)
	}
}
