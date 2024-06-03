package main

import (
	"flag"
	"time"
)

// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
var flagRunAddr string

// частота отправки метрик на сервер
var flagReportInterval time.Duration

// частота опроса метрик из пакета runtime
var flagPollInterval time.Duration

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func parseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")

	flag.DurationVar(&flagReportInterval, "r", 10*time.Second, "interval between report calls")

	flag.DurationVar(&flagPollInterval, "p", 2*time.Second, "interval between polling calls")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
