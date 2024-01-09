package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Config struct {
	LogLevel    string `envconfig:"LOGGER_LEVEL" default:"debug"`
	ServiceBind string `envconfig:"BIND_ADDR" default:"0.0.0.0:9095"`
}

var cfg *Config

func MustGet() Config {
	if cfg == nil {
		cfgPath := filepath.Join("local.env")
		_, err := os.Stat(cfgPath)
		if err == nil {
			if err = godotenv.Load(cfgPath); err != nil {
				panic(errors.Wrap(err, "can't load .env config file"))
			}
		}
		cfg = &Config{}
		if err = envconfig.Process("", cfg); err != nil {
			panic(errors.Wrap(err, "can't get config from env"))
		}
	}
	return *cfg
}

func (cfg Config) Logger() (logger zerolog.Logger) {
	level := zerolog.DebugLevel
	newLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err == nil {
		level = newLevel
	}
	return zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()
}
