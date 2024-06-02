package main

import (
	"log"
	"net/http"

	"github.com/DenisquaP/ya-metrics/internal/server/handlers"
)

func main() {
	log.Println("Starting server...")
	router := handlers.InitRouter()

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
