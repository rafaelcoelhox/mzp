package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server        ServerConfig        `toml:"server"`
	Environment   EnvironmentConfig   `toml:"environment"`
	WhatsappCache WhatsappCacheConfig `toml:"whatsappcache"`
}

type ServerConfig struct {
	Port string `toml:"port"`
}

type EnvironmentConfig struct {
	Env string `toml:"env"`
}

type WhatsappCacheConfig struct {
	Driver     string `toml:"driver"`
	Connection string `toml:"connection"`
}

var cfg *Config

func Load(configPath string) (*Config, error) {
	config := &Config{}

	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Printf("Error loading config: %v", err)
		return nil, err
	}

	cfg = config
	return config, nil
}

func Get() *Config {
	return cfg
}
