// Package for configuration,
// provide basic config struct and utility to load.
package config

// Config for this program.
type ProgramConfig struct {
	DefaultExport string `koanf:"default.export-folder"` // Default export folder, apply to both quick & standard
}

// Default config struct for this program.
func Default() *ProgramConfig {
	return &ProgramConfig{
		DefaultExport: "", // Indicate input folder is used
	}
}
