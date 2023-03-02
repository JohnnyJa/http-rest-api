package apiserver

import (
	"github.com/JohnnyJa/http-rest-api/internal/app/apiclient"
	"github.com/JohnnyJa/http-rest-api/internal/app/store"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
	Client 	 *apiclient.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
		Client: apiclient.NewConfig(),
	}
}
