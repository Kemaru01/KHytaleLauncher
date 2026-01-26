package game

import (
	"KHytaleLauncher/internal/config"
	"KHytaleLauncher/internal/env"
	game_fix "KHytaleLauncher/internal/game/fix"
	"KHytaleLauncher/internal/java"
	"KHytaleLauncher/internal/patcher"
	"KHytaleLauncher/internal/progress"
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var isGameRunning = false

func Launch(ctx context.Context, playerName, gameVersion string) error {
	progress.SetProgressStatus("Oyun calistiriliyor...", -1)

	javaVersion := "25.0.1_8"
	gameBranch := "release"

	if isGameRunning {
		return nil
	}

	isGameRunning = true

	if err := env.EnsurePreFolders(javaVersion, gameBranch, gameVersion); err != nil {
		return err
	}

	butlerPath, err := patcher.EnsureButler()
	if err != nil {
		return err
	}

	javaPath, err := java.EnsureJRE(javaVersion)
	if err != nil {
		return err
	}

	hytalePath, err := EnsureGame(gameBranch, gameVersion, butlerPath)
	if err != nil {
		return err
	}

	err = game_fix.EnsureServerAndClientFix(gameBranch, gameVersion)
	if err != nil {
		return err
	}

	playerUUID := config.AppConf.PlayerUUID

	progress.SetProgressStatus(fmt.Sprintf("Oyun calistiriliyor (Hytale - branch: %s, v%s)", gameBranch, gameVersion), 100)

	cmd := exec.Command(hytalePath,
		"--app-dir", env.GetGameDir(gameBranch, gameVersion),
		"--user-dir", env.GetUserDataDir(),
		"--java-exec", javaPath,
		"--auth-mode", "offline",
		//"--identity-token", utils.GenerateDummyJwt(playerUuid, playerUuid),
		//"--session-token", utils.GenerateDummyJwt(playerUuid, playerUuid),
		"--uuid", playerUUID,
		"--name", playerName)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	runtime.Hide(ctx)

	go func() {
		_ = cmd.Wait()
		isGameRunning = false

		progress.ClearProgress()
		runtime.Show(ctx)
		//runtime.Focus(ctx)
	}()

	return nil
}
