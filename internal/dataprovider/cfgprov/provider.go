// Package to get comicinfo data from config,
// which is the program config that used by application and retrieve from program start.
package cfgprov

import (
	"fmt"
	"path/filepath"

	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/config"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider"
	"github.com/dark-person/comicinfo-parser/internal/files"
)

// Data provider that use default value stated in configuration.
type ConfigProvider struct {
	folderPath string
	cfg        *config.ProgramConfig
}

var _ dataprovider.DataProvider = (*ConfigProvider)(nil)

// Create a new data provider that use default value stated in configuration.
func New(cfg *config.ProgramConfig, folderPath string) *ConfigProvider {
	return &ConfigProvider{cfg: cfg, folderPath: folderPath}
}

// Fill input comicinfo by configuration.
// When comicinfo.xml is already exist, this provider will have no-ops.
// There has two possiblity of fill method:
//
//  1. Fill success, return changed comicinfo and nil error
//  2. Fill failed, return unchanged comicinfo (same as input) and error it self
func (c *ConfigProvider) Fill(input *comicinfo.ComicInfo) (result *comicinfo.ComicInfo, err error) {
	// Ensure configuration is not empty
	if c.cfg == nil {
		return input, fmt.Errorf("configuration cannot be nil")
	}

	// Ignore fill when comicinfo file is exist
	if files.IsFileExist(filepath.Join(c.folderPath, "ComicInfo.xml")) {
		return input, nil
	}

	// Fill values
	if c.cfg.Metadata.Number != "" {
		input.Number = c.cfg.Metadata.Number
	}

	return input, nil
}
