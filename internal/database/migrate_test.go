package database

import (
	"embed"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

//go:embed test_schema/*
var fsTesting embed.FS

// Function to prepare usable AppDB object.
// Purpose of function is to reuse code.
//
// Developer MUST close AppDB after testing is completed,
// to prevent filesystem blocking database removal.
func prepareAppDB_TestOnly(t *testing.T, dbName string) (a *AppDB, err error) {
	// Remove any existing database
	os.Remove(dbName)

	// Create AppDB object
	a, err = NewPathDB(dbName)
	if err != nil {
		t.Fatal("Failed to create db: ", err)
	}

	// Connect to database
	err = a.Connect()
	if err != nil {
		t.Fatal("Failed to connect, return error: ", err)
	}

	// Override migration directory name
	a.MigrateDir = "test_schema"

	return a, err
}

// Function to prepare usable AppDB object.
// Purpose of function is to reuse code.
//
// Developer MUST close AppDB after testing is completed,
// to prevent filesystem blocking database removal.
func prepareAppDB_LatestTestOnly(t *testing.T, dbName string) (*AppDB, error) {
	// Use existing func
	a, err := prepareAppDB_TestOnly(t, dbName)
	if err != nil {
		t.Fatal("Failed to prepare db: ", err)
	}

	// Step to latest
	err = a.StepToLatestWithFs(fsTesting)
	if err != nil {
		t.Fatal("Failed to prepare db as latest: ", err)
	}

	// Return as success
	return a, err
}

// Test migration to latest. Act as smoke test only.
func TestStepToLatest(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	// Get temp directory
	tempDir := t.TempDir()

	// Create AppDB object
	a, err := prepareAppDB_TestOnly(t, filepath.Join(tempDir, "test.db"))
	if err != nil {
		t.Fatal("Failed to prepare db: ", err)
	}

	// Migration
	err = a.StepToLatestWithFs(fsTesting)
	if err != nil {
		t.Error("MigrateUp return error: ", err)
	}

	a.Close()
}

// Test migration downgrade 2 version. Act as smoke test only.
// TODO: Increase coverage of step down function.
func TestStepDown(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	// Get temp directory
	tempDir := t.TempDir()

	// Test case struct, use fsTesting for all tests
	type testCase struct {
		dbName      string // Name of database file
		step        int    // step down how many versions
		wantErr     bool   // Should error appear
		wantVersion int    // user_version that should be appear in migrated database
	}

	tests := []testCase{
		// Normal Input
		{"test0.db", 2, false, 1},
		{"test1.db", 3, false, 0},
		{"test2.db", 0, false, 3},

		// Invalid input
		{"test3.db", -1, true, -1},
		{"test4.db", 4, true, -1},
	}

	// Run tests
	for idx, tt := range tests {
		// Create AppDB object
		a, err := prepareAppDB_LatestTestOnly(t, filepath.Join(tempDir, tt.dbName))

		if err != nil {
			t.Fatal("Failed to prepare db: ", err)
		}
		defer a.Close() // Close to prevent leak

		// Main test scope
		err = a.stepDown(fsTesting, tt.step)

		assert.EqualValuesf(t, tt.wantErr, err != nil, "Unexpected err in case %d: %v", idx, err)

		// Get user version
		if !tt.wantErr {
			ver, _ := getUserVersion(a.db)
			assert.EqualValuesf(t, tt.wantVersion, ver, "user_version not matched in case %d: %d", idx, ver)
		}
	}
}
