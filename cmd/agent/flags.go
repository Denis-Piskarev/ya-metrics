package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type config struct {
	runAddr string `env:"ADDRESS"`
}

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func initConfig() (config, error) {
	var cfg config

	err := env.Parse(&cfg)
	if err != nil {
		return config{}, err
	}

	if cfg.runAddr == "" {
		flag.StringVar(&cfg.runAddr, "a", "localhost:8080", "address and port to run server")
	}

	return cfg, nil
}
