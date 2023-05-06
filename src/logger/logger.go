package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger
var appContextFields []zap.Field

func InitLogger() {
	var err error
	zapConfig := zap.NewProductionConfig()
	enccoderConfig := zap.NewProductionEncoderConfig()
	enccoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.EncoderConfig = enccoderConfig
	appContextFields = []zap.Field{
		zap.String("source", "lists"),
		// zap.String("version", "1.0.0"),
		// zap.String("environment", config.AppConfig.Environment),
	}
	zapLog, err = zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	fields = append(fields, appContextFields...)
	zapLog.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	fields = append(fields, appContextFields...)
	zapLog.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	fields = append(fields, appContextFields...)
	zapLog.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	fields = append(fields, appContextFields...)
	zapLog.Fatal(message, fields...)
}
