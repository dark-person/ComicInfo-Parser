package store

import "fmt"

// Error blocks
var (
	// Error when trying to use nil database in this module.
	ErrDatabaseNil = fmt.Errorf("Database cannot be nil")
)
