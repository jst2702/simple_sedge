package db

import (
	"log"

	kitcfg "simplesedge.com/gokit/config"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func DefaultConfig(path string) *Config {
	config := &Config{}
	err := kitcfg.LoadConfig(path, "app", "env", config)
	if err != nil {
		log.Fatal("cannot loading config ", err)
	}
	return config
}
