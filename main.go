package main

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/services"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assetsFs embed.FS

func main() {
	// Load Config
	cfg := assets.Config()

	// Setup logger
	setupLogger(zerolog.InfoLevel)

	// Get Home Directory
	home, err := os.UserHomeDir()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to get home directory")
	}

	// ------------------------------------------------------

	// Get database path
	path := filepath.Join(home, assets.RootDir, assets.DatabaseFile)
	backupDir := filepath.Join(home, assets.RootDir, assets.BackupDir)

	// Replace path to create database if config specified path
	if cfg.Database.Path != "" {
		path = cfg.Database.Path
	}

	// Prepare lazydb
	l := assets.DefaultDBWithBackup(path, backupDir)

	// ------------------------------------------------------

	// Create an instance of the app structure
	app := services.NewApp(cfg, l)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Comicinfo Parser GUI",
		Width:  1200,
		Height: 810,
		AssetServer: &assetserver.Options{
			Assets: assetsFs,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
		EnumBind: []interface{}{
			services.AllMangaValue,
			services.AllAgeRatingValue,
			services.AllExportMethodValue,
		},
	})

	if err != nil {
		log.Panic().Err(err).Msg("Failed to start application.")
	}
}
