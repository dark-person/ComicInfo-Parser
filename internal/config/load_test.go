package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLoadYaml(t *testing.T) {
	type testCase struct {
		path    string
		want    *ProgramConfig
		wantErr bool
	}

	exPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	t.Log(exPath)

	tests := []testCase{
		{"mock/case-normal.yaml", &ProgramConfig{
			Preference: preferenceConfig{
				ExportDir:         filepath.Join(exPath, "./my-export"),
				ComicDir:          filepath.Join(exPath, "./my-input"),
				ExportMethod:      definitions.ExportMethodOnlyCBZ,
				DeleteAfterExport: true,
			},
			Metadata: metadataConfig{
				Number: "1",
			},
			Database: databaseConfig{
				Path: filepath.Join(exPath, "./my-data.db"),
			},
			TrashBin: trashBinConfig{
				Path: filepath.Join(exPath, "./.trash"),
			},
			LogConfig: logConfig{
				LogLevelStr: "warn",
			},
		}, false},

		{"mock/case-typo1.yaml", Default(), false},
		{"mock/case-typo2.yaml", Default(), false},
		{"mock/case-empty.yaml", Default(), false},
		{"mock/not-exist.yaml", nil, true},
	}

	for idx, tt := range tests {
		// Load YAML file and check result
		c, err := LoadYaml(tt.path)

		if tt.wantErr {
			assert.Errorf(t, err, "Error should be returned in case %d, but return nil", idx)

			// Ensure default is returned
			assert.EqualValuesf(t, Default(), c, "Incorrect values in case %d, should be default config", idx)

		} else {
			assert.EqualValuesf(t, tt.want, c, "Incorrect values in case %d", idx)
			assert.NoErrorf(t, err, "Unexpected error in case %d", idx)
		}
	}
}

func TestLoadYamlExportMethod(t *testing.T) {
	type testCase struct {
		path    string
		want    definitions.ExportMethod
		wantErr bool
	}

	exPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	t.Log(exPath)

	tests := []testCase{
		{"mock/case-normal.yaml", definitions.ExportMethodOnlyCBZ, false},
		{"mock/case-empty.yaml", definitions.ExportMethodDefaultWrapCBZ, false},

		{"mock/case-export0.yaml", definitions.ExportMethodDefaultWrapCBZ, false},
		{"mock/case-export1.yaml", definitions.ExportMethodOnlyCBZ, false},
		{"mock/case-export2.yaml", definitions.ExportMethodCustomWrapCBZ, false},
	}

	for idx, tt := range tests {
		// Load YAML file and check result
		c, err := LoadYaml(tt.path)

		assert.EqualValuesf(t, tt.want, c.Preference.ExportMethod, "Incorrect values in case %d", idx)
		assert.NoErrorf(t, err, "Unexpected error in case %d", idx)
	}
}

func TestLoadYamlLogLevel(t *testing.T) {
	tests := []struct {
		input  string
		wanted zerolog.Level
	}{
		{"trace", zerolog.TraceLevel},
		{"debug", zerolog.DebugLevel},
		{"info", zerolog.InfoLevel},
		{"warn", zerolog.WarnLevel},
		{"error", zerolog.ErrorLevel},
		{"fatal", zerolog.FatalLevel},
		{"panic", zerolog.PanicLevel},
		{"disabled", zerolog.Disabled},
	}

	for _, tt := range tests {
		// Prepare an dummy config file for test case
		path := filepath.Join(t.TempDir(), "config.yaml")
		config := fmt.Appendf(nil, "log:\n  level: %s\n", tt.input)

		if err := os.WriteFile(path, config, 0o600); err != nil {
			t.Fatal(err)
		}

		// Test real function
		c, err := LoadYaml(path)

		assert.NoError(t, err)
		assert.Equal(t, tt.wanted, c.LogConfig.Level())
	}
}
