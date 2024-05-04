package database

import "fmt"

// Error when try to pass nil value to *sql.DB
var ErrNilDatabase = fmt.Errorf("nil value of *sql.DB")
