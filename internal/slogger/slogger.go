package slogger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger setups the default slog.
func NewLogger(file *os.File) *zap.Logger {
	multiCore := multiCore(file, os.Stdout)
	core := zapcore.NewTee(multiCore)

	logger := zap.New(core)

	return logger
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
