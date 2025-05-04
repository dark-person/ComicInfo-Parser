package config

import "path/filepath"

// Convert relative path to absolute path.
// If path passed is empty string, then it perform nothing.
func parsePath(relativePath string) (absPath string, err error) {
	if relativePath == "" {
		return "", nil
	}

	return filepath.Abs(relativePath)
}

// Parse all path in config struct to absolute path.
func (cfg *ProgramConfig) parse() error {
	var err error

	cfg.Folder.ExportDir, err = parsePath(cfg.Folder.ExportDir)
	if err != nil {
		return err
	}

	cfg.Folder.ComicDir, err = parsePath(cfg.Folder.ComicDir)
	if err != nil {
		return err
	}

	cfg.TrashBin.Path, err = parsePath(cfg.TrashBin.Path)
	if err != nil {
		return err
	}

	cfg.Database.Path, err = parsePath(cfg.Database.Path)
	return err
}
