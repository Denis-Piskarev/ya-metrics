package main

import (
	"log"
	"net/http"

	"github.com/DenisquaP/ya-metrics/internal/server/handlers"
)

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on " + cfg.runAddr + "...")
	router := handlers.InitRouter()

	if err := http.ListenAndServe(cfg.runAddr, router); err != nil {
		panic(err)
	}
}
