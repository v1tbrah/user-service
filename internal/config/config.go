package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

const (
	defaultLogLvl = zerolog.InfoLevel
	envNameLogLvl = "LOG_LVL"
)

type Config struct {
	GRPCConfig    GRPCConfig
	StorageConfig StorageConfig
	LogLvl        zerolog.Level
}

func NewDefaultConfig() Config {
	return Config{
		GRPCConfig:    newDefaultGRPCConfig(),
		StorageConfig: newDefaultStorageConfig(),
		LogLvl:        defaultLogLvl,
	}
}

func (c *Config) ParseEnv() error {
	c.GRPCConfig.parseEnv()

	c.StorageConfig.parseEnv()

	if err := c.parseEnvLogLvl(); err != nil {
		return err
	}

	return nil
}

func (c *Config) parseEnvLogLvl() error {
	envLogLvl := os.Getenv(envNameLogLvl)
	if envLogLvl != "" {
		logLevel, err := zerolog.ParseLevel(envLogLvl)
		if err != nil {
			return fmt.Errorf("parse log lvl: %s", envLogLvl)
		}
		c.LogLvl = logLevel
	}
	return nil
}
