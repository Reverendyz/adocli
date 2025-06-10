package common

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	OrganizationUrl string `json:"organizationUrl"`
	ProjectId       string `json:"projectId"`
	TeamId          string `json:"teamId"`
}

func GetOrCreateConfigPath() (string, error) {
	var homeDir string
	var err error

	if runtime.GOOS == "windows" {
		homeDir = os.Getenv("USERPROFILE")
	} else {
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return "", err
		}
	}
	configDir := filepath.Join(homeDir, "config")
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return "", err
	}

	return filepath.Join(configDir, ".adocli"), nil
}

func SetCache(config Config) error {
	configPath, err := GetOrCreateConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func GetFromConfig(key string) (string, error) {
	configPath, err := GetOrCreateConfigPath()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", nil
	}

	var config map[string]string
	if err := json.Unmarshal(data, &config); err != nil {
		return "", err
	}

	if value, exists := config[key]; exists {
		return value, nil
	}

	return "", fmt.Errorf("key '%s' not found in cache", key)
}
