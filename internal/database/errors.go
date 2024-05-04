package database

import "fmt"

// Error when try to pass nil value to *sql.DB
var ErrNilDatabase = fmt.Errorf("nil value of *sql.DB")

// Error represents database path is invalid
var ErrInvalidPath = fmt.Errorf("invalid database path")
