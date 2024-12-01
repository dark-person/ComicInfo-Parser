package assets

import (
	"fmt"

	"github.com/dark-person/comicinfo-parser/internal/config"
)

var cfg *config.ProgramConfig

// Load config from yaml file.
// If any error occur in loading, then a default config will be returned, and no error return.
func Config() *config.ProgramConfig {
	// Return parsed config if any
	if cfg != nil {
		return cfg
	}

	// Load Config from filesystem
	c, err := config.LoadYaml("config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	// Set loaded config to package
	cfg = c
	return cfg
}
