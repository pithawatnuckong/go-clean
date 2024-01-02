package logs

import (
	"github.com/pithawatnuckong/go-clean/environment"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	production = "production"
)

var logs *zap.Logger

func NewCustomerLogger(configuration environment.LoggingEnv, finder environment.Environment) {

	logs = zap.Must(newProduction(loggingLevel(configuration.Level)))
	if strings.ToLower(finder.Get("APP_ENV")) != production {
		logs = zap.Must(zap.NewDevelopment())
	}
}

func Info(message string, fields ...zapcore.Field) {
	logs.Info(message, fields...)
}

func Debug(message string, fields ...zapcore.Field) {
	logs.Debug(message, fields...)
}

func Error(message interface{}, fields ...zapcore.Field) {
	switch t := message.(type) {
	case error:
		logs.Error(t.Error(), fields...)
	case string:
		logs.Error(t, fields...)
	}
}

func Fatal(message string, fields ...zapcore.Field) {
	logs.Fatal(message, fields...)
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
	const (
		INFO  = "INFO"
		DEBUG = "DEBUG"
		ERROR = "ERROR"
		FATAL = "FATAL"
	)
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
