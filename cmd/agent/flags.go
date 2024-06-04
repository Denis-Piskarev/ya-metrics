package main

import (
	"flag"
)

// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
var flagRunAddr string

// частота отправки метрик на сервер
var flagReportInterval int

// частота опроса метрик из пакета runtime
var flagPollInterval int

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func parseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "address and port to run server")

	flag.IntVar(&flagReportInterval, "r", 10, "interval between report calls")

	flag.IntVar(&flagPollInterval, "p", 2, "interval between polling calls")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
