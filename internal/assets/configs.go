package assets

import (
	"github.com/dark-person/comicinfo-parser/internal/config"
	"github.com/rs/zerolog/log"
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
		log.Warn().Err(err).Msg("Failed to load config.yaml, use default instead")
	}

	// Set loaded config to package
	cfg = c
	return cfg
}
