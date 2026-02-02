package env

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	AppEnv  = "prod" // default
	AppName = "KHytaleLauncher"
)

func GetDeviceInfo() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}

func GetDefaultAppDir() string {
	base := os.Getenv("LOCALAPPDATA")

	if base == "" {
		home, _ := os.UserHomeDir()
		base = home
	}

	return filepath.Join(base, "KHytaleLauncher")
}

func GetDefaultCacheDir() string {
	appDir := GetDefaultAppDir()
	return filepath.Join(appDir, "Cache")
}

func GetButlerDir() string {
	appDir := GetDefaultAppDir()
	return filepath.Join(appDir, "packages", "tools", "butler")
}

func GetJavaDir(version string) string {
	appDir := GetDefaultAppDir()
	return filepath.Join(appDir, "packages", "jre", version)
}

func GetUserDataDir() string {
	appDir := GetDefaultAppDir()
	return filepath.Join(appDir, "UserData")
}

func GetLogsDir() string {
	appDir := GetDefaultAppDir()
	return filepath.Join(appDir, "Logs")
}

func GetGameDir(branch, version string) string {
	appDir := GetDefaultAppDir()
	return filepath.Join(appDir, "packages", "game", branch, version)
}

func EnsurePreFolders(javaVersion, gameBranch, gameVersion string) error {
	paths := []string{
		GetDefaultAppDir(),
		GetUserDataDir(),
		GetDefaultCacheDir(),
		GetLogsDir(),
		GetButlerDir(),
		GetJavaDir(javaVersion),
		GetGameDir(gameBranch, gameVersion),
	}

	for _, p := range paths {
		if err := os.MkdirAll(p, 0755); err != nil {
			return err
		}
	}
	return nil
}
