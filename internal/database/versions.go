package database

import (
	"database/sql"
	"fmt"
)

// Get user_version value from database.
//
// This utility function is ONLY for internal use.
func getUserVersion(db *sql.DB) (int, error) {
	if db == nil {
		return -1, ErrNilDatabase
	}

	// Get User Version
	row := db.QueryRow("PRAGMA user_version")
	if row == nil {
		return -1, fmt.Errorf("nil row query")
	}

	// Scan value into var
	var userVersion int
	err := row.Scan(&userVersion)
	if err != nil {
		return -1, err
	}

	// Return
	return userVersion, nil
}
