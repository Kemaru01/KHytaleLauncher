package game_fix

import (
	"KHytaleLauncher/internal/download"
	"KHytaleLauncher/internal/env"
	"fmt"
	"path/filepath"
)

const (
	fixArchiveAssetsURL = "https://raw.githubusercontent.com/Kemaru01/KHytaleLauncher/refs/heads/main/assets/"
)

func ApplyOnlineFixWindows(
	branch,
	version string) error {
	gameDir := env.GetGameDir(branch, version)
	osName, _ := env.GetDeviceInfo()

	downloadUrl := fmt.Sprintf("%s/%s/game/%s/%s", fixArchiveAssetsURL, osName, branch, version)

	serverJarPath := filepath.Join(gameDir, "Server", "HytaleServer.jar")
	_, err := download.FromUrl(
		serverJarPath,
		fmt.Sprintf("%s/Server/HytaleServer.jar", downloadUrl),
		true)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded server jar fix to %s\n", serverJarPath)

	clientExePath := filepath.Join(gameDir, "Client", "HytaleClient.exe")
	_, err = download.FromUrl(
		clientExePath,
		fmt.Sprintf("%s/Client/HytaleClient.exe", downloadUrl),
		true)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded server jar fix to %s\n", serverJarPath)

	serverStartBatPath := filepath.Join(gameDir, "Server", "start-server.bat")
	_, err = download.FromUrl(
		serverStartBatPath,
		fmt.Sprintf("%s/Server/start-server.bat", downloadUrl),
		true)

	fmt.Printf("Downloaded server jar fix to %s\n", serverJarPath)

	if err != nil {
		return err
	}

	return nil
}
