package main

import (
	"log"
	"os"

	"github.com/kropidlowsky/pln-checker/cmd"
	"github.com/kropidlowsky/pln-checker/internal/config"
	"github.com/kropidlowsky/pln-checker/internal/slogger"
)

func main() {
	config := config.LoadConfig()

	file, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalf("error wile opening the log file: %s", err.Error())
	}

	defer file.Close()

	logger := slogger.NewLogger(file)

	cmd.Execute(logger)
}
