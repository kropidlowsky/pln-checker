package slogger

import (
	"io"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger setups the default slog.
func NewLogger(filename string) (*zap.Logger, func()) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error wile opening the log file: %s", err.Error())
	}

	multiCore := multiCore(file, os.Stdout)
	core := zapcore.NewTee(multiCore)

	logger := zap.New(core)

	return logger, func() {
		file.Close()
	}
}

func multiCore(writers ...io.Writer) zapcore.Core {
	multiWriter := io.MultiWriter(writers...)
	multiCore := core(multiWriter)
	core := zapcore.NewTee(multiCore)

	return core
}

func core(w io.Writer) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(w),
		zapcore.InfoLevel,
	)
}
