package main

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/services"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails/v3/pkg/application"
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

	app := application.New(application.Options{
		Name: "My App",
		Services: []application.Service{
			application.NewService(services.NewComicService(cfg, l)),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assetsFs),
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "Comicinfo Parser GUI",
		Width:            1200,
		Height:           810,
		BackgroundColour: application.NewRGBA(27, 38, 54, 1),
	})

	err = app.Run()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to start application.")
	}
}
