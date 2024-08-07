package main

import (
	"embed"
	"gui-comicinfo/internal/application"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := application.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Gui-comicInfo-Parser",
		Width:  1200,
		Height: 810,
		AssetServer: &assetserver.Options{
			Assets: assets,
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
