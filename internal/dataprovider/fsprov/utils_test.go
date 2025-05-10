package fsprov

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSupportedImg(t *testing.T) {
	type testCase struct {
		filename string
		want     bool
	}

	tests := []testCase{
		{"foo/image1.jpg", true},
		{"foo/image2.png", true},
		{"foo/image3.jpeg", true},
		{"foo/image4.webp", true},

		// Not Supported
		{"foo/abc/", false},
		{"foo/abc.txt", false},
		{"foo/abc.docx", false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, isSupportedImg(tt.filename))
	}
}
