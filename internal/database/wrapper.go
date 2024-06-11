package database

/* This file contains wrapper function of *sql.DB.*/

import (
	"database/sql"
)

// Create a prepare statement of database.
//
// This function is a wrapper for *sql.DB.Prepare(query).
// If database is not init in *AppDB, then ErrNilDatabase is returned.
func (a *AppDB) Prepare(query string) (*sql.Stmt, error) {
	// Prevent database is nil value
	if a.db == nil {
		return nil, ErrNilDatabase
	}

	return a.db.Prepare(query)
}
