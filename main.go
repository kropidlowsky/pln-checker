package main

import (
	"github.com/kropidlowsky/pln-checker/cmd"
	"github.com/kropidlowsky/pln-checker/internal/config"
	"github.com/kropidlowsky/pln-checker/internal/slogger"
)

func main() {
	config := config.LoadConfig()
	logger, closer := slogger.NewLogger(config.LogFile)
	defer closer()

	cmd.Execute(logger)
}
