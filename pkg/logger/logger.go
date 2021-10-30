package logger

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// Debug ...
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf ...
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// Info ...
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof ...
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Warn ...
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf ...
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

// Error ...
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf ...
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// DPanic ...
func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

// DPanicf ...
func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}

// Panic ...
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf ...
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

// Fatal ...
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf ...
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
