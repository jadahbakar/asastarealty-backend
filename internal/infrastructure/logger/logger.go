package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugaredLogger *zap.SugaredLogger

func New() *zap.SugaredLogger {
	fileLogger := &lumberjack.Logger{
		Filename:   "./log/exampe.log",
		MaxSize:    10, // MB
		MaxBackups: 5,
		MaxAge:     30,   //days
		Compress:   true, // disabled by default
	}
	fileSyncer := zapcore.AddSync(fileLogger)
	loggerEncoderConfig := zap.NewProductionEncoderConfig()
	loggerEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	loggerEncoder := zapcore.NewConsoleEncoder(loggerEncoderConfig)

	core := zapcore.NewCore(loggerEncoder, fileSyncer, zapcore.DebugLevel)
	logger := zap.New(core)
	defer logger.Sync()

	sugaredLogger = logger.Sugar()
	return sugaredLogger
}
