package services

import (
	"context"
	"errors"

	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/config"
	"github.com/dark-person/comicinfo-parser/internal/definitions"
	"github.com/dark-person/lazydb"
	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// ComicService struct
type ComicService struct {
	app *application.App // wails3 application to control everything

	DB  *lazydb.LazyDB
	ctx context.Context
	cfg *config.ProgramConfig

	lastExportedComic string // last exported comic folder path, for soft delete purpose
}

// Create a new ComicService with specified config.
// This function is designed to run in production.
func NewComicService(cfg *config.ProgramConfig, db *lazydb.LazyDB) *ComicService {
	return &ComicService{DB: db, cfg: cfg}
}

// SetApp will set the wails3 application to this service.
// If service is not provided, then related function will likely throw error, not UI can be displayed.
func (s *ComicService) SetApp(app *application.App) {
	s.app = app
}

// Create a new ComicService with default configuration.
// Suggested to use in testing only.
func NewComicServiceWithDefaultConfig(db *lazydb.LazyDB) *ComicService {
	return &ComicService{DB: db, cfg: assets.Config()}
}

// ServiceStartup is called when wails application call `Run()` method.
// Any error returned will cause the application to stop.
func (s *ComicService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	if s.app == nil {
		return errors.New("wails application was not provided")
	}

	// Init database
	err := s.InitDatabase()
	if err != nil {
		s.app.Dialog.Error().
			SetTitle("Error").
			SetMessage("Database version is corrupted. To fix this issue, you can:\n\n1. Change database by config.yaml.\n2. Backup and remove the database file.\n3. Fix your database (Advanced User ONLY).").
			Show()

		log.Error().Err(err).Msg("Failed to start application due to database corrupted")
		return err
	}

	return nil
}

func (s *ComicService) InitDatabase() error {
	// Perform connect & migration
	err := s.DB.Connect()
	if err != nil {
		return err
	}

	// Perform migration to database if needed
	tmpPath, err := s.DB.Migrate()
	log.Info().Str("backup path", tmpPath).Send()

	return err
}

// LoadSessionData will get session data that will be used in frontend for whole application session.
func (s *ComicService) LoadSessionData() definitions.SessionData {
	return definitions.SessionData{
		ExportMethod:      s.cfg.Preference.ExportMethod,
		DeleteAfterExport: s.cfg.Preference.DeleteAfterExport,
	}
}
