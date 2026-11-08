package config

import (
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Config for this program.
type ProgramConfig struct {
	Preference preferenceConfig `koanf:"default"`   // Config setting for general user preference
	Metadata   metadataConfig   `koanf:"metadata"`  // Config for metadata default values
	Database   databaseConfig   `koanf:"database"`  // Config for database to use
	TrashBin   trashBinConfig   `koanf:"trash-bin"` // Config of trash bin usage
	LogConfig  logConfig        `koanf:"log"`       // Config for log settings
}

// Config for some user preference, e.g. default folder used.
type preferenceConfig struct {
	ComicDir          string                   `koanf:"comic-folder"`  // Default input directory, all folder select dialog will start from here
	ExportDir         string                   `koanf:"export-folder"` // Default export folder, apply to both quick & standard
	ExportMethod      definitions.ExportMethod `koanf:"export-method"`
	DeleteAfterExport bool                     `koanf:"delete-after-export"`
}

// Configuration for default value of metadata.
type metadataConfig struct {
	Number string `koanf:"default-number"`
}

// Config for database setting.
type databaseConfig struct {
	Path string `koanf:"path"` // Database path, empty string imply use default path instead
}

// Config for trash bin location.
type trashBinConfig struct {
	Path string `koanf:"path"` // Trash bin location, empty string imply no trash bin defined
}

// logConfig is config for log settings.
type logConfig struct {
	LogLevelStr string `koanf:"level"` // Log level, empty string imply use INFO. Use Level() to get zerolog level.
}

// Level will convert logConfig value into zerolog.Level for global log level setting.
func (l *logConfig) Level() zerolog.Level {
	// Use default level is not specific
	if l.LogLevelStr == "" {
		return zerolog.InfoLevel
	}

	// Try to parse level from zerolog
	level, err := zerolog.ParseLevel(l.LogLevelStr)

	// Use default level if any error
	if err != nil {
		log.Warn().Err(err).Msg("Failed to parse log. Use INFO instead.")
		return zerolog.InfoLevel
	}

	return level
}

// Default config struct for this program.
func Default() *ProgramConfig {
	return &ProgramConfig{
		Preference: preferenceConfig{
			ComicDir:          "", // Indicate use wails default directory
			ExportDir:         "", // Indicate input folder is used
			ExportMethod:      definitions.ExportMethodDefaultWrapCBZ,
			DeleteAfterExport: false,
		},
		Metadata: metadataConfig{
			Number: "",
		},
		Database: databaseConfig{
			Path: "", // Indicate default database location is used
		},
		TrashBin: trashBinConfig{
			Path: "", // No trash bin location defined
		},
		LogConfig: logConfig{
			LogLevelStr: "",
		},
	}
}
