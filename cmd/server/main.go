package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DenisquaP/ya-metrics/internal/server/handlers"
)

func main() {
	parseFlags()
	log.Println("Starting server on port " + flagRunAddr + "...")
	router := handlers.InitRouter()

	if err := http.ListenAndServe(fmt.Sprintf("%s", flagRunAddr), router); err != nil {
		panic(err)
	}
}
