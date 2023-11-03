package api

import (
	"github.com/killabayte/Go/RESTapi/ServerAndDBV2/storage"
	_ "github.com/lib/pq"
)

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
