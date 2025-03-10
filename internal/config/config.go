package config

import (
	"encoding/json"
	"io"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	Dburl           string `json:"dburl"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homePath + configFileName, nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = json.Unmarshal(data, &data)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c Config) SetUser() error {
	return nil
}
