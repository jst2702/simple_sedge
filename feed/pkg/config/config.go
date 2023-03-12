package db

import (
	"log"
	"time"

	kitcfg "simplesedge.com/gokit/config"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenduration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func DefaultConfig(path string) *Config {
	config := &Config{}
	err := kitcfg.LoadConfig(path, "app", "env", config)
	if err != nil {
		log.Fatal("cannot loading config ", err)
	}
	return config
}
