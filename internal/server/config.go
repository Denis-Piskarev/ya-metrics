package server

import (
	"flag"
	"log"
	"net/http"

	"github.com/DenisquaP/ya-metrics/internal/server/handlers"
	"github.com/caarlos0/env/v11"
)

type config struct {
	RunAddr string `env:"ADDRESS" envDefault:"localhost:8082"`
}

func NewConfig() (config, error) {
	var cfg config

	// Setting values by flags, if env not empty, using env
	flag.StringVar(&cfg.RunAddr, "a", "localhost:8080", "address and port to run server")

	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}

	flag.Parse()
	return cfg, nil
}

func Run() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on " + cfg.RunAddr + "...")
	router := handlers.InitRouter()

	if err := http.ListenAndServe(cfg.RunAddr, router); err != nil {
		log.Fatal(err)
	}
}
