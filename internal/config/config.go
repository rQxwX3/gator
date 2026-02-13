package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBurl string `json:"db_url"`
}

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(homeDir + "/.gatorconfig.json")
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err = json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
