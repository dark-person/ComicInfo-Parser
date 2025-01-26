package application

import (
	"context"
	"fmt"
	"os"

	"github.com/dark-person/comicinfo-parser/internal/assets"
	"github.com/dark-person/comicinfo-parser/internal/config"
	"github.com/dark-person/lazydb"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	DB  *lazydb.LazyDB
	ctx context.Context
	cfg *config.ProgramConfig

	lastExportedComic string // last exported comic folder path, for soft delete purpose
}

// NewApp creates a new App application struct
func NewApp(db *lazydb.LazyDB) *App {
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
	fmt.Println("Backup Path: " + tmpPath)

	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Error",
			Message: "Database version is corrupted, please fix database or remove current database file.",
		})
		os.Exit(1)
	}
}
