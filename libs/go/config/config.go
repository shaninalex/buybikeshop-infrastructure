package config

import (
	"io"
	"os"

	"go.yaml.in/yaml/v2"
)

type Config struct {
	Port     uint         `yaml:"port"`
	Database string       `yaml:"database"`
	Kratos   KratosConfig `yaml:"kratos"`
}

type KratosConfig struct {
	UrlBrowser string `yaml:"url_browser"`
	UrlAdmin   string `yaml:"url_admin"`
}

func (s *Config) KratosUrlBrowser() string {
	return s.Kratos.UrlBrowser
}

func (s *Config) KratosUrlAdmin() string {
	return s.Kratos.UrlBrowser
}

func (s *Config) GetDatabaseUrl() string {
	return s.Database
}

func ReadConfig(path string) *Config {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := yaml.Unmarshal(b, &config); err != nil {
		panic(err)
	}

	return &config
}

func ProvideConfig(configPath string) func() *Config {
	return func() *Config {
		return ReadConfig(configPath)
	}
}
