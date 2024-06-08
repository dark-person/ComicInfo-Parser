package database

import "fmt"

// Error represents program is outdated
var ErrProgramOutdated = fmt.Errorf("outdated program")

// Error represents database schema is outdated
var ErrSchemaOutdated = fmt.Errorf("outdated schema")

// Error represents database user_version is invalid,
// usually appear in empty database
var ErrInvalidVersion = fmt.Errorf("invalid user_version")

// Error when try to pass nil value to *sql.DB
var ErrNilDatabase = fmt.Errorf("nil value of *sql.DB")

// Error represents database path is invalid
var ErrInvalidPath = fmt.Errorf("invalid database path")

// Error represents user_version is negative number
var ErrNegativeVersion = fmt.Errorf("negative user_version")
