package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Load yaml config from given path,
// while no koanf instance will preserved (i.e. every call overwrite previous call).
//
// If failed to load config, then a default config will be returned.
func LoadYaml(path string) (*ProgramConfig, error) {
	var k = koanf.New(".")

	// Check if file exist
	if _, err := os.Stat(path); err != nil {
		return Default(), fmt.Errorf("path %s does not exist", path)
	}

	// Start Load file
	err := k.Load(file.Provider(path), yaml.Parser())
	if err != nil {
		return Default(), err
	}

	// Unmarshal to struct
	var out ProgramConfig
	err = k.UnmarshalWithConf("", &out, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true})
	if err != nil {
		return Default(), err
	}

	// Parse path due to relative path issue
	out.DefaultExport, err = parsePath(out.DefaultExport)
	if err != nil {
		return Default(), err
	}

	return &out, nil
}

// Convert relative path to absolute path.
// If path passed is empty string, then it perform nothing.
func parsePath(relativePath string) (absPath string, err error) {
	if relativePath == "" {
		return "", nil
	}

	return filepath.Abs(relativePath)
}
