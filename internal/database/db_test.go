package database

// Important NOTE:
// This file is for holding s1ome test utility.
//
// It will not include any actual test of any function.

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Prepare testing. This function should be called in 1st line of every test.
func prepareTest() {
	// Set Log Level & Output
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)
}
