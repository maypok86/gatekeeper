package logger

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"errors": zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

type Config struct {
	Level string `envconfig:"LOGGER_LEVEL,default=info"`
	File  string `envconfig:"LOGGER_FILE,default=develop.log"`
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// Configure logger.
func Configure(config *Config) {
	level := getLoggerLevel(strings.ToLower(config.Level))

	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  config.File,
		MaxSize:   1 << 30,
		LocalTime: true,
		Compress:  true,
	})

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05 -0700"))
	}

	multiWriter := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(syncWriter))
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), multiWriter, zap.NewAtomicLevelAt(level))
	logger = zap.New(core).Sugar()

	Info("Logger init success.")
}
