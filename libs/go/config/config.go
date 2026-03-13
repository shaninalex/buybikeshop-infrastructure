package config

import (
	"io"
	"os"

	"github.com/google/uuid"
	"go.yaml.in/yaml/v2"
)

type Config struct {
	Port       uint         `yaml:"port"`
	Database   string       `yaml:"database"`
	Kratos     KratosConfig `yaml:"kratos"`
	OAuth      OAuthConfig  `yaml:"oauth"`
	Datasource Datasource   `yaml:"datasource"`
}

type OAuthConfig struct {
	ClientID         uuid.UUID `yaml:"client_id"`
	RedirectUri      string    `yaml:"redirect_uri"`
	AuthorizationUrl string    `yaml:"authorization_url"`
	TokenUrl         string    `yaml:"token_url"`
	Scopes           []string  `yaml:"scopes"`
}

type Datasource struct {
	GrpcPort uint `yaml:"grpc_port"`
}

func (s *Config) GetOAuthConfig() OAuthConfig {
	return s.OAuth
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
