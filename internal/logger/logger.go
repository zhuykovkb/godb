package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"sync"
)

var (
	logger   *zap.Logger
	initOnce sync.Once
)

type OutputTarget struct {
	Writer io.Writer
}

type Config struct {
	Level   string
	Outputs []OutputTarget
	Format  string
}

func Init(cfg Config) {
	initOnce.Do(func() {
		var cores []zapcore.Core
		level := parseLogLevel(cfg.Level)
		encoder := selectEncoder(cfg.Format)

		for _, target := range cfg.Outputs {
			core := zapcore.NewCore(encoder, zapcore.AddSync(target.Writer), level)
			cores = append(cores, core)
		}

		logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	})
}

func selectEncoder(format string) zapcore.Encoder {
	if format == "json" {
		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
}

func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func Info(msg string) {
	logger.Info(msg)
}

func Warn(msg string) {
	logger.Warn(msg)
}

func Fatal(msg string) {
	logger.Fatal(msg)
	os.Exit(1)
}
