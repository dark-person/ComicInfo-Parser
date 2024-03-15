// Internal Package for file related utilities.
package files

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Check the file is exist or not.
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
