package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type сonfig struct {
	// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
	runAddr string `env:"ADDRESS"`
}

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func initConfig() (сonfig, error) {
	var cfg сonfig

	if err := env.Parse(&cfg); err != nil {
		return сonfig{}, err
	}

	if cfg.runAddr == "" {
		flag.StringVar(&cfg.runAddr, "a", "localhost:8080", "address and port to run server")
	}

	return cfg, nil
}
