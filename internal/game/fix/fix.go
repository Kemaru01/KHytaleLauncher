package game_fix

import (
	"KHytaleLauncher/internal/env"
	"KHytaleLauncher/internal/progress"
	"os"
	"path/filepath"
)

func EnsureServerAndClientFix(
	branch,
	version string) error {
	osName, _ := env.GetDeviceInfo()

	if osName != "windows" {
		return nil
	}

	progress.SetProgressStatus("Hytale fix dosyalari indiriliyor...", 0)

	serverBat := filepath.Join(env.GetGameDir(branch, version), "Server", "start-server.bat")
	if _, err := os.Stat(serverBat); err == nil {
		return nil
	}

	if err := ApplyOnlineFixWindows(branch, version); err != nil {
		return err
	}

	return nil
}
