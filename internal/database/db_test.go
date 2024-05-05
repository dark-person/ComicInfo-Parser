package database

// Important NOTE:
// This file is for holding s1ome test utility.
//
// It will not include any actual test of any function.

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Prepare testing. This function should be called in 1st line of every test.
func prepareTest() {
	// Set Log Level & Output
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)

	// Prepare Testing Directory
	os.MkdirAll("testing", 0755)
}

// Set user_version of a existing database.
//
// This function SHOULD be used in testing ONLY.
func testOnly_setVersion(db *sql.DB, version int) error {
	query := fmt.Sprintf("PRAGMA user_version = %d", version)
	_, err := db.Exec(query)
	return err
}
