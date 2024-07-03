package server

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	// Server address and port
	RunAddr string `env:"ADDRESS" envDefault:"localhost:8080"`
	// Interval between saving metrics to file
	StoreInterval int `env:"STORE_INTERVAL" envDefault:"300"`
	// Path to file with metrics
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"/tmp/metrics-db.json"`
	// Restore metrics from file
	Restore bool `env:"RESTORE" envDefault:"true"`
	// Database address
	DatabaseDsn string `env:"DATABASE_DSN" envDefault:"host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"`
}

func NewConfig() (Config, error) {
	var cfg Config

	// Setting values by flags, if env not empty, using env
	flag.StringVar(&cfg.RunAddr, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&cfg.StoreInterval, "i", 300, "interval between saving metrics to file")
	flag.StringVar(&cfg.FileStoragePath, "f", "/tmp/metrics-db.json", "path to file with metrics")
	flag.BoolVar(&cfg.Restore, "r", true, "restore metrics from file")
	flag.StringVar(&cfg.DatabaseDsn, "d", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", "database address")

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	flag.Parse()
	return cfg, nil
}
