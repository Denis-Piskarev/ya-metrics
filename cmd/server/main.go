package main

import (
	"log"
	"net/http"

	"github.com/DenisquaP/ya-metrics/internal/server/handlers"
)

func main() {
	parseFlags()
	log.Println("Starting server on port " + flagRunPort + "...")
	router := handlers.InitRouter()

	if err := http.ListenAndServe(flagRunHost+":"+flagRunPort, router); err != nil {
		panic(err)
	}
}
