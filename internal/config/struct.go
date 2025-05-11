// Package for configuration,
// provide basic config struct and utility to load.
package config

// Config for this program.
type ProgramConfig struct {
	Folder   folderConfig   `koanf:"default"`   // Config setting for folders
	Metadata metadataConfig `koanf:"metadata"`  // Config for metadata default values
	Database databaseConfig `koanf:"database"`  // Config for database to use
	TrashBin trashBinConfig `koanf:"trash-bin"` // Config of trash bin usage
}

// Config for folder to be used.
type folderConfig struct {
	ComicDir  string `koanf:"comic-folder"`  // Default input directory, all folder select dialog will start from here
	ExportDir string `koanf:"export-folder"` // Default export folder, apply to both quick & standard
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
	Path string `koanf:"path"` // Trash bin location, empty string implment no trash bin defined
}

// Default config struct for this program.
func Default() *ProgramConfig {
	return &ProgramConfig{
		Folder: folderConfig{
			ComicDir:  "", // Indicate use wails default directory
			ExportDir: "", // Indicate input folder is used
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
	}
}
