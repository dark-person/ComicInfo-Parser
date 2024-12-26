// Package for configuration,
// provide basic config struct and utility to load.
package config

// Config for this program.
type ProgramConfig struct {
	DefaultComicDir string `koanf:"default.comic-folder"`  // Default input directory, all folder select dialog will start from here
	DefaultExport   string `koanf:"default.export-folder"` // Default export folder, apply to both quick & standard
}

// Default config struct for this program.
func Default() *ProgramConfig {
	return &ProgramConfig{
		DefaultComicDir: "", // Indicate use wails default directory
		DefaultExport:   "", // Indicate input folder is used
	}
}
