package store

import "fmt"

// Error when trying to use nil database in this module.
var ErrDatabaseNil = fmt.Errorf("database cannot be nil")
