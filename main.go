package main

import (
	"embed"
	"gui-comicinfo/internal/application"
	"gui-comicinfo/internal/assets"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assetsFs embed.FS

func main() {
	// Get Home Directory
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Get database path
	path := filepath.Join(home, assets.RootDir, assets.DatabaseFile)

	// Prepare lazydb
	l := assets.DefaultDb(path)

	// Create an instance of the app structure
	app := application.NewApp(l)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Gui-comicInfo-Parser",
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
			application.AllMangaValue,
			application.AllAgeRatingValue,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
