package config

import (
	"os"
	"path/filepath"
	"testing"

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
			Folder: folderConfig{
				ExportDir: filepath.Join(exPath, "./my-export"),
				ComicDir:  filepath.Join(exPath, "./my-input"),
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
