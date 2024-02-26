package config

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	API API `yaml:"api"`
}

type API struct {
	Port int `yaml:"port"`

	ShutdownTimeoutSec int `yaml:"shutdownTimeoutSec"`
}

func NewConfig(filepath string) (*Config, error) {
	yamlBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := new(Config)

	if err := yaml.Unmarshal(yamlBytes, config); err != nil {
		return nil, err
	}

	return config, nil
}
