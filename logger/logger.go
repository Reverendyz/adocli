package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
)

func InitLogger(debug bool) error {
	var config zap.Config

	if debug {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	config.EncoderConfig.MessageKey = "message"
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	logger, err := config.Build(
		zap.AddCallerSkip(1),
	)
	if err != nil {
		return err
	}

	Log = logger
	return nil
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

func Info(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Info(msg, fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Debug(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Warn(msg, fields...)
	}
}

// Error logs an error message
func Error(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Error(msg, fields...)
	}
}

// Fatal logs a fatal message and exits
func Fatal(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Fatal(msg, fields...)
	}
	os.Exit(1)
}

// InfoF formats and logs an info message (convenience method for CLI output)
func InfoF(template string, args ...interface{}) {
	if Log != nil {
		Log.Sugar().Infof(template, args...)
	}
}

// DebugF formats and logs a debug message
func DebugF(template string, args ...interface{}) {
	if Log != nil {
		Log.Sugar().Debugf(template, args...)
	}
}

// WarnF formats and logs a warning message
func WarnF(template string, args ...interface{}) {
	if Log != nil {
		Log.Sugar().Warnf(template, args...)
	}
}

// ErrorF formats and logs an error message
func ErrorF(template string, args ...interface{}) {
	if Log != nil {
		Log.Sugar().Errorf(template, args...)
	}
}

// FatalF formats and logs a fatal message and exits
func FatalF(template string, args ...interface{}) {
	if Log != nil {
		Log.Sugar().Fatalf(template, args...)
	}
	os.Exit(1)
}
