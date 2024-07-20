package config

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	RunAddr string `env:"ADDRESS" envDefault:"localhost:8080"`

	// частота отправки метрик на сервер
	ReportInterval int `env:"REPORT_INTERVAL" envDefault:"10"`

	// частота опроса метрик из пакета runtime
	PollInterval int `env:"POLL_INTERVAL" envDefault:"2"`
}

func NewConfig() (Config, error) {
	var cfg Config

	// Setting values by flags, if env not empty, using env
	flag.StringVar(&cfg.RunAddr, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&cfg.ReportInterval, "r", 10, "interval between report calls")
	flag.IntVar(&cfg.PollInterval, "p", 2, "interval between polling calls")

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	flag.Parse()
	return cfg, nil
}
