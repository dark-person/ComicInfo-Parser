package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test will only test handling when *AppDB contains nil database.
func TestAppDB_Prepare(t *testing.T) {
	a := &AppDB{db: nil}

	// Test scope
	_, err := a.Prepare("abc")
	assert.NotNilf(t, err, "Expected ErrNilDatabase error return")
}
