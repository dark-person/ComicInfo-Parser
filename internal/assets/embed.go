// Package to control static assets, e.g. databases, schema migrations sql files .etc.
//
// This package also control some core constants, includes:
//   - default file location of database file.
//   - database schema version
//   - database type
//   - default backup location of database file
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
const supportedSchemaVersion = 4

// Get embedded schema from filesystem.
// Used for database schema migrations.
func GetSchema() embed.FS {
	return schema
}

// Return default lazydb.
func DefaultDB(path string) *lazydb.LazyDB {
	return lazydb.New(
		lazydb.DbPath(path),
		lazydb.Migrate(schema, "schema"),
		lazydb.Version(supportedSchemaVersion),
	)
}

// Return default lazydb that allow auto-backup when schema version changed.
func DefaultDBWithBackup(path string, backupDir string) *lazydb.LazyDB {
	return lazydb.New(
		lazydb.DbPath(path),
		lazydb.Migrate(schema, "schema"),
		lazydb.Version(supportedSchemaVersion),
		lazydb.BackupDir(backupDir),
	)
}
