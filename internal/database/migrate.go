package database

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	migrateSqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/sirupsen/logrus"
)

// Directory that storing sql for migration
//
//go:embed schema/*
var fsMigrate embed.FS

// Create a new `migrate.Migrate` instance, which can be used to migrate up/down.
//
// Purpose of this function is to reuse code to create a new migrate object.
func (a *AppDB) migrateInstance(embedFs embed.FS) (*migrate.Migrate, error) {
	// Prevent db is nil
	if a.db == nil {
		logrus.Fatal("Database connection is not opened. Migration failed.")
		return nil, ErrNilDatabase
	}

	// Prevent empty migration directory
	if a.MigrateDir == "" {
		logrus.Fatal("Migration directory is not set.")
		return nil, ErrNilDatabase
	}

	// Get sqlite3 instance
	instance, err := migrateSqlite3.WithInstance(a.db, &migrateSqlite3.Config{})
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	// Create iofs with embedded filesystem, with directory specified
	f, err := iofs.New(embedFs, a.MigrateDir)
	if err != nil {
		logrus.Fatal("Failed to get migration embed fs: ", err)
		return nil, err
	}

	return migrate.NewWithInstance("iofs", f, "sqlite3", instance)
}

// Start migration to latest.
func (a *AppDB) StepToLatest() error {
	return a.StepToLatestWithFs(fsMigrate)
}

// Start migration to latest, with embed fs specified.
func (a *AppDB) StepToLatestWithFs(embedFs embed.FS) error {
	// Prepare migration instance
	m, err := a.migrateInstance(embedFs)
	if err != nil {
		logrus.Errorf("Failed to create instance: %v", err)
		return err
	}

	// modify for Down
	err = m.Up()

	// Early return if no error
	if err == nil {
		return nil
	}

	// No changes applied, which is acceptable
	if err.Error() == "no change" {
		logrus.Warn("No migration change is performed.")
		return nil
	}

	logrus.Errorf("Failed to update db: %v", err)
	return err
}

// Start migration down for n version.
func (a *AppDB) stepDown(embedFs embed.FS, n int) error {
	// Prevent n is negative number
	if n < 0 {
		logrus.Errorf("Invalid version number: %d", n)
		return fmt.Errorf("invalid parameter n")
	}

	// Prepare migration instance
	m, err := a.migrateInstance(embedFs)
	if err != nil {
		logrus.Errorf("Failed to prepare instance: %v", err)
		return err
	}

	// Version check
	ver, _, err := m.Version()
	if err != nil {
		return err
	}
	logrus.Debugf("Current version: %d", ver)

	// Prevent negative version number as result
	if int(ver)-n < 0 {
		logrus.Warnf("Version will be negative if follow migrations. Abort.")
		return ErrNegativeVersion
	}

	// Perform step down
	err = m.Steps(-1 * n)

	// Early return as passed
	if err == nil {
		return nil
	}

	// No changes applied, which is acceptable
	if err.Error() == "no change" {
		logrus.Warn("No migration change is performed.")
		return nil
	}

	// File not exist
	if err.Error() == "file does not exist" {
		logrus.Warnf("Possible cause: cannot downgrade for version <= 0. Failed to run Step(%d).", n)
		return err
	}

	// Unknown error
	logrus.Errorf("Failed to run Step(%d) : %v", n, err)
	return err
}
