package services

import (
	"context"

	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/config"
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/lazydb"
	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	DB  *lazydb.LazyDB
	ctx context.Context
	cfg *config.ProgramConfig

	lastExportedComic string // last exported comic folder path, for soft delete purpose
}

// Create a new app with specified config.
// This function is designed to run in production.
func NewApp(cfg *config.ProgramConfig, db *lazydb.LazyDB) *App {
	return &App{DB: db, cfg: cfg}
}

// Create a new app with default configuration.
// Suggested to use in testing only.
func NewAppWithDefaultConfig(db *lazydb.LazyDB) *App {
	return &App{DB: db, cfg: assets.Config()}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	var err error

	// Perform connect & migration
	err = a.DB.Connect()
	if err != nil {
		panic(err)
	}

	// Perform migration to database if needed
	tmpPath, err := a.DB.Migrate()
	log.Info().Str("backup path", tmpPath).Send()

	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Error",
			Message: "Database version is corrupted. To fix this issue, you can:\n\n1. Change database by config.yaml.\n2. Backup and remove the database file.\n3. Fix your database (Advanced User ONLY).",
		})
		log.Panic().Err(err).Msg("Failed to start application due to database corrupted")
	}
}

// LoadSessionData will get session data that will be used in frontend for whole application session.
func (a *App) LoadSessionData() definitions.SessionData {
	return definitions.SessionData{
		ExportMethod:      a.cfg.Preference.ExportMethod,
		DeleteAfterExport: a.cfg.Preference.DeleteAfterExport,
	}
}
