package main

import (
	"github.com/DenisquaP/ya-metrics/internal/server"
	"log"
	"net/http"

	"github.com/DenisquaP/ya-metrics/internal/server/handlers"
)

func main() {
	cfg, err := server.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on " + cfg.RunAddr + "...")
	router := handlers.InitRouter()

	if err := http.ListenAndServe(cfg.RunAddr, router); err != nil {
		panic(err)
	}
}
