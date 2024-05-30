package logger

import (
	"go.uber.org/zap"
)

func NewLogger(logLevel string) *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.Level.SetLevel(parseLogLevel(logLevel))
	logger, _ := config.Build()
	return logger.Sugar()
}

func parseLogLevel(logLevel string) zap.AtomicLevel {
	switch logLevel {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}
