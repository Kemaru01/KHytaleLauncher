package app

import (
	"KHytaleLauncher/internal/config"
	"KHytaleLauncher/internal/env"
	"KHytaleLauncher/internal/game"
	"KHytaleLauncher/internal/progress"
	"context"
	"fmt"
	"os/exec"
	"runtime"
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

	if err := game.Launch(app.ctx, playerName, gameVersion); err != nil {
		fmt.Printf("Error launching the game: %v\n", err)
	}
}

func (app *App) OpenToDir() error {
	var cmd *exec.Cmd

	path := env.GetDefaultAppDir()
	osName, _ := env.GetDeviceInfo()

	switch osName {
	case "windows":
		cmd = exec.Command("explorer.exe", path)

	case "darwin":
		cmd = exec.Command("open", path)

	case "linux":
		cmd = exec.Command("xdg-open", path)

	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	return cmd.Start()
}
