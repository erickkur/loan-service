package logger

import (
	"go.uber.org/zap"
)

type Interface interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Sync() error
}
