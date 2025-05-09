package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	zapLogger *zap.Logger
}

func New() *Logger {
	var l Logger

	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = true

	logger, err := config.Build()
	if err != nil {
		fmt.Println("Failed to initiate logger", err)
		return nil
	}

	l = Logger{
		zapLogger: logger,
	}

	return &l
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.zapLogger.Fatal(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.zapLogger.Panic(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.zapLogger.Sync()
}
