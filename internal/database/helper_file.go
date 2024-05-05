package database

import (
	"gui-comicinfo/internal/files"
	"os"

	"github.com/sirupsen/logrus"
)

// Create a new database file, WITHOUT apply any schema changes.
//
// If file is already existing, then this function has no effect.
func createFile(path string) error {
	// Prevent invalid database path
	if path == "" {
		logrus.Warnf("Attempt to create new empty path Database")
		return ErrNilDatabase
	}

	// Prevent already existing database file
	if files.IsFileExist(path) {
		logrus.Tracef("Database already exists, skip create database.")
		return nil
	}

	// Start Create Database
	f, err := os.Create(path)
	if err != nil {
		logrus.Warnf("Failed to create database: %v", err)
		return err
	}
	defer f.Close()

	logrus.Tracef("Database %s is created.", path)
	return nil
}
