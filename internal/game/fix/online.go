package game_fix

import (
	"KHytaleLauncher/internal/download"
	"fmt"
	"path/filepath"
)

const (
	fixArchiveURL = "https://raw.githubusercontent.com/ArchDevs/HyLauncher/refs/heads/main/assets"
)

func ApplyOnlineFixWindows(
	gameDir string) error {

	serverJarPath := filepath.Join(gameDir, "Server", "HytaleServer.jar")
	_, err := download.FromUrl(
		serverJarPath,
		fmt.Sprintf("%s/Server/HytaleServer.jar", fixArchiveURL),
		true)
	if err != nil {
		return err
	}

	clientExePath := filepath.Join(gameDir, "Client", "HytaleClient.exe")
	_, err = download.FromUrl(
		clientExePath,
		fmt.Sprintf("%s/Client/HytaleClient.exe", fixArchiveURL),
		true)
	if err != nil {
		return err
	}

	serverStartBatPath := filepath.Join(gameDir, "Server", "start-server.bat")
	_, err = download.FromUrl(
		serverStartBatPath,
		fmt.Sprintf("%s/Server/start-server.bat", fixArchiveURL),
		true)
	if err != nil {
		return err
	}

	return nil
}
