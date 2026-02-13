package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBurl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const (
	configFileName = "/.gatorconfig.json"
)

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return homeDir + configFileName, nil
}

func getConfigData() ([]byte, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Read() (Config, error) {
	data, err := getConfigData()
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err = json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, os.ModeAppend)
	if err != nil {
		return err
	}

	return nil
}
