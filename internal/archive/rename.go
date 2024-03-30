package archive

import (
	"gui-comicinfo/internal/files"
	"os"
	"path/filepath"
)

// Rename zip file to cbz file.
// If user want to wrap the .cbz file with its filename,
// then put true in isWrap parameter.
//
// The reason for wrap is to designed for komga exports,
// when only one book is available,
// this filepath format would be better for komga:
//
//	{bookName}/{bookName}.cbz
func RenameZip(absPath string, isWrap bool) error {
	originalDir := filepath.Dir(absPath)
	originalFile := filepath.Base(absPath)
	name := files.TrimExt(originalFile)

	// If not wrap, then just rename the file extension to .cbz
	if !isWrap {
		return os.Rename(absPath, filepath.Join(originalDir, name+".cbz"))
	}

	// Create Wrap Folder
	wrappedDir := filepath.Join(originalDir, name)
	err := os.Mkdir(wrappedDir, 0755)
	if err != nil {
		return err
	}

	// Rename
	return os.Rename(absPath, filepath.Join(wrappedDir, name+".cbz"))
}
