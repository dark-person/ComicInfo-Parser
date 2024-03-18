// Internal Package for file related utilities.
package files

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Check the file is exist or not.
//
// Please note that this function will not check filepath valid or not.
// Developer should use IsFileValid() instead.
func IsFileExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		logrus.Infof("%v file does not exist\n", path)
		return false
	}

	if err == nil {
		logrus.Infof("%v file exist\n", path)
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
