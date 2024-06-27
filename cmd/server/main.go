package main

import (
	"time"

	"github.com/DenisquaP/ya-metrics/internal/server"
)

func main() {
	time.Sleep(11 * time.Second)
	server.Run()
}
