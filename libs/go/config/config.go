package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func ReadConfig(path string) *Config {
	s := &Config{
		v: viper.New(),
	}
	s.v.AddConfigPath(path)
	s.v.SetConfigType("yml")
	s.v.SetConfigName("config.dev")
	s.v.WatchConfig() //we need it to reload config in containers without restart

	err := s.v.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Can't open config file. %s \n", err))
	}

	return s
}

func ProvideConfig(configPath string) func() *Config {
	return func() *Config {
		return ReadConfig(configPath)
	}
}

func (s *Config) Int(param string) int { return s.v.GetInt(param) }

func (s *Config) String(param string) string { return s.v.GetString(param) }

func (s *Config) Bool(param string) bool { return s.v.GetBool(param) }

func (s *Config) StringSlice(param string) []string { return s.v.GetStringSlice(param) }
