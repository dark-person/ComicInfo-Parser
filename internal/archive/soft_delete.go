package archive

import (
	"fmt"
	"os"
	"path/filepath"
)

// Soft delete original comic directory, by move it to given trash bin path.
// i.e. Given param "C:/foo/bar" and "C:/dest", "C:/foo/bar" will be renamed to "C:/dest/bar"
func SoftDeleteComic(originDir string, trashBin string) error {
	// Ensure no empty parameters
	if originDir == "" || trashBin == "" {
		return fmt.Errorf("original or trash bin path is empty")
	}

	// Ensure original directory and trash bin is different
	if originDir == trashBin {
		return fmt.Errorf("original and trash bin path is the same")
	}

	// Check if original folder exist
	if _, err := os.Stat(originDir); os.IsNotExist(err) {
		return fmt.Errorf("original directory does not exist")
	}

	// Ensure dest folder exist
	err := os.MkdirAll(trashBin, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("unable to create trash bin folder")
	}

	// Move folder to destination
	return os.Rename(originDir, filepath.Join(trashBin, filepath.Base(originDir)))
}
