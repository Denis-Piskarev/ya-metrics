package main

import (
	"flag"
	"strings"
)

// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
var flagRunHost string
var flagRunPort string

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func parseFlags() {
	var addr string
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&addr, "a", "localhost:8080", "address and port to run server")

	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()

	sliceAddr := strings.Split(addr, ":")

	flagRunHost = sliceAddr[0]
	flagRunPort = sliceAddr[1]
}
