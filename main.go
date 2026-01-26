package main

import (
	"KHytaleLauncher/internal/app"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := app.New()

	err := wails.Run(&options.App{
		Title:  "KHytale Launcher",
		Width:  1024,
		Height: 768,

		DisableResize:        true,
		DisablePanicRecovery: true,
		Frameless:            true,

		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},

		Assets:    assets,
		OnStartup: app.Startup,

		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error on Initilaized wails App:", err.Error())
	}
}
