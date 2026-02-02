package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type AppConfig struct {
	Version    string `json:"version"`
	PlayerUUID string `json:"playerUUID"`
}

var AppConf *AppConfig = nil

func RestoreDefaultConfig() (*AppConfig, error) {
	return &AppConfig{
		Version:    "0.2",
		PlayerUUID: uuid.New().String(),
	}, nil
}

func EnsureConfig(filePath string) (*AppConfig, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		defaultConfig, err := RestoreDefaultConfig()
		if err != nil {
			return nil, err
		}

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(defaultConfig); err != nil {
			return nil, err
		}

		return defaultConfig, nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config AppConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return RestoreDefaultConfig()
	}

	return &config, nil
}

func Get(appDir string) *AppConfig {
	configPath := filepath.Join(appDir, "launcher.json")

	config, err := EnsureConfig(configPath)
	if err != nil {
		return &AppConfig{}
	}

	AppConf = config
	return config
}
