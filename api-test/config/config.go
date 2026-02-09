package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
)

type Config struct {
	Database Database `yaml:"databases"`
}

func NewConfig(filePath string) (*Config, error) {
	if filePath == "" {
		return nil, errors.New("config file path is empty")
	}

	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %w", err)
	}

	var cfg Config
	err = yaml.Unmarshal(configBytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &cfg, nil
}
