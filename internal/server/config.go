package server

import (
	"flag"
	"log"
	"net/http"

	"github.com/DenisquaP/ya-metrics/internal/server/handlers"
	"github.com/caarlos0/env/v6"
)

type config struct {
	// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
	RunAddr string `env:"ADDRESS" envDefault:"localhost:8080"`
}

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func NewConfig() (config, error) {
	var cfg config

	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}

	if cfg.RunAddr == "" {
		flag.StringVar(&cfg.RunAddr, "a", "localhost:8080", "address and port to run server")
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
