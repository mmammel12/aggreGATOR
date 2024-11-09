// Package config -
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func getConfigFilePath() (string, error) {
	const configFileName = ".gatorconfig.json"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not find home directory")
	}

	return fmt.Sprintf("%v/%v", homeDir, configFileName), nil
}

// Read -
func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Could not find config file path")
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("Could not read gatorconfig")
	}

	config := Config{}
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return Config{}, fmt.Errorf("Could not unmarshal config json")
	}

	return config, nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Could not find config file path")
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Could not marshal config")
	}

	err = os.WriteFile(filePath, data, 0666)
	if err != nil {
		return fmt.Errorf("Could not write to config file")
	}

	return nil
}

// SetUser -
func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	err := write(*c)
	if err != nil {
		return err
	}

	return nil
}
