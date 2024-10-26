// Package to control static assets, e.g. databases, schema migrations sql files .etc.
//
// This package also control some core constants, includes:
//   - default file location of database file.
//   - database schema version
//   - database type
package assets

import (
	"embed"

	"github.com/dark-person/lazydb"
)

//go:embed all:schema
var schema embed.FS

// Supported database schema version.
//
// Developer should change this value when any schema update performed.
const supportedSchemaVersion = 1

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
		lazydb.Version(supportedSchemaVersion),
	)
}
