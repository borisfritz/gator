package config

import (
	"fmt"
	"errors"
	"os"
	"path/filepath"
	"encoding/json"
)

const configFilename = ".gatorconfig.json"

type Config struct {
	DBURL 			string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c Config) SetUser(user string) error {
	c.CurrentUserName = user
	err := write(c)
	if err != nil {
		return fmt.Errorf("Error setting 'user' in config file: %w", err)
	}
	return nil
}

func Read() (Config, error) {
	home, err := getConfigFilePath()
	if err != nil {
		fmt.Printf("Error opening config: %v\n", err)
	}
	configPath := filepath.Join(home, configFilename)
	file, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, errors.New("Config file does not exist.")
		} else {
			return Config{}, fmt.Errorf("Error opening file: %v\n", err)
		}
	}
	var cfgFile Config
	err = json.Unmarshal(file, &cfgFile)
	if err != nil {
		return Config{}, fmt.Errorf("Error parsing JSON: %v\n", err)
	}
	return cfgFile, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Unable to get local home directory: %w", err)
	}
	return home, nil
}

func write(cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return fmt.Errorf("Error marshaling data: %w", err)
	}
	home, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Error opening config: %w", err)
	}
	configPath := filepath.Join(home, configFilename)
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("Error writing file: %w", err)
	}
	return nil
}
