package java

import (
	"KHytaleLauncher/internal/download"
	"KHytaleLauncher/internal/env"
	"KHytaleLauncher/internal/progress"
	"fmt"
	"os"
	"path/filepath"
)

func EnsureJRE(javaVersion string) (string, error) {
	programDir := env.GetJavaDir(javaVersion)
	osName, osArch := env.GetDeviceInfo()

	programExe := "java.exe"
	if osName != "windows" {
		programExe = "java"
	}

	programPath := filepath.Join(programDir, "bin", programExe)
	if _, err := os.Stat(programPath); err == nil {
		progress.SetProgressStatus(fmt.Sprintf("JRE (Java: v%s) bulundu", javaVersion), -1)
		return programPath, nil
	}

	progress.SetProgressStatus(fmt.Sprintf("JRE (Java: v%s) indiriliyor", javaVersion), 0)

	_, err := download.FromUrlAndExtract(
		programDir,
		fmt.Sprintf("jre-%s.zip", javaVersion),
		fmt.Sprintf("https://launcher.hytale.com/redist/jre/%s/%s/jre-%s.zip",
			osName, osArch, javaVersion))

	return programPath, err
}
