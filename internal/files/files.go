// Internal Package for file related utilities.
package files

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// Check the file is exist or not.
//
// Please note that this function will not check filepath valid or not.
// Developer should use IsFileValid() instead.
func IsFileExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		logrus.Debugf("%v file does not exist\n", path)
		return false
	}

	if err == nil {
		logrus.Debugf("%v file exist\n", path)
		return true
	}

	logrus.Warnf("Unknown error: %s - %v", path, err)
	return false
}

// Check filepath is valid or not.
//
// This method will only ignore is actually exist or not,
// or any permission error,
// focus on filepath is valid character or not.
func IsPathValid(path string) bool {
	_, err := os.Stat(path)

	// Ignore nil error or Not Exist error
	if err == nil || os.IsNotExist(err) || os.IsPermission(err) {
		return true
	}

	// Log Error & return
	logrus.Debugf("invalid path error: %v", err)
	return false
}

// Get the filename, without the file extension.
// This method is a short-hand for getting filename.
//
// The Filename should be Base element of path,
// but not a path which contains parent directory.
func TrimExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// Move File from target to dest.
//
// This function will create all necessary folder & file that necessary to copy.
func MoveFile(target, dest string) error {
	// Already same location
	if target == dest {
		return nil
	}

	// Use rename if volume is same
	if filepath.VolumeName(target) == filepath.VolumeName(dest) {
		logrus.Info("Same volume in copy & dest, use rename")
		return os.Rename(target, dest)
	}

	// -------------- Standard io.Copy ---------------------
	srcFile, err := os.Open(target)
	if err != nil {
		return err
	}

	// Create Folder if necessary
	err = os.MkdirAll(filepath.Dir(dest), 0755)
	if err != nil {
		return err
	}

	// Get file mode
	info, err := srcFile.Stat()
	if err != nil {
		return err
	}

	// creates if file doesn't exist
	destFile, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	defer destFile.Close()

	// check first var for number of bytes copied
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	// ---------------- Standard Copy END -------------------------

	// Delete target file
	srcFile.Close()
	return os.Remove(target)
}
