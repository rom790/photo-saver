package server

import "photo-saver/store"

type Config struct {
	Addr  string        `json:"addr"`
	Store *store.Config `json:"store"`
	// Db   string `json:"db"`
}

func NewConfig() *Config {
	return &Config{
		Addr:  ":8080",
		Store: store.NewConfig(),
		// Db:   "photo_saver",
	}
}
