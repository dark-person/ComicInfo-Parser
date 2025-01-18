package archive

import (
	"os"
	"path/filepath"

	"github.com/dark-person/comicinfo-parser/internal/files"
)

// Option for rename .cbz files. Only one option can passed at a time.
type RenameOption renameOption

// Unexported type, option for rename .cbz files.
type renameOption struct {
	isWrap           bool   // Wrap .cbz file with a folder if true
	isDefaultWrap    bool   // Use cbz filename as wrap folder name if true
	customWrapFolder string // custom wrap folder name
}

// Use default wrap option, which use .cbz filename as wrap folder.
//
// The reason for this wrap is to designed for komga exports,
// when only one book is available,
// this filepath format would be better for komga:
//
//	{bookName}/{bookName}.cbz
func UseDefaultWrap() RenameOption {
	return RenameOption{
		isWrap:           true,
		isDefaultWrap:    true,
		customWrapFolder: "",
	}
}

// Custom wrap folder name, which allow different with .cbz filename.
func UseCustomWrap(folder string) RenameOption {
	return RenameOption{
		isWrap:           true,
		isDefaultWrap:    false,
		customWrapFolder: folder,
	}
}

// Not use any wrap method, only single .cbz file will be created.
func NoWrap() RenameOption {
	return RenameOption{
		isWrap:           false,
		isDefaultWrap:    false,
		customWrapFolder: "",
	}
}

// Rename zip file to cbz file.
//
// Developer can wrap behavior by:
//
//	RenameZip(absPath, NoWrap()) // No Wrap method
//	RenameZip(absPath, UseDefaultWrap()) // Default wrap with .cbz filename
//	RenameZip(absPath, UseCustomWrap("someFolder")) // Wrap with custom folder name
func RenameZip(absPath string, opt RenameOption) error {
	originalDir := filepath.Dir(absPath)
	originalFile := filepath.Base(absPath)
	name := files.TrimExt(originalFile)

	// If not wrap, then just rename the file extension to .cbz
	if !opt.isWrap {
		return os.Rename(absPath, filepath.Join(originalDir, name+".cbz"))
	}

	// Create Wrap Folder
	var wrappedDir string
	if opt.isDefaultWrap {
		wrappedDir = filepath.Join(originalDir, name)
	} else {
		wrappedDir = filepath.Join(originalDir, opt.customWrapFolder)
	}

	err := os.Mkdir(wrappedDir, 0755)
	if err != nil && !os.IsExist(err) { // Folder already exist is not an error
		return err
	}

	// Rename
	return os.Rename(absPath, filepath.Join(wrappedDir, name+".cbz"))
}
