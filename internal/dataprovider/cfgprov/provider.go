// Package to get comicinfo data from config,
// which is the program config that used by application and retrieve from program start.
package cfgprov

import (
	"fmt"

	"github.com/dark-person/comicinfo-parser/internal/comicinfo"
	"github.com/dark-person/comicinfo-parser/internal/config"
	"github.com/dark-person/comicinfo-parser/internal/dataprovider"
)

// Data provider that use default value stated in configuration.
type ConfigProvider struct {
	cfg *config.ProgramConfig
}

var _ dataprovider.DataProvider = (*ConfigProvider)(nil)

// Create a new data provider that use default value stated in configuration.
func New(cfg *config.ProgramConfig) *ConfigProvider {
	return &ConfigProvider{cfg: cfg}
}

// Fill input comicinfo by configuration.
// There has two possiblity of fill method:
//
//  1. Fill success, return changed comicinfo and nil error
//  2. Fill failed, return unchanged comicinfo (same as input) and error it self
func (c *ConfigProvider) Fill(input *comicinfo.ComicInfo) (result *comicinfo.ComicInfo, err error) {
	// Ensure configuration is not empty
	if c.cfg == nil {
		return input, fmt.Errorf("configuration cannot be nil")
	}

	// Fill values
	if c.cfg.Metadata.Number != "" {
		input.Number = c.cfg.Metadata.Number
	}

	return input, nil
}
