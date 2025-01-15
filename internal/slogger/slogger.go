package slogger

import (
	"io"
	"log"
	"log/slog"
	"os"
)

// NewLogger setups the default slog.
func NewLogger(filename string) (*slog.Logger, func()) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error wile opening the log file: %s", err.Error())
	}

	multiWriter := io.MultiWriter(file, os.Stdout)
	logger := slog.New(slog.NewTextHandler(multiWriter, nil))

	return logger, func() {
		file.Close()
	}
}
