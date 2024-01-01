package configuration

import (
	"github.com/pithawatnuckong/go-clean/environment"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	PRODUCTION = "production"

	INFO  = "INFO"
	DEBUG = "DEBUG"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

func NewZapLogging(configuration environment.LoggingEnv, finder environment.Environment) *zap.Logger {
	var logger *zap.Logger

	logger = zap.Must(newProduction(loggingLevel(configuration.Level)))
	if strings.ToLower(finder.Get("APP_ENV")) != PRODUCTION {
		logger = zap.Must(zap.NewDevelopment())
	}

	return logger
}

func newProduction(level zapcore.Level) (*zap.Logger, error) {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "timestamp"
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	configuration := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encodeConfig,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	return configuration.Build()
}

func loggingLevel(level string) zapcore.Level {
	switch level {
	case INFO:
		return zapcore.InfoLevel
	case DEBUG:
		return zapcore.DebugLevel
	case ERROR:
		return zapcore.ErrorLevel
	case FATAL:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
