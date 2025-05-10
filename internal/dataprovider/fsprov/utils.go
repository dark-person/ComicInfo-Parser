package fsprov

import (
	"path/filepath"
	"strings"
)

// Check if image file extension is supported.
func isSupportedImg(filename string) bool {
	ext := filepath.Ext(filename)
	ext = strings.ToLower(ext)
	return ext == ".jpg" || ext == ".png" || ext == ".jpeg" || ext == ".webp"
}
