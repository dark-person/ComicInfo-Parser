package database

import (
	"database/sql"
	"gui-comicinfo/internal/constant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getUserVersion(t *testing.T) {
	prepareTest()

	// Create database with user_ver = 999
	db_999, err := sql.Open(constant.DatabaseType, "testing/test.db")
	if err != nil {
		t.Error(err)
	}
	err = testOnly_setVersion(db_999, 999)
	if err != nil {
		t.Error(err)
	}

	// Prepare Test Case Struct
	type TestCase struct {
		db      *sql.DB
		want    int
		wantErr bool
	}

	// Add test case
	tests := []TestCase{
		// Normal case
		{db_999, 999, false},
		// Case when database is empty (want=0 as init value of `int`)
		{nil, 0, true},
	}

	// Start Tests
	for idx, tt := range tests {
		got, err := getUserVersion(tt.db)

		// Check error is wanted
		assert.Equalf(t, err != nil, tt.wantErr, "Case %d: Expected has any error: %v, but %v", idx, tt.wantErr, err)

		// Check user_version if no error is wanted
		if !tt.wantErr {
			assert.EqualValues(t, got, tt.want, "Case %d: values not equals", idx)
		}
	}
}
