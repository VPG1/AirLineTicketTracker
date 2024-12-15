package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env string `yaml:"env" env-required:"true"
`
	Telegram struct {
		Token string `yaml:"token" env-required:"true"` // TODO: move to environment variables
	} `yaml:"telegram"`

	Database struct {
		Host     string `yaml:"host" env-required:"true"`
		Port     int    `yaml:"port" env-required:"true"`
		Username string `yaml:"username" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		DbName   string `yaml:"dbname" env-required:"true"`
	} `yaml:"database"`

	FlightsAPI struct {
		Token   string `yaml:"token" env-required:"true"` // TODO: move to environment variables
		BaseURL string `yaml:"base_url" env-required:"true"`
	} `yaml:"flights_api"`

	Notification struct {
		Interval time.Duration `yaml:"interval" env-default:"15m"`
	} `yaml:"notification"`
}

// MustLoadConfig function for config parsing
func MustLoadConfig(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("Config file does not exist")
	}

	config := &Config{}
	if err := cleanenv.ReadConfig(path, config); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if config.Env != "prod" && config.Env != "debug" {
		log.Fatal("Config file must be either 'prod' or 'debug'")
	}

	return config
}
