package patcher

import (
	"KHytaleLauncher/internal/download"
	"KHytaleLauncher/internal/env"
	"KHytaleLauncher/internal/progress"
	"fmt"
	"os"
	"path/filepath"
)

func EnsureButler() (string, error) {
	programDir := env.GetButlerDir()
	osName, osArch := env.GetDeviceInfo()

	programExe := "butler.exe"
	if osName != "windows" {
		programExe = "butler"
	}

	programPath := filepath.Join(programDir, programExe)
	if _, err := os.Stat(programPath); err == nil {
		progress.SetProgressStatus("Butler dosyalari bulundu", -1)
		return programPath, nil
	}

	progress.SetProgressStatus("Butler dosyalari indiriliyor", 0)

	_, err := download.FromUrlAndExtract(
		programDir,
		"butler.zip",
		fmt.Sprintf("https://broth.itch.zone/butler/%s-%s/LATEST/archive/default",
			osName, osArch))

	return programPath, err
}
