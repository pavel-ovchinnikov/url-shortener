package config

import (
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	HTTPServer HTTPServerConfig `yaml:"http_server"`
}

type HTTPServerConfig struct {
	Address     string `yaml:"address" env-default:"localhost:8080"`
	Timeout     string `yaml:"timeout" env-default:"10s"`
	IdleTimeout string `yaml:"idle_timeout" env-default:"60s"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("./config/local.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
