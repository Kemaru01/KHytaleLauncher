package app

import (
	"KHytaleLauncher/internal/config"
	"KHytaleLauncher/internal/env"
	"KHytaleLauncher/internal/game"
	"KHytaleLauncher/internal/progress"
	"context"
)

type App struct {
	ctx    context.Context
	appDir string
	config *config.AppConfig
}

func New() *App {
	return &App{}
}

func (app *App) Startup(ctx context.Context) {
	app.ctx = ctx
	app.appDir = env.GetDefaultAppDir()

	app.config = config.Get(app.appDir)
	progress.InitilaizeProgess(ctx)
}

func (app *App) LaunchTheGame(
	playerName,
	gameVersion string) {

	game.Launch(app.ctx, playerName, gameVersion)
}
