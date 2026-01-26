package game

import (
	"KHytaleLauncher/internal/download"
	"KHytaleLauncher/internal/env"
	"KHytaleLauncher/internal/progress"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func EnsureGame(
	branch,
	version,
	butlerPath string) (string, error) {

	gameDir := env.GetGameDir(branch, version)
	osName, osArch := env.GetDeviceInfo()

	programExe := "HytaleClient.exe"
	if osName != "windows" {
		programExe = "HytaleClient"
	}

	programPath := filepath.Join(gameDir, "Client", programExe)
	if _, err := os.Stat(programPath); err == nil {
		progress.SetProgressStatus(fmt.Sprintf("Hytale oyun dosyalari bulundu (v%s)", version), -1)
		return programPath, nil
	}

	progress.SetProgressStatus(fmt.Sprintf("Hytale oyun dosyalari indiriliyor (v%s)", version), 0)

	cacheDir := env.GetDefaultCacheDir()
	cachedPwrPath := filepath.Join(cacheDir, fmt.Sprintf("hytale-data-%s_%s.pwr", branch, version))

	_, err := download.FromUrl(
		cachedPwrPath,
		fmt.Sprintf("https://game-patches.hytale.com/patches/%s/%s/%s/0/%s.pwr",
			osName, osArch, branch, version),
		false)

	if err != nil {
		return "", fmt.Errorf("Download failed: %w", err)
	}

	stagingDir := filepath.Join(cacheDir, fmt.Sprintf("hytale-data-%s_%s", branch, version))
	cmd := exec.Command(
		butlerPath,
		"apply",
		"--staging-dir", stagingDir,
		cachedPwrPath,
		gameDir,
	)

	progress.SetProgressStatus(fmt.Sprintf("Oyun dosyalari cikartiliyor (Hytale v%s)", version), 0)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		progress.SetProgressStatus(fmt.Sprintf("Oyun dosyalari cikarilirlirken hata olustu (Hytale v%s)", version), -1)
		return "", fmt.Errorf("butler apply failed: %w", err)
	}

	progress.SetProgressStatus(fmt.Sprintf("Oyun dosyalari cikarilma basarili (Hytale v%s)", version), 100)

	return programPath, nil
}
