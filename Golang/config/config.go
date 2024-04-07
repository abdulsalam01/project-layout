package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func NewConfigManager(path string) (*Config, error) {
	var (
		config Config
		err    error

		filePath string = path + configFile
	)

	raw, err := os.ReadFile(filePath)
	if err != nil {
		return &config, err
	}

	// Unmarshal.
	err = yaml.Unmarshal([]byte(raw), &config)
	if err != nil {
		return &config, err
	}

	return &config, nil
}
