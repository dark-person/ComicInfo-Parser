package application

import (
	"path/filepath"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestGetDefaultOutputDirectory(t *testing.T) {
	type testCase struct {
		configExportDir string // Default export directory for config files
		inputDir        string // Parameter for function
		want            string // Expected result
	}

	absPath1, err := filepath.Abs("/absPath/")
	if err != nil {
		t.Errorf("Failed to get absolute path")
	}

	absPath2, err := filepath.Abs("relatedPath")
	if err != nil {
		t.Errorf("Failed to get absolute path")
	}

	// Create a dummy APP
	a := &App{}

	// Start test
	tests := []testCase{
		{"/absPath/", "inputDir", absPath1},
		{"relatedPath", "inputDir", absPath2},
		{"", "inputDir", "inputDir"},
	}

	for idx, tt := range tests {
		// Reset config
		a.cfg = config.Default()

		// Plug exported directory to dummy app
		if tt.configExportDir != "" {
			a.cfg.DefaultExport, _ = filepath.Abs(tt.configExportDir)
		}

		// Run function
		result := a.GetDefaultOutputDirectory(tt.inputDir)

		assert.EqualValuesf(t, tt.want, result, "Unexpected directory in case %d", idx)
	}
}
