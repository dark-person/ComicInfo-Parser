package database

import (
	"database/sql"
	"gui-comicinfo/internal/constant"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

// Latest Version of schema supported.
//
// Every database schema changes,
// should also change this value at same time.
const LatestSchema = 5

// The database manager for this application.
type AppDB struct {
	db     *sql.DB // Database connection
	dbPath string  // Database absolute path, for easy reuse
}

// Create a new AppDB,
// with database location in "{user home directory}/gui-comicinfo/storage.db".
func NewDB() (*AppDB, error) {
	// Get Home Directory
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// Get database path
	path := filepath.Join(home, constant.RootDir, constant.DatabaseFile)
	return NewPathDB(path)
}

// Create a AppDB object, with given database path specified.
//
// This function is NOT suggested to access outside this package.
// Reason is NewDB() method, which is another constructor,
// has force database file location, which ensure all user have same settings.
//
// Developer can use mocked filepath in tests with this function.
func NewPathDB(path string) (*AppDB, error) {
	// Prevent Not database file
	if filepath.Ext(path) != ".db" {
		return nil, ErrInvalidPath
	}

	// Prevent no directory is created
	os.MkdirAll(filepath.Dir(path), 0755)

	// Create Database if need
	err := createFile(path)
	if err != nil {
		return nil, err
	}

	// Return
	return &AppDB{dbPath: path}, nil
}

// Connect to database, when path already stored in AppDB.
func (a *AppDB) Connect() error {
	// Prevent Empty Path
	if a.dbPath == "" {
		logrus.Warnf("Attempt to connect to empty path Database")
		return ErrNilDatabase
	}

	logrus.Infof("Connecting to database: %s", a.dbPath)

	// Open database connection, which create file if not exist
	var err error
	a.db, err = sql.Open(constant.DatabaseType, a.dbPath)
	if err != nil {
		logrus.Warnf("Failed to open database: %v", err)
		return err
	}

	// TODO: Test DB connection by user version
	return nil
}

// Close all existing database connection.
//
// If AppDB has no database connected, then this function has no effect,
// with no error returned.
func (a *AppDB) Close() error {
	// Prevent no connection for nil pointer
	if a.db == nil {
		logrus.Tracef("Closing nil database connection, no effect.")
		return nil
	}

	// Close connection
	err := a.db.Close()
	a.db = nil
	logrus.Tracef("Database connection closed successfully.")

	// Return error
	return err
}
