package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(logLevel string) *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(parseLogLevel(logLevel))
	logger, _ := config.Build()
	return logger.Sugar()
}

func parseLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
