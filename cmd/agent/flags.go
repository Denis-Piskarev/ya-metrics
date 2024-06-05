package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type config struct {
	// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
	runAddr string `env:"ADDRESS"`

	// частота отправки метрик на сервер
	reportInterval int `env:"REPORT_INTERVAL"`

	// частота опроса метрик из пакета runtime
	pollInterval int `env:"POLL_INTERVAL"`
}

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func initConfig() (config, error) {
	var cfg config

	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}

	if cfg.runAddr == "" {
		flag.StringVar(&cfg.runAddr, "a", "localhost:8080", "address and port to run server")
	}

	if cfg.reportInterval == 0 {
		flag.IntVar(&cfg.reportInterval, "r", 10, "interval between report calls")
	}

	if cfg.pollInterval == 0 {
		flag.IntVar(&cfg.pollInterval, "p", 2, "interval between polling calls")
	}

	return cfg, nil
}
