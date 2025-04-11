package comicinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		args []string
		want []string
	}{
		{[]string{"item1", "item2"}, []string{"item1", "item2"}},
		{[]string{"item1", "item1"}, []string{"item1"}},
		{[]string{"item1", "item1 "}, []string{"item1"}},
		{[]string{"item1"}, []string{"item1"}},
		{[]string{}, []string{}},
	}

	for idx, tt := range tests {
		assert.EqualValuesf(t, tt.want, removeDuplicates(tt.args), "Unexpected value in case %d", idx)
	}
}

func TestSplitValue(t *testing.T) {
	tests := []struct {
		str  string
		want []string
	}{
		{"item1,item2,item3", []string{"item1", "item2", "item3"}},
		{"item1, item2, item3", []string{"item1", "item2", "item3"}},
		{"item1,,item3", []string{"item1", "item3"}},
		{"item", []string{"item"}},
		{"", []string{}},
	}

	for idx, tt := range tests {
		assert.EqualValuesf(t, tt.want, splitValue(tt.str), "Unexpected result in case %d", idx)
	}
}
