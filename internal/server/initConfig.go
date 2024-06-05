package server

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
	RunAddr string `env:"ADDRESS"`
}

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func InitConfig() (Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	if cfg.RunAddr == "" {
		flag.StringVar(&cfg.RunAddr, "a", "localhost:8080", "address and port to run server")
	}

	return cfg, nil
}
