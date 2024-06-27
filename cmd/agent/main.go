package main

import (
	"log"
	"time"

	"github.com/DenisquaP/ya-metrics/internal/agent"
)

func main() {
	time.Sleep(5 * time.Second)
	log.Println("agent run")
	agent.Run()
}
