package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpecialTags(t *testing.T) {
	// Prepare test case struct
	type testCase struct {
		filename string
		want     []string
	}

	// Prepare tests
	tests := []testCase{
		// No special tags
		{"[author] title", []string{}},

		// One Special tag
		{"[author] title 無修正", []string{"無修正"}},
		{"[author] title [無修正]", []string{"無修正"}},

		// Two Special tags
		{"[author] title 無修正 DL版", []string{"無修正", "DL版"}},
		{"[author] title [無修正][DL版]", []string{"無修正", "DL版"}},

		// No special tags, but has similar string to special tags
		{"[author] title 無1修正", []string{}},

		// Contain special tags with string suffix
		{"[author] title 無修正test", []string{"無修正"}},
	}

	// Run Tests
	for idx, tt := range tests {
		got := GetSpecialTags(tt.filename)
		assert.EqualValuesf(t, tt.want, got, "Case %d: tags not matched", idx)
	}
}
