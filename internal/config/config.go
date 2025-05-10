package config

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	API      API      `yaml:"api"`
	WebAuthn WebAuthn `yaml:"webauthn"`
}

type API struct {
	Port int `yaml:"port"`

	ShutdownTimeoutSec int `yaml:"shutdownTimeoutSec"`
}

type WebAuthn struct {
	RPDisplayName string `yaml:"RPDisplayName"`
	RPID          string `yaml:"RPID"`
	RPOrigin      string `yaml:"RPOrigin"`
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
