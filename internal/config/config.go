// Package config  Writes and read the .gatorconfig.json file
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	cfg := Config{}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return cfg, err
	}

	data, err := os.ReadFile(filepath.Join(homeDir, configFileName))
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func SetUser(username string) error {
	cfg, err := Read()
	if err != nil {
		return err
	}

	cfg.CurrentUserName = username

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	data, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(homeDir, configFileName), data, 0o644)
	if err != nil {
		return err
	}

	return nil
}
