package historyprov

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitKeywords(t *testing.T) {
	type testCase struct {
		bookname string
		keywords []string
	}

	tests := []testCase{
		{
			"Some Name 2 (ABC)｜Translated Name [Translated] [DL]",
			[]string{"Some", "Name", "2", "ABC", "｜Translated", "Name", "Translated", "DL"},
		},
		{
			// cSpell:disable
			"Some Name 7.5 [TranslatorxTranslator]",
			[]string{"Some", "Name", "7.5", "TranslatorxTranslator"},
			// cSpell:enable
		},
		{
			"SomeName2[Translator]",
			[]string{"SomeName2", "Translator"},
		},
		{
			"Some Name Part 1 (Magazine 2025-1) [DL] [Translator]",
			[]string{"Some", "Name", "Part", "1", "Magazine", "2025-1", "DL", "Translator"},
		},
	}

	for idx, tt := range tests {
		assert.EqualValuesf(t, tt.keywords, splitKeywords(tt.bookname), "Unmatched result %d", idx)
	}
}

func TestRelativeParentKeywords(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rootPath    string
		currentPath string
		want        []string
	}{
		// --------------------- Normal Case ---------------------
		{
			name:        "two intermediate dirs",
			rootPath:    "/c/user/home/input/",
			currentPath: "/c/user/home/input/category1/category2/book1",
			want:        []string{"category1", "category2"},
		},
		{
			name:        "one intermediate dir",
			rootPath:    "/c/user/home/input/",
			currentPath: "/c/user/home/input/category1/book1",
			want:        []string{"category1"},
		},
		{
			name:        "no suffix",
			rootPath:    "/c/user/home/input",
			currentPath: "/c/user/home/input/category1/category2/book1",
			want:        []string{"category1", "category2"},
		},

		// --------------------- Special Case ---------------------
		{
			name:        "direct child of root",
			rootPath:    "/c/user/home/input/",
			currentPath: "/c/user/home/input/book1",
			want:        []string{},
		},
		{
			name:        "equals root",
			rootPath:    "/c/user/home/input/",
			currentPath: "/c/user/home/input/",
			want:        []string{},
		},
		{
			name:        "trailing separator",
			rootPath:    "/c/user/home/input////",
			currentPath: "/c/user/home/input/category1///book1",
			want:        []string{"category1"},
		},

		// --------------------- Invalid Case --------------------------
		{
			name:        "reversed",
			rootPath:    "/c/user/home/input/category1/category2/book1",
			currentPath: "/c/user/home/input",
			want:        []string{},
		},
		{
			name:        "unrelated path",
			rootPath:    "/c/user/home/input/",
			currentPath: "/other/path/file",
			want:        []string{},
		},
		{
			name:        "prefix collision",
			rootPath:    "/a/b",
			currentPath: "/a/bc/d/file",
			want:        []string{},
		},
		{
			name:        "empty root",
			rootPath:    "",
			currentPath: "/c/user/home/input/book1",
			want:        []string{},
		},
		{
			name:        "empty current",
			rootPath:    "/c/user/home/input/",
			currentPath: "",
			want:        []string{},
		},
		{
			name:        "empty current and root",
			rootPath:    "",
			currentPath: "",
			want:        []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := relativeParentKeywords(tt.rootPath, tt.currentPath)
			assert.Equal(t, tt.want, got, "Failed : "+tt.name)
		})
	}
}
