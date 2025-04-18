package parser

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
			"Some Name 7.5 [TranslatorxTranslator]",
			[]string{"Some", "Name", "7.5", "TranslatorxTranslator"},
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
		assert.EqualValuesf(t, tt.keywords, SplitKeywords(tt.bookname), "Unmatched result %d", idx)
	}
}
