package application

import (
	"context"
	"gui-comicinfo/internal/database"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	DB  *database.AppDB
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	// Init Database
	var err error
	a.DB, err = database.NewDB()
	if err != nil {
		panic(err)
	}

	// Perform connect & migration
	err = a.DB.Connect()
	if err != nil {
		panic(err)
	}

	// Perform migration to database if needed
	err = a.DB.StepToLatest()
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Error",
			Message: "Database version is corrupted, please fix database or remove current database file.",
		})
		os.Exit(1)
	}
}

// Function that is almost same with `startup()`,
// but different on database handling.
//
// This function MUST not used outside test purposes.
func (a *App) StartUpTest(ctx context.Context, dbPath string) {
	a.ctx = ctx

	// Init Database
	var err error
	a.DB, err = database.NewPathDB(dbPath)
	if err != nil {
		panic(err)
	}

	// Perform connect & migration
	err = a.DB.Connect()
	if err != nil {
		panic(err)
	}

	// Perform migration to database if needed
	err = a.DB.StepToLatest()
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Error",
			Message: "Database version is corrupted, please fix database or remove current database file.",
		})
		os.Exit(1)
	}
}
