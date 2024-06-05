package agent

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
	RunAddr string `env:"ADDRESS"`

	// частота отправки метрик на сервер
	ReportInterval int `env:"REPORT_INTERVAL"`

	// частота опроса метрик из пакета runtime
	PollInterval int `env:"POLL_INTERVAL"`
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

	if cfg.ReportInterval == 0 {
		flag.IntVar(&cfg.ReportInterval, "r", 10, "interval between report calls")
	}

	if cfg.PollInterval == 0 {
		flag.IntVar(&cfg.PollInterval, "p", 2, "interval between polling calls")
	}

	flag.Parse()
	return cfg, nil
}
