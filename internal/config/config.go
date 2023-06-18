package config

import (
	"github.com/spf13/viper"
)

// Config represents the application configuration structure
type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	JWTSecret string
}

// LoadConfig loads the application configuration from environment variables or a config file
func LoadConfig() *Config {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
