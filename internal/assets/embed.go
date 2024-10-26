// Package to control static assets, e.g. databases, schema migrations sql files .etc.
package assets

import (
	"embed"

	"github.com/dark-person/lazydb"
)

//go:embed all:schema
var schema embed.FS

// Get embedded schema from filesystem.
// Used for database schema migrations.
func GetSchema() embed.FS {
	return schema
}

// Return default lazydb.
func DefaultDb(path string) *lazydb.LazyDB {
	return lazydb.New(
		lazydb.DbPath(path),
		lazydb.Migrate(schema, "schema"),
		lazydb.Version(1),
	)
}
