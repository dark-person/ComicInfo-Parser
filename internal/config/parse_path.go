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

	cfg.DefaultExport, err = parsePath(cfg.DefaultExport)
	if err != nil {
		return err
	}

	cfg.DefaultComicDir, err = parsePath(cfg.DefaultComicDir)
	if err != nil {
		return err
	}

	return nil
}
