package config

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
)

const configFilename = ".gatorconfig.json"

type Config struct {
	DBURL 			string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c Config) SetUser(user string) {
	c.CurrentUserName = user
	err := write(c)
	if err != nil {
		fmt.Printf("Error setting 'user' in config file: %v\n", err)
	}
}

func Read() Config {
	home, err := getConfigFilePath()
	if err != nil {
		fmt.Printf("Error opening config: %v\n", err)
	}
	configPath := filepath.Join(home, configFilename)
	file, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Config file does not exist.")
		} else {
			fmt.Printf("Error opening file: %v\n", err)
		}
		return Config{}
	}
	var cfgFile Config
	err = json.Unmarshal(file, &cfgFile)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
	}
	return cfgFile
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
