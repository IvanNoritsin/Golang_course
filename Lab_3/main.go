package main

import (
	"factorial/calculations"
	"flag"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func main() {
	args := os.Args[1:] // Получение аргументов командной строки
	var n int64

	if len(args) > 0 {
		n, _ = strconv.ParseInt(args[0], 10, 64) // Преобразование строки в число
	}

	var useLogging bool
	flag.BoolVar(&useLogging, "log", true, "Enable logging")
	flag.Parse()

	result := calculations.Calculate(n, useLogging)

	if useLogging {
		log.Info("Factorial result:", result)
	} else {
		fmt.Println("Factorial result:", result)
	}
}
