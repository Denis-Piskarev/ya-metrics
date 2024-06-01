package main

import (
	"github.com/DenisquaP/ya-metrics/internal/server/handlers"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")
	handlers := handlers.InitHandlers()

	if err := http.ListenAndServe(":8080", handlers); err != nil {
		panic(err)
	}
}
